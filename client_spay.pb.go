// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_spay.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SpayType int32

const (
	SpayType_Banktranfer SpayType = 1
	SpayType_Momo        SpayType = 2
)

var SpayType_name = map[int32]string{
	1: "Banktranfer",
	2: "Momo",
}
var SpayType_value = map[string]int32{
	"Banktranfer": 1,
	"Momo":        2,
}

func (x SpayType) Enum() *SpayType {
	p := new(SpayType)
	*p = x
	return p
}
func (x SpayType) String() string {
	return proto.EnumName(SpayType_name, int32(x))
}
func (x *SpayType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SpayType_value, data, "SpayType")
	if err != nil {
		return err
	}
	*x = SpayType(value)
	return nil
}
func (SpayType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type SpayOrderStatus int32

const (
	SpayOrderStatus_SpayOrderStatusNone   SpayOrderStatus = 1
	SpayOrderStatus_SpayOrderStatusFailed SpayOrderStatus = 2
	SpayOrderStatus_SpayOrderStatusDone   SpayOrderStatus = 3
	SpayOrderStatus_SpayOrderStatusNext   SpayOrderStatus = 4
	SpayOrderStatus_SpayOrderStatusWait   SpayOrderStatus = 5
	SpayOrderStatus_SpayOrderStatusNew    SpayOrderStatus = 6
)

var SpayOrderStatus_name = map[int32]string{
	1: "SpayOrderStatusNone",
	2: "SpayOrderStatusFailed",
	3: "SpayOrderStatusDone",
	4: "SpayOrderStatusNext",
	5: "SpayOrderStatusWait",
	6: "SpayOrderStatusNew",
}
var SpayOrderStatus_value = map[string]int32{
	"SpayOrderStatusNone":   1,
	"SpayOrderStatusFailed": 2,
	"SpayOrderStatusDone":   3,
	"SpayOrderStatusNext":   4,
	"SpayOrderStatusWait":   5,
	"SpayOrderStatusNew":    6,
}

func (x SpayOrderStatus) Enum() *SpayOrderStatus {
	p := new(SpayOrderStatus)
	*p = x
	return p
}
func (x SpayOrderStatus) String() string {
	return proto.EnumName(SpayOrderStatus_name, int32(x))
}
func (x *SpayOrderStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SpayOrderStatus_value, data, "SpayOrderStatus")
	if err != nil {
		return err
	}
	*x = SpayOrderStatus(value)
	return nil
}
func (SpayOrderStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

// 状态 0-未处理，1-正在处理，2-已处理，3-已拒绝
type SpayRecordStatus int32

const (
	SpayRecordStatus_SpayRecordStatusNot    SpayRecordStatus = 0
	SpayRecordStatus_SpayRecordStatusWait   SpayRecordStatus = 1
	SpayRecordStatus_SpayRecordStatusDone   SpayRecordStatus = 2
	SpayRecordStatus_SpayRecordStatusFailed SpayRecordStatus = 3
)

var SpayRecordStatus_name = map[int32]string{
	0: "SpayRecordStatusNot",
	1: "SpayRecordStatusWait",
	2: "SpayRecordStatusDone",
	3: "SpayRecordStatusFailed",
}
var SpayRecordStatus_value = map[string]int32{
	"SpayRecordStatusNot":    0,
	"SpayRecordStatusWait":   1,
	"SpayRecordStatusDone":   2,
	"SpayRecordStatusFailed": 3,
}

func (x SpayRecordStatus) Enum() *SpayRecordStatus {
	p := new(SpayRecordStatus)
	*p = x
	return p
}
func (x SpayRecordStatus) String() string {
	return proto.EnumName(SpayRecordStatus_name, int32(x))
}
func (x *SpayRecordStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SpayRecordStatus_value, data, "SpayRecordStatus")
	if err != nil {
		return err
	}
	*x = SpayRecordStatus(value)
	return nil
}
func (SpayRecordStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

type SpayBankAccount struct {
	Id                *int32  `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	BankName          *string `protobuf:"bytes,2,req,name=BankName" json:"BankName,omitempty"`
	BankAccountNumber *string `protobuf:"bytes,3,req,name=BankAccountNumber" json:"BankAccountNumber,omitempty"`
	BankAccountName   *string `protobuf:"bytes,4,req,name=BankAccountName" json:"BankAccountName,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *SpayBankAccount) Reset()                    { *m = SpayBankAccount{} }
func (m *SpayBankAccount) String() string            { return proto.CompactTextString(m) }
func (*SpayBankAccount) ProtoMessage()               {}
func (*SpayBankAccount) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SpayBankAccount) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SpayBankAccount) GetBankName() string {
	if m != nil && m.BankName != nil {
		return *m.BankName
	}
	return ""
}

