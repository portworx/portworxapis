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
// source: public/portworx/platform/targetclusterregistrationmanifest/apiv1/targetclusterregistrationmanifest.proto

package targetclusterregistrationmanifest

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TargetClusterRegistrationManifestService_GenerateTargetClusterRegistrationManifest_FullMethodName = "/public.portworx.platform.targetclusterregistrationmanifest.v1.TargetClusterRegistrationManifestService/GenerateTargetClusterRegistrationManifest"
	TargetClusterRegistrationManifestService_ReRegisterTargetCluster_FullMethodName                   = "/public.portworx.platform.targetclusterregistrationmanifest.v1.TargetClusterRegistrationManifestService/ReRegisterTargetCluster"
)

// TargetClusterRegistrationManifestServiceClient is the client API for TargetClusterRegistrationManifestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// TargetClusterRegistrationManifestService provides API to generate registration manifest.
type TargetClusterRegistrationManifestServiceClient interface {
	// GetTargetClusterRegistrationManifest fetches registration manifest for the given request under specified Tenant (Organization).
	GenerateTargetClusterRegistrationManifest(ctx context.Context, in *GenerateTargetClusterRegistrationManifestRequest, opts ...grpc.CallOption) (*TargetClusterRegistrationManifest, error)
	// ReRegisterTargetCluster fetches manifest for re-registration of target cluster.
	ReRegisterTargetCluster(ctx context.Context, in *ReRegisterTargetClusterRequest, opts ...grpc.CallOption) (*TargetClusterRegistrationManifest, error)
}

type targetClusterRegistrationManifestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTargetClusterRegistrationManifestServiceClient(cc grpc.ClientConnInterface) TargetClusterRegistrationManifestServiceClient {
	return &targetClusterRegistrationManifestServiceClient{cc}
}

func (c *targetClusterRegistrationManifestServiceClient) GenerateTargetClusterRegistrationManifest(ctx context.Context, in *GenerateTargetClusterRegistrationManifestRequest, opts ...grpc.CallOption) (*TargetClusterRegistrationManifest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TargetClusterRegistrationManifest)
	err := c.cc.Invoke(ctx, TargetClusterRegistrationManifestService_GenerateTargetClusterRegistrationManifest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetClusterRegistrationManifestServiceClient) ReRegisterTargetCluster(ctx context.Context, in *ReRegisterTargetClusterRequest, opts ...grpc.CallOption) (*TargetClusterRegistrationManifest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TargetClusterRegistrationManifest)
	err := c.cc.Invoke(ctx, TargetClusterRegistrationManifestService_ReRegisterTargetCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TargetClusterRegistrationManifestServiceServer is the server API for TargetClusterRegistrationManifestService service.
// All implementations must embed UnimplementedTargetClusterRegistrationManifestServiceServer
// for forward compatibility
//
// TargetClusterRegistrationManifestService provides API to generate registration manifest.
type TargetClusterRegistrationManifestServiceServer interface {
	// GetTargetClusterRegistrationManifest fetches registration manifest for the given request under specified Tenant (Organization).
	GenerateTargetClusterRegistrationManifest(context.Context, *GenerateTargetClusterRegistrationManifestRequest) (*TargetClusterRegistrationManifest, error)
	// ReRegisterTargetCluster fetches manifest for re-registration of target cluster.
	ReRegisterTargetCluster(context.Context, *ReRegisterTargetClusterRequest) (*TargetClusterRegistrationManifest, error)
	mustEmbedUnimplementedTargetClusterRegistrationManifestServiceServer()
}

// UnimplementedTargetClusterRegistrationManifestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTargetClusterRegistrationManifestServiceServer struct {
}

func (UnimplementedTargetClusterRegistrationManifestServiceServer) GenerateTargetClusterRegistrationManifest(context.Context, *GenerateTargetClusterRegistrationManifestRequest) (*TargetClusterRegistrationManifest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTargetClusterRegistrationManifest not implemented")
}
func (UnimplementedTargetClusterRegistrationManifestServiceServer) ReRegisterTargetCluster(context.Context, *ReRegisterTargetClusterRequest) (*TargetClusterRegistrationManifest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReRegisterTargetCluster not implemented")
}
func (UnimplementedTargetClusterRegistrationManifestServiceServer) mustEmbedUnimplementedTargetClusterRegistrationManifestServiceServer() {
}

// UnsafeTargetClusterRegistrationManifestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TargetClusterRegistrationManifestServiceServer will
// result in compilation errors.
type UnsafeTargetClusterRegistrationManifestServiceServer interface {
	mustEmbedUnimplementedTargetClusterRegistrationManifestServiceServer()
}

func RegisterTargetClusterRegistrationManifestServiceServer(s grpc.ServiceRegistrar, srv TargetClusterRegistrationManifestServiceServer) {
	s.RegisterService(&TargetClusterRegistrationManifestService_ServiceDesc, srv)
}

func _TargetClusterRegistrationManifestService_GenerateTargetClusterRegistrationManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTargetClusterRegistrationManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetClusterRegistrationManifestServiceServer).GenerateTargetClusterRegistrationManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TargetClusterRegistrationManifestService_GenerateTargetClusterRegistrationManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetClusterRegistrationManifestServiceServer).GenerateTargetClusterRegistrationManifest(ctx, req.(*GenerateTargetClusterRegistrationManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetClusterRegistrationManifestService_ReRegisterTargetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReRegisterTargetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetClusterRegistrationManifestServiceServer).ReRegisterTargetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TargetClusterRegistrationManifestService_ReRegisterTargetCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetClusterRegistrationManifestServiceServer).ReRegisterTargetCluster(ctx, req.(*ReRegisterTargetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TargetClusterRegistrationManifestService_ServiceDesc is the grpc.ServiceDesc for TargetClusterRegistrationManifestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TargetClusterRegistrationManifestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.portworx.platform.targetclusterregistrationmanifest.v1.TargetClusterRegistrationManifestService",
	HandlerType: (*TargetClusterRegistrationManifestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTargetClusterRegistrationManifest",
			Handler:    _TargetClusterRegistrationManifestService_GenerateTargetClusterRegistrationManifest_Handler,
		},
		{
			MethodName: "ReRegisterTargetCluster",
			Handler:    _TargetClusterRegistrationManifestService_ReRegisterTargetCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/portworx/platform/targetclusterregistrationmanifest/apiv1/targetclusterregistrationmanifest.proto",
}
