// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: test.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SourceApiClient is the client API for SourceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceApiClient interface {
	Create(ctx context.Context, in *AddSourceRequest, opts ...grpc.CallOption) (*AddSourceResponse, error)
	FindAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FindAllSourceResponse, error)
	FindById(ctx context.Context, in *FindByIdSourceRequest, opts ...grpc.CallOption) (*FindByIdSourceResponse, error)
	Update(ctx context.Context, in *UpdateSourceRequest, opts ...grpc.CallOption) (*UpdateSourceResponse, error)
	Delete(ctx context.Context, in *DeleteSourceRequest, opts ...grpc.CallOption) (*DeleteSourceResponse, error)
}

type sourceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceApiClient(cc grpc.ClientConnInterface) SourceApiClient {
	return &sourceApiClient{cc}
}

func (c *sourceApiClient) Create(ctx context.Context, in *AddSourceRequest, opts ...grpc.CallOption) (*AddSourceResponse, error) {
	out := new(AddSourceResponse)
	err := c.cc.Invoke(ctx, "/test.sourceApi/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceApiClient) FindAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FindAllSourceResponse, error) {
	out := new(FindAllSourceResponse)
	err := c.cc.Invoke(ctx, "/test.sourceApi/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceApiClient) FindById(ctx context.Context, in *FindByIdSourceRequest, opts ...grpc.CallOption) (*FindByIdSourceResponse, error) {
	out := new(FindByIdSourceResponse)
	err := c.cc.Invoke(ctx, "/test.sourceApi/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceApiClient) Update(ctx context.Context, in *UpdateSourceRequest, opts ...grpc.CallOption) (*UpdateSourceResponse, error) {
	out := new(UpdateSourceResponse)
	err := c.cc.Invoke(ctx, "/test.sourceApi/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceApiClient) Delete(ctx context.Context, in *DeleteSourceRequest, opts ...grpc.CallOption) (*DeleteSourceResponse, error) {
	out := new(DeleteSourceResponse)
	err := c.cc.Invoke(ctx, "/test.sourceApi/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceApiServer is the server API for SourceApi service.
// All implementations must embed UnimplementedSourceApiServer
// for forward compatibility
type SourceApiServer interface {
	Create(context.Context, *AddSourceRequest) (*AddSourceResponse, error)
	FindAll(context.Context, *Empty) (*FindAllSourceResponse, error)
	FindById(context.Context, *FindByIdSourceRequest) (*FindByIdSourceResponse, error)
	Update(context.Context, *UpdateSourceRequest) (*UpdateSourceResponse, error)
	Delete(context.Context, *DeleteSourceRequest) (*DeleteSourceResponse, error)
	mustEmbedUnimplementedSourceApiServer()
}

// UnimplementedSourceApiServer must be embedded to have forward compatible implementations.
type UnimplementedSourceApiServer struct {
}

func (UnimplementedSourceApiServer) Create(context.Context, *AddSourceRequest) (*AddSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSourceApiServer) FindAll(context.Context, *Empty) (*FindAllSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedSourceApiServer) FindById(context.Context, *FindByIdSourceRequest) (*FindByIdSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedSourceApiServer) Update(context.Context, *UpdateSourceRequest) (*UpdateSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSourceApiServer) Delete(context.Context, *DeleteSourceRequest) (*DeleteSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSourceApiServer) mustEmbedUnimplementedSourceApiServer() {}

// UnsafeSourceApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceApiServer will
// result in compilation errors.
type UnsafeSourceApiServer interface {
	mustEmbedUnimplementedSourceApiServer()
}

func RegisterSourceApiServer(s grpc.ServiceRegistrar, srv SourceApiServer) {
	s.RegisterService(&SourceApi_ServiceDesc, srv)
}

func _SourceApi_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceApiServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.sourceApi/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceApiServer).Create(ctx, req.(*AddSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceApi_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceApiServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.sourceApi/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceApiServer).FindAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceApi_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceApiServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.sourceApi/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceApiServer).FindById(ctx, req.(*FindByIdSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.sourceApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceApiServer).Update(ctx, req.(*UpdateSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceApi_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceApiServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.sourceApi/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceApiServer).Delete(ctx, req.(*DeleteSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SourceApi_ServiceDesc is the grpc.ServiceDesc for SourceApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SourceApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.sourceApi",
	HandlerType: (*SourceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SourceApi_Create_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _SourceApi_FindAll_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _SourceApi_FindById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SourceApi_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SourceApi_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

// PaperApiClient is the client API for PaperApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaperApiClient interface {
	Start(ctx context.Context, in *StartTestRequest, opts ...grpc.CallOption) (*StartTestResponse, error)
	FindById(ctx context.Context, in *FindByIdPaperRequest, opts ...grpc.CallOption) (*FindByIdPaperResponse, error)
	Update(ctx context.Context, in *UpdatePaperRequest, opts ...grpc.CallOption) (*UpdatePaperResponse, error)
}

type paperApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPaperApiClient(cc grpc.ClientConnInterface) PaperApiClient {
	return &paperApiClient{cc}
}

func (c *paperApiClient) Start(ctx context.Context, in *StartTestRequest, opts ...grpc.CallOption) (*StartTestResponse, error) {
	out := new(StartTestResponse)
	err := c.cc.Invoke(ctx, "/test.paperApi/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paperApiClient) FindById(ctx context.Context, in *FindByIdPaperRequest, opts ...grpc.CallOption) (*FindByIdPaperResponse, error) {
	out := new(FindByIdPaperResponse)
	err := c.cc.Invoke(ctx, "/test.paperApi/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paperApiClient) Update(ctx context.Context, in *UpdatePaperRequest, opts ...grpc.CallOption) (*UpdatePaperResponse, error) {
	out := new(UpdatePaperResponse)
	err := c.cc.Invoke(ctx, "/test.paperApi/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaperApiServer is the server API for PaperApi service.
// All implementations must embed UnimplementedPaperApiServer
// for forward compatibility
type PaperApiServer interface {
	Start(context.Context, *StartTestRequest) (*StartTestResponse, error)
	FindById(context.Context, *FindByIdPaperRequest) (*FindByIdPaperResponse, error)
	Update(context.Context, *UpdatePaperRequest) (*UpdatePaperResponse, error)
	mustEmbedUnimplementedPaperApiServer()
}

// UnimplementedPaperApiServer must be embedded to have forward compatible implementations.
type UnimplementedPaperApiServer struct {
}

func (UnimplementedPaperApiServer) Start(context.Context, *StartTestRequest) (*StartTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedPaperApiServer) FindById(context.Context, *FindByIdPaperRequest) (*FindByIdPaperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedPaperApiServer) Update(context.Context, *UpdatePaperRequest) (*UpdatePaperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPaperApiServer) mustEmbedUnimplementedPaperApiServer() {}

// UnsafePaperApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaperApiServer will
// result in compilation errors.
type UnsafePaperApiServer interface {
	mustEmbedUnimplementedPaperApiServer()
}

func RegisterPaperApiServer(s grpc.ServiceRegistrar, srv PaperApiServer) {
	s.RegisterService(&PaperApi_ServiceDesc, srv)
}

func _PaperApi_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaperApiServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.paperApi/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaperApiServer).Start(ctx, req.(*StartTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaperApi_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdPaperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaperApiServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.paperApi/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaperApiServer).FindById(ctx, req.(*FindByIdPaperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaperApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaperApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.paperApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaperApiServer).Update(ctx, req.(*UpdatePaperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaperApi_ServiceDesc is the grpc.ServiceDesc for PaperApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaperApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.paperApi",
	HandlerType: (*PaperApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _PaperApi_Start_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _PaperApi_FindById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PaperApi_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}