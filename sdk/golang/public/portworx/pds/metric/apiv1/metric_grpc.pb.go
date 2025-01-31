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
// source: public/portworx/pds/metric/apiv1/metric.proto

package metric

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
	MetricService_GetMetrics_FullMethodName = "/public.portworx.pds.metric.v1.MetricService/GetMetrics"
)

// MetricServiceClient is the client API for MetricService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MetricService serves Metric of deployment resource.
type MetricServiceClient interface {
	// (-- api-linter: core::0131::response-message-name=disabled
	//
	//	aip.dev/not-precedent: We need to do this because reasons. --)
	//
	// GetMetrics API returns the Metrics of DataService Deployment.
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type metricServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricServiceClient(cc grpc.ClientConnInterface) MetricServiceClient {
	return &metricServiceClient{cc}
}

func (c *metricServiceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, MetricService_GetMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricServiceServer is the server API for MetricService service.
// All implementations must embed UnimplementedMetricServiceServer
// for forward compatibility
//
// MetricService serves Metric of deployment resource.
type MetricServiceServer interface {
	// (-- api-linter: core::0131::response-message-name=disabled
	//
	//	aip.dev/not-precedent: We need to do this because reasons. --)
	//
	// GetMetrics API returns the Metrics of DataService Deployment.
	GetMetrics(context.Context, *GetMetricsRequest) (*MetricsResponse, error)
	mustEmbedUnimplementedMetricServiceServer()
}

// UnimplementedMetricServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricServiceServer struct {
}

func (UnimplementedMetricServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedMetricServiceServer) mustEmbedUnimplementedMetricServiceServer() {}

// UnsafeMetricServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricServiceServer will
// result in compilation errors.
type UnsafeMetricServiceServer interface {
	mustEmbedUnimplementedMetricServiceServer()
}

func RegisterMetricServiceServer(s grpc.ServiceRegistrar, srv MetricServiceServer) {
	s.RegisterService(&MetricService_ServiceDesc, srv)
}

func _MetricService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricService_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricService_ServiceDesc is the grpc.ServiceDesc for MetricService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.portworx.pds.metric.v1.MetricService",
	HandlerType: (*MetricServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _MetricService_GetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/portworx/pds/metric/apiv1/metric.proto",
}
