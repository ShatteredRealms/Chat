// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	pb "github.com/shatteredrealms/chat/pkg/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FrontendServiceClient is the client API for FrontendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrontendServiceClient interface {
	// Sends a message to a channel
	Send(ctx context.Context, opts ...grpc.CallOption) (FrontendService_SendClient, error)
	// Connects a user to a channel and returns a stream of messages that were sent on the channel
	Connect(ctx context.Context, in *pb.Subscription, opts ...grpc.CallOption) (FrontendService_ConnectClient, error)
}

type frontendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontendServiceClient(cc grpc.ClientConnInterface) FrontendServiceClient {
	return &frontendServiceClient{cc}
}

func (c *frontendServiceClient) Send(ctx context.Context, opts ...grpc.CallOption) (FrontendService_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &FrontendService_ServiceDesc.Streams[0], "/chat.FrontendService/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendServiceSendClient{stream}
	return x, nil
}

type FrontendService_SendClient interface {
	Send(*pb.Message) error
	CloseAndRecv() (*EmptyMessage, error)
	grpc.ClientStream
}

type frontendServiceSendClient struct {
	grpc.ClientStream
}

func (x *frontendServiceSendClient) Send(m *pb.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontendServiceSendClient) CloseAndRecv() (*EmptyMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *frontendServiceClient) Connect(ctx context.Context, in *pb.Subscription, opts ...grpc.CallOption) (FrontendService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &FrontendService_ServiceDesc.Streams[1], "/chat.FrontendService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FrontendService_ConnectClient interface {
	Recv() (*pb.RawMessage, error)
	grpc.ClientStream
}

type frontendServiceConnectClient struct {
	grpc.ClientStream
}

func (x *frontendServiceConnectClient) Recv() (*pb.RawMessage, error) {
	m := new(pb.RawMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontendServiceServer is the server API for FrontendService service.
// All implementations must embed UnimplementedFrontendServiceServer
// for forward compatibility
type FrontendServiceServer interface {
	// Sends a message to a channel
	Send(FrontendService_SendServer) error
	// Connects a user to a channel and returns a stream of messages that were sent on the channel
	Connect(*pb.Subscription, FrontendService_ConnectServer) error
	mustEmbedUnimplementedFrontendServiceServer()
}

// UnimplementedFrontendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFrontendServiceServer struct {
}

func (UnimplementedFrontendServiceServer) Send(FrontendService_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedFrontendServiceServer) Connect(*pb.Subscription, FrontendService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedFrontendServiceServer) mustEmbedUnimplementedFrontendServiceServer() {}

// UnsafeFrontendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrontendServiceServer will
// result in compilation errors.
type UnsafeFrontendServiceServer interface {
	mustEmbedUnimplementedFrontendServiceServer()
}

func RegisterFrontendServiceServer(s grpc.ServiceRegistrar, srv FrontendServiceServer) {
	s.RegisterService(&FrontendService_ServiceDesc, srv)
}

func _FrontendService_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontendServiceServer).Send(&frontendServiceSendServer{stream})
}

type FrontendService_SendServer interface {
	SendAndClose(*EmptyMessage) error
	Recv() (*pb.Message, error)
	grpc.ServerStream
}

type frontendServiceSendServer struct {
	grpc.ServerStream
}

func (x *frontendServiceSendServer) SendAndClose(m *EmptyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontendServiceSendServer) Recv() (*pb.Message, error) {
	m := new(pb.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FrontendService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.Subscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FrontendServiceServer).Connect(m, &frontendServiceConnectServer{stream})
}

type FrontendService_ConnectServer interface {
	Send(*pb.RawMessage) error
	grpc.ServerStream
}

type frontendServiceConnectServer struct {
	grpc.ServerStream
}

func (x *frontendServiceConnectServer) Send(m *pb.RawMessage) error {
	return x.ServerStream.SendMsg(m)
}

// FrontendService_ServiceDesc is the grpc.ServiceDesc for FrontendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FrontendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.FrontendService",
	HandlerType: (*FrontendServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _FrontendService_Send_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Connect",
			Handler:       _FrontendService_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "frontend.proto",
}