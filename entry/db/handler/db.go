package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/micro/services/pkg/tenant"
	cache "github.com/patrickmn/go-cache"
	"github.com/pubgo/lava/clients/orm"
	"github.com/pubgo/lava/errors"
	"github.com/pubgo/lava/logger"
	"github.com/pubgo/xerror"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/lavaxx/services/entry/db/db_pb"
)

const idKey = "id"
const stmt = "create table if not exists %v(id text not null, data jsonb, primary key(id)); alter table %v add created_at timestamptz; alter table %v add updated_at timestamptz"
const truncateStmt = `truncate table "%v"`
const dropTableStmt = `drop table "%v"`
const renameTableStmt = `ALTER TABLE "%v" RENAME TO "%v"`

var re = regexp.MustCompile("^[a-zA-Z0-9_]*$")
var c = cache.New(5*time.Minute, 10*time.Minute)

type Record struct {
	ID   string
	Data datatypes.JSON `json:"data"`
	// private field, ignored from gorm
	table     string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var _ db_pb.DbServer = (*Db)(nil)

type Db struct {
	Db *orm.Client `dix:""`
}

func (e *Db) Init() {
	xerror.Panic(e.Db.Ping())
}

func correctFieldName(s string, isText bool) string {
	operator := "->"
	if isText {
		// https: //stackoverflow.com/questions/27215216/postgres-how-to-convert-a-json-string-to-text
		operator = "->>"
	}
	switch s {
	// top level fields can stay top level
	case "id": // "created_at", "updated_at",  <-- these are not special fields for now
		return s
	}
	if !strings.Contains(s, ".") {
		return fmt.Sprintf("data %v '%v'", operator, s)
	}
	paths := strings.Split(s, ".")
	ret := "data"
	for i, path := range paths {
		if i == len(paths)-1 && isText {
			ret += fmt.Sprintf(" ->> '%v'", path)
			break
		}
		ret += fmt.Sprintf(" -> '%v'", path)
	}
	return ret
}

func (e *Db) tableName(ctx context.Context, t string) (string, error) {
	tenantId, ok := tenant.FromContext(ctx)
	if !ok {
		tenantId = "micro"
	}
	if t == "" {
		t = "default"
	}
	t = strings.ToLower(t)
	t = strings.Replace(t, "-", "_", -1)
	tenantId = strings.Replace(strings.Replace(tenantId, "/", "_", -1), "-", "_", -1)

	tableName := tenantId + "_" + t
	if !re.Match([]byte(tableName)) {
		return "", fmt.Errorf("table name %v is invalid", t)
	}

	return tableName, nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Db) Create(ctx context.Context, req *db_pb.CreateRequest) (*db_pb.CreateResponse, error) {
	var log = logger.GetLog(ctx)
	if len(req.Record.AsMap()) == 0 {
		return nil, errors.BadRequest("db.create", "missing record")
	}

	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}
	log.Infof("Inserting into table '%v'", tableName)

	db := e.Db.WithContext(ctx)
	_, ok := c.Get(tableName)
	if !ok {
		log.Infof("Creating table '%v'", tableName)
		db.Exec(fmt.Sprintf(stmt, tableName, tableName, tableName))
		c.Set(tableName, true, 0)
	}

	m := req.Record.AsMap()
	if _, ok := m[idKey].(string); !ok {
		m[idKey] = uuid.New().String()
	}
	bs, _ := json.Marshal(m)

	err = db.Table(tableName).Create(&Record{
		ID:   m[idKey].(string),
		Data: bs,
	}).Error
	if err != nil {
		return nil, err
	}

	return &db_pb.CreateResponse{Id: m[idKey].(string)}, nil
}

