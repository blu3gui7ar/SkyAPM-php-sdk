// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/JVMMetricsService.proto

package agent

import (
	common "agent/agent/pb6/common"
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

type JVMMetrics struct {
	Metrics               []*common.JVMMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ApplicationInstanceId int32               `protobuf:"varint,2,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}            `json:"-"`
	XXX_unrecognized      []byte              `json:"-"`
	XXX_sizecache         int32               `json:"-"`
}

func (m *JVMMetrics) Reset()         { *m = JVMMetrics{} }
func (m *JVMMetrics) String() string { return proto.CompactTextString(m) }
func (*JVMMetrics) ProtoMessage()    {}
func (*JVMMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_84f947edc73ca4ba, []int{0}
}

func (m *JVMMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetrics.Unmarshal(m, b)
}
func (m *JVMMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetrics.Marshal(b, m, deterministic)
}
func (m *JVMMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetrics.Merge(m, src)
}
func (m *JVMMetrics) XXX_Size() int {
	return xxx_messageInfo_JVMMetrics.Size(m)
}
func (m *JVMMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetrics proto.InternalMessageInfo

func (m *JVMMetrics) GetMetrics() []*common.JVMMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *JVMMetrics) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

func init() {
	proto.RegisterType((*JVMMetrics)(nil), "JVMMetrics")
}

func init() {
	proto.RegisterFile("language-agent/JVMMetricsService.proto", fileDescriptor_84f947edc73ca4ba)
}

var fileDescriptor_84f947edc73ca4ba = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x75, 0x2b, 0x5a, 0xc8, 0x5e, 0x74, 0xa1, 0x50, 0xf6, 0x62, 0x29, 0x2a, 0xbd, 0x18, 0xa5,
	0x8a, 0x07, 0x8f, 0xe2, 0xa5, 0x42, 0xa5, 0xb4, 0xa0, 0xe0, 0x6d, 0x1a, 0x87, 0xdd, 0xb0, 0x49,
	0x26, 0x24, 0xd1, 0xa5, 0xbf, 0xe4, 0x57, 0xca, 0x36, 0x5d, 0x17, 0xb4, 0x97, 0x30, 0xcc, 0x7b,
	0x79, 0xf3, 0xde, 0x63, 0x97, 0x0a, 0x4c, 0xf1, 0x09, 0x05, 0x5e, 0x41, 0x81, 0x26, 0x5c, 0x3f,
	0xbf, 0xce, 0xe7, 0x18, 0x9c, 0x14, 0x7e, 0x85, 0xee, 0x4b, 0x0a, 0xe4, 0xd6, 0x51, 0xa0, 0xfc,
	0xec, 0x0f, 0xef, 0x89, 0x6a, 0xe3, 0x83, 0x43, 0xd0, 0x3b, 0xc2, 0x89, 0x20, 0xad, 0xc9, 0x34,
	0x02, 0x71, 0x33, 0x2e, 0x19, 0xeb, 0xd4, 0xb2, 0x73, 0xd6, 0xd7, 0x71, 0x1c, 0x26, 0xa3, 0xc3,
	0x49, 0x3a, 0x65, 0xfc, 0x17, 0x5d, 0xb6, 0x50, 0x76, 0xc7, 0x06, 0x60, 0xad, 0x92, 0x02, 0x82,
	0x24, 0x33, 0x33, 0x3e, 0x80, 0x11, 0x38, 0xfb, 0x18, 0xf6, 0x46, 0xc9, 0xe4, 0x68, 0xb9, 0x1f,
	0x9c, 0x3e, 0xb0, 0xd3, 0x7f, 0xbe, 0xb3, 0x0b, 0xd6, 0x17, 0xa4, 0x14, 0x8a, 0x90, 0xa5, 0xdd,
	0x29, 0x9f, 0xa7, 0xbc, 0xf3, 0x3e, 0x3e, 0x78, 0x2c, 0xd9, 0x0d, 0xb9, 0x82, 0x83, 0x05, 0x51,
	0x22, 0xf7, 0xd5, 0xa6, 0x06, 0x55, 0x49, 0xd3, 0x6c, 0x34, 0x37, 0x18, 0x6a, 0x72, 0x15, 0x6f,
	0xc3, 0xf3, 0x6d, 0xf8, 0x45, 0xf2, 0x3e, 0x88, 0x2d, 0xc4, 0xd7, 0xae, 0xef, 0xe3, 0xf4, 0xdd,
	0xcb, 0x57, 0xd5, 0xe6, 0x6d, 0x27, 0xf0, 0x12, 0x3f, 0x2f, 0x9a, 0x32, 0x04, 0xa9, 0xf5, 0xf1,
	0xb6, 0x96, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x52, 0xc1, 0x7e, 0x73, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JVMMetricsServiceClient is the client API for JVMMetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JVMMetricsServiceClient interface {
	Collect(ctx context.Context, in *JVMMetrics, opts ...grpc.CallOption) (*Downstream, error)
}

type jVMMetricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewJVMMetricsServiceClient(cc *grpc.ClientConn) JVMMetricsServiceClient {
	return &jVMMetricsServiceClient{cc}
}

func (c *jVMMetricsServiceClient) Collect(ctx context.Context, in *JVMMetrics, opts ...grpc.CallOption) (*Downstream, error) {
	out := new(Downstream)
	err := c.cc.Invoke(ctx, "/JVMMetricsService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JVMMetricsServiceServer is the server API for JVMMetricsService service.
type JVMMetricsServiceServer interface {
	Collect(context.Context, *JVMMetrics) (*Downstream, error)
}

// UnimplementedJVMMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJVMMetricsServiceServer struct {
}

func (*UnimplementedJVMMetricsServiceServer) Collect(ctx context.Context, req *JVMMetrics) (*Downstream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterJVMMetricsServiceServer(s *grpc.Server, srv JVMMetricsServiceServer) {
	s.RegisterService(&_JVMMetricsService_serviceDesc, srv)
}

func _JVMMetricsService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JVMMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JVMMetricsServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JVMMetricsService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JVMMetricsServiceServer).Collect(ctx, req.(*JVMMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

var _JVMMetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JVMMetricsService",
	HandlerType: (*JVMMetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _JVMMetricsService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/JVMMetricsService.proto",
}