func (m *SpayBankAccount) GetBankAccountNumber() string {
	if m != nil && m.BankAccountNumber != nil {
		return *m.BankAccountNumber
	}
	return ""
}

func (m *SpayBankAccount) GetBankAccountName() string {
	if m != nil && m.BankAccountName != nil {
		return *m.BankAccountName
	}
	return ""
}

type SpayOrder struct {
	Status            *string `protobuf:"bytes,1,req,name=Status" json:"Status,omitempty"`
	BankName          *string `protobuf:"bytes,2,req,name=BankName" json:"BankName,omitempty"`
	BankAccountNumber *string `protobuf:"bytes,3,req,name=BankAccountNumber" json:"BankAccountNumber,omitempty"`
	BankAccountName   *string `protobuf:"bytes,4,req,name=BankAccountName" json:"BankAccountName,omitempty"`
	Amount            *int64  `protobuf:"varint,5,req,name=Amount" json:"Amount,omitempty"`
	RefCode           *string `protobuf:"bytes,6,req,name=RefCode" json:"RefCode,omitempty"`
	OrderNo           *string `protobuf:"bytes,7,req,name=OrderNo" json:"OrderNo,omitempty"`
	Timeout           *int32  `protobuf:"varint,8,req,name=Timeout" json:"Timeout,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *SpayOrder) Reset()                    { *m = SpayOrder{} }
func (m *SpayOrder) String() string            { return proto.CompactTextString(m) }
func (*SpayOrder) ProtoMessage()               {}
func (*SpayOrder) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SpayOrder) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *SpayOrder) GetBankName() string {
	if m != nil && m.BankName != nil {
		return *m.BankName
	}
	return ""
}

func (m *SpayOrder) GetBankAccountNumber() string {
	if m != nil && m.BankAccountNumber != nil {
		return *m.BankAccountNumber
	}
	return ""
}

func (m *SpayOrder) GetBankAccountName() string {
	if m != nil && m.BankAccountName != nil {
		return *m.BankAccountName
	}
	return ""
}

func (m *SpayOrder) GetAmount() int64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *SpayOrder) GetRefCode() string {
	if m != nil && m.RefCode != nil {
		return *m.RefCode
	}
	return ""
}

func (m *SpayOrder) GetOrderNo() string {
	if m != nil && m.OrderNo != nil {
		return *m.OrderNo
	}
	return ""
}

func (m *SpayOrder) GetTimeout() int32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

// 获取银行列表请求
type SpayBankAccountListReq struct {
	Type             *SpayType `protobuf:"varint,1,req,name=Type,enum=VK.Proto.SpayType" json:"Type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SpayBankAccountListReq) Reset()                    { *m = SpayBankAccountListReq{} }
func (m *SpayBankAccountListReq) String() string            { return proto.CompactTextString(m) }
func (*SpayBankAccountListReq) ProtoMessage()               {}
func (*SpayBankAccountListReq) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *SpayBankAccountListReq) GetType() SpayType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SpayType_Banktranfer
}

