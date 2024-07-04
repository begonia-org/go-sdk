// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: tenant.proto

package v1

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

const (
	TenantsService_Register_FullMethodName             = "/begonia.org.sdk.TenantsService/Register"
	TenantsService_Update_FullMethodName               = "/begonia.org.sdk.TenantsService/Update"
	TenantsService_Get_FullMethodName                  = "/begonia.org.sdk.TenantsService/Get"
	TenantsService_Delete_FullMethodName               = "/begonia.org.sdk.TenantsService/Delete"
	TenantsService_List_FullMethodName                 = "/begonia.org.sdk.TenantsService/List"
	TenantsService_AddTenantBusiness_FullMethodName    = "/begonia.org.sdk.TenantsService/AddTenantBusiness"
	TenantsService_ListTenantBusiness_FullMethodName   = "/begonia.org.sdk.TenantsService/ListTenantBusiness"
	TenantsService_DeleteTenantBusiness_FullMethodName = "/begonia.org.sdk.TenantsService/DeleteTenantBusiness"
)

// TenantsServiceClient is the client API for TenantsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantsServiceClient interface {
	Register(ctx context.Context, in *PostTenantRequest, opts ...grpc.CallOption) (*Tenants, error)
	Update(ctx context.Context, in *PatchTenantRequest, opts ...grpc.CallOption) (*Tenants, error)
	Get(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenants, error)
	Delete(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error)
	List(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	AddTenantBusiness(ctx context.Context, in *AddTenantBusinessRequest, opts ...grpc.CallOption) (*TenantsBusiness, error)
	ListTenantBusiness(ctx context.Context, in *ListTenantBusinessRequest, opts ...grpc.CallOption) (*ListTenantBusinessResponse, error)
	DeleteTenantBusiness(ctx context.Context, in *DeleteTenantBusinessRequest, opts ...grpc.CallOption) (*DeleteTenantBusinessResponse, error)
}

type tenantsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantsServiceClient(cc grpc.ClientConnInterface) TenantsServiceClient {
	return &tenantsServiceClient{cc}
}

func (c *tenantsServiceClient) Register(ctx context.Context, in *PostTenantRequest, opts ...grpc.CallOption) (*Tenants, error) {
	out := new(Tenants)
	err := c.cc.Invoke(ctx, TenantsService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) Update(ctx context.Context, in *PatchTenantRequest, opts ...grpc.CallOption) (*Tenants, error) {
	out := new(Tenants)
	err := c.cc.Invoke(ctx, TenantsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) Get(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenants, error) {
	out := new(Tenants)
	err := c.cc.Invoke(ctx, TenantsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) Delete(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error) {
	out := new(DeleteTenantResponse)
	err := c.cc.Invoke(ctx, TenantsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) List(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, TenantsService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) AddTenantBusiness(ctx context.Context, in *AddTenantBusinessRequest, opts ...grpc.CallOption) (*TenantsBusiness, error) {
	out := new(TenantsBusiness)
	err := c.cc.Invoke(ctx, TenantsService_AddTenantBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) ListTenantBusiness(ctx context.Context, in *ListTenantBusinessRequest, opts ...grpc.CallOption) (*ListTenantBusinessResponse, error) {
	out := new(ListTenantBusinessResponse)
	err := c.cc.Invoke(ctx, TenantsService_ListTenantBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) DeleteTenantBusiness(ctx context.Context, in *DeleteTenantBusinessRequest, opts ...grpc.CallOption) (*DeleteTenantBusinessResponse, error) {
	out := new(DeleteTenantBusinessResponse)
	err := c.cc.Invoke(ctx, TenantsService_DeleteTenantBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantsServiceServer is the server API for TenantsService service.
// All implementations must embed UnimplementedTenantsServiceServer
// for forward compatibility
type TenantsServiceServer interface {
	Register(context.Context, *PostTenantRequest) (*Tenants, error)
	Update(context.Context, *PatchTenantRequest) (*Tenants, error)
	Get(context.Context, *GetTenantRequest) (*Tenants, error)
	Delete(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error)
	List(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	AddTenantBusiness(context.Context, *AddTenantBusinessRequest) (*TenantsBusiness, error)
	ListTenantBusiness(context.Context, *ListTenantBusinessRequest) (*ListTenantBusinessResponse, error)
	DeleteTenantBusiness(context.Context, *DeleteTenantBusinessRequest) (*DeleteTenantBusinessResponse, error)
	mustEmbedUnimplementedTenantsServiceServer()
}

// UnimplementedTenantsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTenantsServiceServer struct {
}

func (UnimplementedTenantsServiceServer) Register(context.Context, *PostTenantRequest) (*Tenants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedTenantsServiceServer) Update(context.Context, *PatchTenantRequest) (*Tenants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTenantsServiceServer) Get(context.Context, *GetTenantRequest) (*Tenants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTenantsServiceServer) Delete(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTenantsServiceServer) List(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTenantsServiceServer) AddTenantBusiness(context.Context, *AddTenantBusinessRequest) (*TenantsBusiness, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTenantBusiness not implemented")
}
func (UnimplementedTenantsServiceServer) ListTenantBusiness(context.Context, *ListTenantBusinessRequest) (*ListTenantBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenantBusiness not implemented")
}
func (UnimplementedTenantsServiceServer) DeleteTenantBusiness(context.Context, *DeleteTenantBusinessRequest) (*DeleteTenantBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenantBusiness not implemented")
}
func (UnimplementedTenantsServiceServer) mustEmbedUnimplementedTenantsServiceServer() {}

// UnsafeTenantsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantsServiceServer will
// result in compilation errors.
type UnsafeTenantsServiceServer interface {
	mustEmbedUnimplementedTenantsServiceServer()
}

func RegisterTenantsServiceServer(s grpc.ServiceRegistrar, srv TenantsServiceServer) {
	s.RegisterService(&TenantsService_ServiceDesc, srv)
}

func _TenantsService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).Register(ctx, req.(*PostTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).Update(ctx, req.(*PatchTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).Get(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).Delete(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).List(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_AddTenantBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTenantBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).AddTenantBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_AddTenantBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).AddTenantBusiness(ctx, req.(*AddTenantBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_ListTenantBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).ListTenantBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_ListTenantBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).ListTenantBusiness(ctx, req.(*ListTenantBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_DeleteTenantBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).DeleteTenantBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_DeleteTenantBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).DeleteTenantBusiness(ctx, req.(*DeleteTenantBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantsService_ServiceDesc is the grpc.ServiceDesc for TenantsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "begonia.org.sdk.TenantsService",
	HandlerType: (*TenantsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _TenantsService_Register_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TenantsService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TenantsService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TenantsService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TenantsService_List_Handler,
		},
		{
			MethodName: "AddTenantBusiness",
			Handler:    _TenantsService_AddTenantBusiness_Handler,
		},
		{
			MethodName: "ListTenantBusiness",
			Handler:    _TenantsService_ListTenantBusiness_Handler,
		},
		{
			MethodName: "DeleteTenantBusiness",
			Handler:    _TenantsService_DeleteTenantBusiness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant.proto",
}