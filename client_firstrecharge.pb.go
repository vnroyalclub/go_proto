// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_firstrecharge.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 玩家信息
type FRPlayerInfo struct {
	AlreadyRecharge  *int64  `protobuf:"varint,1,req,name=AlreadyRecharge" json:"AlreadyRecharge,omitempty"`
	ReachedTax       *int64  `protobuf:"varint,2,req,name=ReachedTax" json:"ReachedTax,omitempty"`
	TargetTax        *int64  `protobuf:"varint,3,req,name=TargetTax" json:"TargetTax,omitempty"`
	IsJoin           *bool   `protobuf:"varint,4,req,name=IsJoin" json:"IsJoin,omitempty"`
	IsGetReward      *bool   `protobuf:"varint,5,req,name=IsGetReward" json:"IsGetReward,omitempty"`
	WareId           *string `protobuf:"bytes,6,opt,name=WareId" json:"WareId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FRPlayerInfo) Reset()                    { *m = FRPlayerInfo{} }
func (m *FRPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*FRPlayerInfo) ProtoMessage()               {}
func (*FRPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *FRPlayerInfo) GetAlreadyRecharge() int64 {
	if m != nil && m.AlreadyRecharge != nil {
		return *m.AlreadyRecharge
	}
	return 0
}

func (m *FRPlayerInfo) GetReachedTax() int64 {
	if m != nil && m.ReachedTax != nil {
		return *m.ReachedTax
	}
	return 0
}

func (m *FRPlayerInfo) GetTargetTax() int64 {
	if m != nil && m.TargetTax != nil {
		return *m.TargetTax
	}
	return 0
}

func (m *FRPlayerInfo) GetIsJoin() bool {
	if m != nil && m.IsJoin != nil {
		return *m.IsJoin
	}
	return false
}

func (m *FRPlayerInfo) GetIsGetReward() bool {
	if m != nil && m.IsGetReward != nil {
		return *m.IsGetReward
	}
	return false
}

func (m *FRPlayerInfo) GetWareId() string {
	if m != nil && m.WareId != nil {
		return *m.WareId
	}
	return ""
}

// 活动信息
type FRActivityInfo struct {
	RechargeAmount   *int64 `protobuf:"varint,1,req,name=RechargeAmount" json:"RechargeAmount,omitempty"`
	SendAmout        *int64 `protobuf:"varint,2,req,name=SendAmout" json:"SendAmout,omitempty"`
	JoinNumber       *int64 `protobuf:"varint,3,req,name=JoinNumber" json:"JoinNumber,omitempty"`
	StartTime        *int64 `protobuf:"varint,4,req,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64 `protobuf:"varint,5,req,name=EndTime" json:"EndTime,omitempty"`
	OneDayBonus      *int64 `protobuf:"varint,6,req,name=OneDayBonus" json:"OneDayBonus,omitempty"`
	TwoDayBonus      *int64 `protobuf:"varint,7,req,name=TwoDayBonus" json:"TwoDayBonus,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FRActivityInfo) Reset()                    { *m = FRActivityInfo{} }
func (m *FRActivityInfo) String() string            { return proto.CompactTextString(m) }
func (*FRActivityInfo) ProtoMessage()               {}
func (*FRActivityInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *FRActivityInfo) GetRechargeAmount() int64 {
	if m != nil && m.RechargeAmount != nil {
		return *m.RechargeAmount
	}
	return 0
}

func (m *FRActivityInfo) GetSendAmout() int64 {
	if m != nil && m.SendAmout != nil {
		return *m.SendAmout
	}
	return 0
}

func (m *FRActivityInfo) GetJoinNumber() int64 {
	if m != nil && m.JoinNumber != nil {
		return *m.JoinNumber
	}
	return 0
}

func (m *FRActivityInfo) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *FRActivityInfo) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *FRActivityInfo) GetOneDayBonus() int64 {
	if m != nil && m.OneDayBonus != nil {
		return *m.OneDayBonus
	}
	return 0
}

func (m *FRActivityInfo) GetTwoDayBonus() int64 {
	if m != nil && m.TwoDayBonus != nil {
		return *m.TwoDayBonus
	}
	return 0
}

// 首充赠送
type FirstRecharge struct {
	ActivityInfo     *FRActivityInfo `protobuf:"bytes,1,req,name=ActivityInfo" json:"ActivityInfo,omitempty"`
	PlayerInfo       *FRPlayerInfo   `protobuf:"bytes,2,req,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *FirstRecharge) Reset()                    { *m = FirstRecharge{} }
func (m *FirstRecharge) String() string            { return proto.CompactTextString(m) }
func (*FirstRecharge) ProtoMessage()               {}
func (*FirstRecharge) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *FirstRecharge) GetActivityInfo() *FRActivityInfo {
	if m != nil {
		return m.ActivityInfo
	}
	return nil
}

