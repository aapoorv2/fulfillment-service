// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: fulfillment.proto

package fulfillment

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

// FulFillmentClient is the client API for FulFillment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulFillmentClient interface {
	AssignDeliveryAgent(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*AssignResponse, error)
	RegisterDeliveryAgent(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	UpdateDeliveryAgentAvailability(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	FetchAllDeliveriesForAnAgent(ctx context.Context, in *FetchDeliveriesRequest, opts ...grpc.CallOption) (*FetchDeliveriesResponse, error)
}

type fulFillmentClient struct {
	cc grpc.ClientConnInterface
}

func NewFulFillmentClient(cc grpc.ClientConnInterface) FulFillmentClient {
	return &fulFillmentClient{cc}
}

func (c *fulFillmentClient) AssignDeliveryAgent(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*AssignResponse, error) {
	out := new(AssignResponse)
	err := c.cc.Invoke(ctx, "/fulfillment.FulFillment/AssignDeliveryAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulFillmentClient) RegisterDeliveryAgent(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/fulfillment.FulFillment/RegisterDeliveryAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulFillmentClient) UpdateDeliveryAgentAvailability(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/fulfillment.FulFillment/UpdateDeliveryAgentAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulFillmentClient) FetchAllDeliveriesForAnAgent(ctx context.Context, in *FetchDeliveriesRequest, opts ...grpc.CallOption) (*FetchDeliveriesResponse, error) {
	out := new(FetchDeliveriesResponse)
	err := c.cc.Invoke(ctx, "/fulfillment.FulFillment/FetchAllDeliveriesForAnAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulFillmentServer is the server API for FulFillment service.
// All implementations must embed UnimplementedFulFillmentServer
// for forward compatibility
type FulFillmentServer interface {
	AssignDeliveryAgent(context.Context, *AssignRequest) (*AssignResponse, error)
	RegisterDeliveryAgent(context.Context, *RegisterRequest) (*RegisterResponse, error)
	UpdateDeliveryAgentAvailability(context.Context, *UpdateRequest) (*UpdateResponse, error)
	FetchAllDeliveriesForAnAgent(context.Context, *FetchDeliveriesRequest) (*FetchDeliveriesResponse, error)
	mustEmbedUnimplementedFulFillmentServer()
}

// UnimplementedFulFillmentServer must be embedded to have forward compatible implementations.
type UnimplementedFulFillmentServer struct {
}

func (UnimplementedFulFillmentServer) AssignDeliveryAgent(context.Context, *AssignRequest) (*AssignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignDeliveryAgent not implemented")
}
func (UnimplementedFulFillmentServer) RegisterDeliveryAgent(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDeliveryAgent not implemented")
}
func (UnimplementedFulFillmentServer) UpdateDeliveryAgentAvailability(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeliveryAgentAvailability not implemented")
}
func (UnimplementedFulFillmentServer) FetchAllDeliveriesForAnAgent(context.Context, *FetchDeliveriesRequest) (*FetchDeliveriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAllDeliveriesForAnAgent not implemented")
}
func (UnimplementedFulFillmentServer) mustEmbedUnimplementedFulFillmentServer() {}

// UnsafeFulFillmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulFillmentServer will
// result in compilation errors.
type UnsafeFulFillmentServer interface {
	mustEmbedUnimplementedFulFillmentServer()
}

func RegisterFulFillmentServer(s grpc.ServiceRegistrar, srv FulFillmentServer) {
	s.RegisterService(&FulFillment_ServiceDesc, srv)
}

func _FulFillment_AssignDeliveryAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulFillmentServer).AssignDeliveryAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fulfillment.FulFillment/AssignDeliveryAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulFillmentServer).AssignDeliveryAgent(ctx, req.(*AssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulFillment_RegisterDeliveryAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulFillmentServer).RegisterDeliveryAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fulfillment.FulFillment/RegisterDeliveryAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulFillmentServer).RegisterDeliveryAgent(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulFillment_UpdateDeliveryAgentAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulFillmentServer).UpdateDeliveryAgentAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fulfillment.FulFillment/UpdateDeliveryAgentAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulFillmentServer).UpdateDeliveryAgentAvailability(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulFillment_FetchAllDeliveriesForAnAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDeliveriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulFillmentServer).FetchAllDeliveriesForAnAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fulfillment.FulFillment/FetchAllDeliveriesForAnAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulFillmentServer).FetchAllDeliveriesForAnAgent(ctx, req.(*FetchDeliveriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FulFillment_ServiceDesc is the grpc.ServiceDesc for FulFillment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FulFillment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fulfillment.FulFillment",
	HandlerType: (*FulFillmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignDeliveryAgent",
			Handler:    _FulFillment_AssignDeliveryAgent_Handler,
		},
		{
			MethodName: "RegisterDeliveryAgent",
			Handler:    _FulFillment_RegisterDeliveryAgent_Handler,
		},
		{
			MethodName: "UpdateDeliveryAgentAvailability",
			Handler:    _FulFillment_UpdateDeliveryAgentAvailability_Handler,
		},
		{
			MethodName: "FetchAllDeliveriesForAnAgent",
			Handler:    _FulFillment_FetchAllDeliveriesForAnAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulfillment.proto",
}