// 获取银行列表返回
type SpayBankAccountListRsp struct {
	List             []*SpayBankAccount `protobuf:"bytes,1,rep,name=List" json:"List,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SpayBankAccountListRsp) Reset()                    { *m = SpayBankAccountListRsp{} }
func (m *SpayBankAccountListRsp) String() string            { return proto.CompactTextString(m) }
func (*SpayBankAccountListRsp) ProtoMessage()               {}
func (*SpayBankAccountListRsp) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *SpayBankAccountListRsp) GetList() []*SpayBankAccount {
	if m != nil {
		return m.List
	}
	return nil
}

// 下单请求
type SpayOrderReq struct {
	Type              *SpayType `protobuf:"varint,1,req,name=Type,enum=VK.Proto.SpayType" json:"Type,omitempty"`
	Amount            *int64    `protobuf:"varint,2,req,name=Amount" json:"Amount,omitempty"`
	Id                *int32    `protobuf:"varint,3,req,name=Id" json:"Id,omitempty"`
	BankName          *string   `protobuf:"bytes,4,req,name=BankName" json:"BankName,omitempty"`
	BankAccountNumber *string   `protobuf:"bytes,5,req,name=BankAccountNumber" json:"BankAccountNumber,omitempty"`
	BankAccountName   *string   `protobuf:"bytes,6,req,name=BankAccountName" json:"BankAccountName,omitempty"`
	Rewards           *int64    `protobuf:"varint,7,req,name=Rewards" json:"Rewards,omitempty"`
	ExtraData         *string   `protobuf:"bytes,8,opt,name=ExtraData" json:"ExtraData,omitempty"`
	GameId            *int32    `protobuf:"varint,9,opt,name=GameId" json:"GameId,omitempty"`
	PackageChn        *int32    `protobuf:"varint,10,opt,name=PackageChn" json:"PackageChn,omitempty"`
	XXX_unrecognized  []byte    `json:"-"`
}

func (m *SpayOrderReq) Reset()                    { *m = SpayOrderReq{} }
func (m *SpayOrderReq) String() string            { return proto.CompactTextString(m) }
func (*SpayOrderReq) ProtoMessage()               {}
func (*SpayOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *SpayOrderReq) GetType() SpayType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SpayType_Banktranfer
}

func (m *SpayOrderReq) GetAmount() int64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *SpayOrderReq) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SpayOrderReq) GetBankName() string {
	if m != nil && m.BankName != nil {
		return *m.BankName
	}
	return ""
}

func (m *SpayOrderReq) GetBankAccountNumber() string {
	if m != nil && m.BankAccountNumber != nil {
		return *m.BankAccountNumber
	}
	return ""
}

func (m *SpayOrderReq) GetBankAccountName() string {
	if m != nil && m.BankAccountName != nil {
		return *m.BankAccountName
	}
	return ""
}

func (m *SpayOrderReq) GetRewards() int64 {
	if m != nil && m.Rewards != nil {
		return *m.Rewards
	}
	return 0
}

func (m *SpayOrderReq) GetExtraData() string {
	if m != nil && m.ExtraData != nil {
		return *m.ExtraData
	}
	return ""
}

func (m *SpayOrderReq) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *SpayOrderReq) GetPackageChn() int32 {
	if m != nil && m.PackageChn != nil {
		return *m.PackageChn
	}
	return 0
}

// 查询请求
type SpayOrderCheckReq struct {
	Type             *SpayType `protobuf:"varint,1,req,name=type,enum=VK.Proto.SpayType" json:"type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SpayOrderCheckReq) Reset()                    { *m = SpayOrderCheckReq{} }
func (m *SpayOrderCheckReq) String() string            { return proto.CompactTextString(m) }
func (*SpayOrderCheckReq) ProtoMessage()               {}
func (*SpayOrderCheckReq) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *SpayOrderCheckReq) GetType() SpayType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SpayType_Banktranfer
}

