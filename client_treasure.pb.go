// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_treasure.proto

package VK_Proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 夺宝请求
type TreasureOps int32

const (
	// 奖品列表
	TreasureOps_Req_PrizeList TreasureOps = 1
	// 参加夺宝
	TreasureOps_Req_PrizePuarchase TreasureOps = 2
	// 往期幸运儿
	TreasureOps_Req_PrizeLuckyRecord TreasureOps = 3
	// 玩家购买记录
	TreasureOps_Req_PuarchaseRecord TreasureOps = 4
	// 玩家领奖记录
	TreasureOps_Req_PlayerLuckyRecord TreasureOps = 5
	// 请求玩家夺宝码
	TreasureOps_Req_PlayerLuckyCode TreasureOps = 6
	// 请求玩家夺宝码
	TreasureOps_Req_LatelyPlayerLuckyRecord TreasureOps = 7
)

var TreasureOps_name = map[int32]string{
	1: "Req_PrizeList",
	2: "Req_PrizePuarchase",
	3: "Req_PrizeLuckyRecord",
	4: "Req_PuarchaseRecord",
	5: "Req_PlayerLuckyRecord",
	6: "Req_PlayerLuckyCode",
	7: "Req_LatelyPlayerLuckyRecord",
}
var TreasureOps_value = map[string]int32{
	"Req_PrizeList":               1,
	"Req_PrizePuarchase":          2,
	"Req_PrizeLuckyRecord":        3,
	"Req_PuarchaseRecord":         4,
	"Req_PlayerLuckyRecord":       5,
	"Req_PlayerLuckyCode":         6,
	"Req_LatelyPlayerLuckyRecord": 7,
}

func (x TreasureOps) Enum() *TreasureOps {
	p := new(TreasureOps)
	*p = x
	return p
}
func (x TreasureOps) String() string {
	return proto.EnumName(TreasureOps_name, int32(x))
}
func (x *TreasureOps) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TreasureOps_value, data, "TreasureOps")
	if err != nil {
		return err
	}
	*x = TreasureOps(value)
	return nil
}
func (TreasureOps) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

// 奖品限制类型
type TsLimitType int32

const (
	TsLimitType_Temporary TsLimitType = 1
	TsLimitType_Permanent TsLimitType = 2
)

var TsLimitType_name = map[int32]string{
	1: "Temporary",
	2: "Permanent",
}
var TsLimitType_value = map[string]int32{
	"Temporary": 1,
	"Permanent": 2,
}

func (x TsLimitType) Enum() *TsLimitType {
	p := new(TsLimitType)
	*p = x
	return p
}
func (x TsLimitType) String() string {
	return proto.EnumName(TsLimitType_name, int32(x))
}
func (x *TsLimitType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TsLimitType_value, data, "TsLimitType")
	if err != nil {
		return err
	}
	*x = TsLimitType(value)
	return nil
}
func (TsLimitType) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

// 奖品限制类型
type TsOpenType int32

const (
	TsOpenType_Time  TsOpenType = 1
	TsOpenType_Times TsOpenType = 2
)

var TsOpenType_name = map[int32]string{
	1: "Time",
	2: "Times",
}
var TsOpenType_value = map[string]int32{
	"Time":  1,
	"Times": 2,
}

func (x TsOpenType) Enum() *TsOpenType {
	p := new(TsOpenType)
	*p = x
	return p
}
func (x TsOpenType) String() string {
	return proto.EnumName(TsOpenType_name, int32(x))
}
func (x *TsOpenType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TsOpenType_value, data, "TsOpenType")
	if err != nil {
		return err
	}
	*x = TsOpenType(value)
	return nil
}
func (TsOpenType) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

// 奖品每期状态
type TsIssueStatus int32

const (
	TsIssueStatus_IssuePurchase TsIssueStatus = 1
	TsIssueStatus_IssueOpen     TsIssueStatus = 2
	TsIssueStatus_IssueEnd      TsIssueStatus = 3
	TsIssueStatus_IssueRemain   TsIssueStatus = 4
	TsIssueStatus_IssuePrepare  TsIssueStatus = 5
	TsIssueStatus_IssueWaitOpen TsIssueStatus = 6
)

var TsIssueStatus_name = map[int32]string{
	1: "IssuePurchase",
	2: "IssueOpen",
	3: "IssueEnd",
	4: "IssueRemain",
	5: "IssuePrepare",
	6: "IssueWaitOpen",
}
var TsIssueStatus_value = map[string]int32{
	"IssuePurchase": 1,
	"IssueOpen":     2,
	"IssueEnd":      3,
	"IssueRemain":   4,
	"IssuePrepare":  5,
	"IssueWaitOpen": 6,
}

