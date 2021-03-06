// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fishLord.proto

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

//【百人类游戏】
//入场 (场景)
type GameFishLordEnter struct {
	AwardAreas           []int32  `protobuf:"varint,1,rep,packed,name=AwardAreas,proto3" json:"AwardAreas,omitempty"`
	Players              []string `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	Countdown            int32    `protobuf:"varint,3,opt,name=Countdown,proto3" json:"Countdown,omitempty"`
	Chips                []int32  `protobuf:"varint,4,rep,packed,name=Chips,proto3" json:"Chips,omitempty"`
	Odds                 []int32  `protobuf:"varint,5,rep,packed,name=Odds,proto3" json:"Odds,omitempty"`
	BankerScore          int64    `protobuf:"varint,6,opt,name=BankerScore,proto3" json:"BankerScore,omitempty"`
	PlayerScore          int64    `protobuf:"varint,7,opt,name=PlayerScore,proto3" json:"PlayerScore,omitempty"`
	Acquire              int64    `protobuf:"varint,8,opt,name=Acquire,proto3" json:"Acquire,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameFishLordEnter) Reset()         { *m = GameFishLordEnter{} }
func (m *GameFishLordEnter) String() string { return proto.CompactTextString(m) }
func (*GameFishLordEnter) ProtoMessage()    {}
func (*GameFishLordEnter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4cc4ee3f27fac1, []int{0}
}

func (m *GameFishLordEnter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameFishLordEnter.Unmarshal(m, b)
}
func (m *GameFishLordEnter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameFishLordEnter.Marshal(b, m, deterministic)
}
func (m *GameFishLordEnter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameFishLordEnter.Merge(m, src)
}
func (m *GameFishLordEnter) XXX_Size() int {
	return xxx_messageInfo_GameFishLordEnter.Size(m)
}
func (m *GameFishLordEnter) XXX_DiscardUnknown() {
	xxx_messageInfo_GameFishLordEnter.DiscardUnknown(m)
}

var xxx_messageInfo_GameFishLordEnter proto.InternalMessageInfo

func (m *GameFishLordEnter) GetAwardAreas() []int32 {
	if m != nil {
		return m.AwardAreas
	}
	return nil
}

func (m *GameFishLordEnter) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *GameFishLordEnter) GetCountdown() int32 {
	if m != nil {
		return m.Countdown
	}
	return 0
}

func (m *GameFishLordEnter) GetChips() []int32 {
	if m != nil {
		return m.Chips
	}
	return nil
}

func (m *GameFishLordEnter) GetOdds() []int32 {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *GameFishLordEnter) GetBankerScore() int64 {
	if m != nil {
		return m.BankerScore
	}
	return 0
}

func (m *GameFishLordEnter) GetPlayerScore() int64 {
	if m != nil {
		return m.PlayerScore
	}
	return 0
}

func (m *GameFishLordEnter) GetAcquire() int64 {
	if m != nil {
		return m.Acquire
	}
	return 0
}