// 下单和查询返回
type SpayOrderRsp struct {
	Type             *SpayType        `protobuf:"varint,1,req,name=Type,enum=VK.Proto.SpayType" json:"Type,omitempty"`
	Status           *SpayOrderStatus `protobuf:"varint,2,req,name=Status,enum=VK.Proto.SpayOrderStatus" json:"Status,omitempty"`
	Order            *SpayOrder       `protobuf:"bytes,3,opt,name=Order" json:"Order,omitempty"`
	Rewards          *int64           `protobuf:"varint,4,opt,name=Rewards" json:"Rewards,omitempty"`
	CD               *int32           `protobuf:"varint,5,opt,name=CD" json:"CD,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SpayOrderRsp) Reset()                    { *m = SpayOrderRsp{} }
func (m *SpayOrderRsp) String() string            { return proto.CompactTextString(m) }
func (*SpayOrderRsp) ProtoMessage()               {}
func (*SpayOrderRsp) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *SpayOrderRsp) GetType() SpayType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SpayType_Banktranfer
}

func (m *SpayOrderRsp) GetStatus() SpayOrderStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return SpayOrderStatus_SpayOrderStatusNone
}

func (m *SpayOrderRsp) GetOrder() *SpayOrder {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *SpayOrderRsp) GetRewards() int64 {
	if m != nil && m.Rewards != nil {
		return *m.Rewards
	}
	return 0
}

func (m *SpayOrderRsp) GetCD() int32 {
	if m != nil && m.CD != nil {
		return *m.CD
	}
	return 0
}

// 下分请求
type BankExchangeReq struct {
	BankName          *string `protobuf:"bytes,1,req,name=BankName" json:"BankName,omitempty"`
	BankAccountNumber *string `protobuf:"bytes,2,req,name=BankAccountNumber" json:"BankAccountNumber,omitempty"`
	BankAccountName   *string `protobuf:"bytes,3,req,name=BankAccountName" json:"BankAccountName,omitempty"`
	Amount            *int64  `protobuf:"varint,4,req,name=Amount" json:"Amount,omitempty"`
	Chips             *int64  `protobuf:"varint,5,req,name=Chips" json:"Chips,omitempty"`
	Context           *string `protobuf:"bytes,6,req,name=Context" json:"Context,omitempty"`
	OTP               *string `protobuf:"bytes,7,req,name=OTP" json:"OTP,omitempty"`
	BankNo            *int32  `protobuf:"varint,8,req,name=BankNo" json:"BankNo,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *BankExchangeReq) Reset()                    { *m = BankExchangeReq{} }
func (m *BankExchangeReq) String() string            { return proto.CompactTextString(m) }
func (*BankExchangeReq) ProtoMessage()               {}
func (*BankExchangeReq) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

func (m *BankExchangeReq) GetBankName() string {
	if m != nil && m.BankName != nil {
		return *m.BankName
	}
	return ""
}

func (m *BankExchangeReq) GetBankAccountNumber() string {
	if m != nil && m.BankAccountNumber != nil {
		return *m.BankAccountNumber
	}
	return ""
}

func (m *BankExchangeReq) GetBankAccountName() string {
	if m != nil && m.BankAccountName != nil {
		return *m.BankAccountName
	}
	return ""
}

func (m *BankExchangeReq) GetAmount() int64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *BankExchangeReq) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *BankExchangeReq) GetContext() string {
	if m != nil && m.Context != nil {
		return *m.Context
	}
	return ""
}

func (m *BankExchangeReq) GetOTP() string {
	if m != nil && m.OTP != nil {
		return *m.OTP
	}
	return ""
}

func (m *BankExchangeReq) GetBankNo() int32 {
	if m != nil && m.BankNo != nil {
		return *m.BankNo
	}
	return 0
}

