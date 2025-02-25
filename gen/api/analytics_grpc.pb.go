// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: analytics.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AnalyticsData_SaveDoneTasks_FullMethodName            = "/analytics.AnalyticsData/SaveDoneTasks"
	AnalyticsData_FetchWeeklyCompletedTask_FullMethodName = "/analytics.AnalyticsData/FetchWeeklyCompletedTask"
)

// AnalyticsDataClient is the client API for AnalyticsData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsDataClient interface {
	SaveDoneTasks(ctx context.Context, in *TaskDoneItems, opts ...grpc.CallOption) (*ServiceResponse, error)
	FetchWeeklyCompletedTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WeeklyCompletedTasksResponse, error)
}

type analyticsDataClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsDataClient(cc grpc.ClientConnInterface) AnalyticsDataClient {
	return &analyticsDataClient{cc}
}

func (c *analyticsDataClient) SaveDoneTasks(ctx context.Context, in *TaskDoneItems, opts ...grpc.CallOption) (*ServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, AnalyticsData_SaveDoneTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsDataClient) FetchWeeklyCompletedTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WeeklyCompletedTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeeklyCompletedTasksResponse)
	err := c.cc.Invoke(ctx, AnalyticsData_FetchWeeklyCompletedTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsDataServer is the server API for AnalyticsData service.
// All implementations must embed UnimplementedAnalyticsDataServer
// for forward compatibility.
type AnalyticsDataServer interface {
	SaveDoneTasks(context.Context, *TaskDoneItems) (*ServiceResponse, error)
	FetchWeeklyCompletedTask(context.Context, *emptypb.Empty) (*WeeklyCompletedTasksResponse, error)
	mustEmbedUnimplementedAnalyticsDataServer()
}

// UnimplementedAnalyticsDataServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnalyticsDataServer struct{}

func (UnimplementedAnalyticsDataServer) SaveDoneTasks(context.Context, *TaskDoneItems) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveDoneTasks not implemented")
}
func (UnimplementedAnalyticsDataServer) FetchWeeklyCompletedTask(context.Context, *emptypb.Empty) (*WeeklyCompletedTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchWeeklyCompletedTask not implemented")
}
func (UnimplementedAnalyticsDataServer) mustEmbedUnimplementedAnalyticsDataServer() {}
func (UnimplementedAnalyticsDataServer) testEmbeddedByValue()                       {}

// UnsafeAnalyticsDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsDataServer will
// result in compilation errors.
type UnsafeAnalyticsDataServer interface {
	mustEmbedUnimplementedAnalyticsDataServer()
}

func RegisterAnalyticsDataServer(s grpc.ServiceRegistrar, srv AnalyticsDataServer) {
	// If the following call pancis, it indicates UnimplementedAnalyticsDataServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AnalyticsData_ServiceDesc, srv)
}

func _AnalyticsData_SaveDoneTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDoneItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsDataServer).SaveDoneTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsData_SaveDoneTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsDataServer).SaveDoneTasks(ctx, req.(*TaskDoneItems))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsData_FetchWeeklyCompletedTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsDataServer).FetchWeeklyCompletedTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsData_FetchWeeklyCompletedTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsDataServer).FetchWeeklyCompletedTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalyticsData_ServiceDesc is the grpc.ServiceDesc for AnalyticsData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.AnalyticsData",
	HandlerType: (*AnalyticsDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveDoneTasks",
			Handler:    _AnalyticsData_SaveDoneTasks_Handler,
		},
		{
			MethodName: "FetchWeeklyCompletedTask",
			Handler:    _AnalyticsData_FetchWeeklyCompletedTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analytics.proto",
}
