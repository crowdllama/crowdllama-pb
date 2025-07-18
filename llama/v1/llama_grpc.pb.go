// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: llama/v1/llama.proto

package llamav1

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
	LlamaService_Generate_FullMethodName = "/llama.v1.LlamaService/Generate"
)

// LlamaServiceClient is the client API for LlamaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// LlamaService defines the service interface for Llama API
type LlamaServiceClient interface {
	// Generate generates a response for the given request
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type llamaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLlamaServiceClient(cc grpc.ClientConnInterface) LlamaServiceClient {
	return &llamaServiceClient{cc}
}

func (c *llamaServiceClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, LlamaService_Generate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LlamaServiceServer is the server API for LlamaService service.
// All implementations must embed UnimplementedLlamaServiceServer
// for forward compatibility.
//
// LlamaService defines the service interface for Llama API
type LlamaServiceServer interface {
	// Generate generates a response for the given request
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	mustEmbedUnimplementedLlamaServiceServer()
}

// UnimplementedLlamaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLlamaServiceServer struct{}

func (UnimplementedLlamaServiceServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedLlamaServiceServer) mustEmbedUnimplementedLlamaServiceServer() {}
func (UnimplementedLlamaServiceServer) testEmbeddedByValue()                      {}

// UnsafeLlamaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LlamaServiceServer will
// result in compilation errors.
type UnsafeLlamaServiceServer interface {
	mustEmbedUnimplementedLlamaServiceServer()
}

func RegisterLlamaServiceServer(s grpc.ServiceRegistrar, srv LlamaServiceServer) {
	// If the following call pancis, it indicates UnimplementedLlamaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LlamaService_ServiceDesc, srv)
}

func _LlamaService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LlamaServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LlamaService_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LlamaServiceServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LlamaService_ServiceDesc is the grpc.ServiceDesc for LlamaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LlamaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llama.v1.LlamaService",
	HandlerType: (*LlamaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _LlamaService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "llama/v1/llama.proto",
}