// 银行操作记录
type SpayBankRecord struct {
	BankName         *string           `protobuf:"bytes,1,req,name=BankName" json:"BankName,omitempty"`
	Amount           *int64            `protobuf:"varint,2,req,name=Amount" json:"Amount,omitempty"`
	Chips            *int64            `protobuf:"varint,3,req,name=Chips" json:"Chips,omitempty"`
	Stamp            *int64            `protobuf:"varint,4,req,name=Stamp" json:"Stamp,omitempty"`
	Status           *SpayRecordStatus `protobuf:"varint,5,req,name=Status,enum=VK.Proto.SpayRecordStatus" json:"Status,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SpayBankRecord) Reset()                    { *m = SpayBankRecord{} }
func (m *SpayBankRecord) String() string            { return proto.CompactTextString(m) }
func (*SpayBankRecord) ProtoMessage()               {}
func (*SpayBankRecord) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *SpayBankRecord) GetBankName() string {
	if m != nil && m.BankName != nil {
		return *m.BankName
	}
	return ""
}

func (m *SpayBankRecord) GetAmount() int64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *SpayBankRecord) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *SpayBankRecord) GetStamp() int64 {
	if m != nil && m.Stamp != nil {
		return *m.Stamp
	}
	return 0
}

func (m *SpayBankRecord) GetStatus() SpayRecordStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return SpayRecordStatus_SpayRecordStatusNot
}

// 银行操作记录返回
type SpayBankRecordRsp struct {
	Records          []*SpayBankRecord `protobuf:"bytes,1,rep,name=Records" json:"Records,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SpayBankRecordRsp) Reset()                    { *m = SpayBankRecordRsp{} }
func (m *SpayBankRecordRsp) String() string            { return proto.CompactTextString(m) }
func (*SpayBankRecordRsp) ProtoMessage()               {}
func (*SpayBankRecordRsp) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

func (m *SpayBankRecordRsp) GetRecords() []*SpayBankRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*SpayBankAccount)(nil), "VK.Proto.SpayBankAccount")
	proto.RegisterType((*SpayOrder)(nil), "VK.Proto.SpayOrder")
	proto.RegisterType((*SpayBankAccountListReq)(nil), "VK.Proto.SpayBankAccountListReq")
	proto.RegisterType((*SpayBankAccountListRsp)(nil), "VK.Proto.SpayBankAccountListRsp")
	proto.RegisterType((*SpayOrderReq)(nil), "VK.Proto.SpayOrderReq")
	proto.RegisterType((*SpayOrderCheckReq)(nil), "VK.Proto.SpayOrderCheckReq")
	proto.RegisterType((*SpayOrderRsp)(nil), "VK.Proto.SpayOrderRsp")
	proto.RegisterType((*BankExchangeReq)(nil), "VK.Proto.BankExchangeReq")
	proto.RegisterType((*SpayBankRecord)(nil), "VK.Proto.SpayBankRecord")
	proto.RegisterType((*SpayBankRecordRsp)(nil), "VK.Proto.SpayBankRecordRsp")
	proto.RegisterEnum("VK.Proto.SpayType", SpayType_name, SpayType_value)
	proto.RegisterEnum("VK.Proto.SpayOrderStatus", SpayOrderStatus_name, SpayOrderStatus_value)
	proto.RegisterEnum("VK.Proto.SpayRecordStatus", SpayRecordStatus_name, SpayRecordStatus_value)
}

