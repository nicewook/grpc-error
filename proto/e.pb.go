// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/e.proto

package sqrtpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SqrtReq struct {
	SqrtInput            int32    `protobuf:"varint,1,opt,name=SqrtInput,proto3" json:"SqrtInput,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtReq) Reset()         { *m = SqrtReq{} }
func (m *SqrtReq) String() string { return proto.CompactTextString(m) }
func (*SqrtReq) ProtoMessage()    {}
func (*SqrtReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f42d666f28ae8b, []int{0}
}

func (m *SqrtReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtReq.Unmarshal(m, b)
}
func (m *SqrtReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtReq.Marshal(b, m, deterministic)
}
func (m *SqrtReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtReq.Merge(m, src)
}
func (m *SqrtReq) XXX_Size() int {
	return xxx_messageInfo_SqrtReq.Size(m)
}
func (m *SqrtReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtReq.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtReq proto.InternalMessageInfo

func (m *SqrtReq) GetSqrtInput() int32 {
	if m != nil {
		return m.SqrtInput
	}
	return 0
}

type SqrtRes struct {
	SqrtResult           float32  `protobuf:"fixed32,1,opt,name=SqrtResult,proto3" json:"SqrtResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtRes) Reset()         { *m = SqrtRes{} }
func (m *SqrtRes) String() string { return proto.CompactTextString(m) }
func (*SqrtRes) ProtoMessage()    {}
func (*SqrtRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f42d666f28ae8b, []int{1}
}

func (m *SqrtRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtRes.Unmarshal(m, b)
}
func (m *SqrtRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtRes.Marshal(b, m, deterministic)
}
func (m *SqrtRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtRes.Merge(m, src)
}
func (m *SqrtRes) XXX_Size() int {
	return xxx_messageInfo_SqrtRes.Size(m)
}
func (m *SqrtRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtRes.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtRes proto.InternalMessageInfo

func (m *SqrtRes) GetSqrtResult() float32 {
	if m != nil {
		return m.SqrtResult
	}
	return 0
}

func init() {
	proto.RegisterType((*SqrtReq)(nil), "sqrt.SqrtReq")
	proto.RegisterType((*SqrtRes)(nil), "sqrt.SqrtRes")
}

func init() { proto.RegisterFile("proto/e.proto", fileDescriptor_b0f42d666f28ae8b) }

var fileDescriptor_b0f42d666f28ae8b = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xd5, 0x03, 0xd3, 0x42, 0x2c, 0xc5, 0x85, 0x45, 0x25, 0x4a, 0xea, 0x5c, 0xec,
	0xc1, 0x85, 0x45, 0x25, 0x41, 0xa9, 0x85, 0x42, 0x32, 0x5c, 0x9c, 0x20, 0xa6, 0x67, 0x5e, 0x41,
	0x69, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x42, 0x40, 0x49, 0x13, 0xa6, 0xb0, 0x58,
	0x48, 0x8e, 0x8b, 0x0b, 0xca, 0x2c, 0xcd, 0x81, 0xa8, 0x64, 0x0a, 0x42, 0x12, 0x31, 0x32, 0xe5,
	0xe2, 0x06, 0xf1, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xd4, 0xb8, 0x58, 0x40, 0x5c,
	0x21, 0x5e, 0x3d, 0x90, 0x8d, 0x7a, 0x50, 0xeb, 0xa4, 0x50, 0xb8, 0xc5, 0x4a, 0x0c, 0x4e, 0x1c,
	0x51, 0x6c, 0x20, 0x91, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x0b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x96, 0xa3, 0x51, 0xb2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SqrtServiceClient is the client API for SqrtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqrtServiceClient interface {
	Sqrt(ctx context.Context, in *SqrtReq, opts ...grpc.CallOption) (*SqrtRes, error)
}

type sqrtServiceClient struct {
	cc *grpc.ClientConn
}

func NewSqrtServiceClient(cc *grpc.ClientConn) SqrtServiceClient {
	return &sqrtServiceClient{cc}
}

func (c *sqrtServiceClient) Sqrt(ctx context.Context, in *SqrtReq, opts ...grpc.CallOption) (*SqrtRes, error) {
	out := new(SqrtRes)
	err := c.cc.Invoke(ctx, "/sqrt.SqrtService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqrtServiceServer is the server API for SqrtService service.
type SqrtServiceServer interface {
	Sqrt(context.Context, *SqrtReq) (*SqrtRes, error)
}

// UnimplementedSqrtServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSqrtServiceServer struct {
}

func (*UnimplementedSqrtServiceServer) Sqrt(ctx context.Context, req *SqrtReq) (*SqrtRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}

func RegisterSqrtServiceServer(s *grpc.Server, srv SqrtServiceServer) {
	s.RegisterService(&_SqrtService_serviceDesc, srv)
}

func _SqrtService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqrtServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqrt.SqrtService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqrtServiceServer).Sqrt(ctx, req.(*SqrtReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SqrtService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sqrt.SqrtService",
	HandlerType: (*SqrtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sqrt",
			Handler:    _SqrtService_Sqrt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/e.proto",
}
