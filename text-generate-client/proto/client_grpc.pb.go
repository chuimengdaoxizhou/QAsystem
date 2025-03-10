// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: client.proto

package streaming_model

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
	ModelService_GenerateContentStream_FullMethodName = "/streaming_model.ModelService/GenerateContentStream"
)

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 ModelService 服务
type ModelServiceClient interface {
	// 服务器端流式 RPC 方法：客户端发送一个字符串请求，服务器流式返回 token
	GenerateContentStream(ctx context.Context, in *GenerationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TokenResponse], error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) GenerateContentStream(ctx context.Context, in *GenerationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TokenResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ModelService_ServiceDesc.Streams[0], ModelService_GenerateContentStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GenerationRequest, TokenResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ModelService_GenerateContentStreamClient = grpc.ServerStreamingClient[TokenResponse]

// ModelServiceServer is the server API for ModelService service.
// All implementations must embed UnimplementedModelServiceServer
// for forward compatibility.
//
// 定义 ModelService 服务
type ModelServiceServer interface {
	// 服务器端流式 RPC 方法：客户端发送一个字符串请求，服务器流式返回 token
	GenerateContentStream(*GenerationRequest, grpc.ServerStreamingServer[TokenResponse]) error
	mustEmbedUnimplementedModelServiceServer()
}

// UnimplementedModelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModelServiceServer struct{}

func (UnimplementedModelServiceServer) GenerateContentStream(*GenerationRequest, grpc.ServerStreamingServer[TokenResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateContentStream not implemented")
}
func (UnimplementedModelServiceServer) mustEmbedUnimplementedModelServiceServer() {}
func (UnimplementedModelServiceServer) testEmbeddedByValue()                      {}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	// If the following call pancis, it indicates UnimplementedModelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_GenerateContentStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GenerationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelServiceServer).GenerateContentStream(m, &grpc.GenericServerStream[GenerationRequest, TokenResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ModelService_GenerateContentStreamServer = grpc.ServerStreamingServer[TokenResponse]

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streaming_model.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateContentStream",
			Handler:       _ModelService_GenerateContentStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "client.proto",
}