func (m *FirstRecharge) GetPlayerInfo() *FRPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// 首充领取奖励
type FirstRechargeReward struct {
	Reward           *PropExchangeData `protobuf:"bytes,1,req,name=Reward" json:"Reward,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *FirstRechargeReward) Reset()                    { *m = FirstRechargeReward{} }
func (m *FirstRechargeReward) String() string            { return proto.CompactTextString(m) }
func (*FirstRechargeReward) ProtoMessage()               {}
func (*FirstRechargeReward) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *FirstRechargeReward) GetReward() *PropExchangeData {
	if m != nil {
		return m.Reward
	}
	return nil
}

type FRGiftPackage struct {
	RechargeAmount   *int64  `protobuf:"varint,1,req,name=RechargeAmount" json:"RechargeAmount,omitempty"`
	SendAmout        *int64  `protobuf:"varint,2,req,name=SendAmout" json:"SendAmout,omitempty"`
	OneDayBonus      *int64  `protobuf:"varint,3,req,name=OneDayBonus" json:"OneDayBonus,omitempty"`
	TwoDayBonus      *int64  `protobuf:"varint,4,req,name=TwoDayBonus" json:"TwoDayBonus,omitempty"`
	WareId           *string `protobuf:"bytes,5,req,name=WareId" json:"WareId,omitempty"`
	Ratio            *int32  `protobuf:"varint,6,req,name=Ratio" json:"Ratio,omitempty"`
	Lock             *int32  `protobuf:"varint,7,req,name=Lock" json:"Lock,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FRGiftPackage) Reset()                    { *m = FRGiftPackage{} }
func (m *FRGiftPackage) String() string            { return proto.CompactTextString(m) }
func (*FRGiftPackage) ProtoMessage()               {}
func (*FRGiftPackage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *FRGiftPackage) GetRechargeAmount() int64 {
	if m != nil && m.RechargeAmount != nil {
		return *m.RechargeAmount
	}
	return 0
}

func (m *FRGiftPackage) GetSendAmout() int64 {
	if m != nil && m.SendAmout != nil {
		return *m.SendAmout
	}
	return 0
}

func (m *FRGiftPackage) GetOneDayBonus() int64 {
	if m != nil && m.OneDayBonus != nil {
		return *m.OneDayBonus
	}
	return 0
}

func (m *FRGiftPackage) GetTwoDayBonus() int64 {
	if m != nil && m.TwoDayBonus != nil {
		return *m.TwoDayBonus
	}
	return 0
}

func (m *FRGiftPackage) GetWareId() string {
	if m != nil && m.WareId != nil {
		return *m.WareId
	}
	return ""
}

func (m *FRGiftPackage) GetRatio() int32 {
	if m != nil && m.Ratio != nil {
		return *m.Ratio
	}
	return 0
}

func (m *FRGiftPackage) GetLock() int32 {
	if m != nil && m.Lock != nil {
		return *m.Lock
	}
	return 0
}

type FirstRechargeNew struct {
	StartTime        *int64           `protobuf:"varint,1,req,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64           `protobuf:"varint,2,req,name=EndTime" json:"EndTime,omitempty"`
	PlayerInfo       *FRPlayerInfo    `protobuf:"bytes,3,req,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	GiftPacakgeList  []*FRGiftPackage `protobuf:"bytes,4,rep,name=GiftPacakgeList" json:"GiftPacakgeList,omitempty"`
	Ratio            *int32           `protobuf:"varint,5,req,name=Ratio" json:"Ratio,omitempty"`
	Lock             *int32           `protobuf:"varint,6,req,name=Lock" json:"Lock,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *FirstRechargeNew) Reset()                    { *m = FirstRechargeNew{} }
func (m *FirstRechargeNew) String() string            { return proto.CompactTextString(m) }
func (*FirstRechargeNew) ProtoMessage()               {}
func (*FirstRechargeNew) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *FirstRechargeNew) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *FirstRechargeNew) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *FirstRechargeNew) GetPlayerInfo() *FRPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *FirstRechargeNew) GetGiftPacakgeList() []*FRGiftPackage {
	if m != nil {
		return m.GiftPacakgeList
	}
	return nil
}

func (m *FirstRechargeNew) GetRatio() int32 {
	if m != nil && m.Ratio != nil {
		return *m.Ratio
	}
	return 0
}

func (m *FirstRechargeNew) GetLock() int32 {
	if m != nil && m.Lock != nil {
		return *m.Lock
	}
	return 0
}

func init() {
	proto.RegisterType((*FRPlayerInfo)(nil), "VK.Proto.FRPlayerInfo")
	proto.RegisterType((*FRActivityInfo)(nil), "VK.Proto.FRActivityInfo")
	proto.RegisterType((*FirstRecharge)(nil), "VK.Proto.FirstRecharge")
	proto.RegisterType((*FirstRechargeReward)(nil), "VK.Proto.FirstRechargeReward")
	proto.RegisterType((*FRGiftPackage)(nil), "VK.Proto.FRGiftPackage")
	proto.RegisterType((*FirstRechargeNew)(nil), "VK.Proto.FirstRechargeNew")
}

func init() { proto.RegisterFile("client_firstrecharge.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0xd8, 0x6d, 0x27, 0xa5, 0x45, 0x5b, 0x54, 0xac, 0x0a, 0x21, 0xcb, 0x07, 0xe4,
	0x53, 0x0e, 0x39, 0x70, 0xe2, 0x12, 0xd4, 0xa6, 0x32, 0x54, 0x25, 0x5a, 0x22, 0x38, 0x56, 0x8b,
	0x3d, 0x49, 0xac, 0xc4, 0xbb, 0xd5, 0x7a, 0x43, 0x9a, 0x0f, 0xe0, 0xdf, 0xf8, 0x07, 0x3e, 0x02,
	0x89, 0x2f, 0x40, 0xbb, 0xd9, 0xc4, 0xeb, 0x80, 0x80, 0x03, 0xb7, 0x9d, 0x37, 0x6f, 0x56, 0xf3,
	0xde, 0x1b, 0xb8, 0xc8, 0x16, 0x05, 0x72, 0x75, 0x37, 0x29, 0x64, 0xa5, 0x24, 0x66, 0x33, 0x26,
	0xa7, 0xd8, 0xbb, 0x97, 0x42, 0x09, 0x72, 0xf8, 0xe1, 0x6d, 0x6f, 0xa4, 0x5f, 0x17, 0x67, 0xd5,
	0x8c, 0x49, 0xcc, 0xef, 0x32, 0x51, 0x96, 0x82, 0x6f, 0xda, 0xf1, 0x57, 0x0f, 0x8e, 0x87, 0x74,
	0xb4, 0x60, 0x6b, 0x94, 0x29, 0x9f, 0x08, 0x92, 0xc0, 0xe9, 0x60, 0x21, 0x91, 0xe5, 0x6b, 0x6a,
	0x3f, 0x0a, 0xbd, 0xa8, 0x95, 0xb4, 0xe9, 0x3e, 0x4c, 0x9e, 0x03, 0x50, 0x64, 0xd9, 0x0c, 0xf3,
	0x31, 0x7b, 0x08, 0x5b, 0x86, 0xe4, 0x20, 0xe4, 0x19, 0x1c, 0x8d, 0x35, 0x51, 0xe9, 0x76, 0xdb,
	0xb4, 0x6b, 0x80, 0x9c, 0x43, 0x90, 0x56, 0x6f, 0x44, 0xc1, 0xc3, 0x4e, 0xd4, 0x4a, 0x0e, 0xa9,
	0xad, 0x48, 0x04, 0xdd, 0xb4, 0xba, 0x46, 0x45, 0x71, 0xc5, 0x64, 0x1e, 0xfa, 0xa6, 0xe9, 0x42,
	0x7a, 0xf2, 0x23, 0x93, 0x98, 0xe6, 0x61, 0x10, 0x79, 0xc9, 0x11, 0xb5, 0x55, 0xfc, 0xc3, 0x83,
	0x93, 0x21, 0x1d, 0x64, 0xaa, 0xf8, 0x5c, 0xa8, 0xb5, 0x11, 0xf3, 0x02, 0x4e, 0xb6, 0xeb, 0x0e,
	0x4a, 0xb1, 0xe4, 0xca, 0x6a, 0xd9, 0x43, 0xf5, 0xaa, 0xef, 0x91, 0xe7, 0xba, 0x52, 0x56, 0x49,
	0x0d, 0x68, 0xa1, 0x7a, 0xb5, 0xdb, 0x65, 0xf9, 0x09, 0xa5, 0x55, 0xe2, 0x20, 0x66, 0x5a, 0x31,
	0xa9, 0xc6, 0x45, 0x89, 0x46, 0x8d, 0x9e, 0xde, 0x02, 0x24, 0x84, 0x83, 0x2b, 0x9e, 0x9b, 0x9e,
	0x6f, 0x7a, 0xdb, 0x52, 0x4b, 0x7d, 0xc7, 0xf1, 0x92, 0xad, 0x5f, 0x0b, 0xbe, 0xac, 0xc2, 0xc0,
	0x74, 0x5d, 0x48, 0x33, 0xc6, 0x2b, 0xb1, 0x63, 0x1c, 0x6c, 0x18, 0x0e, 0x14, 0x7f, 0xf1, 0xe0,
	0xd1, 0x50, 0xc7, 0xbe, 0x8b, 0xe5, 0x15, 0x1c, 0xbb, 0x1e, 0x18, 0xc5, 0xdd, 0x7e, 0xd8, 0xdb,
	0xde, 0x41, 0xaf, 0xe9, 0x11, 0x6d, 0xb0, 0xc9, 0x4b, 0x80, 0xfa, 0x18, 0x8c, 0x15, 0xdd, 0xfe,
	0xb9, 0x3b, 0x5b, 0x77, 0xa9, 0xc3, 0x8c, 0x53, 0x38, 0x6b, 0xac, 0x61, 0xb3, 0xea, 0x43, 0x60,
	0x83, 0xdc, 0xac, 0x71, 0x51, 0x7f, 0x35, 0x92, 0xe2, 0xfe, 0xea, 0x21, 0x9b, 0x31, 0x3e, 0xc5,
	0x4b, 0xa6, 0x18, 0xb5, 0xcc, 0xf8, 0x9b, 0x96, 0x44, 0xaf, 0x8b, 0x89, 0x1a, 0xb1, 0x6c, 0xce,
	0xa6, 0xf8, 0x9f, 0x62, 0xdc, 0xb3, 0xbb, 0xfd, 0x57, 0xbb, 0x3b, 0xbf, 0xd8, 0xed, 0xdc, 0x9e,
	0xce, 0x72, 0x77, 0x7b, 0xe4, 0x09, 0xf8, 0x94, 0xa9, 0x42, 0x98, 0x10, 0x7d, 0xba, 0x29, 0x08,
	0x81, 0xce, 0x8d, 0xc8, 0xe6, 0x26, 0x37, 0x9f, 0x9a, 0x77, 0xfc, 0xdd, 0x83, 0xc7, 0x0d, 0xa7,
	0x6e, 0x71, 0xd5, 0xbc, 0x20, 0xef, 0x0f, 0x17, 0xd4, 0x6a, 0x5e, 0x50, 0x33, 0xad, 0xf6, 0xbf,
	0xa6, 0x45, 0x06, 0x70, 0x6a, 0xfd, 0x65, 0xf3, 0x29, 0xde, 0x14, 0x95, 0x0a, 0x3b, 0x51, 0x3b,
	0xe9, 0xf6, 0x9f, 0xba, 0xc3, 0x4e, 0x04, 0x74, 0x9f, 0x5f, 0x2b, 0xf6, 0x7f, 0xa7, 0x38, 0xa8,
	0x15, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xa0, 0xae, 0xd7, 0x9e, 0x04, 0x00, 0x00,
}
