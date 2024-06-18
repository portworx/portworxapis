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
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: public/portworx/pds/dataservicedeploymentevents/apiv1/dataservicedeploymentevents.proto

package dataservicedeploymentevents

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
	DataServiceDeploymentEventService_ListDataServiceDeploymentEvents_FullMethodName = "/public.portworx.pds.dataservicedeploymentevents.v1.DataServiceDeploymentEventService/ListDataServiceDeploymentEvents"
)

// DataServiceDeploymentEventServiceClient is the client API for DataServiceDeploymentEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceDeploymentEventServiceClient interface {
	// ListDataServiceDeploymentEvents API returns the list of kubernetes events related to a DataService deployment.
	ListDataServiceDeploymentEvents(ctx context.Context, in *ListDataServiceDeploymentEventsRequest, opts ...grpc.CallOption) (*ListDataServiceDeploymentEventsResponse, error)
}

type dataServiceDeploymentEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceDeploymentEventServiceClient(cc grpc.ClientConnInterface) DataServiceDeploymentEventServiceClient {
	return &dataServiceDeploymentEventServiceClient{cc}
}

func (c *dataServiceDeploymentEventServiceClient) ListDataServiceDeploymentEvents(ctx context.Context, in *ListDataServiceDeploymentEventsRequest, opts ...grpc.CallOption) (*ListDataServiceDeploymentEventsResponse, error) {
	out := new(ListDataServiceDeploymentEventsResponse)
	err := c.cc.Invoke(ctx, DataServiceDeploymentEventService_ListDataServiceDeploymentEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceDeploymentEventServiceServer is the server API for DataServiceDeploymentEventService service.
// All implementations must embed UnimplementedDataServiceDeploymentEventServiceServer
// for forward compatibility
type DataServiceDeploymentEventServiceServer interface {
	// ListDataServiceDeploymentEvents API returns the list of kubernetes events related to a DataService deployment.
	ListDataServiceDeploymentEvents(context.Context, *ListDataServiceDeploymentEventsRequest) (*ListDataServiceDeploymentEventsResponse, error)
	mustEmbedUnimplementedDataServiceDeploymentEventServiceServer()
}

// UnimplementedDataServiceDeploymentEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceDeploymentEventServiceServer struct {
}

func (UnimplementedDataServiceDeploymentEventServiceServer) ListDataServiceDeploymentEvents(context.Context, *ListDataServiceDeploymentEventsRequest) (*ListDataServiceDeploymentEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataServiceDeploymentEvents not implemented")
}
func (UnimplementedDataServiceDeploymentEventServiceServer) mustEmbedUnimplementedDataServiceDeploymentEventServiceServer() {
}

// UnsafeDataServiceDeploymentEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceDeploymentEventServiceServer will
// result in compilation errors.
type UnsafeDataServiceDeploymentEventServiceServer interface {
	mustEmbedUnimplementedDataServiceDeploymentEventServiceServer()
}

func RegisterDataServiceDeploymentEventServiceServer(s grpc.ServiceRegistrar, srv DataServiceDeploymentEventServiceServer) {
	s.RegisterService(&DataServiceDeploymentEventService_ServiceDesc, srv)
}

func _DataServiceDeploymentEventService_ListDataServiceDeploymentEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataServiceDeploymentEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceDeploymentEventServiceServer).ListDataServiceDeploymentEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataServiceDeploymentEventService_ListDataServiceDeploymentEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceDeploymentEventServiceServer).ListDataServiceDeploymentEvents(ctx, req.(*ListDataServiceDeploymentEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataServiceDeploymentEventService_ServiceDesc is the grpc.ServiceDesc for DataServiceDeploymentEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataServiceDeploymentEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.portworx.pds.dataservicedeploymentevents.v1.DataServiceDeploymentEventService",
	HandlerType: (*DataServiceDeploymentEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDataServiceDeploymentEvents",
			Handler:    _DataServiceDeploymentEventService_ListDataServiceDeploymentEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/portworx/pds/dataservicedeploymentevents/apiv1/dataservicedeploymentevents.proto",
}
