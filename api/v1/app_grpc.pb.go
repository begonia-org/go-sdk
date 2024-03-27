// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: app.proto

package v1

import (
	context "context"
	v1 "github.com/begonia-org/begonia-go-sdk/common/api/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AppsService_AddApps_FullMethodName   = "/begonia.org.begonia.AppsService/AddApps"
	AppsService_CreateApp_FullMethodName = "/begonia.org.begonia.AppsService/CreateApp"
	AppsService_GetApps_FullMethodName   = "/begonia.org.begonia.AppsService/GetApps"
)

// AppsServiceClient is the client API for AppsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppsServiceClient interface {
	// @gotags: doc:"添加app服务"
	AddApps(ctx context.Context, in *AddAppsRequest, opts ...grpc.CallOption) (*v1.APIResponse, error)
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	// @gotags: doc:"获取app服务"
	GetApps(ctx context.Context, in *AppsListRequest, opts ...grpc.CallOption) (*v1.APIResponse, error)
}

type appsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppsServiceClient(cc grpc.ClientConnInterface) AppsServiceClient {
	return &appsServiceClient{cc}
}

func (c *appsServiceClient) AddApps(ctx context.Context, in *AddAppsRequest, opts ...grpc.CallOption) (*v1.APIResponse, error) {
	out := new(v1.APIResponse)
	err := c.cc.Invoke(ctx, AppsService_AddApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	out := new(CreateAppResponse)
	err := c.cc.Invoke(ctx, AppsService_CreateApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) GetApps(ctx context.Context, in *AppsListRequest, opts ...grpc.CallOption) (*v1.APIResponse, error) {
	out := new(v1.APIResponse)
	err := c.cc.Invoke(ctx, AppsService_GetApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppsServiceServer is the server API for AppsService service.
// All implementations must embed UnimplementedAppsServiceServer
// for forward compatibility
type AppsServiceServer interface {
	// @gotags: doc:"添加app服务"
	AddApps(context.Context, *AddAppsRequest) (*v1.APIResponse, error)
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	// @gotags: doc:"获取app服务"
	GetApps(context.Context, *AppsListRequest) (*v1.APIResponse, error)
	mustEmbedUnimplementedAppsServiceServer()
}

// UnimplementedAppsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppsServiceServer struct {
}

func (UnimplementedAppsServiceServer) AddApps(context.Context, *AddAppsRequest) (*v1.APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddApps not implemented")
}
func (UnimplementedAppsServiceServer) CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (UnimplementedAppsServiceServer) GetApps(context.Context, *AppsListRequest) (*v1.APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedAppsServiceServer) mustEmbedUnimplementedAppsServiceServer() {}

// UnsafeAppsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppsServiceServer will
// result in compilation errors.
type UnsafeAppsServiceServer interface {
	mustEmbedUnimplementedAppsServiceServer()
}

func RegisterAppsServiceServer(s grpc.ServiceRegistrar, srv AppsServiceServer) {
	s.RegisterService(&AppsService_ServiceDesc, srv)
}

func _AppsService_AddApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).AddApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_AddApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).AddApps(ctx, req.(*AddAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_CreateApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_GetApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).GetApps(ctx, req.(*AppsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppsService_ServiceDesc is the grpc.ServiceDesc for AppsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "begonia.org.begonia.AppsService",
	HandlerType: (*AppsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddApps",
			Handler:    _AppsService_AddApps_Handler,
		},
		{
			MethodName: "CreateApp",
			Handler:    _AppsService_CreateApp_Handler,
		},
		{
			MethodName: "GetApps",
			Handler:    _AppsService_GetApps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
