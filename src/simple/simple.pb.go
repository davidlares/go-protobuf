// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple/simple.proto

package simple

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

type SimpleMessage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple             bool     `protobuf:"varint,2,opt,name=is_simple,json=isSimple,proto3" json:"is_simple,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SampleList           []int32  `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList,proto3" json:"sample_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3c868e94d57426, []int{0}
}

func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SimpleMessage) GetIsSimple() bool {
	if m != nil {
		return m.IsSimple
	}
	return false
}

func (m *SimpleMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleMessage) GetSampleList() []int32 {
	if m != nil {
		return m.SampleList
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "simple.SimpleMessage")
}

func init() {
	proto.RegisterFile("simple/simple.proto", fileDescriptor_9b3c868e94d57426)
}

var fileDescriptor_9b3c868e94d57426 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x52,
	0x21, 0x17, 0x6f, 0x30, 0x98, 0xe5, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc4, 0xc7, 0xc5,
	0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xcd, 0xc5,
	0x99, 0x59, 0x1c, 0x0f, 0x51, 0x2d, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x91, 0x59, 0x0c,
	0xd1, 0x23, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x66, 0x0b, 0xc9, 0x73, 0x71, 0x17, 0x27, 0x82, 0x64, 0xe3, 0x73, 0x32, 0x8b, 0x4b, 0x24,
	0x58, 0x14, 0x98, 0x35, 0x58, 0x83, 0xb8, 0x20, 0x42, 0x3e, 0x99, 0xc5, 0x25, 0x49, 0x6c, 0x60,
	0x17, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xc3, 0x36, 0xdb, 0x98, 0x00, 0x00, 0x00,
}
