// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package mwitkow_testproto is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Empty
	PingRequest
	PingResponse
*/
package mwitkow_testproto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingRequest struct {
	Value             string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	SleepTimeMs       int32  `protobuf:"varint,2,opt,name=sleep_time_ms,json=sleepTimeMs" json:"sleep_time_ms,omitempty"`
	ErrorCodeReturned uint32 `protobuf:"varint,3,opt,name=error_code_returned,json=errorCodeReturned" json:"error_code_returned,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetSleepTimeMs() int32 {
	if m != nil {
		return m.SleepTimeMs
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

type PingResponse struct {
	Value   string `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
	Counter int32  `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "mwitkow.testproto.Empty")
	proto.RegisterType((*PingRequest)(nil), "mwitkow.testproto.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "mwitkow.testproto.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
	PingStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingStreamClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/mwitkow.testproto.TestService/PingEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/mwitkow.testproto.TestService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/mwitkow.testproto.TestService/PingError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/mwitkow.testproto.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/mwitkow.testproto.TestService/PingStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingStreamClient{stream}
	return x, nil
}

type TestService_PingStreamClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingStreamClient struct {
	grpc.ClientStream
}

func (x *testServicePingStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingStreamClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	PingEmpty(context.Context, *Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
	PingStream(TestService_PingStreamServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mwitkow.testproto.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mwitkow.testproto.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mwitkow.testproto.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_PingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingStream(&testServicePingStreamServer{stream})
}

type TestService_PingStreamServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingStreamServer struct {
	grpc.ServerStream
}

func (x *testServicePingStreamServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mwitkow.testproto.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingStream",
			Handler:       _TestService_PingStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xed, 0xfe, 0xfa, 0x8b, 0xb5, 0x13, 0x7b, 0xe8, 0xea, 0x21, 0x78, 0xd0, 0xb0, 0xa7, 0x9c,
	0x42, 0xd1, 0xbb, 0x17, 0x11, 0x15, 0x14, 0x25, 0x29, 0x5e, 0x43, 0x4d, 0x06, 0x59, 0xec, 0x66,
	0xe3, 0xee, 0xa4, 0xc1, 0x8f, 0xe1, 0x37, 0x96, 0xdd, 0x46, 0x28, 0x68, 0xd1, 0x43, 0x8f, 0xef,
	0xbd, 0xe1, 0xfd, 0x19, 0x00, 0x42, 0x4b, 0x69, 0x63, 0x34, 0x69, 0x3e, 0x55, 0x9d, 0xa4, 0x57,
	0xdd, 0xa5, 0x8e, 0xf3, 0x94, 0x18, 0x41, 0x70, 0xa5, 0x1a, 0x7a, 0x17, 0x1d, 0x84, 0x8f, 0xb2,
	0x7e, 0xc9, 0xf0, 0xad, 0x45, 0x4b, 0xfc, 0x08, 0x82, 0xd5, 0x62, 0xd9, 0x62, 0xc4, 0x62, 0x96,
	0x8c, 0xb3, 0x35, 0xe0, 0x02, 0x26, 0x76, 0x89, 0xd8, 0x14, 0x24, 0x15, 0x16, 0xca, 0x46, 0xff,
	0x62, 0x96, 0x04, 0x59, 0xe8, 0xc9, 0xb9, 0x54, 0x78, 0x6f, 0x79, 0x0a, 0x87, 0x68, 0x8c, 0x36,
	0x45, 0xa9, 0x2b, 0x2c, 0x0c, 0x52, 0x6b, 0x6a, 0xac, 0xa2, 0x61, 0xcc, 0x92, 0x49, 0x36, 0xf5,
	0xd2, 0xa5, 0xae, 0x30, 0xeb, 0x05, 0x71, 0x01, 0x07, 0xeb, 0x60, 0xdb, 0xe8, 0xda, 0xa2, 0x4b,
	0x7e, 0xda, 0x4c, 0xf6, 0x80, 0x47, 0x30, 0x2a, 0x75, 0x5b, 0x13, 0x9a, 0x3e, 0xf3, 0x0b, 0x9e,
	0x7d, 0x0c, 0x21, 0x9c, 0xa3, 0xa5, 0x1c, 0xcd, 0x4a, 0x96, 0xc8, 0x6f, 0x60, 0xec, 0xfc, 0xfc,
	0x2a, 0x1e, 0xa5, 0xdf, 0x26, 0xa7, 0x5e, 0x39, 0x3e, 0xfd, 0x41, 0xd9, 0xec, 0x21, 0x06, 0xfc,
	0x16, 0xfe, 0x3b, 0x86, 0x9f, 0x6c, 0x3d, 0xf5, 0xbf, 0xfa, 0x8b, 0xd5, 0x75, 0x5f, 0xca, 0xad,
	0xff, 0xd5, 0x6f, 0x6b, 0x69, 0x31, 0xe0, 0x0f, 0xb0, 0xef, 0x4e, 0xef, 0xa4, 0xa5, 0x1d, 0xf4,
	0x9a, 0x31, 0x9e, 0x03, 0x38, 0x2e, 0x27, 0x83, 0x0b, 0xb5, 0x03, 0xcb, 0x84, 0xcd, 0xd8, 0xf3,
	0x9e, 0x57, 0xce, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x75, 0xf0, 0x5c, 0x7d, 0x02, 0x00,
	0x00,
}
