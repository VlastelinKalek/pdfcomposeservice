// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/api-stream.proto

package composeserver

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
	Server_ComposePDF_FullMethodName = "/composeserver.Server/ComposePDF"
)

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	ComposePDF(ctx context.Context, opts ...grpc.CallOption) (Server_ComposePDFClient, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) ComposePDF(ctx context.Context, opts ...grpc.CallOption) (Server_ComposePDFClient, error) {
	stream, err := c.cc.NewStream(ctx, &Server_ServiceDesc.Streams[0], Server_ComposePDF_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serverComposePDFClient{stream}
	return x, nil
}

type Server_ComposePDFClient interface {
	Send(*FileList) error
	Recv() (*FilePDF, error)
	grpc.ClientStream
}

type serverComposePDFClient struct {
	grpc.ClientStream
}

func (x *serverComposePDFClient) Send(m *FileList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverComposePDFClient) Recv() (*FilePDF, error) {
	m := new(FilePDF)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	ComposePDF(Server_ComposePDFServer) error
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) ComposePDF(Server_ComposePDFServer) error {
	return status.Errorf(codes.Unimplemented, "method ComposePDF not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_ComposePDF_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerServer).ComposePDF(&serverComposePDFServer{stream})
}

type Server_ComposePDFServer interface {
	Send(*FilePDF) error
	Recv() (*FileList, error)
	grpc.ServerStream
}

type serverComposePDFServer struct {
	grpc.ServerStream
}

func (x *serverComposePDFServer) Send(m *FilePDF) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverComposePDFServer) Recv() (*FileList, error) {
	m := new(FileList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "composeserver.Server",
	HandlerType: (*ServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComposePDF",
			Handler:       _Server_ComposePDF_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api-stream.proto",
}
