// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamecomm.proto

package _go

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//玩家信息
type PlayerInfo struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Sex                  int32    `protobuf:"varint,4,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Gold                 int64    `protobuf:"varint,5,opt,name=Gold,proto3" json:"Gold,omitempty"`
	VipLevel             int32    `protobuf:"varint,6,opt,name=VipLevel,proto3" json:"VipLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{0}
}

func (m *PlayerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerInfo.Unmarshal(m, b)
}
func (m *PlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerInfo.Marshal(b, m, deterministic)
}
func (m *PlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerInfo.Merge(m, src)
}
func (m *PlayerInfo) XXX_Size() int {
	return xxx_messageInfo_PlayerInfo.Size(m)
}
func (m *PlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerInfo proto.InternalMessageInfo

func (m *PlayerInfo) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *PlayerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlayerInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PlayerInfo) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *PlayerInfo) GetGold() int64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *PlayerInfo) GetVipLevel() int32 {
	if m != nil {
		return m.VipLevel
	}
	return 0
}

//玩家列表
type UserList struct {
	AllInfos             []*PlayerInfo `protobuf:"bytes,1,rep,name=AllInfos,proto3" json:"AllInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetAllInfos() []*PlayerInfo {
	if m != nil {
		return m.AllInfos
	}
	return nil
}

//玩家记录（从数据库中获取）
type PlayerRecord struct {
	User                 *PlayerInfo `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Twice                int32       `protobuf:"varint,2,opt,name=Twice,proto3" json:"Twice,omitempty"`
	Ranking              int32       `protobuf:"varint,3,opt,name=Ranking,proto3" json:"Ranking,omitempty"`
	Bankroll             int32       `protobuf:"varint,4,opt,name=Bankroll,proto3" json:"Bankroll,omitempty"`
	WinLos               int64       `protobuf:"varint,5,opt,name=WinLos,proto3" json:"WinLos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PlayerRecord) Reset()         { *m = PlayerRecord{} }
func (m *PlayerRecord) String() string { return proto.CompactTextString(m) }
func (*PlayerRecord) ProtoMessage()    {}
func (*PlayerRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{2}
}

func (m *PlayerRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerRecord.Unmarshal(m, b)
}
func (m *PlayerRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerRecord.Marshal(b, m, deterministic)
}
func (m *PlayerRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerRecord.Merge(m, src)
}
func (m *PlayerRecord) XXX_Size() int {
	return xxx_messageInfo_PlayerRecord.Size(m)
}
func (m *PlayerRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerRecord proto.InternalMessageInfo

func (m *PlayerRecord) GetUser() *PlayerInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PlayerRecord) GetTwice() int32 {
	if m != nil {
		return m.Twice
	}
	return 0
}

func (m *PlayerRecord) GetRanking() int32 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func (m *PlayerRecord) GetBankroll() int32 {
	if m != nil {
		return m.Bankroll
	}
	return 0
}

func (m *PlayerRecord) GetWinLos() int64 {
	if m != nil {
		return m.WinLos
	}
	return 0
}

type GameReady struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IsReady              bool     `protobuf:"varint,2,opt,name=IsReady,proto3" json:"IsReady,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameReady) Reset()         { *m = GameReady{} }
func (m *GameReady) String() string { return proto.CompactTextString(m) }
func (*GameReady) ProtoMessage()    {}
func (*GameReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{3}
}

func (m *GameReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameReady.Unmarshal(m, b)
}
func (m *GameReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameReady.Marshal(b, m, deterministic)
}
func (m *GameReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameReady.Merge(m, src)
}
func (m *GameReady) XXX_Size() int {
	return xxx_messageInfo_GameReady.Size(m)
}
func (m *GameReady) XXX_DiscardUnknown() {
	xxx_messageInfo_GameReady.DiscardUnknown(m)
}

var xxx_messageInfo_GameReady proto.InternalMessageInfo

func (m *GameReady) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GameReady) GetIsReady() bool {
	if m != nil {
		return m.IsReady
	}
	return false
}

