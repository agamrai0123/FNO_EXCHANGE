// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: order.proto

package proto

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
	RandomOrderGenerator_SendOrder_FullMethodName = "/order.random_order_generator/SendOrder"
)

// RandomOrderGeneratorClient is the client API for RandomOrderGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomOrderGeneratorClient interface {
	SendOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OrderResponse, error)
}

type randomOrderGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomOrderGeneratorClient(cc grpc.ClientConnInterface) RandomOrderGeneratorClient {
	return &randomOrderGeneratorClient{cc}
}

func (c *randomOrderGeneratorClient) SendOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, RandomOrderGenerator_SendOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomOrderGeneratorServer is the server API for RandomOrderGenerator service.
// All implementations must embed UnimplementedRandomOrderGeneratorServer
// for forward compatibility.
type RandomOrderGeneratorServer interface {
	SendOrder(context.Context, *Order) (*OrderResponse, error)
	mustEmbedUnimplementedRandomOrderGeneratorServer()
}

// UnimplementedRandomOrderGeneratorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRandomOrderGeneratorServer struct{}

func (UnimplementedRandomOrderGeneratorServer) SendOrder(context.Context, *Order) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrder not implemented")
}
func (UnimplementedRandomOrderGeneratorServer) mustEmbedUnimplementedRandomOrderGeneratorServer() {}
func (UnimplementedRandomOrderGeneratorServer) testEmbeddedByValue()                              {}

// UnsafeRandomOrderGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomOrderGeneratorServer will
// result in compilation errors.
type UnsafeRandomOrderGeneratorServer interface {
	mustEmbedUnimplementedRandomOrderGeneratorServer()
}

func RegisterRandomOrderGeneratorServer(s grpc.ServiceRegistrar, srv RandomOrderGeneratorServer) {
	// If the following call pancis, it indicates UnimplementedRandomOrderGeneratorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RandomOrderGenerator_ServiceDesc, srv)
}

func _RandomOrderGenerator_SendOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomOrderGeneratorServer).SendOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomOrderGenerator_SendOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomOrderGeneratorServer).SendOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

// RandomOrderGenerator_ServiceDesc is the grpc.ServiceDesc for RandomOrderGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandomOrderGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.random_order_generator",
	HandlerType: (*RandomOrderGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOrder",
			Handler:    _RandomOrderGenerator_SendOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
