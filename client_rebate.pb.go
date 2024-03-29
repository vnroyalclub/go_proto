// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_rebate.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 越南每天返利信息
type Rebate struct {
	Id               *int32            `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Ratio            *int32            `protobuf:"varint,2,req,name=Ratio" json:"Ratio,omitempty"`
	Reward           *PropExchangeData `protobuf:"bytes,3,req,name=Reward" json:"Reward,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Rebate) Reset()                    { *m = Rebate{} }
func (m *Rebate) String() string            { return proto.CompactTextString(m) }
func (*Rebate) ProtoMessage()               {}
func (*Rebate) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Rebate) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Rebate) GetRatio() int32 {
	if m != nil && m.Ratio != nil {
		return *m.Ratio
	}
	return 0
}

func (m *Rebate) GetReward() *PropExchangeData {
	if m != nil {
		return m.Reward
	}
	return nil
}

// 越南每天返利信息
type FutureRebate struct {
	RebateInfo       []*Rebate         `protobuf:"bytes,1,rep,name=RebateInfo" json:"RebateInfo,omitempty"`
	TotalReward      *PropExchangeData `protobuf:"bytes,2,req,name=TotalReward" json:"TotalReward,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *FutureRebate) Reset()                    { *m = FutureRebate{} }
func (m *FutureRebate) String() string            { return proto.CompactTextString(m) }
func (*FutureRebate) ProtoMessage()               {}
func (*FutureRebate) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *FutureRebate) GetRebateInfo() []*Rebate {
	if m != nil {
		return m.RebateInfo
	}
	return nil
}

func (m *FutureRebate) GetTotalReward() *PropExchangeData {
	if m != nil {
		return m.TotalReward
	}
	return nil
}

// 越南3天返利信息
type RebateInfo struct {
	Rebates          []*Rebate         `protobuf:"bytes,1,rep,name=Rebates" json:"Rebates,omitempty"`
	StartTime        *int64            `protobuf:"varint,2,req,name=StartTime" json:"StartTime,omitempty"`
	StopTime         *int64            `protobuf:"varint,3,req,name=StopTime" json:"StopTime,omitempty"`
	Status           *int32            `protobuf:"varint,4,req,name=Status" json:"Status,omitempty"`
	Reward           *PropExchangeData `protobuf:"bytes,5,req,name=Reward" json:"Reward,omitempty"`
	ActualReward     *PropExchangeData `protobuf:"bytes,6,req,name=ActualReward" json:"ActualReward,omitempty"`
	FutureReward     *FutureRebate     `protobuf:"bytes,7,req,name=FutureReward" json:"FutureReward,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RebateInfo) Reset()                    { *m = RebateInfo{} }
func (m *RebateInfo) String() string            { return proto.CompactTextString(m) }
func (*RebateInfo) ProtoMessage()               {}
func (*RebateInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *RebateInfo) GetRebates() []*Rebate {
	if m != nil {
		return m.Rebates
	}
	return nil
}

func (m *RebateInfo) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *RebateInfo) GetStopTime() int64 {
	if m != nil && m.StopTime != nil {
		return *m.StopTime
	}
	return 0
}

func (m *RebateInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *RebateInfo) GetReward() *PropExchangeData {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *RebateInfo) GetActualReward() *PropExchangeData {
	if m != nil {
		return m.ActualReward
	}
	return nil
}

func (m *RebateInfo) GetFutureReward() *FutureRebate {
	if m != nil {
		return m.FutureReward
	}
	return nil
}

