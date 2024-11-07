// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: raspi_services.proto

package raspiservices

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
	WeatherService_Start_FullMethodName        = "/WeatherService/Start"
	WeatherService_Stop_FullMethodName         = "/WeatherService/Stop"
	WeatherService_GetStatus_FullMethodName    = "/WeatherService/GetStatus"
	WeatherService_GetFullInfo_FullMethodName  = "/WeatherService/GetFullInfo"
	WeatherService_GetConfig_FullMethodName    = "/WeatherService/GetConfig"
	WeatherService_UpdateConfig_FullMethodName = "/WeatherService/UpdateConfig"
)

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	Start(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Stop(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetStatus(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetFullInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*FullInfoResponse, error)
	GetConfig(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	UpdateConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) Start(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, WeatherService_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) Stop(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, WeatherService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetStatus(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, WeatherService_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetFullInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*FullInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FullInfoResponse)
	err := c.cc.Invoke(ctx, WeatherService_GetFullInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetConfig(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, WeatherService_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) UpdateConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, WeatherService_UpdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility.
type WeatherServiceServer interface {
	Start(context.Context, *EmptyRequest) (*StatusResponse, error)
	Stop(context.Context, *EmptyRequest) (*StatusResponse, error)
	GetStatus(context.Context, *EmptyRequest) (*StatusResponse, error)
	GetFullInfo(context.Context, *EmptyRequest) (*FullInfoResponse, error)
	GetConfig(context.Context, *EmptyRequest) (*ConfigResponse, error)
	UpdateConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeatherServiceServer struct{}

func (UnimplementedWeatherServiceServer) Start(context.Context, *EmptyRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedWeatherServiceServer) Stop(context.Context, *EmptyRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedWeatherServiceServer) GetStatus(context.Context, *EmptyRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedWeatherServiceServer) GetFullInfo(context.Context, *EmptyRequest) (*FullInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullInfo not implemented")
}
func (UnimplementedWeatherServiceServer) GetConfig(context.Context, *EmptyRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedWeatherServiceServer) UpdateConfig(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}
func (UnimplementedWeatherServiceServer) testEmbeddedByValue()                        {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	// If the following call pancis, it indicates UnimplementedWeatherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).Start(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).Stop(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetStatus(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetFullInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetFullInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_GetFullInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetFullInfo(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetConfig(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_UpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).UpdateConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _WeatherService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _WeatherService_Stop_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _WeatherService_GetStatus_Handler,
		},
		{
			MethodName: "GetFullInfo",
			Handler:    _WeatherService_GetFullInfo_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _WeatherService_GetConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _WeatherService_UpdateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raspi_services.proto",
}
