// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deploy.proto

package deploy

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

type Deploy struct {
	HttpAddr             string   `protobuf:"bytes,1,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deploy) Reset()         { *m = Deploy{} }
func (m *Deploy) String() string { return proto.CompactTextString(m) }
func (*Deploy) ProtoMessage()    {}
func (*Deploy) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{0}
}

func (m *Deploy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deploy.Unmarshal(m, b)
}
func (m *Deploy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deploy.Marshal(b, m, deterministic)
}
func (m *Deploy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deploy.Merge(m, src)
}
func (m *Deploy) XXX_Size() int {
	return xxx_messageInfo_Deploy.Size(m)
}
func (m *Deploy) XXX_DiscardUnknown() {
	xxx_messageInfo_Deploy.DiscardUnknown(m)
}

var xxx_messageInfo_Deploy proto.InternalMessageInfo

func (m *Deploy) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*Deploy)(nil), "deploy.Deploy")
}

func init() { proto.RegisterFile("deploy.proto", fileDescriptor_05f09e103004e384) }

var fileDescriptor_05f09e103004e384 = []byte{
	// 74 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xc8,
	0xc9, 0xaf, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x54, 0xb8, 0xd8,
	0x5c, 0xc0, 0x2c, 0x21, 0x29, 0x2e, 0x8e, 0x8c, 0x92, 0x92, 0x02, 0xc7, 0x94, 0x94, 0x22, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x3f, 0x89, 0x0d, 0xac, 0xc9, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0x0e, 0x30, 0x8f, 0x44, 0x00, 0x00, 0x00,
}