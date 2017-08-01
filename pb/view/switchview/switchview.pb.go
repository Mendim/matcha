// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/switchview/switchview.proto

/*
Package switchview is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/pb/view/switchview/switchview.proto

It has these top-level messages:
	View
	Event
*/
package switchview

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type View struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *View) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Event struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*View)(nil), "switchview.View")
	proto.RegisterType((*Event)(nil), "switchview.Event")
}

func init() {
	proto.RegisterFile("gomatcha.io/matcha/pb/view/switchview/switchview.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x92, 0xf4, 0xcb, 0x32,
	0x53, 0xcb, 0xf5, 0x8b, 0xcb, 0x33, 0x4b, 0x92, 0x33, 0xd0, 0x98, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x19, 0x2e, 0x96, 0xb0, 0xcc, 0xd4, 0x72, 0x21, 0x11, 0x2e,
	0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0x47, 0x49,
	0x96, 0x8b, 0xd5, 0xb5, 0x2c, 0x35, 0xaf, 0x04, 0xbb, 0xb4, 0x93, 0x62, 0x14, 0x92, 0x51, 0x8b,
	0x98, 0x84, 0x7c, 0xc1, 0x6e, 0x08, 0x70, 0x0a, 0x06, 0x0b, 0x82, 0xcc, 0x4d, 0x62, 0x03, 0x5b,
	0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x67, 0x22, 0x28, 0x91, 0xac, 0x00, 0x00, 0x00,
}