//下注
type GameBet struct {
	BetArea              int32    `protobuf:"varint,1,opt,name=BetArea,proto3" json:"BetArea,omitempty"`
	BetScore             int64    `protobuf:"varint,2,opt,name=BetScore,proto3" json:"BetScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameBet) Reset()         { *m = GameBet{} }
func (m *GameBet) String() string { return proto.CompactTextString(m) }
func (*GameBet) ProtoMessage()    {}
func (*GameBet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{4}
}

func (m *GameBet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameBet.Unmarshal(m, b)
}
func (m *GameBet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameBet.Marshal(b, m, deterministic)
}
func (m *GameBet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameBet.Merge(m, src)
}
func (m *GameBet) XXX_Size() int {
	return xxx_messageInfo_GameBet.Size(m)
}
func (m *GameBet) XXX_DiscardUnknown() {
	xxx_messageInfo_GameBet.DiscardUnknown(m)
}

var xxx_messageInfo_GameBet proto.InternalMessageInfo

func (m *GameBet) GetBetArea() int32 {
	if m != nil {
		return m.BetArea
	}
	return 0
}

func (m *GameBet) GetBetScore() int64 {
	if m != nil {
		return m.BetScore
	}
	return 0
}

//下注结果
type GameBetResult struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	State                int32    `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	Hints                string   `protobuf:"bytes,3,opt,name=Hints,proto3" json:"Hints,omitempty"`
	BetArea              int32    `protobuf:"varint,4,opt,name=BetArea,proto3" json:"BetArea,omitempty"`
	BetScore             int64    `protobuf:"varint,5,opt,name=BetScore,proto3" json:"BetScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameBetResult) Reset()         { *m = GameBetResult{} }
func (m *GameBetResult) String() string { return proto.CompactTextString(m) }
func (*GameBetResult) ProtoMessage()    {}
func (*GameBetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{5}
}

func (m *GameBetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameBetResult.Unmarshal(m, b)
}
func (m *GameBetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameBetResult.Marshal(b, m, deterministic)
}
func (m *GameBetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameBetResult.Merge(m, src)
}
func (m *GameBetResult) XXX_Size() int {
	return xxx_messageInfo_GameBetResult.Size(m)
}
func (m *GameBetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GameBetResult.DiscardUnknown(m)
}

var xxx_messageInfo_GameBetResult proto.InternalMessageInfo

func (m *GameBetResult) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GameBetResult) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *GameBetResult) GetHints() string {
	if m != nil {
		return m.Hints
	}
	return ""
}

func (m *GameBetResult) GetBetArea() int32 {
	if m != nil {
		return m.BetArea
	}
	return 0
}

func (m *GameBetResult) GetBetScore() int64 {
	if m != nil {
		return m.BetScore
	}
	return 0
}

//抢庄
type GameHost struct {
	IsWant               bool     `protobuf:"varint,1,opt,name=IsWant,proto3" json:"IsWant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameHost) Reset()         { *m = GameHost{} }
func (m *GameHost) String() string { return proto.CompactTextString(m) }
func (*GameHost) ProtoMessage()    {}
func (*GameHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{6}
}

func (m *GameHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameHost.Unmarshal(m, b)
}
func (m *GameHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameHost.Marshal(b, m, deterministic)
}
func (m *GameHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameHost.Merge(m, src)
}
func (m *GameHost) XXX_Size() int {
	return xxx_messageInfo_GameHost.Size(m)
}
func (m *GameHost) XXX_DiscardUnknown() {
	xxx_messageInfo_GameHost.DiscardUnknown(m)
}

var xxx_messageInfo_GameHost proto.InternalMessageInfo

func (m *GameHost) GetIsWant() bool {
	if m != nil {
		return m.IsWant
	}
	return false
}

//超级抢庄
type GameSuperHost struct {
	IsWant               bool     `protobuf:"varint,1,opt,name=IsWant,proto3" json:"IsWant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameSuperHost) Reset()         { *m = GameSuperHost{} }
func (m *GameSuperHost) String() string { return proto.CompactTextString(m) }
func (*GameSuperHost) ProtoMessage()    {}
func (*GameSuperHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{7}
}

func (m *GameSuperHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameSuperHost.Unmarshal(m, b)
}
func (m *GameSuperHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameSuperHost.Marshal(b, m, deterministic)
}
func (m *GameSuperHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameSuperHost.Merge(m, src)
}
func (m *GameSuperHost) XXX_Size() int {
	return xxx_messageInfo_GameSuperHost.Size(m)
}
func (m *GameSuperHost) XXX_DiscardUnknown() {
	xxx_messageInfo_GameSuperHost.DiscardUnknown(m)
}

var xxx_messageInfo_GameSuperHost proto.InternalMessageInfo

func (m *GameSuperHost) GetIsWant() bool {
	if m != nil {
		return m.IsWant
	}
	return false
}

//游戏记录
type GameRecord struct {
	Pork                 []byte   `protobuf:"bytes,1,opt,name=Pork,proto3" json:"Pork,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	IsWon                bool     `protobuf:"varint,3,opt,name=isWon,proto3" json:"isWon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameRecord) Reset()         { *m = GameRecord{} }
func (m *GameRecord) String() string { return proto.CompactTextString(m) }
func (*GameRecord) ProtoMessage()    {}
func (*GameRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{8}
}

func (m *GameRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameRecord.Unmarshal(m, b)
}
func (m *GameRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameRecord.Marshal(b, m, deterministic)
}
func (m *GameRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameRecord.Merge(m, src)
}
func (m *GameRecord) XXX_Size() int {
	return xxx_messageInfo_GameRecord.Size(m)
}
func (m *GameRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_GameRecord.DiscardUnknown(m)
}

var xxx_messageInfo_GameRecord proto.InternalMessageInfo

func (m *GameRecord) GetPork() []byte {
	if m != nil {
		return m.Pork
	}
	return nil
}

func (m *GameRecord) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GameRecord) GetIsWon() bool {
	if m != nil {
		return m.IsWon
	}
	return false
}