//游戏消息
//游戏中
type GameFishLordPlaying struct {
	BetArea              int32    `protobuf:"varint,1,opt,name=BetArea,proto3" json:"BetArea,omitempty"`
	BetScore             int64    `protobuf:"varint,2,opt,name=BetScore,proto3" json:"BetScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameFishLordPlaying) Reset()         { *m = GameFishLordPlaying{} }
func (m *GameFishLordPlaying) String() string { return proto.CompactTextString(m) }
func (*GameFishLordPlaying) ProtoMessage()    {}
func (*GameFishLordPlaying) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4cc4ee3f27fac1, []int{1}
}

func (m *GameFishLordPlaying) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameFishLordPlaying.Unmarshal(m, b)
}
func (m *GameFishLordPlaying) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameFishLordPlaying.Marshal(b, m, deterministic)
}
func (m *GameFishLordPlaying) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameFishLordPlaying.Merge(m, src)
}
func (m *GameFishLordPlaying) XXX_Size() int {
	return xxx_messageInfo_GameFishLordPlaying.Size(m)
}
func (m *GameFishLordPlaying) XXX_DiscardUnknown() {
	xxx_messageInfo_GameFishLordPlaying.DiscardUnknown(m)
}

var xxx_messageInfo_GameFishLordPlaying proto.InternalMessageInfo

func (m *GameFishLordPlaying) GetBetArea() int32 {
	if m != nil {
		return m.BetArea
	}
	return 0
}

func (m *GameFishLordPlaying) GetBetScore() int64 {
	if m != nil {
		return m.BetScore
	}
	return 0
}

//下注结果
type GameFishLordBetResult struct {
	State                int32    `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	Hints                string   `protobuf:"bytes,2,opt,name=Hints,proto3" json:"Hints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameFishLordBetResult) Reset()         { *m = GameFishLordBetResult{} }
func (m *GameFishLordBetResult) String() string { return proto.CompactTextString(m) }
func (*GameFishLordBetResult) ProtoMessage()    {}
func (*GameFishLordBetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4cc4ee3f27fac1, []int{2}
}

func (m *GameFishLordBetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameFishLordBetResult.Unmarshal(m, b)
}
func (m *GameFishLordBetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameFishLordBetResult.Marshal(b, m, deterministic)
}
func (m *GameFishLordBetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameFishLordBetResult.Merge(m, src)
}
func (m *GameFishLordBetResult) XXX_Size() int {
	return xxx_messageInfo_GameFishLordBetResult.Size(m)
}
func (m *GameFishLordBetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GameFishLordBetResult.DiscardUnknown(m)
}

var xxx_messageInfo_GameFishLordBetResult proto.InternalMessageInfo

func (m *GameFishLordBetResult) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *GameFishLordBetResult) GetHints() string {
	if m != nil {
		return m.Hints
	}
	return ""
}

//结束
type GameFishLordOver struct {
	AwardArea            []int32  `protobuf:"varint,1,rep,packed,name=AwardArea,proto3" json:"AwardArea,omitempty"`
	PlayerCard           []int32  `protobuf:"varint,2,rep,packed,name=PlayerCard,proto3" json:"PlayerCard,omitempty"`
	BankerCard           []int32  `protobuf:"varint,3,rep,packed,name=BankerCard,proto3" json:"BankerCard,omitempty"`
	Acquire              int64    `protobuf:"varint,4,opt,name=Acquire,proto3" json:"Acquire,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameFishLordOver) Reset()         { *m = GameFishLordOver{} }
func (m *GameFishLordOver) String() string { return proto.CompactTextString(m) }
func (*GameFishLordOver) ProtoMessage()    {}
func (*GameFishLordOver) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4cc4ee3f27fac1, []int{3}
}

func (m *GameFishLordOver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameFishLordOver.Unmarshal(m, b)
}
func (m *GameFishLordOver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameFishLordOver.Marshal(b, m, deterministic)
}
func (m *GameFishLordOver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameFishLordOver.Merge(m, src)
}
func (m *GameFishLordOver) XXX_Size() int {
	return xxx_messageInfo_GameFishLordOver.Size(m)
}
func (m *GameFishLordOver) XXX_DiscardUnknown() {
	xxx_messageInfo_GameFishLordOver.DiscardUnknown(m)
}

var xxx_messageInfo_GameFishLordOver proto.InternalMessageInfo

func (m *GameFishLordOver) GetAwardArea() []int32 {
	if m != nil {
		return m.AwardArea
	}
	return nil
}

func (m *GameFishLordOver) GetPlayerCard() []int32 {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *GameFishLordOver) GetBankerCard() []int32 {
	if m != nil {
		return m.BankerCard
	}
	return nil
}

func (m *GameFishLordOver) GetAcquire() int64 {
	if m != nil {
		return m.Acquire
	}
	return 0
}

func init() {
	proto.RegisterType((*GameFishLordEnter)(nil), "go.GameFishLordEnter")
	proto.RegisterType((*GameFishLordPlaying)(nil), "go.GameFishLordPlaying")
	proto.RegisterType((*GameFishLordBetResult)(nil), "go.GameFishLordBetResult")
	proto.RegisterType((*GameFishLordOver)(nil), "go.GameFishLordOver")
}

func init() { proto.RegisterFile("fishLord.proto", fileDescriptor_dd4cc4ee3f27fac1) }

