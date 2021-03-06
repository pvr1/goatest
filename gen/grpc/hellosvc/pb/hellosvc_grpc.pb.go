// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hellosvcpb

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

// HellosvcClient is the client API for Hellosvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HellosvcClient interface {
	// Greet implements greet.
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type hellosvcClient struct {
	cc grpc.ClientConnInterface
}

func NewHellosvcClient(cc grpc.ClientConnInterface) HellosvcClient {
	return &hellosvcClient{cc}
}

func (c *hellosvcClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/hellosvc.Hellosvc/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HellosvcServer is the server API for Hellosvc service.
// All implementations must embed UnimplementedHellosvcServer
// for forward compatibility
type HellosvcServer interface {
	// Greet implements greet.
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	mustEmbedUnimplementedHellosvcServer()
}

// UnimplementedHellosvcServer must be embedded to have forward compatible implementations.
type UnimplementedHellosvcServer struct {
}

func (UnimplementedHellosvcServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedHellosvcServer) mustEmbedUnimplementedHellosvcServer() {}

// UnsafeHellosvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HellosvcServer will
// result in compilation errors.
type UnsafeHellosvcServer interface {
	mustEmbedUnimplementedHellosvcServer()
}

func RegisterHellosvcServer(s grpc.ServiceRegistrar, srv HellosvcServer) {
	s.RegisterService(&Hellosvc_ServiceDesc, srv)
}

func _Hellosvc_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HellosvcServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellosvc.Hellosvc/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HellosvcServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hellosvc_ServiceDesc is the grpc.ServiceDesc for Hellosvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hellosvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hellosvc.Hellosvc",
	HandlerType: (*HellosvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Hellosvc_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hellosvc.proto",
}
