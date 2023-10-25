// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_periodic_sign.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 签到信息
type SignInfo struct {
	Sign             *int32 `protobuf:"varint,1,req,name=Sign" json:"Sign,omitempty"`
	Resign           *int32 `protobuf:"varint,2,req,name=Resign" json:"Resign,omitempty"`
	Offset           *int32 `protobuf:"varint,3,req,name=Offset" json:"Offset,omitempty"`
	SignedToday      *bool  `protobuf:"varint,4,req,name=SignedToday" json:"SignedToday,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SignInfo) Reset()                    { *m = SignInfo{} }
func (m *SignInfo) String() string            { return proto.CompactTextString(m) }
func (*SignInfo) ProtoMessage()               {}
func (*SignInfo) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *SignInfo) GetSign() int32 {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return 0
}

func (m *SignInfo) GetResign() int32 {
	if m != nil && m.Resign != nil {
		return *m.Resign
	}
	return 0
}

func (m *SignInfo) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *SignInfo) GetSignedToday() bool {
	if m != nil && m.SignedToday != nil {
		return *m.SignedToday
	}
	return false
}

// 签到结果
type SignV1Rsp struct {
	Info             *SignInfo           `protobuf:"bytes,1,req,name=Info" json:"Info,omitempty"`
	Items            []*PropExchangeData `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *SignV1Rsp) Reset()                    { *m = SignV1Rsp{} }
func (m *SignV1Rsp) String() string            { return proto.CompactTextString(m) }
func (*SignV1Rsp) ProtoMessage()               {}
func (*SignV1Rsp) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *SignV1Rsp) GetInfo() *SignInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SignV1Rsp) GetItems() []*PropExchangeData {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*SignInfo)(nil), "VK.Proto.SignInfo")
	proto.RegisterType((*SignV1Rsp)(nil), "VK.Proto.SignV1Rsp")
}

func init() { proto.RegisterFile("client_periodic_sign.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xb1, 0xdb, 0x95, 0x3a, 0xbd, 0x8d, 0x20, 0xa1, 0xa7, 0xb2, 0x07, 0xe9, 0xa9, 0xe8,
	0xfe, 0x06, 0x3d, 0x2c, 0x1e, 0x5c, 0xa2, 0xec, 0xb5, 0x84, 0x66, 0xda, 0x0d, 0xd8, 0x4c, 0x48,
	0x72, 0xd0, 0x7f, 0x2f, 0x49, 0x77, 0xd1, 0xdb, 0x7b, 0xdf, 0xbc, 0xe4, 0x3d, 0x68, 0xc6, 0x2f,
	0x43, 0x36, 0x0e, 0x8e, 0xbc, 0x61, 0x6d, 0xc6, 0x21, 0x98, 0xd9, 0xf6, 0xce, 0x73, 0x64, 0xac,
	0x4e, 0x6f, 0xfd, 0x31, 0xa9, 0xe6, 0x3e, 0x9c, 0x95, 0x27, 0x3d, 0x8c, 0xbc, 0x2c, 0x7c, 0x39,
	0xef, 0x1c, 0x54, 0x1f, 0x66, 0xb6, 0x07, 0x3b, 0x31, 0x22, 0x94, 0x49, 0x8b, 0x9b, 0xb6, 0xe8,
	0xb6, 0x32, 0x6b, 0x7c, 0x80, 0x5b, 0x49, 0xe9, 0x3b, 0x51, 0x64, 0x7a, 0x71, 0x89, 0xbf, 0x4f,
	0x53, 0xa0, 0x28, 0x36, 0x2b, 0x5f, 0x1d, 0xb6, 0x50, 0xa7, 0x77, 0xa4, 0x3f, 0x59, 0xab, 0x1f,
	0x51, 0xb6, 0x45, 0x57, 0xc9, 0xff, 0x68, 0x47, 0x70, 0x97, 0xec, 0xe9, 0x59, 0x06, 0x87, 0x8f,
	0x50, 0xa6, 0xea, 0x5c, 0x59, 0xef, 0xb1, 0xbf, 0x8e, 0xed, 0xaf, 0xa3, 0x64, 0xbe, 0xe3, 0x13,
	0x6c, 0x0f, 0x91, 0x96, 0x20, 0x8a, 0x76, 0xd3, 0xd5, 0xfb, 0xe6, 0x2f, 0x78, 0xf4, 0xec, 0x5e,
	0xbf, 0xc7, 0xb3, 0xb2, 0x33, 0xbd, 0xa8, 0xa8, 0xe4, 0x1a, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0x7c, 0x2d, 0xc3, 0x14, 0x01, 0x00, 0x00,
}
