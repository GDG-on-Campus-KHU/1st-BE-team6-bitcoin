// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: monitoring.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Monitoring_Service_SayHello_FullMethodName = "/monitoring.Monitoring_Service/SayHello"
)

// Monitoring_ServiceClient is the client API for Monitoring_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Monitoring_ServiceClient interface {
	SayHello(ctx context.Context, in *MonitoringRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MonitoringResponse], error)
}

type monitoring_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoring_ServiceClient(cc grpc.ClientConnInterface) Monitoring_ServiceClient {
	return &monitoring_ServiceClient{cc}
}

func (c *monitoring_ServiceClient) SayHello(ctx context.Context, in *MonitoringRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MonitoringResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Monitoring_Service_ServiceDesc.Streams[0], Monitoring_Service_SayHello_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MonitoringRequest, MonitoringResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Monitoring_Service_SayHelloClient = grpc.ServerStreamingClient[MonitoringResponse]

// Monitoring_ServiceServer is the server API for Monitoring_Service service.
// All implementations must embed UnimplementedMonitoring_ServiceServer
// for forward compatibility.
type Monitoring_ServiceServer interface {
	SayHello(*MonitoringRequest, grpc.ServerStreamingServer[MonitoringResponse]) error
	mustEmbedUnimplementedMonitoring_ServiceServer()
}

// UnimplementedMonitoring_ServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMonitoring_ServiceServer struct{}

func (UnimplementedMonitoring_ServiceServer) SayHello(*MonitoringRequest, grpc.ServerStreamingServer[MonitoringResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMonitoring_ServiceServer) mustEmbedUnimplementedMonitoring_ServiceServer() {}
func (UnimplementedMonitoring_ServiceServer) testEmbeddedByValue()                            {}

// UnsafeMonitoring_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Monitoring_ServiceServer will
// result in compilation errors.
type UnsafeMonitoring_ServiceServer interface {
	mustEmbedUnimplementedMonitoring_ServiceServer()
}

func RegisterMonitoring_ServiceServer(s grpc.ServiceRegistrar, srv Monitoring_ServiceServer) {
	// If the following call pancis, it indicates UnimplementedMonitoring_ServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Monitoring_Service_ServiceDesc, srv)
}

func _Monitoring_Service_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitoringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Monitoring_ServiceServer).SayHello(m, &grpc.GenericServerStream[MonitoringRequest, MonitoringResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Monitoring_Service_SayHelloServer = grpc.ServerStreamingServer[MonitoringResponse]

// Monitoring_Service_ServiceDesc is the grpc.ServiceDesc for Monitoring_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitoring_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.Monitoring_Service",
	HandlerType: (*Monitoring_ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Monitoring_Service_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "monitoring.proto",
}