//历史记录(保存十二条记录)
type GameRecordList struct {
	List                 []*GameRecord `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameRecordList) Reset()         { *m = GameRecordList{} }
func (m *GameRecordList) String() string { return proto.CompactTextString(m) }
func (*GameRecordList) ProtoMessage()    {}
func (*GameRecordList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{9}
}

func (m *GameRecordList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameRecordList.Unmarshal(m, b)
}
func (m *GameRecordList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameRecordList.Marshal(b, m, deterministic)
}
func (m *GameRecordList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameRecordList.Merge(m, src)
}
func (m *GameRecordList) XXX_Size() int {
	return xxx_messageInfo_GameRecordList.Size(m)
}
func (m *GameRecordList) XXX_DiscardUnknown() {
	xxx_messageInfo_GameRecordList.DiscardUnknown(m)
}

var xxx_messageInfo_GameRecordList proto.InternalMessageInfo

func (m *GameRecordList) GetList() []*GameRecord {
	if m != nil {
		return m.List
	}
	return nil
}

//反馈结果
type GameResult struct {
	Flag                 int32    `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
	Reason               []byte   `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameResult) Reset()         { *m = GameResult{} }
func (m *GameResult) String() string { return proto.CompactTextString(m) }
func (*GameResult) ProtoMessage()    {}
func (*GameResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0affc1bc90b8bda7, []int{10}
}

func (m *GameResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameResult.Unmarshal(m, b)
}
func (m *GameResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameResult.Marshal(b, m, deterministic)
}
func (m *GameResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameResult.Merge(m, src)
}
func (m *GameResult) XXX_Size() int {
	return xxx_messageInfo_GameResult.Size(m)
}
func (m *GameResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GameResult.DiscardUnknown(m)
}

var xxx_messageInfo_GameResult proto.InternalMessageInfo

func (m *GameResult) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *GameResult) GetReason() []byte {
	if m != nil {
		return m.Reason
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerInfo)(nil), "go.PlayerInfo")
	proto.RegisterType((*UserList)(nil), "go.UserList")
	proto.RegisterType((*PlayerRecord)(nil), "go.PlayerRecord")
	proto.RegisterType((*GameReady)(nil), "go.GameReady")
	proto.RegisterType((*GameBet)(nil), "go.GameBet")
	proto.RegisterType((*GameBetResult)(nil), "go.GameBetResult")
	proto.RegisterType((*GameHost)(nil), "go.GameHost")
	proto.RegisterType((*GameSuperHost)(nil), "go.GameSuperHost")
	proto.RegisterType((*GameRecord)(nil), "go.GameRecord")
	proto.RegisterType((*GameRecordList)(nil), "go.GameRecordList")
	proto.RegisterType((*GameResult)(nil), "go.GameResult")
}

func init() { proto.RegisterFile("gamecomm.proto", fileDescriptor_0affc1bc90b8bda7) }

var fileDescriptor_0affc1bc90b8bda7 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5d, 0x8b, 0xd3, 0x4c,
	0x14, 0xc7, 0xc9, 0x36, 0x69, 0xd3, 0xf3, 0xf4, 0x29, 0x32, 0x2c, 0x12, 0xbc, 0x2a, 0x73, 0x63,
	0xf0, 0xa2, 0x17, 0x2a, 0xe2, 0x8d, 0x48, 0x8b, 0xb8, 0x5b, 0x29, 0xb2, 0x9c, 0xac, 0xf6, 0x7a,
	0x6c, 0x8f, 0x21, 0x74, 0x92, 0x29, 0x99, 0x59, 0xdd, 0x7e, 0x04, 0xc1, 0x6b, 0x3f, 0xaf, 0xcc,
	0xc9, 0x4b, 0x17, 0xa5, 0xde, 0x9d, 0xff, 0x79, 0x99, 0xfc, 0xfe, 0x67, 0x26, 0x30, 0xcd, 0x55,
	0x49, 0x5b, 0x53, 0x96, 0xf3, 0x43, 0x6d, 0x9c, 0x11, 0x17, 0xb9, 0x91, 0x3f, 0x03, 0x80, 0x1b,
	0xad, 0x8e, 0x54, 0xaf, 0xaa, 0xaf, 0x46, 0x3c, 0x86, 0xe1, 0x27, 0x4b, 0xf5, 0xea, 0x5d, 0x12,
	0xcc, 0x82, 0x34, 0xc4, 0x56, 0x09, 0x01, 0xe1, 0x47, 0x55, 0x52, 0x72, 0x31, 0x0b, 0xd2, 0x31,
	0x72, 0x2c, 0x1e, 0xc1, 0x60, 0x91, 0x53, 0x32, 0x98, 0x05, 0x69, 0x84, 0x3e, 0xf4, 0x99, 0x8c,
	0xee, 0x93, 0xb0, 0xc9, 0x64, 0x74, 0xef, 0xe7, 0xae, 0x8c, 0xde, 0x25, 0xd1, 0x2c, 0x48, 0x07,
	0xc8, 0xb1, 0x78, 0x02, 0xf1, 0xe7, 0xe2, 0xb0, 0xa6, 0x6f, 0xa4, 0x93, 0x21, 0xb7, 0xf6, 0x5a,
	0xbe, 0x82, 0xd8, 0x7f, 0x71, 0x5d, 0x58, 0x27, 0x9e, 0x41, 0xbc, 0xd0, 0xda, 0x63, 0xd9, 0x24,
	0x98, 0x0d, 0xd2, 0xff, 0x9e, 0x4f, 0xe7, 0xb9, 0x99, 0x9f, 0x68, 0xb1, 0xaf, 0xcb, 0x5f, 0x01,
	0x4c, 0x9a, 0x02, 0xd2, 0xd6, 0xd4, 0x3b, 0x21, 0x21, 0xf4, 0x07, 0xb1, 0x8d, 0xbf, 0x07, 0xb9,
	0x26, 0x2e, 0x21, 0xba, 0xfd, 0x5e, 0x6c, 0x1b, 0x57, 0x11, 0x36, 0x42, 0x24, 0x30, 0x42, 0x55,
	0xed, 0x8b, 0x2a, 0x6f, 0xad, 0x75, 0xd2, 0x83, 0x2f, 0x55, 0xb5, 0xaf, 0x8d, 0xd6, 0xad, 0xc7,
	0x5e, 0xfb, 0xc5, 0x6d, 0x8a, 0x6a, 0x6d, 0x6c, 0x6b, 0xb5, 0x55, 0xf2, 0x0d, 0x8c, 0xaf, 0x54,
	0x49, 0x48, 0x6a, 0x77, 0x3c, 0xbb, 0xdd, 0x04, 0x46, 0x2b, 0xcb, 0x2d, 0x8c, 0x12, 0x63, 0x27,
	0xe5, 0x5b, 0x18, 0xf9, 0xf1, 0x25, 0x39, 0xdf, 0xb4, 0x24, 0xb7, 0xa8, 0x49, 0xf1, 0x74, 0x84,
	0x9d, 0x64, 0x2e, 0x72, 0xd9, 0xd6, 0xd4, 0x8d, 0x95, 0x01, 0xf6, 0x5a, 0xfe, 0x08, 0xe0, 0xff,
	0xf6, 0x04, 0x24, 0x7b, 0xa7, 0xdd, 0x59, 0x88, 0x4b, 0x88, 0x32, 0xa7, 0x5c, 0xbf, 0x0d, 0x16,
	0x3e, 0x7b, 0x5d, 0x54, 0xce, 0xf2, 0x2e, 0xc6, 0xd8, 0x88, 0x87, 0x2c, 0xe1, 0x79, 0x96, 0xe8,
	0x0f, 0x16, 0x09, 0xb1, 0x47, 0xb9, 0x36, 0x96, 0x29, 0x56, 0x76, 0xa3, 0x2a, 0xc7, 0x14, 0x31,
	0xb6, 0x4a, 0x3e, 0x6d, 0x70, 0xb3, 0xbb, 0x03, 0xd5, 0xff, 0x6c, 0xfc, 0x00, 0xd0, 0x2c, 0x96,
	0xaf, 0x5b, 0x40, 0x78, 0x63, 0xea, 0x3d, 0xf7, 0x4c, 0x90, 0x63, 0x9f, 0xbb, 0x3d, 0x1e, 0x3a,
	0x3f, 0x1c, 0x7b, 0x3b, 0x85, 0xdd, 0x98, 0x8a, 0xed, 0xc4, 0xd8, 0x08, 0xf9, 0x12, 0xa6, 0xa7,
	0xb3, 0xf8, 0xed, 0x49, 0x08, 0x75, 0x61, 0xdd, 0xc3, 0x77, 0x77, 0xea, 0x40, 0xae, 0xc9, 0xd7,
	0x1d, 0x01, 0xaf, 0x55, 0x40, 0xf8, 0x5e, 0xab, 0xbc, 0xbd, 0x1b, 0x8e, 0x3d, 0x3b, 0x92, 0xb2,
	0xa6, 0x62, 0x86, 0x09, 0xb6, 0xea, 0xcb, 0x90, 0xff, 0xbf, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x9c, 0x47, 0x39, 0xe3, 0x91, 0x03, 0x00, 0x00,
}
