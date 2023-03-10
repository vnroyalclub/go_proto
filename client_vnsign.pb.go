// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_vnsign.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 越南每天签到信息
type VnSign struct {
	Id               *int32              `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Reward           []*PropExchangeData `protobuf:"bytes,2,rep,name=Reward" json:"Reward,omitempty"`
	Status           *int32              `protobuf:"varint,3,req,name=Status" json:"Status,omitempty"`
	SignTime         *int64              `protobuf:"varint,4,opt,name=SignTime" json:"SignTime,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *VnSign) Reset()                    { *m = VnSign{} }
func (m *VnSign) String() string            { return proto.CompactTextString(m) }
func (*VnSign) ProtoMessage()               {}
func (*VnSign) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *VnSign) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *VnSign) GetReward() []*PropExchangeData {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *VnSign) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *VnSign) GetSignTime() int64 {
	if m != nil && m.SignTime != nil {
		return *m.SignTime
	}
	return 0
}

// 越南7天签到信息
type VnSignInfo struct {
	SignInfos        []*VnSign `protobuf:"bytes,1,rep,name=SignInfos" json:"SignInfos,omitempty"`
	StartTime        *int64    `protobuf:"varint,2,req,name=StartTime" json:"StartTime,omitempty"`
	StopTime         *int64    `protobuf:"varint,3,req,name=StopTime" json:"StopTime,omitempty"`
	RechargeStatus   *int32    `protobuf:"varint,4,req,name=RechargeStatus" json:"RechargeStatus,omitempty"`
	IsNewPlayer      *int32    `protobuf:"varint,5,req,name=IsNewPlayer" json:"IsNewPlayer,omitempty"`
	GoldSignInfos    []*VnSign `protobuf:"bytes,6,rep,name=GoldSignInfos" json:"GoldSignInfos,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *VnSignInfo) Reset()                    { *m = VnSignInfo{} }
func (m *VnSignInfo) String() string            { return proto.CompactTextString(m) }
func (*VnSignInfo) ProtoMessage()               {}
func (*VnSignInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *VnSignInfo) GetSignInfos() []*VnSign {
	if m != nil {
		return m.SignInfos
	}
	return nil
}

func (m *VnSignInfo) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *VnSignInfo) GetStopTime() int64 {
	if m != nil && m.StopTime != nil {
		return *m.StopTime
	}
	return 0
}

func (m *VnSignInfo) GetRechargeStatus() int32 {
	if m != nil && m.RechargeStatus != nil {
		return *m.RechargeStatus
	}
	return 0
}

func (m *VnSignInfo) GetIsNewPlayer() int32 {
	if m != nil && m.IsNewPlayer != nil {
		return *m.IsNewPlayer
	}
	return 0
}

func (m *VnSignInfo) GetGoldSignInfos() []*VnSign {
	if m != nil {
		return m.GoldSignInfos
	}
	return nil
}

// 越南签到回复
type ReqVnSign struct {
	Day              *int32 `protobuf:"varint,1,req,name=Day" json:"Day,omitempty"`
	Type             *int32 `protobuf:"varint,2,req,name=Type" json:"Type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReqVnSign) Reset()                    { *m = ReqVnSign{} }
func (m *ReqVnSign) String() string            { return proto.CompactTextString(m) }
func (*ReqVnSign) ProtoMessage()               {}
func (*ReqVnSign) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{2} }

func (m *ReqVnSign) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

func (m *ReqVnSign) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

