// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/gateway/gateway.proto

package gateway

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
	GatewayService_ApexDriveStart_FullMethodName    = "/gateway.GatewayService/ApexDriveStart"
	GatewayService_MaleniaStart_FullMethodName      = "/gateway.GatewayService/MaleniaStart"
	GatewayService_TimeSquaredStart_FullMethodName  = "/gateway.GatewayService/TimeSquaredStart"
	GatewayService_IndriyasStart_FullMethodName     = "/gateway.GatewayService/IndriyasStart"
	GatewayService_NeithStart_FullMethodName        = "/gateway.GatewayService/NeithStart"
	GatewayService_ApexDriveStatus_FullMethodName   = "/gateway.GatewayService/ApexDriveStatus"
	GatewayService_MaleniaStatus_FullMethodName     = "/gateway.GatewayService/MaleniaStatus"
	GatewayService_TimeSquaredStatus_FullMethodName = "/gateway.GatewayService/TimeSquaredStatus"
	GatewayService_IndriyasStatus_FullMethodName    = "/gateway.GatewayService/IndriyasStatus"
	GatewayService_NeithStatus_FullMethodName       = "/gateway.GatewayService/NeithStatus"
)

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	// Start operation for Apexdrive
	ApexDriveStart(ctx context.Context, in *ApexDriveStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Start operation for Malenia
	MaleniaStart(ctx context.Context, in *MaleniaStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Start operation for TimeSquared
	TimeSquaredStart(ctx context.Context, in *TimeSquaredStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Start operation for Indriyas
	IndriyasStart(ctx context.Context, in *IndriyasStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Start operation for Neith
	NeithStart(ctx context.Context, in *NeithStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Status operation for Apexdrive
	ApexDriveStatus(ctx context.Context, in *ApexDriveStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Status operation for Malenia
	MaleniaStatus(ctx context.Context, in *MaleniaStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Status operation for TimeSquared
	TimeSquaredStatus(ctx context.Context, in *TimeSquaredStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Status operation for Indriyas
	IndriyasStatus(ctx context.Context, in *IndriyasStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
	// Status operation for Neith
	NeithStatus(ctx context.Context, in *NeithStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) ApexDriveStart(ctx context.Context, in *ApexDriveStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_ApexDriveStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) MaleniaStart(ctx context.Context, in *MaleniaStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_MaleniaStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) TimeSquaredStart(ctx context.Context, in *TimeSquaredStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_TimeSquaredStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) IndriyasStart(ctx context.Context, in *IndriyasStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_IndriyasStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) NeithStart(ctx context.Context, in *NeithStartRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_NeithStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) ApexDriveStatus(ctx context.Context, in *ApexDriveStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_ApexDriveStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) MaleniaStatus(ctx context.Context, in *MaleniaStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_MaleniaStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) TimeSquaredStatus(ctx context.Context, in *TimeSquaredStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_TimeSquaredStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) IndriyasStatus(ctx context.Context, in *IndriyasStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_IndriyasStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) NeithStatus(ctx context.Context, in *NeithStatusRequest, opts ...grpc.CallOption) (*GatewayResponse, error) {
	out := new(GatewayResponse)
	err := c.cc.Invoke(ctx, GatewayService_NeithStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations must embed UnimplementedGatewayServiceServer
// for forward compatibility
type GatewayServiceServer interface {
	// Start operation for Apexdrive
	ApexDriveStart(context.Context, *ApexDriveStartRequest) (*GatewayResponse, error)
	// Start operation for Malenia
	MaleniaStart(context.Context, *MaleniaStartRequest) (*GatewayResponse, error)
	// Start operation for TimeSquared
	TimeSquaredStart(context.Context, *TimeSquaredStartRequest) (*GatewayResponse, error)
	// Start operation for Indriyas
	IndriyasStart(context.Context, *IndriyasStartRequest) (*GatewayResponse, error)
	// Start operation for Neith
	NeithStart(context.Context, *NeithStartRequest) (*GatewayResponse, error)
	// Status operation for Apexdrive
	ApexDriveStatus(context.Context, *ApexDriveStatusRequest) (*GatewayResponse, error)
	// Status operation for Malenia
	MaleniaStatus(context.Context, *MaleniaStatusRequest) (*GatewayResponse, error)
	// Status operation for TimeSquared
	TimeSquaredStatus(context.Context, *TimeSquaredStatusRequest) (*GatewayResponse, error)
	// Status operation for Indriyas
	IndriyasStatus(context.Context, *IndriyasStatusRequest) (*GatewayResponse, error)
	// Status operation for Neith
	NeithStatus(context.Context, *NeithStatusRequest) (*GatewayResponse, error)
	mustEmbedUnimplementedGatewayServiceServer()
}

// UnimplementedGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (UnimplementedGatewayServiceServer) ApexDriveStart(context.Context, *ApexDriveStartRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApexDriveStart not implemented")
}
func (UnimplementedGatewayServiceServer) MaleniaStart(context.Context, *MaleniaStartRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaleniaStart not implemented")
}
func (UnimplementedGatewayServiceServer) TimeSquaredStart(context.Context, *TimeSquaredStartRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeSquaredStart not implemented")
}
func (UnimplementedGatewayServiceServer) IndriyasStart(context.Context, *IndriyasStartRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndriyasStart not implemented")
}
func (UnimplementedGatewayServiceServer) NeithStart(context.Context, *NeithStartRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NeithStart not implemented")
}
func (UnimplementedGatewayServiceServer) ApexDriveStatus(context.Context, *ApexDriveStatusRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApexDriveStatus not implemented")
}
func (UnimplementedGatewayServiceServer) MaleniaStatus(context.Context, *MaleniaStatusRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaleniaStatus not implemented")
}
func (UnimplementedGatewayServiceServer) TimeSquaredStatus(context.Context, *TimeSquaredStatusRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeSquaredStatus not implemented")
}
func (UnimplementedGatewayServiceServer) IndriyasStatus(context.Context, *IndriyasStatusRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndriyasStatus not implemented")
}
func (UnimplementedGatewayServiceServer) NeithStatus(context.Context, *NeithStatusRequest) (*GatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NeithStatus not implemented")
}
func (UnimplementedGatewayServiceServer) mustEmbedUnimplementedGatewayServiceServer() {}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_ApexDriveStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApexDriveStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ApexDriveStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_ApexDriveStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ApexDriveStart(ctx, req.(*ApexDriveStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_MaleniaStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaleniaStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).MaleniaStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_MaleniaStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).MaleniaStart(ctx, req.(*MaleniaStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_TimeSquaredStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeSquaredStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).TimeSquaredStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_TimeSquaredStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).TimeSquaredStart(ctx, req.(*TimeSquaredStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_IndriyasStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndriyasStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).IndriyasStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_IndriyasStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).IndriyasStart(ctx, req.(*IndriyasStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_NeithStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NeithStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).NeithStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_NeithStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).NeithStart(ctx, req.(*NeithStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_ApexDriveStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApexDriveStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ApexDriveStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_ApexDriveStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ApexDriveStatus(ctx, req.(*ApexDriveStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_MaleniaStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaleniaStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).MaleniaStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_MaleniaStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).MaleniaStatus(ctx, req.(*MaleniaStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_TimeSquaredStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeSquaredStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).TimeSquaredStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_TimeSquaredStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).TimeSquaredStatus(ctx, req.(*TimeSquaredStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_IndriyasStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndriyasStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).IndriyasStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_IndriyasStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).IndriyasStatus(ctx, req.(*IndriyasStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_NeithStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NeithStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).NeithStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_NeithStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).NeithStatus(ctx, req.(*NeithStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApexDriveStart",
			Handler:    _GatewayService_ApexDriveStart_Handler,
		},
		{
			MethodName: "MaleniaStart",
			Handler:    _GatewayService_MaleniaStart_Handler,
		},
		{
			MethodName: "TimeSquaredStart",
			Handler:    _GatewayService_TimeSquaredStart_Handler,
		},
		{
			MethodName: "IndriyasStart",
			Handler:    _GatewayService_IndriyasStart_Handler,
		},
		{
			MethodName: "NeithStart",
			Handler:    _GatewayService_NeithStart_Handler,
		},
		{
			MethodName: "ApexDriveStatus",
			Handler:    _GatewayService_ApexDriveStatus_Handler,
		},
		{
			MethodName: "MaleniaStatus",
			Handler:    _GatewayService_MaleniaStatus_Handler,
		},
		{
			MethodName: "TimeSquaredStatus",
			Handler:    _GatewayService_TimeSquaredStatus_Handler,
		},
		{
			MethodName: "IndriyasStatus",
			Handler:    _GatewayService_IndriyasStatus_Handler,
		},
		{
			MethodName: "NeithStatus",
			Handler:    _GatewayService_NeithStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gateway/gateway.proto",
}
