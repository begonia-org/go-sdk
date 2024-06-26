// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: abac.proto

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
	ABACService_Auth_FullMethodName         = "/begonia.org.go.access.control.api.v1.ABACService/Auth"
	ABACService_PolicyPut_FullMethodName    = "/begonia.org.go.access.control.api.v1.ABACService/PolicyPut"
	ABACService_PolicyPatch_FullMethodName  = "/begonia.org.go.access.control.api.v1.ABACService/PolicyPatch"
	ABACService_PolicyDelete_FullMethodName = "/begonia.org.go.access.control.api.v1.ABACService/PolicyDelete"
	ABACService_PolicyGet_FullMethodName    = "/begonia.org.go.access.control.api.v1.ABACService/PolicyGet"
	ABACService_PolicyList_FullMethodName   = "/begonia.org.go.access.control.api.v1.ABACService/PolicyList"
)

// ABACServiceClient is the client API for ABACService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ABACServiceClient interface {
	Auth(ctx context.Context, in *AccessContext, opts ...grpc.CallOption) (*AccessResponse, error)
	PolicyPut(ctx context.Context, in *PutPolicyRequest, opts ...grpc.CallOption) (*PutPolicyResponse, error)
	PolicyPatch(ctx context.Context, in *PutPolicyRequest, opts ...grpc.CallOption) (*PatchPolicyResponse, error)
	PolicyDelete(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	PolicyGet(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	PolicyList(ctx context.Context, in *Policy, opts ...grpc.CallOption) (*Policy, error)
}

type aBACServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewABACServiceClient(cc grpc.ClientConnInterface) ABACServiceClient {
	return &aBACServiceClient{cc}
}

func (c *aBACServiceClient) Auth(ctx context.Context, in *AccessContext, opts ...grpc.CallOption) (*AccessResponse, error) {
	out := new(AccessResponse)
	err := c.cc.Invoke(ctx, ABACService_Auth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBACServiceClient) PolicyPut(ctx context.Context, in *PutPolicyRequest, opts ...grpc.CallOption) (*PutPolicyResponse, error) {
	out := new(PutPolicyResponse)
	err := c.cc.Invoke(ctx, ABACService_PolicyPut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBACServiceClient) PolicyPatch(ctx context.Context, in *PutPolicyRequest, opts ...grpc.CallOption) (*PatchPolicyResponse, error) {
	out := new(PatchPolicyResponse)
	err := c.cc.Invoke(ctx, ABACService_PolicyPatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBACServiceClient) PolicyDelete(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, ABACService_PolicyDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBACServiceClient) PolicyGet(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, ABACService_PolicyGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBACServiceClient) PolicyList(ctx context.Context, in *Policy, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, ABACService_PolicyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABACServiceServer is the server API for ABACService service.
// All implementations must embed UnimplementedABACServiceServer
// for forward compatibility
type ABACServiceServer interface {
	Auth(context.Context, *AccessContext) (*AccessResponse, error)
	PolicyPut(context.Context, *PutPolicyRequest) (*PutPolicyResponse, error)
	PolicyPatch(context.Context, *PutPolicyRequest) (*PatchPolicyResponse, error)
	PolicyDelete(context.Context, *PolicyRequest) (*Policy, error)
	PolicyGet(context.Context, *PolicyRequest) (*Policy, error)
	PolicyList(context.Context, *Policy) (*Policy, error)
	mustEmbedUnimplementedABACServiceServer()
}

// UnimplementedABACServiceServer must be embedded to have forward compatible implementations.
type UnimplementedABACServiceServer struct {
}

func (UnimplementedABACServiceServer) Auth(context.Context, *AccessContext) (*AccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedABACServiceServer) PolicyPut(context.Context, *PutPolicyRequest) (*PutPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyPut not implemented")
}
func (UnimplementedABACServiceServer) PolicyPatch(context.Context, *PutPolicyRequest) (*PatchPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyPatch not implemented")
}
func (UnimplementedABACServiceServer) PolicyDelete(context.Context, *PolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyDelete not implemented")
}
func (UnimplementedABACServiceServer) PolicyGet(context.Context, *PolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyGet not implemented")
}
func (UnimplementedABACServiceServer) PolicyList(context.Context, *Policy) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyList not implemented")
}
func (UnimplementedABACServiceServer) mustEmbedUnimplementedABACServiceServer() {}

// UnsafeABACServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ABACServiceServer will
// result in compilation errors.
type UnsafeABACServiceServer interface {
	mustEmbedUnimplementedABACServiceServer()
}

func RegisterABACServiceServer(s grpc.ServiceRegistrar, srv ABACServiceServer) {
	s.RegisterService(&ABACService_ServiceDesc, srv)
}

func _ABACService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABACServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABACService_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABACServiceServer).Auth(ctx, req.(*AccessContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABACService_PolicyPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABACServiceServer).PolicyPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABACService_PolicyPut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABACServiceServer).PolicyPut(ctx, req.(*PutPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABACService_PolicyPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABACServiceServer).PolicyPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABACService_PolicyPatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABACServiceServer).PolicyPatch(ctx, req.(*PutPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABACService_PolicyDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABACServiceServer).PolicyDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABACService_PolicyDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABACServiceServer).PolicyDelete(ctx, req.(*PolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABACService_PolicyGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABACServiceServer).PolicyGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABACService_PolicyGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABACServiceServer).PolicyGet(ctx, req.(*PolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABACService_PolicyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Policy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABACServiceServer).PolicyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABACService_PolicyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABACServiceServer).PolicyList(ctx, req.(*Policy))
	}
	return interceptor(ctx, in, info, handler)
}

// ABACService_ServiceDesc is the grpc.ServiceDesc for ABACService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ABACService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "begonia.org.go.access.control.api.v1.ABACService",
	HandlerType: (*ABACServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _ABACService_Auth_Handler,
		},
		{
			MethodName: "PolicyPut",
			Handler:    _ABACService_PolicyPut_Handler,
		},
		{
			MethodName: "PolicyPatch",
			Handler:    _ABACService_PolicyPatch_Handler,
		},
		{
			MethodName: "PolicyDelete",
			Handler:    _ABACService_PolicyDelete_Handler,
		},
		{
			MethodName: "PolicyGet",
			Handler:    _ABACService_PolicyGet_Handler,
		},
		{
			MethodName: "PolicyList",
			Handler:    _ABACService_PolicyList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "abac.proto",
}
