// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_msign.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏服务请求
type MsignOps int32

const (
	// 询问是否可以签到
	MsignOps_Req_AskSign MsignOps = 1
	// 签到
	MsignOps_Req_Sign MsignOps = 2
	// 询问是否可以补签到
	MsignOps_Req_AskReplenish MsignOps = 3
	// 补签
	MsignOps_Req_Replenish MsignOps = 4
	// 询问是否可以开启宝箱
	MsignOps_Req_AskBox MsignOps = 5
	// 开启次日宝箱
	MsignOps_Req_OpenBoxDay MsignOps = 6
	// 开启七日宝箱
	MsignOps_Req_OpenBoxWeek MsignOps = 7
	// 开启15日宝箱
	MsignOps_Req_OpenBoxHalfMonth MsignOps = 8
	// 开启30日宝箱
	MsignOps_Req_OpenBoxMonth MsignOps = 9
	// 查询滚动弹幕
	MsignOps_Req_AskRoll MsignOps = 10
	// 查询排行榜
	MsignOps_Req_AskRank MsignOps = 11
)

var MsignOps_name = map[int32]string{
	1:  "Req_AskSign",
	2:  "Req_Sign",
	3:  "Req_AskReplenish",
	4:  "Req_Replenish",
	5:  "Req_AskBox",
	6:  "Req_OpenBoxDay",
	7:  "Req_OpenBoxWeek",
	8:  "Req_OpenBoxHalfMonth",
	9:  "Req_OpenBoxMonth",
	10: "Req_AskRoll",
	11: "Req_AskRank",
}
var MsignOps_value = map[string]int32{
	"Req_AskSign":          1,
	"Req_Sign":             2,
	"Req_AskReplenish":     3,
	"Req_Replenish":        4,
	"Req_AskBox":           5,
	"Req_OpenBoxDay":       6,
	"Req_OpenBoxWeek":      7,
	"Req_OpenBoxHalfMonth": 8,
	"Req_OpenBoxMonth":     9,
	"Req_AskRoll":          10,
	"Req_AskRank":          11,
}

func (x MsignOps) Enum() *MsignOps {
	p := new(MsignOps)
	*p = x
	return p
}
func (x MsignOps) String() string {
	return proto.EnumName(MsignOps_name, int32(x))
}
func (x *MsignOps) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsignOps_value, data, "MsignOps")
	if err != nil {
		return err
	}
	*x = MsignOps(value)
	return nil
}
func (MsignOps) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

// 是否签到标识
type MsMark int32

const (
	MsMark_Sign      MsMark = 1
	MsMark_Replenish MsMark = 2
	MsMark_Delay     MsMark = 3
)

var MsMark_name = map[int32]string{
	1: "Sign",
	2: "Replenish",
	3: "Delay",
}
var MsMark_value = map[string]int32{
	"Sign":      1,
	"Replenish": 2,
	"Delay":     3,
}

func (x MsMark) Enum() *MsMark {
	p := new(MsMark)
	*p = x
	return p
}
func (x MsMark) String() string {
	return proto.EnumName(MsMark_name, int32(x))
}
func (x *MsMark) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsMark_value, data, "MsMark")
	if err != nil {
		return err
	}
	*x = MsMark(value)
	return nil
}
func (MsMark) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

// 是否签到标识
type MsBoxType int32

const (
	MsBoxType_BT_Day       MsBoxType = 1
	MsBoxType_BT_Week      MsBoxType = 2
	MsBoxType_BT_HalfMonth MsBoxType = 3
	MsBoxType_BT_Month     MsBoxType = 4
)

var MsBoxType_name = map[int32]string{
	1: "BT_Day",
	2: "BT_Week",
	3: "BT_HalfMonth",
	4: "BT_Month",
}
var MsBoxType_value = map[string]int32{
	"BT_Day":       1,
	"BT_Week":      2,
	"BT_HalfMonth": 3,
	"BT_Month":     4,
}