func (x TsIssueStatus) Enum() *TsIssueStatus {
	p := new(TsIssueStatus)
	*p = x
	return p
}
func (x TsIssueStatus) String() string {
	return proto.EnumName(TsIssueStatus_name, int32(x))
}
func (x *TsIssueStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TsIssueStatus_value, data, "TsIssueStatus")
	if err != nil {
		return err
	}
	*x = TsIssueStatus(value)
	return nil
}
func (TsIssueStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

// 奖品信息
type TsPrize struct {
	PrizeId          *int32         `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Issue            *int32         `protobuf:"varint,2,req,name=Issue" json:"Issue,omitempty"`
	VipLimit         *int32         `protobuf:"varint,3,req,name=VipLimit" json:"VipLimit,omitempty"`
	Currency         *int32         `protobuf:"varint,4,req,name=Currency" json:"Currency,omitempty"`
	Price            *int32         `protobuf:"varint,5,req,name=Price" json:"Price,omitempty"`
	LimitType        *TsLimitType   `protobuf:"varint,6,req,name=LimitType,enum=VK.Proto.TsLimitType" json:"LimitType,omitempty"`
	SellStartTime    *int64         `protobuf:"varint,7,opt,name=SellStartTime" json:"SellStartTime,omitempty"`
	SellEndTime      *int64         `protobuf:"varint,8,opt,name=SellEndTime" json:"SellEndTime,omitempty"`
	OpenType         *TsOpenType    `protobuf:"varint,9,req,name=OpenType,enum=VK.Proto.TsOpenType" json:"OpenType,omitempty"`
	CountDown        *int32         `protobuf:"varint,10,opt,name=CountDown" json:"CountDown,omitempty"`
	PurchasedQuota   *int32         `protobuf:"varint,11,req,name=PurchasedQuota" json:"PurchasedQuota,omitempty"`
	TotalQuota       *int32         `protobuf:"varint,12,req,name=TotalQuota" json:"TotalQuota,omitempty"`
	PropId           *int32         `protobuf:"varint,13,req,name=PropId" json:"PropId,omitempty"`
	PropCount        *int32         `protobuf:"varint,14,req,name=PropCount" json:"PropCount,omitempty"`
	Status           *TsIssueStatus `protobuf:"varint,15,req,name=Status,enum=VK.Proto.TsIssueStatus" json:"Status,omitempty"`
	LuckyPlayer      *TsLuckyPlayer `protobuf:"bytes,16,opt,name=LuckyPlayer" json:"LuckyPlayer,omitempty"`
	SoldQuota        *int32         `protobuf:"varint,17,req,name=SoldQuota" json:"SoldQuota,omitempty"`
	LimitQuota       *int32         `protobuf:"varint,18,req,name=LimitQuota" json:"LimitQuota,omitempty"`
	WaitOpen         *bool          `protobuf:"varint,19,req,name=WaitOpen" json:"WaitOpen,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TsPrize) Reset()                    { *m = TsPrize{} }
func (m *TsPrize) String() string            { return proto.CompactTextString(m) }
func (*TsPrize) ProtoMessage()               {}
func (*TsPrize) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *TsPrize) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsPrize) GetIssue() int32 {
	if m != nil && m.Issue != nil {
		return *m.Issue
	}
	return 0
}

func (m *TsPrize) GetVipLimit() int32 {
	if m != nil && m.VipLimit != nil {
		return *m.VipLimit
	}
	return 0
}

func (m *TsPrize) GetCurrency() int32 {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return 0
}

func (m *TsPrize) GetPrice() int32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *TsPrize) GetLimitType() TsLimitType {
	if m != nil && m.LimitType != nil {
		return *m.LimitType
	}
	return TsLimitType_Temporary
}

func (m *TsPrize) GetSellStartTime() int64 {
	if m != nil && m.SellStartTime != nil {
		return *m.SellStartTime
	}
	return 0
}

func (m *TsPrize) GetSellEndTime() int64 {
	if m != nil && m.SellEndTime != nil {
		return *m.SellEndTime
	}
	return 0
}

func (m *TsPrize) GetOpenType() TsOpenType {
	if m != nil && m.OpenType != nil {
		return *m.OpenType
	}
	return TsOpenType_Time
}

func (m *TsPrize) GetCountDown() int32 {
	if m != nil && m.CountDown != nil {
		return *m.CountDown
	}
	return 0
}

func (m *TsPrize) GetPurchasedQuota() int32 {
	if m != nil && m.PurchasedQuota != nil {
		return *m.PurchasedQuota
	}
	return 0
}

func (m *TsPrize) GetTotalQuota() int32 {
	if m != nil && m.TotalQuota != nil {
		return *m.TotalQuota
	}
	return 0
}

func (m *TsPrize) GetPropId() int32 {
	if m != nil && m.PropId != nil {
		return *m.PropId
	}
	return 0
}

func (m *TsPrize) GetPropCount() int32 {
	if m != nil && m.PropCount != nil {
		return *m.PropCount
	}
	return 0
}

func (m *TsPrize) GetStatus() TsIssueStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return TsIssueStatus_IssuePurchase
}

func (m *TsPrize) GetLuckyPlayer() *TsLuckyPlayer {
	if m != nil {
		return m.LuckyPlayer
	}
	return nil
}

func (m *TsPrize) GetSoldQuota() int32 {
	if m != nil && m.SoldQuota != nil {
		return *m.SoldQuota
	}
	return 0
}

func (m *TsPrize) GetLimitQuota() int32 {
	if m != nil && m.LimitQuota != nil {
		return *m.LimitQuota
	}
	return 0
}

func (m *TsPrize) GetWaitOpen() bool {
	if m != nil && m.WaitOpen != nil {
		return *m.WaitOpen
	}
	return false
}

