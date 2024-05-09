// Please use the following editor setup for this file:
// Tab size=2; Tabs as spaces; Clean up trailing whitespace
//
// In vim add: au FileType proto setl sw=2 ts=2 expandtab list
//
// In vscode install vscode-proto3 extension and add this to your settings.json:
//    "[proto3]": {
//        "editor.tabSize": 2,
//        "editor.insertSpaces": true,
//        "editor.rulers": [80],
//        "editor.detectIndentation": true,
//        "files.trimTrailingWhitespace": true
//    }
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: public/portworx/platform/tenant/apiv1/tenant.proto

package tenant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TenantService_GetTenant_FullMethodName    = "/public.portworx.platform.tenant.v1.TenantService/GetTenant"
	TenantService_CreateTenant_FullMethodName = "/public.portworx.platform.tenant.v1.TenantService/CreateTenant"
	TenantService_UpdateTenant_FullMethodName = "/public.portworx.platform.tenant.v1.TenantService/UpdateTenant"
	TenantService_DeleteTenant_FullMethodName = "/public.portworx.platform.tenant.v1.TenantService/DeleteTenant"
	TenantService_ListTenants_FullMethodName  = "/public.portworx.platform.tenant.v1.TenantService/ListTenants"
)

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantServiceClient interface {
	// GetTenant API returns the info about  for given tenant id.
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// CreateTenant API creates a new Tenant.
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// UpdateTenant API updates tenant.
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Delete tenant removes a tenant record.
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListTenants API lists the tenants visible to the caller for the current account.
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, TenantService_GetTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, TenantService_CreateTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, TenantService_UpdateTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TenantService_DeleteTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, TenantService_ListTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations must embed UnimplementedTenantServiceServer
// for forward compatibility
type TenantServiceServer interface {
	// GetTenant API returns the info about  for given tenant id.
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	// CreateTenant API creates a new Tenant.
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	// UpdateTenant API updates tenant.
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	// Delete tenant removes a tenant record.
	DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error)
	// ListTenants API lists the tenants visible to the caller for the current account.
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	mustEmbedUnimplementedTenantServiceServer()
}

// UnimplementedTenantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (UnimplementedTenantServiceServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantServiceServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (UnimplementedTenantServiceServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantServiceServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}
func (UnimplementedTenantServiceServer) mustEmbedUnimplementedTenantServiceServer() {}

// UnsafeTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServiceServer will
// result in compilation errors.
type UnsafeTenantServiceServer interface {
	mustEmbedUnimplementedTenantServiceServer()
}

func RegisterTenantServiceServer(s grpc.ServiceRegistrar, srv TenantServiceServer) {
	s.RegisterService(&TenantService_ServiceDesc, srv)
}

func _TenantService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_GetTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_CreateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_UpdateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_DeleteTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_ListTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantService_ServiceDesc is the grpc.ServiceDesc for TenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.portworx.platform.tenant.v1.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenant",
			Handler:    _TenantService_GetTenant_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _TenantService_CreateTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _TenantService_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantService_DeleteTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _TenantService_ListTenants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/portworx/platform/tenant/apiv1/tenant.proto",
}