func (x MsBoxType) Enum() *MsBoxType {
	p := new(MsBoxType)
	*p = x
	return p
}
func (x MsBoxType) String() string {
	return proto.EnumName(MsBoxType_name, int32(x))
}
func (x *MsBoxType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsBoxType_value, data, "MsBoxType")
	if err != nil {
		return err
	}
	*x = MsBoxType(value)
	return nil
}
func (MsBoxType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

// 查询签到返回
type MsAskSignResp struct {
	IsSign           *bool  `protobuf:"varint,1,req,name=IsSign" json:"IsSign,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsAskSignResp) Reset()                    { *m = MsAskSignResp{} }
func (m *MsAskSignResp) String() string            { return proto.CompactTextString(m) }
func (*MsAskSignResp) ProtoMessage()               {}
func (*MsAskSignResp) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *MsAskSignResp) GetIsSign() bool {
	if m != nil && m.IsSign != nil {
		return *m.IsSign
	}
	return false
}

// 签到返回
type MsSignResp struct {
	ConfigId         *int32 `protobuf:"varint,1,req,name=ConfigId" json:"ConfigId,omitempty"`
	Count            *int64 `protobuf:"varint,2,req,name=Count" json:"Count,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsSignResp) Reset()                    { *m = MsSignResp{} }
func (m *MsSignResp) String() string            { return proto.CompactTextString(m) }
func (*MsSignResp) ProtoMessage()               {}
func (*MsSignResp) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *MsSignResp) GetConfigId() int32 {
	if m != nil && m.ConfigId != nil {
		return *m.ConfigId
	}
	return 0
}

func (m *MsSignResp) GetCount() int64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

// 补签查询返回
type MsAskReplenishResp struct {
	Year             *int32  `protobuf:"varint,1,req,name=Year" json:"Year,omitempty"`
	Month            *int32  `protobuf:"varint,2,req,name=Month" json:"Month,omitempty"`
	Day              *int32  `protobuf:"varint,3,req,name=Day" json:"Day,omitempty"`
	Date             []int32 `protobuf:"varint,4,rep,name=Date" json:"Date,omitempty"`
	RemainingTimes   *int32  `protobuf:"varint,5,req,name=RemainingTimes" json:"RemainingTimes,omitempty"`
	Expend           *int32  `protobuf:"varint,6,req,name=Expend" json:"Expend,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsAskReplenishResp) Reset()                    { *m = MsAskReplenishResp{} }
func (m *MsAskReplenishResp) String() string            { return proto.CompactTextString(m) }
func (*MsAskReplenishResp) ProtoMessage()               {}
func (*MsAskReplenishResp) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *MsAskReplenishResp) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *MsAskReplenishResp) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *MsAskReplenishResp) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

func (m *MsAskReplenishResp) GetDate() []int32 {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *MsAskReplenishResp) GetRemainingTimes() int32 {
	if m != nil && m.RemainingTimes != nil {
		return *m.RemainingTimes
	}
	return 0
}

func (m *MsAskReplenishResp) GetExpend() int32 {
	if m != nil && m.Expend != nil {
		return *m.Expend
	}
	return 0
}

// 补签请求
type MsReplenishReq struct {
	Month            *int32 `protobuf:"varint,1,req,name=Month" json:"Month,omitempty"`
	Day              *int32 `protobuf:"varint,2,req,name=Day" json:"Day,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsReplenishReq) Reset()                    { *m = MsReplenishReq{} }
func (m *MsReplenishReq) String() string            { return proto.CompactTextString(m) }
func (*MsReplenishReq) ProtoMessage()               {}
func (*MsReplenishReq) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *MsReplenishReq) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *MsReplenishReq) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

