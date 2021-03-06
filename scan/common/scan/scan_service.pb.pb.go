// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scan_service.pb

package scan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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

// The request message containing the start key and fetch num
type FindRequest struct {
	Start                *string  `protobuf:"bytes,1,req,name=start" json:"start,omitempty"`
	Num                  *int32   `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scan_service_684d02b83cb4c6cd, []int{0}
}
func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (dst *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(dst, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetStart() string {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return ""
}

func (m *FindRequest) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// The request message containing the start index and fetch num
type FindRequestByIndex struct {
	Start                *int32   `protobuf:"varint,1,req,name=start" json:"start,omitempty"`
	Num                  *int32   `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequestByIndex) Reset()         { *m = FindRequestByIndex{} }
func (m *FindRequestByIndex) String() string { return proto.CompactTextString(m) }
func (*FindRequestByIndex) ProtoMessage()    {}
func (*FindRequestByIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_scan_service_684d02b83cb4c6cd, []int{1}
}
func (m *FindRequestByIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequestByIndex.Unmarshal(m, b)
}
func (m *FindRequestByIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequestByIndex.Marshal(b, m, deterministic)
}
func (dst *FindRequestByIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequestByIndex.Merge(dst, src)
}
func (m *FindRequestByIndex) XXX_Size() int {
	return xxx_messageInfo_FindRequestByIndex.Size(m)
}
func (m *FindRequestByIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequestByIndex.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequestByIndex proto.InternalMessageInfo

func (m *FindRequestByIndex) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *FindRequestByIndex) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// The response message containing the errcode, data and the next fetch index
type FindResult struct {
	Err                  *int32   `protobuf:"varint,1,req,name=err" json:"err,omitempty"`
	Kvpairs              []string `protobuf:"bytes,2,rep,name=kvpairs" json:"kvpairs,omitempty"`
	Nextindex            *int32   `protobuf:"varint,3,req,name=nextindex" json:"nextindex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindResult) Reset()         { *m = FindResult{} }
func (m *FindResult) String() string { return proto.CompactTextString(m) }
func (*FindResult) ProtoMessage()    {}
func (*FindResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_scan_service_684d02b83cb4c6cd, []int{2}
}
func (m *FindResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResult.Unmarshal(m, b)
}
func (m *FindResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResult.Marshal(b, m, deterministic)
}
func (dst *FindResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResult.Merge(dst, src)
}
func (m *FindResult) XXX_Size() int {
	return xxx_messageInfo_FindResult.Size(m)
}
func (m *FindResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResult.DiscardUnknown(m)
}

var xxx_messageInfo_FindResult proto.InternalMessageInfo

func (m *FindResult) GetErr() int32 {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return 0
}

func (m *FindResult) GetKvpairs() []string {
	if m != nil {
		return m.Kvpairs
	}
	return nil
}

func (m *FindResult) GetNextindex() int32 {
	if m != nil && m.Nextindex != nil {
		return *m.Nextindex
	}
	return 0
}

func init() {
	proto.RegisterType((*FindRequest)(nil), "scan.FindRequest")
	proto.RegisterType((*FindRequestByIndex)(nil), "scan.FindRequestByIndex")
	proto.RegisterType((*FindResult)(nil), "scan.FindResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScanerClient is the client API for Scaner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScanerClient interface {
	// Sends a greeting
	FindByKey(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResult, error)
	FindByIndex(ctx context.Context, in *FindRequestByIndex, opts ...grpc.CallOption) (*FindResult, error)
}

type scanerClient struct {
	cc *grpc.ClientConn
}

func NewScanerClient(cc *grpc.ClientConn) ScanerClient {
	return &scanerClient{cc}
}

func (c *scanerClient) FindByKey(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResult, error) {
	out := new(FindResult)
	err := c.cc.Invoke(ctx, "/scan.Scaner/FindByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanerClient) FindByIndex(ctx context.Context, in *FindRequestByIndex, opts ...grpc.CallOption) (*FindResult, error) {
	out := new(FindResult)
	err := c.cc.Invoke(ctx, "/scan.Scaner/FindByIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanerServer is the server API for Scaner service.
type ScanerServer interface {
	// Sends a greeting
	FindByKey(context.Context, *FindRequest) (*FindResult, error)
	FindByIndex(context.Context, *FindRequestByIndex) (*FindResult, error)
}

func RegisterScanerServer(s *grpc.Server, srv ScanerServer) {
	s.RegisterService(&_Scaner_serviceDesc, srv)
}

func _Scaner_FindByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanerServer).FindByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scan.Scaner/FindByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanerServer).FindByKey(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scaner_FindByIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequestByIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanerServer).FindByIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scan.Scaner/FindByIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanerServer).FindByIndex(ctx, req.(*FindRequestByIndex))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scaner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scan.Scaner",
	HandlerType: (*ScanerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByKey",
			Handler:    _Scaner_FindByKey_Handler,
		},
		{
			MethodName: "FindByIndex",
			Handler:    _Scaner_FindByIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scan_service.pb",
}

func init() { proto.RegisterFile("scan_service.pb", fileDescriptor_scan_service_684d02b83cb4c6cd) }

var fileDescriptor_scan_service_684d02b83cb4c6cd = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4e, 0x4e, 0xcc,
	0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x48, 0x12, 0x62, 0x01, 0x09, 0x28,
	0x69, 0x72, 0x71, 0xbb, 0x65, 0xe6, 0xa5, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1,
	0x72, 0xb1, 0x16, 0x97, 0x24, 0x16, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x0a, 0x71, 0x73,
	0x31, 0xe7, 0x95, 0xe6, 0x4a, 0x30, 0x29, 0x30, 0x69, 0xb0, 0x2a, 0x19, 0x70, 0x09, 0x21, 0x29,
	0x75, 0xaa, 0xf4, 0xcc, 0x4b, 0x49, 0xad, 0x40, 0xd5, 0xc1, 0x8a, 0xaa, 0xc3, 0x96, 0x8b, 0x0b,
	0xa2, 0xa3, 0xb8, 0x34, 0xa7, 0x04, 0x24, 0x95, 0x5a, 0x54, 0x04, 0x55, 0xc7, 0xcf, 0xc5, 0x9e,
	0x5d, 0x56, 0x90, 0x98, 0x59, 0x54, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x29, 0x24, 0xc8, 0xc5,
	0x99, 0x97, 0x5a, 0x51, 0x92, 0x09, 0x32, 0x54, 0x82, 0x19, 0xa4, 0xc6, 0xa8, 0x92, 0x8b, 0x2d,
	0x38, 0x39, 0x31, 0x2f, 0xb5, 0x48, 0xc8, 0x88, 0x8b, 0x13, 0x64, 0x90, 0x53, 0xa5, 0x77, 0x6a,
	0xa5, 0x90, 0xa0, 0x1e, 0xc8, 0xe5, 0x7a, 0x48, 0x6e, 0x91, 0x12, 0x40, 0x16, 0x02, 0x59, 0xa6,
	0xc4, 0x20, 0x64, 0x0d, 0xf1, 0x19, 0xcc, 0x9d, 0x12, 0x18, 0xba, 0xa0, 0x32, 0xd8, 0x34, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x67, 0x18, 0x7f, 0x2e, 0x01, 0x00, 0x00,
}
