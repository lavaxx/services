package db

import (
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/entry/grpcEntry"

	"github.com/lavaxx/services/entry/db/handler"
)

var Name = "db"

func GetEntry() entry.Entry {
	ent := grpcEntry.New(Name)
	ent.Description("db services")
	ent.Register(new(handler.Db))
	return ent
}