type MsBoxContent struct {
	EntityId         *int32 `protobuf:"varint,1,req,name=EntityId" json:"EntityId,omitempty"`
	Value            *int64 `protobuf:"varint,2,req,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsBoxContent) Reset()                    { *m = MsBoxContent{} }
func (m *MsBoxContent) String() string            { return proto.CompactTextString(m) }
func (*MsBoxContent) ProtoMessage()               {}
func (*MsBoxContent) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *MsBoxContent) GetEntityId() int32 {
	if m != nil && m.EntityId != nil {
		return *m.EntityId
	}
	return 0
}

func (m *MsBoxContent) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type MsBox struct {
	SignTimes        *int32        `protobuf:"varint,1,req,name=SignTimes" json:"SignTimes,omitempty"`
	NeedTimes        *int32        `protobuf:"varint,2,req,name=NeedTimes" json:"NeedTimes,omitempty"`
	IsOpen           *bool         `protobuf:"varint,3,req,name=IsOpen" json:"IsOpen,omitempty"`
	CanOpen          *bool         `protobuf:"varint,4,req,name=CanOpen" json:"CanOpen,omitempty"`
	NextOpenTime     *int64        `protobuf:"varint,5,req,name=NextOpenTime" json:"NextOpenTime,omitempty"`
	BoxType          *int32        `protobuf:"varint,6,req,name=BoxType" json:"BoxType,omitempty"`
	Content          *MsBoxContent `protobuf:"bytes,7,req,name=Content" json:"Content,omitempty"`
	Expire           *bool         `protobuf:"varint,8,req,name=Expire" json:"Expire,omitempty"`
	NextAtivityTime  *int64        `protobuf:"varint,9,req,name=NextAtivityTime" json:"NextAtivityTime,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *MsBox) Reset()                    { *m = MsBox{} }
func (m *MsBox) String() string            { return proto.CompactTextString(m) }
func (*MsBox) ProtoMessage()               {}
func (*MsBox) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *MsBox) GetSignTimes() int32 {
	if m != nil && m.SignTimes != nil {
		return *m.SignTimes
	}
	return 0
}

func (m *MsBox) GetNeedTimes() int32 {
	if m != nil && m.NeedTimes != nil {
		return *m.NeedTimes
	}
	return 0
}

func (m *MsBox) GetIsOpen() bool {
	if m != nil && m.IsOpen != nil {
		return *m.IsOpen
	}
	return false
}

func (m *MsBox) GetCanOpen() bool {
	if m != nil && m.CanOpen != nil {
		return *m.CanOpen
	}
	return false
}

func (m *MsBox) GetNextOpenTime() int64 {
	if m != nil && m.NextOpenTime != nil {
		return *m.NextOpenTime
	}
	return 0
}

func (m *MsBox) GetBoxType() int32 {
	if m != nil && m.BoxType != nil {
		return *m.BoxType
	}
	return 0
}

func (m *MsBox) GetContent() *MsBoxContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MsBox) GetExpire() bool {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return false
}

func (m *MsBox) GetNextAtivityTime() int64 {
	if m != nil && m.NextAtivityTime != nil {
		return *m.NextAtivityTime
	}
	return 0
}

type MsAskBoxResp struct {
	Value            []*MsBox `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsAskBoxResp) Reset()                    { *m = MsAskBoxResp{} }
func (m *MsAskBoxResp) String() string            { return proto.CompactTextString(m) }
func (*MsAskBoxResp) ProtoMessage()               {}
func (*MsAskBoxResp) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *MsAskBoxResp) GetValue() []*MsBox {
	if m != nil {
		return m.Value
	}
	return nil
}

type MsBarrage struct {
	PlayerNick       *string `protobuf:"bytes,1,req,name=PlayerNick" json:"PlayerNick,omitempty"`
	EntityId         *int32  `protobuf:"varint,2,req,name=EntityId" json:"EntityId,omitempty"`
	Value            *int64  `protobuf:"varint,3,req,name=Value" json:"Value,omitempty"`
	OpenTime         *int64  `protobuf:"varint,4,opt,name=OpenTime" json:"OpenTime,omitempty"`
	PlayerId         *int64  `protobuf:"varint,5,opt,name=PlayerId" json:"PlayerId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsBarrage) Reset()                    { *m = MsBarrage{} }
func (m *MsBarrage) String() string            { return proto.CompactTextString(m) }
func (*MsBarrage) ProtoMessage()               {}
func (*MsBarrage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *MsBarrage) GetPlayerNick() string {
	if m != nil && m.PlayerNick != nil {
		return *m.PlayerNick
	}
	return ""
}

func (m *MsBarrage) GetEntityId() int32 {
	if m != nil && m.EntityId != nil {
		return *m.EntityId
	}
	return 0
}

func (m *MsBarrage) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *MsBarrage) GetOpenTime() int64 {
	if m != nil && m.OpenTime != nil {
		return *m.OpenTime
	}
	return 0
}