func init() { proto.RegisterFile("client_spay.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xff, 0x76, 0x6d, 0xa7, 0xc9, 0xf4, 0x53, 0xea, 0x6e, 0xfb, 0xf5, 0xdb, 0x56, 0x08, 0x59,
	0x96, 0x40, 0xa6, 0x82, 0x48, 0xe4, 0xca, 0x85, 0x92, 0x94, 0xaa, 0x02, 0xd2, 0xca, 0xad, 0xe0,
	0x88, 0x96, 0x78, 0xdb, 0x44, 0xad, 0xff, 0x60, 0x6f, 0xd4, 0x54, 0x3c, 0x07, 0x6f, 0xc0, 0x81,
	0xa7, 0xe0, 0xc0, 0x4b, 0x71, 0x42, 0x42, 0xbb, 0x5e, 0xbb, 0x8e, 0x6b, 0x68, 0x7b, 0xe2, 0xe6,
	0xdf, 0xfc, 0x66, 0x46, 0x33, 0xf3, 0x9b, 0x59, 0xc3, 0xea, 0xf8, 0x7c, 0xca, 0x23, 0xf1, 0x3e,
	0x4b, 0xd8, 0x65, 0x2f, 0x49, 0x63, 0x11, 0x93, 0xf6, 0xdb, 0x57, 0xbd, 0x43, 0xf9, 0xe5, 0x7e,
	0x46, 0xb0, 0x72, 0x94, 0xb0, 0xcb, 0x17, 0x2c, 0x3a, 0xdb, 0x19, 0x8f, 0xe3, 0x59, 0x24, 0x48,
	0x17, 0xf0, 0x7e, 0x40, 0x91, 0x83, 0x3d, 0xcb, 0xc7, 0xfb, 0x01, 0xd9, 0x82, 0xb6, 0xa4, 0x47,
	0x2c, 0xe4, 0x14, 0x3b, 0xd8, 0xeb, 0xf8, 0x25, 0x26, 0x8f, 0x61, 0xb5, 0x12, 0x3a, 0x9a, 0x85,
	0x1f, 0x78, 0x4a, 0x0d, 0xe5, 0x74, 0x9d, 0x20, 0x1e, 0xac, 0x54, 0x8d, 0x32, 0xa1, 0xa9, 0x7c,
	0xeb, 0x66, 0xf7, 0x27, 0x82, 0x8e, 0xac, 0xeb, 0x20, 0x0d, 0x78, 0x4a, 0x36, 0xa0, 0x75, 0x24,
	0x98, 0x98, 0x65, 0xaa, 0xaa, 0x8e, 0xaf, 0xd1, 0xdf, 0xa8, 0x4c, 0xd6, 0xb2, 0x13, 0x4a, 0x44,
	0x2d, 0x07, 0x7b, 0x86, 0xaf, 0x11, 0xa1, 0xb0, 0xe4, 0xf3, 0x93, 0x41, 0x1c, 0x70, 0xda, 0x52,
	0x91, 0x05, 0x94, 0x8c, 0x6a, 0x63, 0x14, 0xd3, 0xa5, 0x9c, 0xd1, 0x50, 0x32, 0xc7, 0xd3, 0x90,
	0xc7, 0x33, 0x41, 0xdb, 0x6a, 0xdc, 0x05, 0x74, 0x9f, 0xc3, 0x46, 0x4d, 0x96, 0xd7, 0xd3, 0x4c,
	0xf8, 0xfc, 0x23, 0x79, 0x08, 0xe6, 0xf1, 0x65, 0xc2, 0xd5, 0x24, 0xba, 0x7d, 0xd2, 0x2b, 0xa4,
	0xec, 0x49, 0x7f, 0xc9, 0xf8, 0x8a, 0x77, 0xf7, 0x9a, 0x33, 0x64, 0x09, 0x79, 0x02, 0xa6, 0xfc,
	0xa4, 0xc8, 0x31, 0xbc, 0xe5, 0xfe, 0xe6, 0x62, 0x86, 0x8a, 0xbf, 0xaf, 0xdc, 0xdc, 0xef, 0x18,
	0xfe, 0x2d, 0xa5, 0xb8, 0x43, 0x05, 0x95, 0x49, 0xe1, 0x85, 0x49, 0xe5, 0xfb, 0x65, 0x34, 0xee,
	0x97, 0x79, 0x1b, 0x15, 0xad, 0x3b, 0xa8, 0xd8, 0x6a, 0x56, 0x51, 0xa9, 0x75, 0xc1, 0xd2, 0x20,
	0x53, 0x9a, 0x18, 0x7e, 0x01, 0xc9, 0x3d, 0xe8, 0xec, 0xce, 0x45, 0xca, 0x86, 0x4c, 0x30, 0xda,
	0x76, 0x90, 0xd7, 0xf1, 0xaf, 0x0c, 0xb2, 0xa7, 0x3d, 0x16, 0xf2, 0xfd, 0x80, 0x76, 0x1c, 0xe4,
	0x59, 0xbe, 0x46, 0xe4, 0x3e, 0xc0, 0x21, 0x1b, 0x9f, 0xb1, 0x53, 0x3e, 0x98, 0x44, 0x14, 0x14,
	0x57, 0xb1, 0xb8, 0xcf, 0x60, 0xb5, 0x9c, 0xe1, 0x60, 0xc2, 0xc7, 0x67, 0x7a, 0x90, 0xe2, 0x86,
	0x41, 0x4a, 0xde, 0xfd, 0x86, 0xaa, 0x0a, 0x64, 0xc9, 0xad, 0x15, 0x78, 0x5a, 0xde, 0x0d, 0x56,
	0x9e, 0x35, 0xad, 0x55, 0xbe, 0xdc, 0xa1, 0x3c, 0xa9, 0x47, 0x60, 0x29, 0x33, 0x35, 0x1c, 0xe4,
	0x2d, 0xf7, 0xd7, 0x1a, 0x22, 0xfc, 0xdc, 0xa3, 0x3a, 0x43, 0xd3, 0x41, 0xd5, 0x19, 0x76, 0x01,
	0x0f, 0x86, 0xd4, 0x52, 0x53, 0xc0, 0x83, 0xa1, 0xfb, 0x03, 0xe5, 0xc2, 0xec, 0xce, 0xc7, 0x13,
	0x16, 0x9d, 0x72, 0xd9, 0x7c, 0x55, 0x75, 0x74, 0x1b, 0xd5, 0xf1, 0x1d, 0x54, 0x37, 0x6e, 0xba,
	0x5d, 0x73, 0x61, 0x23, 0xd7, 0xc1, 0x1a, 0x4c, 0xa6, 0x49, 0xa6, 0x4f, 0x3a, 0x07, 0xb2, 0xbf,
	0x41, 0x1c, 0x09, 0x3e, 0x17, 0xc5, 0x45, 0x6b, 0x48, 0x6c, 0x30, 0x0e, 0x8e, 0x0f, 0xf5, 0x35,
	0xcb, 0x4f, 0x99, 0x59, 0x55, 0x1f, 0xeb, 0x43, 0xd6, 0xc8, 0xfd, 0x82, 0xa0, 0x5b, 0x9c, 0x95,
	0xcf, 0xc7, 0x71, 0x1a, 0xfc, 0xb1, 0xf1, 0xdf, 0x9d, 0x4c, 0x59, 0xa0, 0x51, 0x2d, 0x70, 0x1d,
	0xac, 0x23, 0xc1, 0xc2, 0x44, 0x77, 0x93, 0x03, 0xd2, 0x2f, 0x45, 0xb7, 0x94, 0xe8, 0x5b, 0x8b,
	0x12, 0xe6, 0x55, 0x2c, 0xaa, 0xee, 0xee, 0xe5, 0xeb, 0x79, 0x55, 0xa5, 0xdc, 0xb2, 0xbe, 0xd4,
	0x57, 0x82, 0x4c, 0x3f, 0x15, 0xf4, 0xfa, 0x53, 0xa1, 0xbd, 0x0b, 0xc7, 0xed, 0x07, 0xd0, 0x2e,
	0x76, 0x90, 0xac, 0xc0, 0xb2, 0x74, 0x11, 0x29, 0x8b, 0x4e, 0x78, 0x6a, 0x23, 0xd2, 0x06, 0xf3,
	0x4d, 0x1c, 0xc6, 0x36, 0xde, 0xfe, 0xaa, 0x7f, 0x3b, 0x95, 0x0d, 0x24, 0xff, 0xc3, 0x5a, 0xcd,
	0x34, 0x8a, 0x23, 0x6e, 0x23, 0xb2, 0x09, 0xff, 0xd5, 0x88, 0x97, 0x6c, 0x7a, 0xce, 0x03, 0x1b,
	0x37, 0xc4, 0x0c, 0x65, 0x8c, 0xd1, 0x94, 0x8c, 0xcf, 0x85, 0x6d, 0x36, 0x10, 0xef, 0xd8, 0x54,
	0xd8, 0x16, 0xd9, 0x00, 0x72, 0x2d, 0xe2, 0xc2, 0x6e, 0x6d, 0x7f, 0x02, 0xbb, 0x3e, 0xb6, 0x22,
	0x49, 0xd5, 0x36, 0x8a, 0x85, 0xfd, 0x0f, 0xa1, 0xb0, 0x5e, 0x27, 0x54, 0x7a, 0xd4, 0xc4, 0xa8,
	0x52, 0x31, 0xd9, 0xca, 0x1f, 0xea, 0x2a, 0xa3, 0xfb, 0x33, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0x7d, 0x0b, 0x75, 0xbc, 0x07, 0x00, 0x00,
}
