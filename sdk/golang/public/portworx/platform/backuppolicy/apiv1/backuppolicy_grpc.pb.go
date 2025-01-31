// Please use the following editor setup for this file:
// Tab size=2; Tabs as spaces; Clean up trailing whitepsace
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
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: public/portworx/platform/backuppolicy/apiv1/backuppolicy.proto

package backuppolicy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BackupPolicyService_GetBackupPolicy_FullMethodName        = "/public.portworx.platform.backuppolicy.v1.BackupPolicyService/GetBackupPolicy"
	BackupPolicyService_CreateBackupPolicy_FullMethodName     = "/public.portworx.platform.backuppolicy.v1.BackupPolicyService/CreateBackupPolicy"
	BackupPolicyService_UpdateBackupPolicyMeta_FullMethodName = "/public.portworx.platform.backuppolicy.v1.BackupPolicyService/UpdateBackupPolicyMeta"
	BackupPolicyService_ListBackupPolicies_FullMethodName     = "/public.portworx.platform.backuppolicy.v1.BackupPolicyService/ListBackupPolicies"
	BackupPolicyService_DeleteBackupPolicy_FullMethodName     = "/public.portworx.platform.backuppolicy.v1.BackupPolicyService/DeleteBackupPolicy"
)

// BackupPolicyServiceClient is the client API for BackupPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// BackupPolicyService serves and manages the backup policies.
type BackupPolicyServiceClient interface {
	// Get API returns the backup policy for the given ID.
	GetBackupPolicy(ctx context.Context, in *GetBackupPolicyRequest, opts ...grpc.CallOption) (*BackupPolicy, error)
	// Create API creates a backup policy for a Tenant (Organization).
	CreateBackupPolicy(ctx context.Context, in *CreateBackupPolicyRequest, opts ...grpc.CallOption) (*BackupPolicy, error)
	// Update API updates a the meta data of a backup policy.
	// (-- api-linter: core::0134::response-message-name=disabled
	//
	//	aip.dev/not-precedent: We need to do this because we need to return the whole Schedule after the update. --)
	//
	// (-- api-linter: core::0136::http-method=disabled
	//
	//	aip.dev/not-precedent: We need to do this because we are updating only the meta data of the schedule and not config. --)
	UpdateBackupPolicyMeta(ctx context.Context, in *UpdateBackupPolicyMetaRequest, opts ...grpc.CallOption) (*BackupPolicy, error)
	// (-- api-linter: core::0132::http-body=disabled
	//
	//	api-linter: core::0132::http-method=disabled
	//	aip.dev/not-precedent: We need to do this because
	//
	// we can't have advance filters as query params.
	// --)
	// List API lists all the backup policies for a Tenant (Organization).
	ListBackupPolicies(ctx context.Context, in *ListBackupPoliciesRequest, opts ...grpc.CallOption) (*ListBackupPoliciesResponse, error)
	// Delete API deletes a backup policy.
	DeleteBackupPolicy(ctx context.Context, in *DeleteBackupPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type backupPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupPolicyServiceClient(cc grpc.ClientConnInterface) BackupPolicyServiceClient {
	return &backupPolicyServiceClient{cc}
}

func (c *backupPolicyServiceClient) GetBackupPolicy(ctx context.Context, in *GetBackupPolicyRequest, opts ...grpc.CallOption) (*BackupPolicy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackupPolicy)
	err := c.cc.Invoke(ctx, BackupPolicyService_GetBackupPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupPolicyServiceClient) CreateBackupPolicy(ctx context.Context, in *CreateBackupPolicyRequest, opts ...grpc.CallOption) (*BackupPolicy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackupPolicy)
	err := c.cc.Invoke(ctx, BackupPolicyService_CreateBackupPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupPolicyServiceClient) UpdateBackupPolicyMeta(ctx context.Context, in *UpdateBackupPolicyMetaRequest, opts ...grpc.CallOption) (*BackupPolicy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackupPolicy)
	err := c.cc.Invoke(ctx, BackupPolicyService_UpdateBackupPolicyMeta_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupPolicyServiceClient) ListBackupPolicies(ctx context.Context, in *ListBackupPoliciesRequest, opts ...grpc.CallOption) (*ListBackupPoliciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBackupPoliciesResponse)
	err := c.cc.Invoke(ctx, BackupPolicyService_ListBackupPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupPolicyServiceClient) DeleteBackupPolicy(ctx context.Context, in *DeleteBackupPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BackupPolicyService_DeleteBackupPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupPolicyServiceServer is the server API for BackupPolicyService service.
// All implementations must embed UnimplementedBackupPolicyServiceServer
// for forward compatibility
//
// BackupPolicyService serves and manages the backup policies.
type BackupPolicyServiceServer interface {
	// Get API returns the backup policy for the given ID.
	GetBackupPolicy(context.Context, *GetBackupPolicyRequest) (*BackupPolicy, error)
	// Create API creates a backup policy for a Tenant (Organization).
	CreateBackupPolicy(context.Context, *CreateBackupPolicyRequest) (*BackupPolicy, error)
	// Update API updates a the meta data of a backup policy.
	// (-- api-linter: core::0134::response-message-name=disabled
	//
	//	aip.dev/not-precedent: We need to do this because we need to return the whole Schedule after the update. --)
	//
	// (-- api-linter: core::0136::http-method=disabled
	//
	//	aip.dev/not-precedent: We need to do this because we are updating only the meta data of the schedule and not config. --)
	UpdateBackupPolicyMeta(context.Context, *UpdateBackupPolicyMetaRequest) (*BackupPolicy, error)
	// (-- api-linter: core::0132::http-body=disabled
	//
	//	api-linter: core::0132::http-method=disabled
	//	aip.dev/not-precedent: We need to do this because
	//
	// we can't have advance filters as query params.
	// --)
	// List API lists all the backup policies for a Tenant (Organization).
	ListBackupPolicies(context.Context, *ListBackupPoliciesRequest) (*ListBackupPoliciesResponse, error)
	// Delete API deletes a backup policy.
	DeleteBackupPolicy(context.Context, *DeleteBackupPolicyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBackupPolicyServiceServer()
}

// UnimplementedBackupPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBackupPolicyServiceServer struct {
}

func (UnimplementedBackupPolicyServiceServer) GetBackupPolicy(context.Context, *GetBackupPolicyRequest) (*BackupPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackupPolicy not implemented")
}
func (UnimplementedBackupPolicyServiceServer) CreateBackupPolicy(context.Context, *CreateBackupPolicyRequest) (*BackupPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBackupPolicy not implemented")
}
func (UnimplementedBackupPolicyServiceServer) UpdateBackupPolicyMeta(context.Context, *UpdateBackupPolicyMetaRequest) (*BackupPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBackupPolicyMeta not implemented")
}
func (UnimplementedBackupPolicyServiceServer) ListBackupPolicies(context.Context, *ListBackupPoliciesRequest) (*ListBackupPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBackupPolicies not implemented")
}
func (UnimplementedBackupPolicyServiceServer) DeleteBackupPolicy(context.Context, *DeleteBackupPolicyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBackupPolicy not implemented")
}
func (UnimplementedBackupPolicyServiceServer) mustEmbedUnimplementedBackupPolicyServiceServer() {}

// UnsafeBackupPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupPolicyServiceServer will
// result in compilation errors.
type UnsafeBackupPolicyServiceServer interface {
	mustEmbedUnimplementedBackupPolicyServiceServer()
}

func RegisterBackupPolicyServiceServer(s grpc.ServiceRegistrar, srv BackupPolicyServiceServer) {
	s.RegisterService(&BackupPolicyService_ServiceDesc, srv)
}

func _BackupPolicyService_GetBackupPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupPolicyServiceServer).GetBackupPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupPolicyService_GetBackupPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupPolicyServiceServer).GetBackupPolicy(ctx, req.(*GetBackupPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupPolicyService_CreateBackupPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBackupPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupPolicyServiceServer).CreateBackupPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupPolicyService_CreateBackupPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupPolicyServiceServer).CreateBackupPolicy(ctx, req.(*CreateBackupPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupPolicyService_UpdateBackupPolicyMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBackupPolicyMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupPolicyServiceServer).UpdateBackupPolicyMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupPolicyService_UpdateBackupPolicyMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupPolicyServiceServer).UpdateBackupPolicyMeta(ctx, req.(*UpdateBackupPolicyMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupPolicyService_ListBackupPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupPolicyServiceServer).ListBackupPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupPolicyService_ListBackupPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupPolicyServiceServer).ListBackupPolicies(ctx, req.(*ListBackupPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupPolicyService_DeleteBackupPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBackupPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupPolicyServiceServer).DeleteBackupPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupPolicyService_DeleteBackupPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupPolicyServiceServer).DeleteBackupPolicy(ctx, req.(*DeleteBackupPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupPolicyService_ServiceDesc is the grpc.ServiceDesc for BackupPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.portworx.platform.backuppolicy.v1.BackupPolicyService",
	HandlerType: (*BackupPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBackupPolicy",
			Handler:    _BackupPolicyService_GetBackupPolicy_Handler,
		},
		{
			MethodName: "CreateBackupPolicy",
			Handler:    _BackupPolicyService_CreateBackupPolicy_Handler,
		},
		{
			MethodName: "UpdateBackupPolicyMeta",
			Handler:    _BackupPolicyService_UpdateBackupPolicyMeta_Handler,
		},
		{
			MethodName: "ListBackupPolicies",
			Handler:    _BackupPolicyService_ListBackupPolicies_Handler,
		},
		{
			MethodName: "DeleteBackupPolicy",
			Handler:    _BackupPolicyService_DeleteBackupPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/portworx/platform/backuppolicy/apiv1/backuppolicy.proto",
}
