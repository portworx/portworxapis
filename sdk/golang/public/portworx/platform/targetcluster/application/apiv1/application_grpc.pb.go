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
// source: public/portworx/platform/targetcluster/application/apiv1/application.proto

package application

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
	ApplicationService_ListApplications_FullMethodName     = "/public.portworx.platform.targetcluster.application.v1.ApplicationService/ListApplications"
	ApplicationService_GetApplication_FullMethodName       = "/public.portworx.platform.targetcluster.application.v1.ApplicationService/GetApplication"
	ApplicationService_InstallApplication_FullMethodName   = "/public.portworx.platform.targetcluster.application.v1.ApplicationService/InstallApplication"
	ApplicationService_UpdateApplication_FullMethodName    = "/public.portworx.platform.targetcluster.application.v1.ApplicationService/UpdateApplication"
	ApplicationService_UninstallApplication_FullMethodName = "/public.portworx.platform.targetcluster.application.v1.ApplicationService/UninstallApplication"
)

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Application service provides APIs to interact with the Application entity
type ApplicationServiceClient interface {
	// ListApplications API lists the applications installed on the target cluster.
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
	// GetApplication API returns the info about application with given id.
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// InstallApplication API installs specified application on the target cluster.
	InstallApplication(ctx context.Context, in *InstallApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// UpdateApplication API updates specified application on the target cluster.
	UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// (-- api-linter: core::0136::http-method=disabled
	//
	//	   aip.dev/not-precedent: We need to do this because
	//	   delete method is required for deleting resource. --)
	//	UninstallApplication API uninstalls the specified application on the target cluster.
	UninstallApplication(ctx context.Context, in *UninstallApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type applicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationServiceClient(cc grpc.ClientConnInterface) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, ApplicationService_ListApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Application)
	err := c.cc.Invoke(ctx, ApplicationService_GetApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) InstallApplication(ctx context.Context, in *InstallApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Application)
	err := c.cc.Invoke(ctx, ApplicationService_InstallApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Application)
	err := c.cc.Invoke(ctx, ApplicationService_UpdateApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) UninstallApplication(ctx context.Context, in *UninstallApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ApplicationService_UninstallApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
// All implementations must embed UnimplementedApplicationServiceServer
// for forward compatibility
//
// Application service provides APIs to interact with the Application entity
type ApplicationServiceServer interface {
	// ListApplications API lists the applications installed on the target cluster.
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
	// GetApplication API returns the info about application with given id.
	GetApplication(context.Context, *GetApplicationRequest) (*Application, error)
	// InstallApplication API installs specified application on the target cluster.
	InstallApplication(context.Context, *InstallApplicationRequest) (*Application, error)
	// UpdateApplication API updates specified application on the target cluster.
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error)
	// (-- api-linter: core::0136::http-method=disabled
	//
	//	   aip.dev/not-precedent: We need to do this because
	//	   delete method is required for deleting resource. --)
	//	UninstallApplication API uninstalls the specified application on the target cluster.
	UninstallApplication(context.Context, *UninstallApplicationRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedApplicationServiceServer()
}

// UnimplementedApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServiceServer struct {
}

func (UnimplementedApplicationServiceServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedApplicationServiceServer) GetApplication(context.Context, *GetApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplication not implemented")
}
func (UnimplementedApplicationServiceServer) InstallApplication(context.Context, *InstallApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallApplication not implemented")
}
func (UnimplementedApplicationServiceServer) UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedApplicationServiceServer) UninstallApplication(context.Context, *UninstallApplicationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallApplication not implemented")
}
func (UnimplementedApplicationServiceServer) mustEmbedUnimplementedApplicationServiceServer() {}

// UnsafeApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServiceServer will
// result in compilation errors.
type UnsafeApplicationServiceServer interface {
	mustEmbedUnimplementedApplicationServiceServer()
}

func RegisterApplicationServiceServer(s grpc.ServiceRegistrar, srv ApplicationServiceServer) {
	s.RegisterService(&ApplicationService_ServiceDesc, srv)
}

func _ApplicationService_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_ListApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_GetApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_InstallApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).InstallApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_InstallApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).InstallApplication(ctx, req.(*InstallApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_UpdateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).UpdateApplication(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_UninstallApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).UninstallApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_UninstallApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).UninstallApplication(ctx, req.(*UninstallApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApplicationService_ServiceDesc is the grpc.ServiceDesc for ApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.portworx.platform.targetcluster.application.v1.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplications",
			Handler:    _ApplicationService_ListApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _ApplicationService_GetApplication_Handler,
		},
		{
			MethodName: "InstallApplication",
			Handler:    _ApplicationService_InstallApplication_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _ApplicationService_UpdateApplication_Handler,
		},
		{
			MethodName: "UninstallApplication",
			Handler:    _ApplicationService_UninstallApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/portworx/platform/targetcluster/application/apiv1/application.proto",
}
