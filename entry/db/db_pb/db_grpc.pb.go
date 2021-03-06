// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package db_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DbClient is the client API for Db service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*TruncateResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
	RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*RenameTableResponse, error)
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	DropTable(ctx context.Context, in *DropTableRequest, opts ...grpc.CallOption) (*DropTableResponse, error)
}

type dbClient struct {
	cc grpc.ClientConnInterface
}

func NewDbClient(cc grpc.ClientConnInterface) DbClient {
	return &dbClient{cc}
}

func (c *dbClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/db.Db/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/db.Db/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/db.Db/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/db.Db/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*TruncateResponse, error) {
	out := new(TruncateResponse)
	err := c.cc.Invoke(ctx, "/db.Db/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/db.Db/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*RenameTableResponse, error) {
	out := new(RenameTableResponse)
	err := c.cc.Invoke(ctx, "/db.Db/RenameTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := c.cc.Invoke(ctx, "/db.Db/ListTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) DropTable(ctx context.Context, in *DropTableRequest, opts ...grpc.CallOption) (*DropTableResponse, error) {
	out := new(DropTableResponse)
	err := c.cc.Invoke(ctx, "/db.Db/DropTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbServer is the server API for Db service.
// All implementations should embed UnimplementedDbServer
// for forward compatibility
type DbServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Truncate(context.Context, *TruncateRequest) (*TruncateResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
	RenameTable(context.Context, *RenameTableRequest) (*RenameTableResponse, error)
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	DropTable(context.Context, *DropTableRequest) (*DropTableResponse, error)
}

// UnimplementedDbServer should be embedded to have forward compatible implementations.
type UnimplementedDbServer struct {
}

func (UnimplementedDbServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDbServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedDbServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDbServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDbServer) Truncate(context.Context, *TruncateRequest) (*TruncateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Truncate not implemented")
}
func (UnimplementedDbServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedDbServer) RenameTable(context.Context, *RenameTableRequest) (*RenameTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameTable not implemented")
}
func (UnimplementedDbServer) ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTables not implemented")
}
func (UnimplementedDbServer) DropTable(context.Context, *DropTableRequest) (*DropTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTable not implemented")
}

// UnsafeDbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbServer will
// result in compilation errors.
type UnsafeDbServer interface {
	mustEmbedUnimplementedDbServer()
}

func RegisterDbServer(s grpc.ServiceRegistrar, srv DbServer) {
	s.RegisterService(&_Db_serviceDesc, srv)
}

func _Db_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruncateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Truncate(ctx, req.(*TruncateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_RenameTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).RenameTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/RenameTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).RenameTable(ctx, req.(*RenameTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_DropTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).DropTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Db/DropTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).DropTable(ctx, req.(*DropTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Db_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.Db",
	HandlerType: (*DbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Db_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Db_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Db_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Db_Delete_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _Db_Truncate_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _Db_Count_Handler,
		},
		{
			MethodName: "RenameTable",
			Handler:    _Db_RenameTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _Db_ListTables_Handler,
		},
		{
			MethodName: "DropTable",
			Handler:    _Db_DropTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entry/db/db_pb/db.proto",
}