var fileDescriptor_dd4cc4ee3f27fac1 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdd, 0x4a, 0x03, 0x31,
	0x10, 0x85, 0x49, 0xb7, 0xdb, 0x9f, 0x11, 0x44, 0xa3, 0x42, 0x10, 0x91, 0x65, 0xaf, 0xf6, 0xca,
	0x1b, 0x9f, 0xa0, 0xbb, 0xf8, 0x03, 0x0a, 0x95, 0xf4, 0x09, 0x62, 0x33, 0xb6, 0x8b, 0x35, 0xa9,
	0x49, 0x6a, 0xf1, 0x15, 0x7c, 0x5d, 0x5f, 0x40, 0x92, 0xb4, 0xdd, 0x78, 0xb7, 0xe7, 0x9b, 0xc9,
	0x49, 0xe6, 0xcc, 0xc2, 0xf1, 0x5b, 0x6b, 0x97, 0xcf, 0xda, 0xc8, 0x9b, 0xb5, 0xd1, 0x4e, 0xd3,
	0xde, 0x42, 0x97, 0xbf, 0x04, 0x4e, 0x1f, 0xc4, 0x07, 0xde, 0xef, 0x4a, 0x77, 0xca, 0xa1, 0xa1,
	0xd7, 0x00, 0x93, 0xad, 0x30, 0x72, 0x62, 0x50, 0x58, 0x46, 0x8a, 0xac, 0xca, 0x79, 0x42, 0x28,
	0x83, 0xe1, 0x7a, 0x25, 0xbe, 0xd1, 0x58, 0xd6, 0x2b, 0xb2, 0x6a, 0xcc, 0xf7, 0x92, 0x5e, 0xc1,
	0xb8, 0xd1, 0x1b, 0xe5, 0xa4, 0xde, 0x2a, 0x96, 0x15, 0xa4, 0xca, 0x79, 0x07, 0xe8, 0x39, 0xe4,
	0xcd, 0xb2, 0x5d, 0x5b, 0xd6, 0x0f, 0x96, 0x51, 0x50, 0x0a, 0xfd, 0xa9, 0x94, 0x96, 0xe5, 0x01,
	0x86, 0x6f, 0x5a, 0xc0, 0x51, 0x2d, 0xd4, 0x3b, 0x9a, 0xd9, 0x5c, 0x1b, 0x64, 0x83, 0x82, 0x54,
	0x19, 0x4f, 0x91, 0xef, 0x78, 0x09, 0x97, 0xc6, 0x8e, 0x61, 0xec, 0x48, 0x90, 0x7f, 0xe5, 0x64,
	0xfe, 0xb9, 0x69, 0x0d, 0xb2, 0x51, 0xa8, 0xee, 0x65, 0xf9, 0x04, 0x67, 0xe9, 0xd0, 0xfe, 0x50,
	0xab, 0x16, 0xfe, 0x40, 0x8d, 0xce, 0x8f, 0xc8, 0x48, 0x78, 0xfa, 0x5e, 0xd2, 0x4b, 0x18, 0xd5,
	0xe8, 0xe2, 0x4d, 0xbd, 0xe0, 0x75, 0xd0, 0x65, 0x03, 0x17, 0xa9, 0x59, 0x8d, 0x8e, 0xa3, 0xdd,
	0xac, 0x9c, 0x9f, 0x76, 0xe6, 0x84, 0xc3, 0x9d, 0x59, 0x14, 0x9e, 0x3e, 0xb6, 0xca, 0xd9, 0xe0,
	0x33, 0xe6, 0x51, 0x94, 0x3f, 0x04, 0x4e, 0x52, 0x97, 0xe9, 0x17, 0x1a, 0x1f, 0xe6, 0x21, 0xf4,
	0xdd, 0x16, 0x3a, 0xe0, 0x97, 0x14, 0xa7, 0x6d, 0x84, 0x91, 0x61, 0x0f, 0x39, 0x4f, 0x88, 0xaf,
	0xc7, 0xbc, 0x42, 0x3d, 0x8b, 0xf5, 0x8e, 0xa4, 0xf1, 0xf4, 0xff, 0xc5, 0xf3, 0x3a, 0x08, 0xff,
	0xc7, 0xed, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x62, 0x3b, 0x4e, 0x31, 0x02, 0x00, 0x00,
}