// 越南返利领取回复
type ReqRebateRsp struct {
	Reward           *PropExchangeData `protobuf:"bytes,1,req,name=Reward" json:"Reward,omitempty"`
	ActualReward     *PropExchangeData `protobuf:"bytes,2,req,name=ActualReward" json:"ActualReward,omitempty"`
	Status           *int32            `protobuf:"varint,3,req,name=Status" json:"Status,omitempty"`
	CanReward        *PropExchangeData `protobuf:"bytes,4,req,name=CanReward" json:"CanReward,omitempty"`
	FutureReward     *FutureRebate     `protobuf:"bytes,5,req,name=FutureReward" json:"FutureReward,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ReqRebateRsp) Reset()                    { *m = ReqRebateRsp{} }
func (m *ReqRebateRsp) String() string            { return proto.CompactTextString(m) }
func (*ReqRebateRsp) ProtoMessage()               {}
func (*ReqRebateRsp) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *ReqRebateRsp) GetReward() *PropExchangeData {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *ReqRebateRsp) GetActualReward() *PropExchangeData {
	if m != nil {
		return m.ActualReward
	}
	return nil
}

func (m *ReqRebateRsp) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *ReqRebateRsp) GetCanReward() *PropExchangeData {
	if m != nil {
		return m.CanReward
	}
	return nil
}

func (m *ReqRebateRsp) GetFutureReward() *FutureRebate {
	if m != nil {
		return m.FutureReward
	}
	return nil
}

func init() {
	proto.RegisterType((*Rebate)(nil), "VK.Proto.Rebate")
	proto.RegisterType((*FutureRebate)(nil), "VK.Proto.FutureRebate")
	proto.RegisterType((*RebateInfo)(nil), "VK.Proto.RebateInfo")
	proto.RegisterType((*ReqRebateRsp)(nil), "VK.Proto.ReqRebateRsp")
}

func init() { proto.RegisterFile("client_rebate.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0x69, 0xb6, 0xdb, 0x97, 0x69, 0xf9, 0xf3, 0x27, 0x95, 0xb2, 0x14, 0x0f, 0x65, 0x4f,
	0xc5, 0x43, 0x91, 0x9e, 0x44, 0x44, 0x10, 0x5f, 0xa0, 0x78, 0x91, 0xb4, 0x78, 0x2d, 0xe9, 0x6e,
	0xb4, 0x85, 0x76, 0xb3, 0x66, 0x67, 0xd1, 0x93, 0x9f, 0xc1, 0xcf, 0xe3, 0xa7, 0x93, 0x66, 0xd3,
	0x26, 0x45, 0x64, 0xab, 0x97, 0x90, 0xc9, 0xf3, 0x9b, 0x79, 0x98, 0x27, 0xd0, 0x89, 0x56, 0x4b,
	0x91, 0xe0, 0x4c, 0x89, 0x39, 0x47, 0x31, 0x4c, 0x95, 0x44, 0x49, 0x1b, 0x8f, 0xf7, 0xc3, 0x87,
	0xcd, 0xad, 0xd7, 0xc9, 0x16, 0x5c, 0x89, 0x78, 0x16, 0xc9, 0xf5, 0x5a, 0x26, 0x85, 0x1c, 0xce,
	0xa1, 0xc6, 0x34, 0x4e, 0xff, 0x01, 0x19, 0xc7, 0x41, 0xa5, 0x4f, 0x06, 0x3e, 0x23, 0xe3, 0x98,
	0x1e, 0x81, 0xcf, 0x38, 0x2e, 0x65, 0x40, 0xf4, 0x53, 0x51, 0xd0, 0xd1, 0x86, 0x7f, 0xe5, 0x2a,
	0x0e, 0xbc, 0x3e, 0x19, 0xb4, 0x46, 0xbd, 0xe1, 0x76, 0xfe, 0xe6, 0x4c, 0x6f, 0xdf, 0xa2, 0x05,
	0x4f, 0x9e, 0xc5, 0x0d, 0x47, 0xce, 0x0c, 0x19, 0xbe, 0x43, 0xfb, 0x2e, 0xc7, 0x5c, 0x09, 0xe3,
	0x74, 0x0a, 0x50, 0xdc, 0xc6, 0xc9, 0x93, 0x0c, 0x2a, 0x7d, 0x6f, 0xd0, 0x1a, 0xfd, 0xb7, 0x73,
	0x0a, 0x8d, 0x39, 0x0c, 0xbd, 0x80, 0xd6, 0x54, 0x22, 0x5f, 0x19, 0x6b, 0x52, 0x6a, 0xed, 0xe2,
	0xe1, 0x27, 0x71, 0x0d, 0xe9, 0x09, 0xd4, 0x8b, 0x2a, 0xfb, 0xd1, 0x7b, 0x0b, 0xd0, 0x63, 0x68,
	0x4e, 0x90, 0x2b, 0x9c, 0x2e, 0xd7, 0x42, 0xdb, 0x7a, 0xcc, 0x3e, 0xd0, 0x1e, 0x34, 0x26, 0x28,
	0x53, 0x2d, 0x7a, 0x5a, 0xdc, 0xd5, 0xb4, 0x0b, 0xb5, 0x09, 0x72, 0xcc, 0xb3, 0xa0, 0xaa, 0xf3,
	0x33, 0x95, 0x13, 0xa0, 0x7f, 0x68, 0x80, 0xf4, 0x12, 0xda, 0x57, 0x11, 0xe6, 0xbb, 0xfd, 0x6b,
	0xa5, 0x9d, 0x7b, 0x3c, 0x3d, 0xb7, 0x1f, 0xa0, 0xfb, 0xeb, 0xba, 0xbf, 0x6b, 0xfb, 0xdd, 0xef,
	0x61, 0x7b, 0x6c, 0xf8, 0x41, 0xa0, 0xcd, 0xc4, 0x8b, 0xd1, 0xb2, 0xd4, 0x59, 0xa0, 0xf2, 0xe7,
	0x05, 0xc8, 0x2f, 0x17, 0xb0, 0x61, 0x7a, 0x7b, 0x61, 0x9e, 0x41, 0xf3, 0x9a, 0x27, 0x66, 0x68,
	0xb5, 0x74, 0xa8, 0x85, 0xbf, 0x45, 0xe2, 0x1f, 0x1e, 0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x92, 0x94, 0x43, 0xcf, 0x68, 0x03, 0x00, 0x00,
}
