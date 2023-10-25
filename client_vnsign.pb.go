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
func (*VnSign) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

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
func (*VnSignInfo) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

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
func (*ReqVnSign) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{2} }

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
	Sign             *VnSign `protobuf:"bytes,1,opt,name=Sign" json:"Sign,omitempty"`
	RechargeStatus   *int32  `protobuf:"varint,2,req,name=RechargeStatus" json:"RechargeStatus,omitempty"`
	IsNewPlayer      *int32  `protobuf:"varint,3,req,name=IsNewPlayer" json:"IsNewPlayer,omitempty"`
	GoldSign         *VnSign `protobuf:"bytes,4,opt,name=GoldSign" json:"GoldSign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReqVnSignRsp) Reset()                    { *m = ReqVnSignRsp{} }
func (m *ReqVnSignRsp) String() string            { return proto.CompactTextString(m) }
func (*ReqVnSignRsp) ProtoMessage()               {}
func (*ReqVnSignRsp) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{3} }

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

func init() { proto.RegisterFile("client_vnsign.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x6b, 0xeb, 0x30,
	0x10, 0x44, 0x92, 0x63, 0x92, 0xcd, 0x7b, 0x21, 0x28, 0xf0, 0x10, 0xe1, 0x1d, 0x8c, 0x29, 0xc5,
	0x87, 0x62, 0x68, 0x0e, 0xfd, 0x05, 0x29, 0xc5, 0x14, 0x4a, 0x50, 0x42, 0xae, 0x41, 0xd8, 0xaa,
	0x63, 0x48, 0x24, 0xd7, 0x56, 0x9b, 0xe6, 0xd6, 0x1f, 0xd4, 0x1f, 0xd8, 0x63, 0xb1, 0xe2, 0xd8,
	0xfd, 0x48, 0xa1, 0x17, 0x33, 0xbb, 0x3b, 0xde, 0x99, 0xd1, 0xc2, 0x28, 0xde, 0x64, 0x52, 0x99,
	0xd5, 0x93, 0x2a, 0xb3, 0x54, 0x85, 0x79, 0xa1, 0x8d, 0xa6, 0xdd, 0xe5, 0x6d, 0x38, 0xab, 0xd0,
	0x78, 0x54, 0xae, 0x45, 0x21, 0x93, 0x55, 0xac, 0xb7, 0x5b, 0x5d, 0x8f, 0xfd, 0x17, 0x04, 0xee,
	0x52, 0xcd, 0xb3, 0x54, 0xd1, 0x01, 0xe0, 0x28, 0x61, 0xc8, 0xc3, 0x41, 0x87, 0xe3, 0x28, 0xa1,
	0x13, 0x70, 0xb9, 0xdc, 0x89, 0x22, 0x61, 0xd8, 0x23, 0x41, 0x7f, 0x32, 0x0e, 0x8f, 0xab, 0xaa,
	0x6f, 0x7e, 0xfd, 0x1c, 0xaf, 0x85, 0x4a, 0xe5, 0x54, 0x18, 0xc1, 0x6b, 0x26, 0xfd, 0x07, 0xee,
	0xdc, 0x08, 0xf3, 0x58, 0x32, 0x62, 0xf7, 0xd4, 0x15, 0x1d, 0x43, 0xb7, 0xd2, 0x58, 0x64, 0x5b,
	0xc9, 0x1c, 0x0f, 0x05, 0x84, 0x37, 0xb5, 0xff, 0x86, 0x00, 0x0e, 0x16, 0x22, 0x75, 0xaf, 0x69,
	0x08, 0xbd, 0x23, 0x2e, 0x19, 0xb2, 0xca, 0xc3, 0x56, 0xf9, 0x40, 0xe4, 0x2d, 0x85, 0xfe, 0x87,
	0xde, 0xdc, 0x88, 0xc2, 0xd8, 0xdd, 0xd8, 0xc3, 0x01, 0xe1, 0x6d, 0xc3, 0x0a, 0x1b, 0x9d, 0xdb,
	0x21, 0xb1, 0xc3, 0xa6, 0xa6, 0xe7, 0x30, 0xe0, 0x32, 0x5e, 0x8b, 0x22, 0x95, 0xb5, 0x69, 0xc7,
	0x9a, 0xfe, 0xd2, 0xa5, 0x1e, 0xf4, 0xa3, 0xf2, 0x4e, 0xee, 0x66, 0x1b, 0xb1, 0x97, 0x05, 0xeb,
	0x58, 0xd2, 0xc7, 0x16, 0xbd, 0x82, 0xbf, 0x37, 0x7a, 0x93, 0xb4, 0xbe, 0xdd, 0x1f, 0x7c, 0x7f,
	0xa6, 0xf9, 0x97, 0xd0, 0xe3, 0xf2, 0xa1, 0x7e, 0xff, 0x21, 0x90, 0xa9, 0xd8, 0xd7, 0x07, 0xa8,
	0x20, 0xa5, 0xe0, 0x2c, 0xf6, 0xf9, 0x21, 0x55, 0x87, 0x5b, 0xec, 0xbf, 0x22, 0xf8, 0xd3, 0xfc,
	0xc3, 0xcb, 0x9c, 0x9e, 0x81, 0x53, 0x41, 0x86, 0x3c, 0x74, 0x52, 0xd2, 0x4e, 0x4f, 0x64, 0xc5,
	0xbf, 0xc9, 0x4a, 0xbe, 0x67, 0xbd, 0x80, 0xee, 0x31, 0x84, 0x3d, 0xe5, 0x29, 0xcd, 0x86, 0xf1,
	0x1e, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x19, 0xc3, 0xdf, 0x94, 0x02, 0x00, 0x00,
}