// 奖品返回列表
type TsPrizeListResp struct {
	PrizeList        []*TsPrize `protobuf:"bytes,1,rep,name=PrizeList" json:"PrizeList,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *TsPrizeListResp) Reset()                    { *m = TsPrizeListResp{} }
func (m *TsPrizeListResp) String() string            { return proto.CompactTextString(m) }
func (*TsPrizeListResp) ProtoMessage()               {}
func (*TsPrizeListResp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *TsPrizeListResp) GetPrizeList() []*TsPrize {
	if m != nil {
		return m.PrizeList
	}
	return nil
}

// 参加夺宝请求
type TsPrizePuarchaseReq struct {
	PrizeId          *int32 `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Issue            *int32 `protobuf:"varint,2,req,name=Issue" json:"Issue,omitempty"`
	Times            *int32 `protobuf:"varint,3,req,name=Times" json:"Times,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsPrizePuarchaseReq) Reset()                    { *m = TsPrizePuarchaseReq{} }
func (m *TsPrizePuarchaseReq) String() string            { return proto.CompactTextString(m) }
func (*TsPrizePuarchaseReq) ProtoMessage()               {}
func (*TsPrizePuarchaseReq) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *TsPrizePuarchaseReq) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsPrizePuarchaseReq) GetIssue() int32 {
	if m != nil && m.Issue != nil {
		return *m.Issue
	}
	return 0
}

func (m *TsPrizePuarchaseReq) GetTimes() int32 {
	if m != nil && m.Times != nil {
		return *m.Times
	}
	return 0
}

// 幸运儿数据请求
type TsLuckyRecordReq struct {
	PrizeId          *int32 `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Start            *int32 `protobuf:"varint,2,req,name=Start" json:"Start,omitempty"`
	End              *int32 `protobuf:"varint,3,req,name=End" json:"End,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsLuckyRecordReq) Reset()                    { *m = TsLuckyRecordReq{} }
func (m *TsLuckyRecordReq) String() string            { return proto.CompactTextString(m) }
func (*TsLuckyRecordReq) ProtoMessage()               {}
func (*TsLuckyRecordReq) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *TsLuckyRecordReq) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsLuckyRecordReq) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *TsLuckyRecordReq) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// 幸运儿数据
type TsLuckyRecord struct {
	PrizeId          *int32         `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Issue            *int32         `protobuf:"varint,2,req,name=Issue" json:"Issue,omitempty"`
	StartTime        *int64         `protobuf:"varint,3,req,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,4,req,name=EndTime" json:"EndTime,omitempty"`
	Remain           *bool          `protobuf:"varint,5,req,name=Remain" json:"Remain,omitempty"`
	LuckyPlayer      *TsLuckyPlayer `protobuf:"bytes,6,opt,name=LuckyPlayer" json:"LuckyPlayer,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TsLuckyRecord) Reset()                    { *m = TsLuckyRecord{} }
func (m *TsLuckyRecord) String() string            { return proto.CompactTextString(m) }
func (*TsLuckyRecord) ProtoMessage()               {}
func (*TsLuckyRecord) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *TsLuckyRecord) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsLuckyRecord) GetIssue() int32 {
	if m != nil && m.Issue != nil {
		return *m.Issue
	}
	return 0
}

func (m *TsLuckyRecord) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *TsLuckyRecord) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *TsLuckyRecord) GetRemain() bool {
	if m != nil && m.Remain != nil {
		return *m.Remain
	}
	return false
}

func (m *TsLuckyRecord) GetLuckyPlayer() *TsLuckyPlayer {
	if m != nil {
		return m.LuckyPlayer
	}
	return nil
}

// 幸运儿返回数据
type TsLuckyRecordResp struct {
	LuckyRecordList  []*TsLuckyRecord `protobuf:"bytes,1,rep,name=LuckyRecordList" json:"LuckyRecordList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *TsLuckyRecordResp) Reset()                    { *m = TsLuckyRecordResp{} }
func (m *TsLuckyRecordResp) String() string            { return proto.CompactTextString(m) }
func (*TsLuckyRecordResp) ProtoMessage()               {}
func (*TsLuckyRecordResp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *TsLuckyRecordResp) GetLuckyRecordList() []*TsLuckyRecord {
	if m != nil {
		return m.LuckyRecordList
	}
	return nil
}

// 购买记录请求
type TsPurchaseRecordReq struct {
	Start            *int32 `protobuf:"varint,1,req,name=Start" json:"Start,omitempty"`
	End              *int32 `protobuf:"varint,2,req,name=End" json:"End,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsPurchaseRecordReq) Reset()                    { *m = TsPurchaseRecordReq{} }
func (m *TsPurchaseRecordReq) String() string            { return proto.CompactTextString(m) }
func (*TsPurchaseRecordReq) ProtoMessage()               {}
func (*TsPurchaseRecordReq) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *TsPurchaseRecordReq) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *TsPurchaseRecordReq) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// 购买记录数据
type TsPurchaseRecord struct {
	PrizeId          *int32         `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Issue            *int32         `protobuf:"varint,2,req,name=Issue" json:"Issue,omitempty"`
	StartTime        *int64         `protobuf:"varint,3,req,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,4,req,name=EndTime" json:"EndTime,omitempty"`
	PropId           *int32         `protobuf:"varint,5,req,name=PropId" json:"PropId,omitempty"`
	PropCout         *int32         `protobuf:"varint,6,req,name=PropCout" json:"PropCout,omitempty"`
	Lucky            *bool          `protobuf:"varint,7,req,name=Lucky" json:"Lucky,omitempty"`
	PurchasedTimes   *int32         `protobuf:"varint,8,req,name=PurchasedTimes" json:"PurchasedTimes,omitempty"`
	Remain           *bool          `protobuf:"varint,9,req,name=Remain" json:"Remain,omitempty"`
	LuckyPlayer      *TsLuckyPlayer `protobuf:"bytes,10,opt,name=LuckyPlayer" json:"LuckyPlayer,omitempty"`
	Proceed          *bool          `protobuf:"varint,11,req,name=proceed" json:"proceed,omitempty"`
	SoldQuota        *int32         `protobuf:"varint,12,opt,name=SoldQuota" json:"SoldQuota,omitempty"`
	OpenNeedQuota    *int32         `protobuf:"varint,13,opt,name=OpenNeedQuota" json:"OpenNeedQuota,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TsPurchaseRecord) Reset()                    { *m = TsPurchaseRecord{} }
func (m *TsPurchaseRecord) String() string            { return proto.CompactTextString(m) }
func (*TsPurchaseRecord) ProtoMessage()               {}
func (*TsPurchaseRecord) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{7} }

func (m *TsPurchaseRecord) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsPurchaseRecord) GetIssue() int32 {
	if m != nil && m.Issue != nil {
		return *m.Issue
	}
	return 0
}

func (m *TsPurchaseRecord) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *TsPurchaseRecord) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *TsPurchaseRecord) GetPropId() int32 {
	if m != nil && m.PropId != nil {
		return *m.PropId
	}
	return 0
}

func (m *TsPurchaseRecord) GetPropCout() int32 {
	if m != nil && m.PropCout != nil {
		return *m.PropCout
	}
	return 0
}

func (m *TsPurchaseRecord) GetLucky() bool {
	if m != nil && m.Lucky != nil {
		return *m.Lucky
	}
	return false
}

func (m *TsPurchaseRecord) GetPurchasedTimes() int32 {
	if m != nil && m.PurchasedTimes != nil {
		return *m.PurchasedTimes
	}
	return 0
}

func (m *TsPurchaseRecord) GetRemain() bool {
	if m != nil && m.Remain != nil {
		return *m.Remain
	}
	return false
}

func (m *TsPurchaseRecord) GetLuckyPlayer() *TsLuckyPlayer {
	if m != nil {
		return m.LuckyPlayer
	}
	return nil
}

func (m *TsPurchaseRecord) GetProceed() bool {
	if m != nil && m.Proceed != nil {
		return *m.Proceed
	}
	return false
}

func (m *TsPurchaseRecord) GetSoldQuota() int32 {
	if m != nil && m.SoldQuota != nil {
		return *m.SoldQuota
	}
	return 0
}

func (m *TsPurchaseRecord) GetOpenNeedQuota() int32 {
	if m != nil && m.OpenNeedQuota != nil {
		return *m.OpenNeedQuota
	}
	return 0
}

// 购买记录返回数据
type TsPurchaseRecordResp struct {
	PurchaseRecordList []*TsPurchaseRecord `protobuf:"bytes,1,rep,name=PurchaseRecordList" json:"PurchaseRecordList,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *TsPurchaseRecordResp) Reset()                    { *m = TsPurchaseRecordResp{} }
func (m *TsPurchaseRecordResp) String() string            { return proto.CompactTextString(m) }
func (*TsPurchaseRecordResp) ProtoMessage()               {}
func (*TsPurchaseRecordResp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{8} }

func (m *TsPurchaseRecordResp) GetPurchaseRecordList() []*TsPurchaseRecord {
	if m != nil {
		return m.PurchaseRecordList
	}
	return nil
}

// 领奖记录请求
type TsPlayerLuckyRecordReq struct {
	Start            *int32 `protobuf:"varint,1,req,name=Start" json:"Start,omitempty"`
	End              *int32 `protobuf:"varint,2,req,name=End" json:"End,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsPlayerLuckyRecordReq) Reset()                    { *m = TsPlayerLuckyRecordReq{} }
func (m *TsPlayerLuckyRecordReq) String() string            { return proto.CompactTextString(m) }
func (*TsPlayerLuckyRecordReq) ProtoMessage()               {}
func (*TsPlayerLuckyRecordReq) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{9} }

func (m *TsPlayerLuckyRecordReq) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *TsPlayerLuckyRecordReq) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// 领奖记录数据
type TsPlayerLuckyRecord struct {
	PrizeId          *int32 `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Issue            *int32 `protobuf:"varint,2,req,name=Issue" json:"Issue,omitempty"`
	StartTime        *int64 `protobuf:"varint,3,req,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64 `protobuf:"varint,4,req,name=EndTime" json:"EndTime,omitempty"`
	PropId           *int32 `protobuf:"varint,5,req,name=PropId" json:"PropId,omitempty"`
	PropCount        *int32 `protobuf:"varint,6,req,name=PropCount" json:"PropCount,omitempty"`
	WinninerNumber   *int32 `protobuf:"varint,7,req,name=WinninerNumber" json:"WinninerNumber,omitempty"`
	PurchasedTimes   *int32 `protobuf:"varint,8,req,name=PurchasedTimes" json:"PurchasedTimes,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsPlayerLuckyRecord) Reset()                    { *m = TsPlayerLuckyRecord{} }
func (m *TsPlayerLuckyRecord) String() string            { return proto.CompactTextString(m) }
func (*TsPlayerLuckyRecord) ProtoMessage()               {}
func (*TsPlayerLuckyRecord) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{10} }

func (m *TsPlayerLuckyRecord) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetIssue() int32 {
	if m != nil && m.Issue != nil {
		return *m.Issue
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetPropId() int32 {
	if m != nil && m.PropId != nil {
		return *m.PropId
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetPropCount() int32 {
	if m != nil && m.PropCount != nil {
		return *m.PropCount
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetWinninerNumber() int32 {
	if m != nil && m.WinninerNumber != nil {
		return *m.WinninerNumber
	}
	return 0
}

func (m *TsPlayerLuckyRecord) GetPurchasedTimes() int32 {
	if m != nil && m.PurchasedTimes != nil {
		return *m.PurchasedTimes
	}
	return 0
}

// 领奖记录数据返回
type TsPlayerLuckyRecordResp struct {
	PlayerLuckyRecord []*TsPlayerLuckyRecord `protobuf:"bytes,1,rep,name=PlayerLuckyRecord" json:"PlayerLuckyRecord,omitempty"`
	XXX_unrecognized  []byte                 `json:"-"`
}

func (m *TsPlayerLuckyRecordResp) Reset()                    { *m = TsPlayerLuckyRecordResp{} }
func (m *TsPlayerLuckyRecordResp) String() string            { return proto.CompactTextString(m) }
func (*TsPlayerLuckyRecordResp) ProtoMessage()               {}
func (*TsPlayerLuckyRecordResp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{11} }

func (m *TsPlayerLuckyRecordResp) GetPlayerLuckyRecord() []*TsPlayerLuckyRecord {
	if m != nil {
		return m.PlayerLuckyRecord
	}
	return nil
}

// 获取玩家夺宝码请求体
type TsPlayerLuckyCodeReq struct {
	PrizeId          *int32 `protobuf:"varint,1,req,name=PrizeId" json:"PrizeId,omitempty"`
	Issue            *int32 `protobuf:"varint,2,req,name=Issue" json:"Issue,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsPlayerLuckyCodeReq) Reset()                    { *m = TsPlayerLuckyCodeReq{} }
func (m *TsPlayerLuckyCodeReq) String() string            { return proto.CompactTextString(m) }
func (*TsPlayerLuckyCodeReq) ProtoMessage()               {}
func (*TsPlayerLuckyCodeReq) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{12} }

func (m *TsPlayerLuckyCodeReq) GetPrizeId() int32 {
	if m != nil && m.PrizeId != nil {
		return *m.PrizeId
	}
	return 0
}

func (m *TsPlayerLuckyCodeReq) GetIssue() int32 {
	if m != nil && m.Issue != nil {
		return *m.Issue
	}
	return 0
}

// 玩家夺宝码返回体
type TsPlayerLuckyCodeResp struct {
	Numbers          []int32 `protobuf:"varint,1,rep,name=Numbers" json:"Numbers,omitempty"`
	RequestTimes     *int32  `protobuf:"varint,2,opt,name=RequestTimes" json:"RequestTimes,omitempty"`
	SuccessTimes     *int32  `protobuf:"varint,3,opt,name=SuccessTimes" json:"SuccessTimes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsPlayerLuckyCodeResp) Reset()                    { *m = TsPlayerLuckyCodeResp{} }
func (m *TsPlayerLuckyCodeResp) String() string            { return proto.CompactTextString(m) }
func (*TsPlayerLuckyCodeResp) ProtoMessage()               {}
func (*TsPlayerLuckyCodeResp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{13} }

func (m *TsPlayerLuckyCodeResp) GetNumbers() []int32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

func (m *TsPlayerLuckyCodeResp) GetRequestTimes() int32 {
	if m != nil && m.RequestTimes != nil {
		return *m.RequestTimes
	}
	return 0
}

func (m *TsPlayerLuckyCodeResp) GetSuccessTimes() int32 {
	if m != nil && m.SuccessTimes != nil {
		return *m.SuccessTimes
	}
	return 0
}

// 中奖玩家
type TsLuckyPlayer struct {
	PlayerId         *int64  `protobuf:"varint,1,req,name=PlayerId" json:"PlayerId,omitempty"`
	NickName         *string `protobuf:"bytes,2,req,name=NickName" json:"NickName,omitempty"`
	Vip              *int32  `protobuf:"varint,3,req,name=Vip" json:"Vip,omitempty"`
	Portrait         *string `protobuf:"bytes,4,req,name=Portrait" json:"Portrait,omitempty"`
	WinninerNumber   *int32  `protobuf:"varint,5,req,name=WinninerNumber" json:"WinninerNumber,omitempty"`
	PurchaseTimes    *int32  `protobuf:"varint,6,opt,name=PurchaseTimes" json:"PurchaseTimes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsLuckyPlayer) Reset()                    { *m = TsLuckyPlayer{} }
func (m *TsLuckyPlayer) String() string            { return proto.CompactTextString(m) }
func (*TsLuckyPlayer) ProtoMessage()               {}
func (*TsLuckyPlayer) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{14} }

func (m *TsLuckyPlayer) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *TsLuckyPlayer) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *TsLuckyPlayer) GetVip() int32 {
	if m != nil && m.Vip != nil {
		return *m.Vip
	}
	return 0
}

func (m *TsLuckyPlayer) GetPortrait() string {
	if m != nil && m.Portrait != nil {
		return *m.Portrait
	}
	return ""
}

func (m *TsLuckyPlayer) GetWinninerNumber() int32 {
	if m != nil && m.WinninerNumber != nil {
		return *m.WinninerNumber
	}
	return 0
}

func (m *TsLuckyPlayer) GetPurchaseTimes() int32 {
	if m != nil && m.PurchaseTimes != nil {
		return *m.PurchaseTimes
	}
	return 0
}

// 中奖玩家
type TsLatelyLuckyPlayer struct {
	PlayerId         *int64  `protobuf:"varint,1,req,name=PlayerId" json:"PlayerId,omitempty"`
	NickName         *string `protobuf:"bytes,2,req,name=NickName" json:"NickName,omitempty"`
	Portrait         *string `protobuf:"bytes,3,req,name=Portrait" json:"Portrait,omitempty"`
	PropId           *int32  `protobuf:"varint,4,req,name=PropId" json:"PropId,omitempty"`
	PropCount        *int32  `protobuf:"varint,5,req,name=PropCount" json:"PropCount,omitempty"`
	OpenTime         *int64  `protobuf:"varint,6,req,name=OpenTime" json:"OpenTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsLatelyLuckyPlayer) Reset()                    { *m = TsLatelyLuckyPlayer{} }
func (m *TsLatelyLuckyPlayer) String() string            { return proto.CompactTextString(m) }
func (*TsLatelyLuckyPlayer) ProtoMessage()               {}
func (*TsLatelyLuckyPlayer) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{15} }

func (m *TsLatelyLuckyPlayer) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *TsLatelyLuckyPlayer) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *TsLatelyLuckyPlayer) GetPortrait() string {
	if m != nil && m.Portrait != nil {
		return *m.Portrait
	}
	return ""
}

func (m *TsLatelyLuckyPlayer) GetPropId() int32 {
	if m != nil && m.PropId != nil {
		return *m.PropId
	}
	return 0
}

func (m *TsLatelyLuckyPlayer) GetPropCount() int32 {
	if m != nil && m.PropCount != nil {
		return *m.PropCount
	}
	return 0
}

func (m *TsLatelyLuckyPlayer) GetOpenTime() int64 {
	if m != nil && m.OpenTime != nil {
		return *m.OpenTime
	}
	return 0
}

// 中奖玩家
type TsLatelyLuckyPlayerResp struct {
	LatelyLuckys     []*TsLatelyLuckyPlayer `protobuf:"bytes,1,rep,name=LatelyLuckys" json:"LatelyLuckys,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TsLatelyLuckyPlayerResp) Reset()                    { *m = TsLatelyLuckyPlayerResp{} }
func (m *TsLatelyLuckyPlayerResp) String() string            { return proto.CompactTextString(m) }
func (*TsLatelyLuckyPlayerResp) ProtoMessage()               {}
func (*TsLatelyLuckyPlayerResp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{16} }

func (m *TsLatelyLuckyPlayerResp) GetLatelyLuckys() []*TsLatelyLuckyPlayer {
	if m != nil {
		return m.LatelyLuckys
	}
	return nil
}

func init() {
	proto.RegisterType((*TsPrize)(nil), "VK.Proto.TsPrize")
	proto.RegisterType((*TsPrizeListResp)(nil), "VK.Proto.TsPrizeListResp")
	proto.RegisterType((*TsPrizePuarchaseReq)(nil), "VK.Proto.TsPrizePuarchaseReq")
	proto.RegisterType((*TsLuckyRecordReq)(nil), "VK.Proto.TsLuckyRecordReq")
	proto.RegisterType((*TsLuckyRecord)(nil), "VK.Proto.TsLuckyRecord")
	proto.RegisterType((*TsLuckyRecordResp)(nil), "VK.Proto.TsLuckyRecordResp")
	proto.RegisterType((*TsPurchaseRecordReq)(nil), "VK.Proto.TsPurchaseRecordReq")
	proto.RegisterType((*TsPurchaseRecord)(nil), "VK.Proto.TsPurchaseRecord")
	proto.RegisterType((*TsPurchaseRecordResp)(nil), "VK.Proto.TsPurchaseRecordResp")
	proto.RegisterType((*TsPlayerLuckyRecordReq)(nil), "VK.Proto.TsPlayerLuckyRecordReq")
	proto.RegisterType((*TsPlayerLuckyRecord)(nil), "VK.Proto.TsPlayerLuckyRecord")
	proto.RegisterType((*TsPlayerLuckyRecordResp)(nil), "VK.Proto.TsPlayerLuckyRecordResp")
	proto.RegisterType((*TsPlayerLuckyCodeReq)(nil), "VK.Proto.TsPlayerLuckyCodeReq")
	proto.RegisterType((*TsPlayerLuckyCodeResp)(nil), "VK.Proto.TsPlayerLuckyCodeResp")
	proto.RegisterType((*TsLuckyPlayer)(nil), "VK.Proto.TsLuckyPlayer")
	proto.RegisterType((*TsLatelyLuckyPlayer)(nil), "VK.Proto.TsLatelyLuckyPlayer")
	proto.RegisterType((*TsLatelyLuckyPlayerResp)(nil), "VK.Proto.TsLatelyLuckyPlayerResp")
	proto.RegisterEnum("VK.Proto.TreasureOps", TreasureOps_name, TreasureOps_value)
	proto.RegisterEnum("VK.Proto.TsLimitType", TsLimitType_name, TsLimitType_value)
	proto.RegisterEnum("VK.Proto.TsOpenType", TsOpenType_name, TsOpenType_value)
	proto.RegisterEnum("VK.Proto.TsIssueStatus", TsIssueStatus_name, TsIssueStatus_value)
}

func init() { proto.RegisterFile("client_treasure.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 1094 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0x9d, 0x38, 0x71, 0x4e, 0x92, 0xd6, 0x99, 0x4d, 0x5b, 0x53, 0x6e, 0xc1, 0xaa, 0x56,
	0x51, 0x91, 0xba, 0x68, 0x79, 0xe2, 0x01, 0x09, 0x28, 0x45, 0x2a, 0x5b, 0x75, 0xc3, 0x24, 0xca,
	0x3e, 0x80, 0xb4, 0xf2, 0x3a, 0x83, 0xb0, 0x36, 0xb1, 0xdd, 0x19, 0x5b, 0x28, 0xbc, 0xf2, 0x27,
	0x78, 0xe1, 0x97, 0xf0, 0x82, 0xc4, 0x1b, 0xbf, 0x0a, 0xcd, 0x99, 0xf1, 0x35, 0x01, 0x5a, 0x84,
	0xe0, 0xcd, 0xe7, 0x3b, 0x17, 0x9f, 0xcb, 0x77, 0x66, 0x06, 0x8e, 0x82, 0x75, 0xc8, 0xa2, 0xf4,
	0x65, 0xca, 0x99, 0x2f, 0x32, 0xce, 0x2e, 0x12, 0x1e, 0xa7, 0x31, 0xb1, 0x97, 0xcf, 0x2e, 0x66,
	0xf2, 0xcb, 0xfb, 0xc9, 0x82, 0xee, 0x42, 0xcc, 0x78, 0xf8, 0x03, 0x23, 0x2e, 0x74, 0xf1, 0xe3,
	0x7a, 0xe5, 0x1a, 0x13, 0x73, 0x6a, 0xd1, 0x5c, 0x24, 0x63, 0xb0, 0xae, 0x85, 0xc8, 0x98, 0x6b,
	0x22, 0xae, 0x04, 0x72, 0x0a, 0xf6, 0x32, 0x4c, 0x6e, 0xc2, 0x4d, 0x98, 0xba, 0x2d, 0x54, 0x14,
	0xb2, 0xd4, 0x5d, 0x66, 0x9c, 0xb3, 0x28, 0xd8, 0xba, 0x6d, 0xa5, 0xcb, 0x65, 0x19, 0x6d, 0xc6,
	0xc3, 0x80, 0xb9, 0x96, 0x8a, 0x86, 0x02, 0xf9, 0x10, 0x7a, 0xe8, 0xba, 0xd8, 0x26, 0xcc, 0xed,
	0x4c, 0xcc, 0xe9, 0xc1, 0xd3, 0xa3, 0x8b, 0x3c, 0xcf, 0x8b, 0x85, 0x28, 0x94, 0xb4, 0xb4, 0x23,
	0x67, 0x30, 0x9c, 0xb3, 0xf5, 0x7a, 0x9e, 0xfa, 0x3c, 0x5d, 0x84, 0x1b, 0xe6, 0x76, 0x27, 0xc6,
	0xb4, 0x45, 0xeb, 0x20, 0x99, 0x40, 0x5f, 0x02, 0x57, 0xd1, 0x0a, 0x6d, 0x6c, 0xb4, 0xa9, 0x42,
	0xe4, 0x03, 0xb0, 0x9f, 0x27, 0x2c, 0xc2, 0x7f, 0xf7, 0xf0, 0xdf, 0xe3, 0xea, 0xbf, 0x73, 0x1d,
	0x2d, 0xac, 0xc8, 0x5b, 0xd0, 0xbb, 0x8c, 0xb3, 0x28, 0xfd, 0x3c, 0xfe, 0x3e, 0x72, 0x61, 0x62,
	0x4c, 0x2d, 0x5a, 0x02, 0xe4, 0x31, 0x1c, 0xcc, 0x32, 0x1e, 0x7c, 0xe7, 0x0b, 0xb6, 0xfa, 0x2a,
	0x8b, 0x53, 0xdf, 0xed, 0x63, 0xad, 0x0d, 0x94, 0xbc, 0x03, 0xb0, 0x88, 0x53, 0x7f, 0xad, 0x6c,
	0x06, 0x68, 0x53, 0x41, 0xc8, 0x31, 0x74, 0x66, 0x3c, 0x4e, 0xae, 0x57, 0xee, 0x10, 0x75, 0x5a,
	0x92, 0x7f, 0x97, 0x5f, 0xf8, 0x43, 0xf7, 0x00, 0x55, 0x25, 0x40, 0x9e, 0x40, 0x67, 0x9e, 0xfa,
	0x69, 0x26, 0xdc, 0x43, 0xac, 0xe5, 0xa4, 0x5a, 0x0b, 0xce, 0x4e, 0xa9, 0xa9, 0x36, 0x23, 0x1f,
	0x41, 0xff, 0x26, 0x0b, 0x5e, 0x6f, 0x67, 0x6b, 0x7f, 0xcb, 0xb8, 0xeb, 0x4c, 0x8c, 0x69, 0xbf,
	0xee, 0x55, 0x51, 0xd3, 0xaa, 0xad, 0xcc, 0x64, 0x1e, 0xaf, 0x75, 0x91, 0x23, 0x95, 0x49, 0x01,
	0xc8, 0xfa, 0x70, 0x58, 0x4a, 0x4d, 0x54, 0x7d, 0x25, 0x22, 0x69, 0xf2, 0xc2, 0x0f, 0x53, 0xd9,
	0x55, 0xf7, 0xd1, 0xc4, 0x9c, 0xda, 0xb4, 0x90, 0xbd, 0xcf, 0xe0, 0x50, 0x33, 0xf3, 0x26, 0x14,
	0x29, 0x65, 0x22, 0x21, 0x4f, 0x64, 0xd9, 0x1a, 0x70, 0x8d, 0x49, 0x6b, 0xda, 0x7f, 0x3a, 0xaa,
	0x66, 0x89, 0x4a, 0x5a, 0xda, 0x78, 0x5f, 0xc3, 0x23, 0x8d, 0xce, 0x32, 0x5f, 0x75, 0x9e, 0xb2,
	0xbb, 0x07, 0x33, 0x7d, 0x0c, 0x96, 0xa4, 0x89, 0xd0, 0x34, 0x57, 0x82, 0xb7, 0x00, 0x47, 0x37,
	0x86, 0xb2, 0x20, 0xe6, 0xab, 0xbf, 0x8d, 0x8c, 0x8c, 0xcc, 0x23, 0xa3, 0x40, 0x1c, 0x68, 0x5d,
	0x45, 0x2b, 0x1d, 0x57, 0x7e, 0x7a, 0xbf, 0x1b, 0x30, 0xac, 0x85, 0x7d, 0x70, 0xb6, 0x72, 0x24,
	0xc5, 0x42, 0xc8, 0xc8, 0x2d, 0x5a, 0x02, 0x32, 0x5a, 0xbe, 0x08, 0x6d, 0xd4, 0xe5, 0xa2, 0x24,
	0x1b, 0x65, 0x1b, 0x3f, 0x8c, 0x70, 0x31, 0x6d, 0xaa, 0xa5, 0x26, 0x3b, 0x3a, 0xf7, 0x67, 0x87,
	0xb7, 0x84, 0x51, 0xa3, 0x45, 0x22, 0x21, 0x9f, 0xc2, 0x61, 0x05, 0xaa, 0xcc, 0x72, 0x37, 0xa6,
	0xf6, 0x6a, 0xda, 0x7b, 0x1f, 0xe3, 0x5c, 0xb3, 0x7c, 0xa2, 0x79, 0xf7, 0x8b, 0x1e, 0x1b, 0x7b,
	0x7a, 0x6c, 0x96, 0x3d, 0xfe, 0xb9, 0x25, 0x47, 0x57, 0xf7, 0xff, 0x2f, 0xdb, 0xac, 0x77, 0xda,
	0xaa, 0xed, 0xf4, 0x29, 0xd8, 0x7a, 0x85, 0x53, 0x3c, 0xff, 0x2c, 0x5a, 0xc8, 0x32, 0x03, 0x6c,
	0x81, 0xdb, 0xc5, 0xc9, 0x28, 0xa1, 0x76, 0xca, 0x28, 0x7e, 0xda, 0x8d, 0x53, 0x06, 0xd1, 0xca,
	0x60, 0x7b, 0x7f, 0x35, 0x58, 0x78, 0xc0, 0xda, 0xbb, 0xd0, 0x4d, 0x78, 0x1c, 0x30, 0xb6, 0xc2,
	0x93, 0xcd, 0xa6, 0xb9, 0x58, 0x3f, 0x10, 0x06, 0xea, 0x60, 0x2c, 0x0f, 0x84, 0x33, 0x18, 0xca,
	0xe5, 0xbe, 0x65, 0xf9, 0xb9, 0x38, 0x44, 0x8b, 0x3a, 0xe8, 0xbd, 0x82, 0xf1, 0xee, 0x78, 0x45,
	0x42, 0xbe, 0x04, 0x52, 0x47, 0x2b, 0xe4, 0x39, 0xad, 0x1d, 0x04, 0x75, 0xdf, 0x3d, 0x5e, 0xde,
	0x27, 0x70, 0xbc, 0x10, 0xaa, 0x9a, 0xc6, 0x0e, 0xdf, 0x97, 0x45, 0x3f, 0x9a, 0xc8, 0xc2, 0x66,
	0x88, 0xff, 0x9d, 0x48, 0xb5, 0xcb, 0xa1, 0xd3, 0xbc, 0x1c, 0x1e, 0xc3, 0xc1, 0x8b, 0x30, 0x8a,
	0xc2, 0x88, 0xf1, 0xdb, 0x6c, 0xf3, 0x8a, 0x71, 0xe4, 0x94, 0x45, 0x1b, 0xe8, 0x7d, 0xc9, 0xe5,
	0x7d, 0x0b, 0x27, 0x7b, 0xfb, 0x28, 0x12, 0xf2, 0x0c, 0x46, 0x3b, 0x0a, 0x3d, 0xad, 0xb7, 0x6b,
	0xd3, 0xda, 0xf1, 0xde, 0xf5, 0xf3, 0xbe, 0x40, 0x4e, 0x94, 0xf0, 0x65, 0xbc, 0xfa, 0x27, 0x67,
	0xb9, 0xb7, 0x85, 0xa3, 0x3d, 0x71, 0x44, 0x22, 0x03, 0xa9, 0xd2, 0x05, 0xe6, 0x68, 0xd1, 0x5c,
	0x24, 0x1e, 0x0c, 0x28, 0xbb, 0xcb, 0x98, 0x48, 0x55, 0x23, 0x4c, 0xe4, 0x6c, 0x0d, 0x93, 0x36,
	0xf3, 0x2c, 0x08, 0x98, 0x10, 0xf9, 0x4d, 0x81, 0x36, 0x55, 0xcc, 0xfb, 0xad, 0x3c, 0xda, 0xf5,
	0x1a, 0xc9, 0x9d, 0xc7, 0x2f, 0x9d, 0x7d, 0x8b, 0x16, 0xb2, 0xd4, 0xdd, 0x86, 0xc1, 0xeb, 0x5b,
	0x7f, 0xa3, 0x2a, 0xe8, 0xd1, 0x42, 0x96, 0x64, 0x5c, 0x86, 0x49, 0x7e, 0x6d, 0x2c, 0xc3, 0x04,
	0x23, 0xc5, 0x3c, 0xe5, 0x7e, 0x98, 0x22, 0x4f, 0x7a, 0xb4, 0x90, 0xf7, 0x8c, 0xdc, 0xda, 0x3b,
	0xf2, 0x33, 0x18, 0xe6, 0xc3, 0x55, 0x45, 0x74, 0xd4, 0x72, 0xd6, 0x40, 0xef, 0x17, 0x43, 0xd2,
	0xfe, 0xc6, 0x4f, 0xd9, 0x7a, 0xfb, 0x6f, 0xd4, 0x52, 0xcd, 0xbc, 0xd5, 0xc8, 0xbc, 0xa4, 0x78,
	0xfb, 0xcf, 0x29, 0x6e, 0x35, 0x29, 0x7e, 0xaa, 0x5f, 0x73, 0x72, 0x67, 0x3a, 0x2a, 0x93, 0x5c,
	0xf6, 0xbe, 0x91, 0x74, 0xdd, 0x49, 0x5e, 0xdf, 0x4b, 0x83, 0x8a, 0x42, 0xec, 0x63, 0xea, 0xae,
	0x63, 0xcd, 0xe5, 0xfc, 0x57, 0x03, 0xfa, 0x0b, 0xfd, 0xd6, 0x7e, 0x9e, 0x08, 0x32, 0x82, 0x21,
	0x65, 0x77, 0x2f, 0x8b, 0x07, 0x89, 0x63, 0x90, 0x63, 0x20, 0x05, 0x54, 0x3c, 0x4a, 0x1c, 0x93,
	0xb8, 0x30, 0x2e, 0x4d, 0x4b, 0xde, 0x3b, 0x2d, 0x72, 0x02, 0x8f, 0x50, 0x53, 0xbe, 0x60, 0x50,
	0xd1, 0x26, 0x6f, 0xc0, 0x11, 0x2a, 0x9a, 0xbb, 0xe2, 0x58, 0x85, 0x4f, 0x9d, 0xe7, 0x4e, 0x87,
	0xbc, 0x0b, 0x6f, 0x4a, 0x85, 0xca, 0x7a, 0xd7, 0xb3, 0x7b, 0xfe, 0x3e, 0xf4, 0x2b, 0x8f, 0x6d,
	0x32, 0x84, 0xde, 0x82, 0x6d, 0x92, 0x98, 0xfb, 0x7c, 0xeb, 0x18, 0x52, 0x9c, 0x31, 0xbe, 0xf1,
	0x23, 0x16, 0xa5, 0x8e, 0x79, 0xfe, 0x1e, 0x40, 0xf9, 0x3a, 0x26, 0x36, 0xb4, 0x65, 0x8f, 0x1d,
	0x83, 0xf4, 0xf4, 0x83, 0xc9, 0x31, 0xcf, 0x33, 0xc9, 0xf9, 0xca, 0xa3, 0x53, 0xf6, 0x04, 0xc5,
	0x9c, 0x55, 0x2a, 0x2a, 0x42, 0x32, 0x92, 0x63, 0x92, 0x01, 0xd8, 0x28, 0x5e, 0x45, 0xb2, 0xfc,
	0x43, 0xe8, 0xa3, 0xa4, 0x2e, 0x2d, 0xa7, 0x4d, 0x1c, 0x18, 0xa8, 0x00, 0x9c, 0x25, 0x3e, 0x67,
	0x8e, 0x55, 0x84, 0xcc, 0xdf, 0x8e, 0x4e, 0xe7, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x67,
	0xa0, 0x7b, 0xfa, 0x0c, 0x00, 0x00,
}