func (e *Db) Update(ctx context.Context, req *db_pb.UpdateRequest) (*db_pb.UpdateResponse, error) {
	var log = logger.GetLog(ctx)

	if len(req.Record.AsMap()) == 0 {
		return nil, errors.BadRequest("db.update", "missing record")
	}
	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}
	log.Infof("Updating table '%v'", tableName)

	db := e.Db.WithContext(ctx)
	m := req.Record.AsMap()

	// where ID is specified do a single update record update
	id := req.Id
	if v, ok := m[idKey].(string); ok && id == "" {
		id = v
	}

	// if the id is blank then check the data
	if len(req.Id) == 0 {
		var ok bool
		id, ok = m[idKey].(string)
		if !ok {
			return nil, fmt.Errorf("update failed: missing id")
		}
	}

	return nil, db.Transaction(func(tx *gorm.DB) error {
		rec := []Record{}
		err = tx.Table(tableName).Where("id = ?", id).Find(&rec).Error
		if err != nil {
			return err
		}
		if len(rec) == 0 {
			return fmt.Errorf("update failed: not found")
		}
		old := map[string]interface{}{}
		err = json.Unmarshal(rec[0].Data, &old)
		if err != nil {
			return err
		}
		for k, v := range m {
			old[k] = v
		}
		bs, _ := json.Marshal(old)

		return tx.Table(tableName).Save(&Record{
			ID:   id,
			Data: bs,
		}).Error
	})
}

func (e *Db) Read(ctx context.Context, req *db_pb.ReadRequest) (*db_pb.ReadResponse, error) {
	var log = logger.GetLog(ctx)

	var recs []Record
	queries, err := Parse(req.Query)
	if err != nil {
		return nil, err
	}
	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}

	db := e.Db.WithContext(ctx)
	_, ok := c.Get(tableName)
	if !ok {
		log.Infof("Creating table '%v'", tableName)
		db.Exec(fmt.Sprintf(stmt, tableName, tableName, tableName))
		c.Set(tableName, true, 0)
	}

	if req.Limit > 1000 {
		return nil, errors.BadRequest("db.read", fmt.Sprintf("limit over 1000 is invalid, you specified %v", req.Limit))
	}
	if req.Limit == 0 {
		req.Limit = 25
	}

	db = db.Table(tableName)
	if req.Id != "" {
		log.Infof("Query by id: %v", req.Id)
		db = db.Where("id = ?", req.Id)
	} else {
		for _, query := range queries {
			log.Infof("Query field: %v, op: %v, value: %v", query.Field, query.Op, query.Value)
			typ := "text"
			switch query.Value.(type) {
			case int64:
				typ = "int"
			case bool:
				typ = "boolean"
			}
			op := ""
			switch query.Op {
			case itemEquals:
				op = "="
			case itemGreaterThan:
				op = ">"
			case itemGreaterThanEquals:
				op = ">="
			case itemLessThan:
				op = "<"
			case itemLessThanEquals:
				op = "<="
			case itemNotEquals:
				op = "!="
			}
			queryField := correctFieldName(query.Field, typ == "text")
			db = db.Where(fmt.Sprintf("(%v)::%v %v ?", queryField, typ, op), query.Value)
		}
	}

	orderField := "created_at"
	if req.OrderBy != "" {
		orderField = req.OrderBy
	}
	orderField = correctFieldName(orderField, false)

	ordering := "asc"
	if req.Order != "" {
		switch strings.ToLower(req.Order) {
		case "asc":
			ordering = "asc"
		case "", "desc":
			ordering = "desc"
		default:
			return nil, errors.BadRequest("db.read", "invalid ordering: "+req.Order)
		}
	}

	db = db.Order(orderField + " " + ordering).Offset(int(req.Offset)).Limit(int(req.Limit))
	err = db.Debug().Find(&recs).Error
	if err != nil {
		return nil, err
	}

	rsp := new(db_pb.ReadResponse)
	rsp.Records = []*structpb.Struct{}
	for _, rec := range recs {
		m, err := rec.Data.MarshalJSON()
		if err != nil {
			return nil, err
		}
		ma := map[string]interface{}{}
		json.Unmarshal(m, &ma)
		ma[idKey] = rec.ID
		m, _ = json.Marshal(ma)
		s := &structpb.Struct{}
		err = s.UnmarshalJSON(m)
		if err != nil {
			return nil, err
		}
		rsp.Records = append(rsp.Records, s)
	}

	return rsp, nil
}

