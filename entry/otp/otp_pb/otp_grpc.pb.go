// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package otp_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// OtpClient is the client API for Otp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtpClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type otpClient struct {
	cc grpc.ClientConnInterface
}

func NewOtpClient(cc grpc.ClientConnInterface) OtpClient {
	return &otpClient{cc}
}

func (c *otpClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/otp.Otp/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otpClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/otp.Otp/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtpServer is the server API for Otp service.
// All implementations should embed UnimplementedOtpServer
// for forward compatibility
type OtpServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

// UnimplementedOtpServer should be embedded to have forward compatible implementations.
type UnimplementedOtpServer struct {
}

func (UnimplementedOtpServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedOtpServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

// UnsafeOtpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtpServer will
// result in compilation errors.
type UnsafeOtpServer interface {
	mustEmbedUnimplementedOtpServer()
}

func RegisterOtpServer(s grpc.ServiceRegistrar, srv OtpServer) {
	s.RegisterService(&_Otp_serviceDesc, srv)
}

func _Otp_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otp.Otp/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Otp_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otp.Otp/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Otp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "otp.Otp",
	HandlerType: (*OtpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Otp_Generate_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Otp_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entry/otp/otp_pb/otp.proto",
}
