// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/course.proto

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

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error)
	CreateCourseStream(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamClient, error)
	CreateCourseStreamBidirecional(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamBidirecionalClient, error)
	ListCourses(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*CourseList, error)
	GetCourse(ctx context.Context, in *CourseGetRequest, opts ...grpc.CallOption) (*CourseResponse, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) CreateCourseStream(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[0], "/pb.CourseService/CreateCourseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceCreateCourseStreamClient{stream}
	return x, nil
}

type CourseService_CreateCourseStreamClient interface {
	Send(*CreateCourseRequest) error
	CloseAndRecv() (*CourseList, error)
	grpc.ClientStream
}

type courseServiceCreateCourseStreamClient struct {
	grpc.ClientStream
}

func (x *courseServiceCreateCourseStreamClient) Send(m *CreateCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamClient) CloseAndRecv() (*CourseList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CourseList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courseServiceClient) CreateCourseStreamBidirecional(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamBidirecionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[1], "/pb.CourseService/CreateCourseStreamBidirecional", opts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceCreateCourseStreamBidirecionalClient{stream}
	return x, nil
}

type CourseService_CreateCourseStreamBidirecionalClient interface {
	Send(*CreateCourseRequest) error
	Recv() (*Course, error)
	grpc.ClientStream
}

type courseServiceCreateCourseStreamBidirecionalClient struct {
	grpc.ClientStream
}

func (x *courseServiceCreateCourseStreamBidirecionalClient) Send(m *CreateCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamBidirecionalClient) Recv() (*Course, error) {
	m := new(Course)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courseServiceClient) ListCourses(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*CourseList, error) {
	out := new(CourseList)
	err := c.cc.Invoke(ctx, "/pb.CourseService/ListCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourse(ctx context.Context, in *CourseGetRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	CreateCourse(context.Context, *CreateCourseRequest) (*CourseResponse, error)
	CreateCourseStream(CourseService_CreateCourseStreamServer) error
	CreateCourseStreamBidirecional(CourseService_CreateCourseStreamBidirecionalServer) error
	ListCourses(context.Context, *Blank) (*CourseList, error)
	GetCourse(context.Context, *CourseGetRequest) (*CourseResponse, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) CreateCourse(context.Context, *CreateCourseRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourseStream(CourseService_CreateCourseStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCourseStream not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourseStreamBidirecional(CourseService_CreateCourseStreamBidirecionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCourseStreamBidirecional not implemented")
}
func (UnimplementedCourseServiceServer) ListCourses(context.Context, *Blank) (*CourseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedCourseServiceServer) GetCourse(context.Context, *CourseGetRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_CreateCourseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourseServiceServer).CreateCourseStream(&courseServiceCreateCourseStreamServer{stream})
}

type CourseService_CreateCourseStreamServer interface {
	SendAndClose(*CourseList) error
	Recv() (*CreateCourseRequest, error)
	grpc.ServerStream
}

type courseServiceCreateCourseStreamServer struct {
	grpc.ServerStream
}

func (x *courseServiceCreateCourseStreamServer) SendAndClose(m *CourseList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamServer) Recv() (*CreateCourseRequest, error) {
	m := new(CreateCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CourseService_CreateCourseStreamBidirecional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourseServiceServer).CreateCourseStreamBidirecional(&courseServiceCreateCourseStreamBidirecionalServer{stream})
}

type CourseService_CreateCourseStreamBidirecionalServer interface {
	Send(*Course) error
	Recv() (*CreateCourseRequest, error)
	grpc.ServerStream
}

type courseServiceCreateCourseStreamBidirecionalServer struct {
	grpc.ServerStream
}

func (x *courseServiceCreateCourseStreamBidirecionalServer) Send(m *Course) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamBidirecionalServer) Recv() (*CreateCourseRequest, error) {
	m := new(CreateCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CourseService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/ListCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).ListCourses(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourse(ctx, req.(*CourseGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourse",
			Handler:    _CourseService_CreateCourse_Handler,
		},
		{
			MethodName: "ListCourses",
			Handler:    _CourseService_ListCourses_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _CourseService_GetCourse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCourseStream",
			Handler:       _CourseService_CreateCourseStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateCourseStreamBidirecional",
			Handler:       _CourseService_CreateCourseStreamBidirecional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/course.proto",
}