func (m *MsBarrage) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type MsRollResp struct {
	Value            []*MsBarrage `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MsRollResp) Reset()                    { *m = MsRollResp{} }
func (m *MsRollResp) String() string            { return proto.CompactTextString(m) }
func (*MsRollResp) ProtoMessage()               {}
func (*MsRollResp) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *MsRollResp) GetValue() []*MsBarrage {
	if m != nil {
		return m.Value
	}
	return nil
}

type MsRankResp struct {
	Value            []*MsBarrage `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MsRankResp) Reset()                    { *m = MsRankResp{} }
func (m *MsRankResp) String() string            { return proto.CompactTextString(m) }
func (*MsRankResp) ProtoMessage()               {}
func (*MsRankResp) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *MsRankResp) GetValue() []*MsBarrage {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MsAskSignResp)(nil), "VK.Proto.MsAskSignResp")
	proto.RegisterType((*MsSignResp)(nil), "VK.Proto.MsSignResp")
	proto.RegisterType((*MsAskReplenishResp)(nil), "VK.Proto.MsAskReplenishResp")
	proto.RegisterType((*MsReplenishReq)(nil), "VK.Proto.MsReplenishReq")
	proto.RegisterType((*MsBoxContent)(nil), "VK.Proto.MsBoxContent")
	proto.RegisterType((*MsBox)(nil), "VK.Proto.MsBox")
	proto.RegisterType((*MsAskBoxResp)(nil), "VK.Proto.MsAskBoxResp")
	proto.RegisterType((*MsBarrage)(nil), "VK.Proto.MsBarrage")
	proto.RegisterType((*MsRollResp)(nil), "VK.Proto.MsRollResp")
	proto.RegisterType((*MsRankResp)(nil), "VK.Proto.MsRankResp")
	proto.RegisterEnum("VK.Proto.MsignOps", MsignOps_name, MsignOps_value)
	proto.RegisterEnum("VK.Proto.MsMark", MsMark_name, MsMark_value)
	proto.RegisterEnum("VK.Proto.MsBoxType", MsBoxType_name, MsBoxType_value)
}

