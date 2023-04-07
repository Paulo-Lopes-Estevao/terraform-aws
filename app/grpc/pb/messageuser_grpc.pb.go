// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: messageuser.proto

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

// MessageUserServiceClient is the client API for MessageUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageUserServiceClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Login(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (MessageUserService_SendMessageClient, error)
	FindByMessageUser(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*MessageResponse, error)
	ListAllUsersOntheNetwork(ctx context.Context, in *ListUser, opts ...grpc.CallOption) (MessageUserService_ListAllUsersOntheNetworkClient, error)
}

type messageUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageUserServiceClient(cc grpc.ClientConnInterface) MessageUserServiceClient {
	return &messageUserServiceClient{cc}
}

func (c *messageUserServiceClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/app.grpc.protofile.MessageUserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageUserServiceClient) Login(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/app.grpc.protofile.MessageUserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageUserServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (MessageUserService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageUserService_ServiceDesc.Streams[0], "/app.grpc.protofile.MessageUserService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageUserServiceSendMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageUserService_SendMessageClient interface {
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type messageUserServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *messageUserServiceSendMessageClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageUserServiceClient) FindByMessageUser(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/app.grpc.protofile.MessageUserService/FindByMessageUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageUserServiceClient) ListAllUsersOntheNetwork(ctx context.Context, in *ListUser, opts ...grpc.CallOption) (MessageUserService_ListAllUsersOntheNetworkClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageUserService_ServiceDesc.Streams[1], "/app.grpc.protofile.MessageUserService/listAllUsersOntheNetwork", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageUserServiceListAllUsersOntheNetworkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageUserService_ListAllUsersOntheNetworkClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type messageUserServiceListAllUsersOntheNetworkClient struct {
	grpc.ClientStream
}

func (x *messageUserServiceListAllUsersOntheNetworkClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageUserServiceServer is the server API for MessageUserService service.
// All implementations must embed UnimplementedMessageUserServiceServer
// for forward compatibility
type MessageUserServiceServer interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
	Login(context.Context, *UserRequest) (*UserResponse, error)
	SendMessage(*MessageRequest, MessageUserService_SendMessageServer) error
	FindByMessageUser(context.Context, *GetUserById) (*MessageResponse, error)
	ListAllUsersOntheNetwork(*ListUser, MessageUserService_ListAllUsersOntheNetworkServer) error
	mustEmbedUnimplementedMessageUserServiceServer()
}

// UnimplementedMessageUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageUserServiceServer struct {
}

func (UnimplementedMessageUserServiceServer) CreateUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMessageUserServiceServer) Login(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMessageUserServiceServer) SendMessage(*MessageRequest, MessageUserService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessageUserServiceServer) FindByMessageUser(context.Context, *GetUserById) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByMessageUser not implemented")
}
func (UnimplementedMessageUserServiceServer) ListAllUsersOntheNetwork(*ListUser, MessageUserService_ListAllUsersOntheNetworkServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllUsersOntheNetwork not implemented")
}
func (UnimplementedMessageUserServiceServer) mustEmbedUnimplementedMessageUserServiceServer() {}

// UnsafeMessageUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageUserServiceServer will
// result in compilation errors.
type UnsafeMessageUserServiceServer interface {
	mustEmbedUnimplementedMessageUserServiceServer()
}

func RegisterMessageUserServiceServer(s grpc.ServiceRegistrar, srv MessageUserServiceServer) {
	s.RegisterService(&MessageUserService_ServiceDesc, srv)
}

func _MessageUserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageUserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.grpc.protofile.MessageUserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageUserServiceServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageUserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageUserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.grpc.protofile.MessageUserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageUserServiceServer).Login(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageUserService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageUserServiceServer).SendMessage(m, &messageUserServiceSendMessageServer{stream})
}

type MessageUserService_SendMessageServer interface {
	Send(*MessageResponse) error
	grpc.ServerStream
}

type messageUserServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *messageUserServiceSendMessageServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageUserService_FindByMessageUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageUserServiceServer).FindByMessageUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.grpc.protofile.MessageUserService/FindByMessageUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageUserServiceServer).FindByMessageUser(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageUserService_ListAllUsersOntheNetwork_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUser)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageUserServiceServer).ListAllUsersOntheNetwork(m, &messageUserServiceListAllUsersOntheNetworkServer{stream})
}

type MessageUserService_ListAllUsersOntheNetworkServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type messageUserServiceListAllUsersOntheNetworkServer struct {
	grpc.ServerStream
}

func (x *messageUserServiceListAllUsersOntheNetworkServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MessageUserService_ServiceDesc is the grpc.ServiceDesc for MessageUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.grpc.protofile.MessageUserService",
	HandlerType: (*MessageUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MessageUserService_CreateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _MessageUserService_Login_Handler,
		},
		{
			MethodName: "FindByMessageUser",
			Handler:    _MessageUserService_FindByMessageUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _MessageUserService_SendMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listAllUsersOntheNetwork",
			Handler:       _MessageUserService_ListAllUsersOntheNetwork_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messageuser.proto",
}