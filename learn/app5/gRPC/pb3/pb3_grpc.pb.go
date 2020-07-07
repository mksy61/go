// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DinoServiceClient is the client API for DinoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DinoServiceClient interface {
	GetAnimal(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Animal, error)
	GetAllAninaks(ctx context.Context, in *Request, opts ...grpc.CallOption) (DinoService_GetAllAninaksClient, error)
}

type dinoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDinoServiceClient(cc grpc.ClientConnInterface) DinoServiceClient {
	return &dinoServiceClient{cc}
}

func (c *dinoServiceClient) GetAnimal(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Animal, error) {
	out := new(Animal)
	err := c.cc.Invoke(ctx, "/pb3.DinoService/GetAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dinoServiceClient) GetAllAninaks(ctx context.Context, in *Request, opts ...grpc.CallOption) (DinoService_GetAllAninaksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DinoService_serviceDesc.Streams[0], "/pb3.DinoService/GetAllAninaks", opts...)
	if err != nil {
		return nil, err
	}
	x := &dinoServiceGetAllAninaksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DinoService_GetAllAninaksClient interface {
	Recv() (*Animal, error)
	grpc.ClientStream
}

type dinoServiceGetAllAninaksClient struct {
	grpc.ClientStream
}

func (x *dinoServiceGetAllAninaksClient) Recv() (*Animal, error) {
	m := new(Animal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DinoServiceServer is the server API for DinoService service.
// All implementations must embed UnimplementedDinoServiceServer
// for forward compatibility
type DinoServiceServer interface {
	GetAnimal(context.Context, *Request) (*Animal, error)
	GetAllAninaks(*Request, DinoService_GetAllAninaksServer) error
	mustEmbedUnimplementedDinoServiceServer()
}

// UnimplementedDinoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDinoServiceServer struct {
}

func (*UnimplementedDinoServiceServer) GetAnimal(context.Context, *Request) (*Animal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimal not implemented")
}
func (*UnimplementedDinoServiceServer) GetAllAninaks(*Request, DinoService_GetAllAninaksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAninaks not implemented")
}
func (*UnimplementedDinoServiceServer) mustEmbedUnimplementedDinoServiceServer() {}

func RegisterDinoServiceServer(s *grpc.Server, srv DinoServiceServer) {
	s.RegisterService(&_DinoService_serviceDesc, srv)
}

func _DinoService_GetAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DinoServiceServer).GetAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb3.DinoService/GetAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DinoServiceServer).GetAnimal(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DinoService_GetAllAninaks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DinoServiceServer).GetAllAninaks(m, &dinoServiceGetAllAninaksServer{stream})
}

type DinoService_GetAllAninaksServer interface {
	Send(*Animal) error
	grpc.ServerStream
}

type dinoServiceGetAllAninaksServer struct {
	grpc.ServerStream
}

func (x *dinoServiceGetAllAninaksServer) Send(m *Animal) error {
	return x.ServerStream.SendMsg(m)
}

var _DinoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb3.DinoService",
	HandlerType: (*DinoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnimal",
			Handler:    _DinoService_GetAnimal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAninaks",
			Handler:       _DinoService_GetAllAninaks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb3.proto",
}