func init() { proto.RegisterFile("client_msign.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6a, 0xdb, 0x4a,
	0x10, 0x45, 0x92, 0x65, 0xcb, 0x63, 0x27, 0xd6, 0xdd, 0x84, 0x20, 0x2e, 0x97, 0x8b, 0x11, 0xdc,
	0x5b, 0x37, 0x14, 0x53, 0x02, 0xa5, 0x7d, 0x2a, 0x8d, 0xec, 0x40, 0x43, 0x51, 0x12, 0x54, 0x93,
	0xd2, 0x27, 0xb1, 0xc4, 0x1b, 0x47, 0x48, 0x59, 0x29, 0x5a, 0xa5, 0x48, 0x1f, 0xd2, 0x3f, 0x28,
	0xf4, 0xab, 0xfa, 0x2f, 0x65, 0x76, 0x25, 0x4b, 0x0e, 0xed, 0x43, 0xdf, 0xe6, 0x9c, 0xd9, 0x19,
	0x9d, 0x39, 0xb3, 0x2b, 0x20, 0x37, 0x49, 0xc4, 0x78, 0x11, 0xde, 0x8b, 0x68, 0xc3, 0xe7, 0x59,
	0x9e, 0x16, 0x29, 0xb1, 0xae, 0x3f, 0xcc, 0xaf, 0x30, 0x72, 0x9f, 0xc1, 0x9e, 0x2f, 0x4e, 0x45,
	0xfc, 0x31, 0xda, 0xf0, 0x80, 0x89, 0x8c, 0x1c, 0x41, 0xff, 0x5c, 0x20, 0x72, 0xb4, 0xa9, 0x3e,
	0xb3, 0x82, 0x1a, 0xb9, 0x6f, 0x01, 0x7c, 0xb1, 0x3d, 0xf5, 0x37, 0x58, 0x8b, 0x94, 0xdf, 0x46,
	0x9b, 0xf3, 0xb5, 0x3c, 0x67, 0x06, 0x5b, 0x4c, 0x0e, 0xc1, 0x5c, 0xa4, 0x8f, 0xbc, 0x70, 0xf4,
	0xa9, 0x3e, 0x33, 0x02, 0x05, 0xdc, 0x6f, 0x1a, 0x10, 0xf9, 0xa5, 0x80, 0x65, 0x09, 0xe3, 0x91,
	0xb8, 0x93, 0x8d, 0x08, 0xf4, 0x3e, 0x33, 0x9a, 0xd7, 0x4d, 0x64, 0x8c, 0x0d, 0xfc, 0x94, 0x17,
	0x77, 0xb2, 0x81, 0x19, 0x28, 0x40, 0x6c, 0x30, 0x96, 0xb4, 0x72, 0x0c, 0xc9, 0x61, 0x88, 0xb5,
	0x4b, 0x5a, 0x30, 0xa7, 0x37, 0x35, 0xb0, 0x16, 0x63, 0xf2, 0x3f, 0xec, 0x07, 0xec, 0x9e, 0x46,
	0x3c, 0xe2, 0x9b, 0x55, 0x74, 0xcf, 0x84, 0x63, 0xca, 0x82, 0x27, 0x2c, 0x8e, 0x79, 0x56, 0x66,
	0x8c, 0xaf, 0x9d, 0xbe, 0xcc, 0xd7, 0xc8, 0x7d, 0x03, 0xfb, 0xbe, 0xe8, 0x48, 0x7c, 0x68, 0xd5,
	0x68, 0xbf, 0x50, 0xa3, 0x6f, 0xd5, 0xb8, 0xef, 0x60, 0xec, 0x0b, 0x2f, 0x2d, 0x17, 0x29, 0x2f,
	0x18, 0x2f, 0xd0, 0xa2, 0x33, 0x5e, 0x44, 0x45, 0xd5, 0x5a, 0xd4, 0x60, 0xec, 0x79, 0x4d, 0x93,
	0x47, 0xd6, 0x58, 0x24, 0x81, 0xfb, 0x5d, 0x07, 0x53, 0xb6, 0x20, 0xff, 0xc0, 0x10, 0xad, 0x56,
	0x03, 0xa8, 0xe2, 0x96, 0xc0, 0xec, 0x05, 0x63, 0x6b, 0x95, 0x55, 0x0a, 0x5a, 0x42, 0x2d, 0xf0,
	0x32, 0x63, 0x5c, 0x5a, 0x25, 0x17, 0x88, 0x88, 0x38, 0x30, 0x58, 0x50, 0x2e, 0x13, 0x3d, 0x99,
	0x68, 0x20, 0x71, 0x61, 0x7c, 0xc1, 0xca, 0x02, 0x63, 0x6c, 0x21, 0x1d, 0x33, 0x82, 0x1d, 0x0e,
	0xab, 0xbd, 0xb4, 0x5c, 0x55, 0x19, 0xab, 0x0d, 0x6b, 0x20, 0x79, 0x09, 0x83, 0x7a, 0x64, 0x67,
	0x30, 0xd5, 0x67, 0xa3, 0x93, 0xa3, 0x79, 0x73, 0xbb, 0xe6, 0x5d, 0x43, 0x82, 0xe6, 0x58, 0xed,
	0x7d, 0x94, 0x33, 0xc7, 0x52, 0x0a, 0x15, 0x22, 0x33, 0x98, 0xe0, 0x37, 0x4f, 0x8b, 0xe8, 0x4b,
	0x54, 0x54, 0x52, 0xca, 0x50, 0x4a, 0x79, 0x4a, 0xbb, 0xaf, 0xd0, 0xeb, 0x53, 0x11, 0x7b, 0x69,
	0x29, 0x6f, 0xd1, 0x7f, 0x8d, 0x9f, 0xda, 0xd4, 0x98, 0x8d, 0x4e, 0x26, 0x4f, 0x14, 0x34, 0x06,
	0x7f, 0xd5, 0x60, 0xe8, 0x0b, 0x8f, 0xe6, 0x39, 0xdd, 0x30, 0xf2, 0x2f, 0xc0, 0x55, 0x42, 0x2b,
	0x96, 0x5f, 0x44, 0x37, 0xb1, 0x74, 0x79, 0x18, 0x74, 0x98, 0x9d, 0x05, 0xea, 0xbf, 0x5b, 0xa0,
	0xd1, 0x59, 0x20, 0x56, 0x6c, 0x4d, 0xec, 0x4d, 0xb5, 0x99, 0x11, 0x6c, 0x31, 0xe6, 0x54, 0xef,
	0xf3, 0xb5, 0x63, 0xaa, 0x5c, 0x83, 0xdd, 0xd7, 0xf8, 0xb6, 0x82, 0x34, 0x49, 0xe4, 0x30, 0xcf,
	0x77, 0x87, 0x39, 0xd8, 0x19, 0x46, 0x69, 0x6f, 0x06, 0x52, 0x85, 0x94, 0xc7, 0x7f, 0x58, 0x78,
	0xfc, 0x43, 0x03, 0xcb, 0xc7, 0x1f, 0xc2, 0x65, 0x26, 0xc8, 0x04, 0x46, 0x01, 0x7b, 0x08, 0xeb,
	0xbf, 0x80, 0xad, 0x91, 0x31, 0x58, 0x48, 0x48, 0xa4, 0x93, 0x43, 0xb0, 0xeb, 0xf4, 0xf6, 0x5d,
	0xd8, 0x06, 0xf9, 0x0b, 0xf6, 0x90, 0x6d, 0xa9, 0x1e, 0xd9, 0x07, 0xa8, 0x0f, 0x7a, 0x69, 0x69,
	0x9b, 0x84, 0xe0, 0x5b, 0x7c, 0x08, 0xd1, 0x02, 0x2f, 0x2d, 0x97, 0xb4, 0xb2, 0xfb, 0xe4, 0x00,
	0x26, 0x1d, 0xee, 0x13, 0x63, 0xb1, 0x3d, 0x20, 0x0e, 0x1c, 0x76, 0xc8, 0xf7, 0x34, 0xb9, 0x95,
	0x8f, 0xcc, 0xb6, 0x9a, 0x6f, 0xd7, 0x19, 0xc5, 0x0e, 0x3b, 0x82, 0xd1, 0x34, 0x1b, 0xba, 0x04,
	0xe5, 0xb1, 0x3d, 0x3a, 0x7e, 0x01, 0x7d, 0x5f, 0xf8, 0x34, 0x8f, 0x89, 0x05, 0xbd, 0x7a, 0xaa,
	0x3d, 0x18, 0xb6, 0x6a, 0x75, 0x32, 0x04, 0x73, 0xc9, 0x12, 0x5a, 0xd9, 0xc6, 0xf1, 0x52, 0x5e,
	0x8b, 0xfa, 0x3e, 0x03, 0xf4, 0xbd, 0x55, 0x88, 0x6a, 0x35, 0x32, 0x82, 0x81, 0xb7, 0x0a, 0xa5,
	0x4a, 0x9d, 0xd8, 0x30, 0xf6, 0x56, 0x61, 0xab, 0xce, 0x40, 0x9f, 0xbc, 0x55, 0xa8, 0x50, 0xef,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x0e, 0xb9, 0x57, 0x69, 0x05, 0x00, 0x00,
}