func (e *Db) Delete(ctx context.Context, req *db_pb.DeleteRequest) (*db_pb.DeleteResponse, error) {
	var log = logger.GetLog(ctx)
	if len(req.Id) == 0 {
		return nil, errors.BadRequest("db.delete", "missing id")
	}
	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}
	log.Infof("Deleting from table '%v'", tableName)

	db := e.Db.WithContext(ctx)
	return nil, db.Table(tableName).Delete(Record{
		ID: req.Id,
	}).Error
}

func (e *Db) Truncate(ctx context.Context, req *db_pb.TruncateRequest) (*db_pb.TruncateResponse, error) {
	var log = logger.GetLog(ctx)
	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}
	log.Infof("Truncating table '%v'", tableName)

	db := e.Db.WithContext(ctx)
	return nil, db.Exec(fmt.Sprintf(truncateStmt, tableName)).Error
}

func (e *Db) DropTable(ctx context.Context, req *db_pb.DropTableRequest) (*db_pb.DropTableResponse, error) {
	var log = logger.GetLog(ctx)
	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}
	log.Infof("Dropping table '%v'", tableName)

	db := e.Db.WithContext(ctx)
	return nil, db.Exec(fmt.Sprintf(dropTableStmt, tableName)).Error
}

func (e *Db) Count(ctx context.Context, req *db_pb.CountRequest) (*db_pb.CountResponse, error) {
	if req.Table == "" {
		req.Table = "default"
	}

	tableName, err := e.tableName(ctx, req.Table)
	if err != nil {
		return nil, err
	}

	db := e.Db.WithContext(ctx)

	var a int64
	err = db.Table(tableName).Model(Record{}).Count(&a).Error
	if err != nil {
		return nil, err
	}

	rsp := new(db_pb.CountResponse)
	rsp.Count = int32(a)
	return rsp, nil
}

func (e *Db) RenameTable(ctx context.Context, req *db_pb.RenameTableRequest) (*db_pb.RenameTableResponse, error) {
	var log = logger.GetLog(ctx)
	if req.From == "" || req.To == "" {
		return nil, errors.BadRequest("db.renameTable", "must provide table names")
	}

	oldtableName, err := e.tableName(ctx, req.From)
	if err != nil {
		return nil, err
	}

	newtableName, err := e.tableName(ctx, req.To)
	if err != nil {
		return nil, err
	}

	db := e.Db.WithContext(ctx)

	stmt := fmt.Sprintf(renameTableStmt, oldtableName, newtableName)
	log.Info(stmt)
	return nil, db.Debug().Exec(stmt).Error
}

func (e *Db) ListTables(ctx context.Context, req *db_pb.ListTablesRequest) (*db_pb.ListTablesResponse, error) {
	tenantId, ok := tenant.FromContext(ctx)
	if !ok {
		tenantId = "micro"
	}
	tenantId = strings.Replace(strings.Replace(tenantId, "/", "_", -1), "-", "_", -1)

	db := e.Db.WithContext(ctx)

	var tables []string
	if err := db.Table("information_schema.tables").Select("table_name").Where("table_schema = ?", "public").Find(&tables).Error; err != nil {
		return nil, err
	}

	rsp := new(db_pb.ListTablesResponse)
	rsp.Tables = []string{}
	for _, v := range tables {
		if strings.HasPrefix(v, tenantId) {
			rsp.Tables = append(rsp.Tables, strings.Replace(v, tenantId+"_", "", -1))
		}
	}
	return rsp, nil
}
