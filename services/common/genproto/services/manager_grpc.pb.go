// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: manager.proto

package manager

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
	ServiceManager_GetServices_FullMethodName    = "/ServiceManager/GetServices"
	ServiceManager_GetServiceData_FullMethodName = "/ServiceManager/GetServiceData"
	ServiceManager_StartService_FullMethodName   = "/ServiceManager/StartService"
	ServiceManager_StopService_FullMethodName    = "/ServiceManager/StopService"
)

// ServiceManagerClient is the client API for ServiceManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceManagerClient interface {
	// TODO: Eliminar `RegisterService`, no tiene sentido ya que el cargado ocurre alrevés
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetServiceData(ctx context.Context, in *ServiceIdRequest, opts ...grpc.CallOption) (*RaspiService, error)
	StartService(ctx context.Context, in *ServiceIdRequest, opts ...grpc.CallOption) (*ServiceStatusResponse, error)
	StopService(ctx context.Context, in *ServiceIdRequest, opts ...grpc.CallOption) (*ServiceStatusResponse, error)
}

type serviceManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceManagerClient(cc grpc.ClientConnInterface) ServiceManagerClient {
	return &serviceManagerClient{cc}
}

func (c *serviceManagerClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, ServiceManager_GetServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceManagerClient) GetServiceData(ctx context.Context, in *ServiceIdRequest, opts ...grpc.CallOption) (*RaspiService, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RaspiService)
	err := c.cc.Invoke(ctx, ServiceManager_GetServiceData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceManagerClient) StartService(ctx context.Context, in *ServiceIdRequest, opts ...grpc.CallOption) (*ServiceStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceStatusResponse)
	err := c.cc.Invoke(ctx, ServiceManager_StartService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceManagerClient) StopService(ctx context.Context, in *ServiceIdRequest, opts ...grpc.CallOption) (*ServiceStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceStatusResponse)
	err := c.cc.Invoke(ctx, ServiceManager_StopService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceManagerServer is the server API for ServiceManager service.
// All implementations must embed UnimplementedServiceManagerServer
// for forward compatibility.
type ServiceManagerServer interface {
	// TODO: Eliminar `RegisterService`, no tiene sentido ya que el cargado ocurre alrevés
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	GetServiceData(context.Context, *ServiceIdRequest) (*RaspiService, error)
	StartService(context.Context, *ServiceIdRequest) (*ServiceStatusResponse, error)
	StopService(context.Context, *ServiceIdRequest) (*ServiceStatusResponse, error)
	mustEmbedUnimplementedServiceManagerServer()
}

// UnimplementedServiceManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceManagerServer struct{}

func (UnimplementedServiceManagerServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedServiceManagerServer) GetServiceData(context.Context, *ServiceIdRequest) (*RaspiService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceData not implemented")
}
func (UnimplementedServiceManagerServer) StartService(context.Context, *ServiceIdRequest) (*ServiceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (UnimplementedServiceManagerServer) StopService(context.Context, *ServiceIdRequest) (*ServiceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopService not implemented")
}
func (UnimplementedServiceManagerServer) mustEmbedUnimplementedServiceManagerServer() {}
func (UnimplementedServiceManagerServer) testEmbeddedByValue()                        {}

// UnsafeServiceManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceManagerServer will
// result in compilation errors.
type UnsafeServiceManagerServer interface {
	mustEmbedUnimplementedServiceManagerServer()
}

func RegisterServiceManagerServer(s grpc.ServiceRegistrar, srv ServiceManagerServer) {
	// If the following call pancis, it indicates UnimplementedServiceManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceManager_ServiceDesc, srv)
}

func _ServiceManager_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceManagerServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceManager_GetServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceManagerServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceManager_GetServiceData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceManagerServer).GetServiceData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceManager_GetServiceData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceManagerServer).GetServiceData(ctx, req.(*ServiceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceManager_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceManagerServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceManager_StartService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceManagerServer).StartService(ctx, req.(*ServiceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceManager_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceManagerServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceManager_StopService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceManagerServer).StopService(ctx, req.(*ServiceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceManager_ServiceDesc is the grpc.ServiceDesc for ServiceManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceManager",
	HandlerType: (*ServiceManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _ServiceManager_GetServices_Handler,
		},
		{
			MethodName: "GetServiceData",
			Handler:    _ServiceManager_GetServiceData_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _ServiceManager_StartService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _ServiceManager_StopService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
