// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comand.proto

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

//数据包
type PacketData struct {
	MainID               uint32   `protobuf:"varint,1,opt,name=MainID,proto3" json:"MainID,omitempty"`
	SubID                uint32   `protobuf:"varint,2,opt,name=SubID,proto3" json:"SubID,omitempty"`
	TransData            []byte   `protobuf:"bytes,3,opt,name=TransData,proto3" json:"TransData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketData) Reset()         { *m = PacketData{} }
func (m *PacketData) String() string { return proto.CompactTextString(m) }
func (*PacketData) ProtoMessage()    {}
func (*PacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_960d5946b8165923, []int{0}
}

func (m *PacketData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketData.Unmarshal(m, b)
}
func (m *PacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketData.Marshal(b, m, deterministic)
}
func (m *PacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketData.Merge(m, src)
}
func (m *PacketData) XXX_Size() int {
	return xxx_messageInfo_PacketData.Size(m)
}
func (m *PacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketData.DiscardUnknown(m)
}

var xxx_messageInfo_PacketData proto.InternalMessageInfo

func (m *PacketData) GetMainID() uint32 {
	if m != nil {
		return m.MainID
	}
	return 0
}

func (m *PacketData) GetSubID() uint32 {
	if m != nil {
		return m.SubID
	}
	return 0
}

func (m *PacketData) GetTransData() []byte {
	if m != nil {
		return m.TransData
	}
	return nil
}

func init() {
	proto.RegisterType((*PacketData)(nil), "go.PacketData")
}

func init() { proto.RegisterFile("comand.proto", fileDescriptor_960d5946b8165923) }

var fileDescriptor_960d5946b8165923 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4d,
	0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0xcf, 0x57, 0x8a, 0xe0, 0xe2,
	0x0a, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe3, 0x62, 0xf3, 0x4d,
	0xcc, 0xcc, 0xf3, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0xf2, 0x84, 0x44, 0xb8,
	0x58, 0x83, 0x4b, 0x93, 0x3c, 0x5d, 0x24, 0x98, 0xc0, 0xc2, 0x10, 0x8e, 0x90, 0x0c, 0x17, 0x67,
	0x48, 0x51, 0x62, 0x5e, 0x31, 0x48, 0xab, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x42, 0x20,
	0x89, 0x0d, 0x6c, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x27, 0x05, 0xa1, 0xe3, 0x74, 0x00,
	0x00, 0x00,
}
