// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// BackendServiceClient is the client API for BackendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendServiceClient interface {
	// Gets all available channels
	GetChannels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Channel, error)
	// Registers a new channel with the database
	CreateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*GenericResponse, error)
	// Deletes a channel
	DeleteChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*GenericResponse, error)
}

type backendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendServiceClient(cc grpc.ClientConnInterface) BackendServiceClient {
	return &backendServiceClient{cc}
}

func (c *backendServiceClient) GetChannels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/chat.BackendService/GetChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServiceClient) CreateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/chat.BackendService/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServiceClient) DeleteChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/chat.BackendService/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServiceServer is the server API for BackendService service.
// All implementations must embed UnimplementedBackendServiceServer
// for forward compatibility
type BackendServiceServer interface {
	// Gets all available channels
	GetChannels(context.Context, *Empty) (*Channel, error)
	// Registers a new channel with the database
	CreateChannel(context.Context, *Channel) (*GenericResponse, error)
	// Deletes a channel
	DeleteChannel(context.Context, *Channel) (*GenericResponse, error)
	mustEmbedUnimplementedBackendServiceServer()
}

// UnimplementedBackendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServiceServer struct {
}

func (UnimplementedBackendServiceServer) GetChannels(context.Context, *Empty) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannels not implemented")
}
func (UnimplementedBackendServiceServer) CreateChannel(context.Context, *Channel) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedBackendServiceServer) DeleteChannel(context.Context, *Channel) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedBackendServiceServer) mustEmbedUnimplementedBackendServiceServer() {}

// UnsafeBackendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServiceServer will
// result in compilation errors.
type UnsafeBackendServiceServer interface {
	mustEmbedUnimplementedBackendServiceServer()
}

func RegisterBackendServiceServer(s grpc.ServiceRegistrar, srv BackendServiceServer) {
	s.RegisterService(&BackendService_ServiceDesc, srv)
}

func _BackendService_GetChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).GetChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.BackendService/GetChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).GetChannels(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendService_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.BackendService/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).CreateChannel(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.BackendService/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).DeleteChannel(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

// BackendService_ServiceDesc is the grpc.ServiceDesc for BackendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.BackendService",
	HandlerType: (*BackendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChannels",
			Handler:    _BackendService_GetChannels_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _BackendService_CreateChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _BackendService_DeleteChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}