// 越南签到回复
type ReqVnSignRsp struct {
	Sign             *VnSign `protobuf:"bytes,1,req,name=Sign" json:"Sign,omitempty"`
	RechargeStatus   *int32  `protobuf:"varint,2,req,name=RechargeStatus" json:"RechargeStatus,omitempty"`
	IsNewPlayer      *int32  `protobuf:"varint,3,req,name=IsNewPlayer" json:"IsNewPlayer,omitempty"`
	GoldSign         *VnSign `protobuf:"bytes,4,req,name=GoldSign" json:"GoldSign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReqVnSignRsp) Reset()                    { *m = ReqVnSignRsp{} }
func (m *ReqVnSignRsp) String() string            { return proto.CompactTextString(m) }
func (*ReqVnSignRsp) ProtoMessage()               {}
func (*ReqVnSignRsp) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{3} }

func (m *ReqVnSignRsp) GetSign() *VnSign {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *ReqVnSignRsp) GetRechargeStatus() int32 {
	if m != nil && m.RechargeStatus != nil {
		return *m.RechargeStatus
	}
	return 0
}

func (m *ReqVnSignRsp) GetIsNewPlayer() int32 {
	if m != nil && m.IsNewPlayer != nil {
		return *m.IsNewPlayer
	}
	return 0
}

func (m *ReqVnSignRsp) GetGoldSign() *VnSign {
	if m != nil {
		return m.GoldSign
	}
	return nil
}

func init() {
	proto.RegisterType((*VnSign)(nil), "VK.Proto.VnSign")
	proto.RegisterType((*VnSignInfo)(nil), "VK.Proto.VnSignInfo")
	proto.RegisterType((*ReqVnSign)(nil), "VK.Proto.ReqVnSign")
	proto.RegisterType((*ReqVnSignRsp)(nil), "VK.Proto.ReqVnSignRsp")
}

func init() { proto.RegisterFile("client_vnsign.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0x25, 0x33, 0x31, 0xe8, 0xf5, 0x3d, 0x91, 0x2b, 0x3c, 0x82, 0xbc, 0x45, 0x08, 0xa5, 0x64,
	0x51, 0x02, 0x75, 0xd1, 0x5f, 0x60, 0x29, 0xa1, 0x50, 0x64, 0x14, 0xb7, 0x32, 0x24, 0xd3, 0x18,
	0xd0, 0x99, 0x34, 0x99, 0xd6, 0xba, 0xeb, 0x0f, 0xea, 0x0f, 0xec, 0xb2, 0x64, 0x8c, 0x49, 0x3f,
	0x2c, 0x74, 0x13, 0xce, 0xbd, 0xf7, 0xcc, 0x3d, 0xe7, 0xdc, 0xc0, 0x28, 0xde, 0x64, 0x42, 0xea,
	0xd5, 0x93, 0x2c, 0xb3, 0x54, 0x86, 0x79, 0xa1, 0xb4, 0xc2, 0xee, 0xf2, 0x36, 0x9c, 0x55, 0x68,
	0x3c, 0x2a, 0xd7, 0xbc, 0x10, 0xc9, 0x2a, 0x56, 0xdb, 0xad, 0xaa, 0xc7, 0xfe, 0x8b, 0x05, 0xce,
	0x52, 0xce, 0xb3, 0x54, 0xe2, 0x00, 0x48, 0x94, 0xb8, 0x96, 0x47, 0x82, 0x0e, 0x23, 0x51, 0x82,
	0x13, 0x70, 0x98, 0xd8, 0xf1, 0x22, 0x71, 0x89, 0x47, 0x83, 0xfe, 0x64, 0x1c, 0x1e, 0x57, 0x55,
	0xdf, 0xfc, 0xfa, 0x39, 0x5e, 0x73, 0x99, 0x8a, 0x29, 0xd7, 0x9c, 0xd5, 0x4c, 0xfc, 0x07, 0xce,
	0x5c, 0x73, 0xfd, 0x58, 0xba, 0xd4, 0xec, 0xa9, 0x2b, 0x1c, 0x43, 0xb7, 0xd2, 0x58, 0x64, 0x5b,
	0xe1, 0xda, 0x9e, 0x15, 0x50, 0xd6, 0xd4, 0xfe, 0x9b, 0x05, 0x70, 0xb0, 0x10, 0xc9, 0x7b, 0x85,
	0x21, 0xf4, 0x8e, 0xb8, 0x74, 0x2d, 0xa3, 0x3c, 0x6c, 0x95, 0x0f, 0x44, 0xd6, 0x52, 0xf0, 0x3f,
	0xf4, 0xe6, 0x9a, 0x17, 0xda, 0xec, 0x26, 0x1e, 0x09, 0x28, 0x6b, 0x1b, 0x46, 0x58, 0xab, 0xdc,
	0x0c, 0xa9, 0x19, 0x36, 0x35, 0x9e, 0xc3, 0x80, 0x89, 0x78, 0xcd, 0x8b, 0x54, 0xd4, 0xa6, 0x6d,
	0x63, 0xfa, 0x4b, 0x17, 0x3d, 0xe8, 0x47, 0xe5, 0x9d, 0xd8, 0xcd, 0x36, 0x7c, 0x2f, 0x0a, 0xb7,
	0x63, 0x48, 0x1f, 0x5b, 0x78, 0x05, 0x7f, 0x6f, 0xd4, 0x26, 0x69, 0x7d, 0x3b, 0x3f, 0xf8, 0xfe,
	0x4c, 0xf3, 0x2f, 0xa1, 0xc7, 0xc4, 0x43, 0x7d, 0xff, 0x21, 0xd0, 0x29, 0xdf, 0xd7, 0x3f, 0xa0,
	0x82, 0x88, 0x60, 0x2f, 0xf6, 0xf9, 0x21, 0x55, 0x87, 0x19, 0xec, 0xbf, 0x5a, 0xf0, 0xa7, 0x79,
	0xc3, 0xca, 0x1c, 0xcf, 0xc0, 0xae, 0xa0, 0x79, 0x77, 0x4a, 0xd2, 0x4c, 0x4f, 0x64, 0x25, 0xbf,
	0xc9, 0x4a, 0xbf, 0x67, 0xbd, 0x80, 0xee, 0x31, 0x84, 0xb9, 0xd7, 0x29, 0xcd, 0x86, 0xf1, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0xa5, 0x06, 0x9f, 0x68, 0x94, 0x02, 0x00, 0x00,
}
