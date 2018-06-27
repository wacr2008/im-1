// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Msg struct {
	Time        string `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	Sid         string `protobuf:"bytes,2,opt,name=sid" json:"sid,omitempty"`
	Id          string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	From        int32  `protobuf:"varint,4,opt,name=from" json:"from,omitempty"`
	To          string `protobuf:"bytes,5,opt,name=to" json:"to,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	ContentType string `protobuf:"bytes,7,opt,name=contentType" json:"contentType,omitempty"`
	Content     string `protobuf:"bytes,8,opt,name=content" json:"content,omitempty"`
	Size        int32  `protobuf:"varint,9,opt,name=size" json:"size,omitempty"`
	W           int32  `protobuf:"varint,10,opt,name=w" json:"w,omitempty"`
	H           int32  `protobuf:"varint,11,opt,name=h" json:"h,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto1.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Msg) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Msg) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *Msg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Msg) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Msg) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Msg) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Msg) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Msg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Msg) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Msg) GetW() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Msg) GetH() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func init() {
	proto1.RegisterType((*Msg)(nil), "proto.msg")
}

func init() { proto1.RegisterFile("msg.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcd, 0x0e, 0x82, 0x30,
	0x10, 0x84, 0x53, 0x7e, 0x65, 0x31, 0xc6, 0xec, 0x69, 0x8f, 0xc4, 0x13, 0x27, 0x2f, 0xbe, 0x09,
	0xf1, 0x09, 0x94, 0x0a, 0x3d, 0x94, 0x12, 0xda, 0x84, 0xe0, 0xcb, 0xfa, 0x2a, 0x66, 0xb7, 0x98,
	0x78, 0xea, 0xcc, 0x37, 0xd3, 0x69, 0x0a, 0x95, 0xf5, 0xc3, 0x75, 0x5e, 0x5c, 0x70, 0x98, 0xcb,
	0x71, 0xf9, 0x28, 0x48, 0xad, 0x1f, 0x10, 0x21, 0x0b, 0xc6, 0x6a, 0x52, 0x8d, 0x6a, 0xab, 0x4e,
	0x34, 0x9e, 0x21, 0xf5, 0xa6, 0xa7, 0x44, 0x10, 0x4b, 0x3c, 0x41, 0x62, 0x7a, 0x4a, 0x05, 0x24,
	0xa6, 0xe7, 0x5b, 0xaf, 0xc5, 0x59, 0xca, 0x1a, 0xd5, 0xe6, 0x9d, 0x68, 0xee, 0x04, 0x47, 0x79,
	0xec, 0x04, 0x27, 0xcb, 0xdb, 0xac, 0xa9, 0xd8, 0x97, 0xb7, 0x59, 0x63, 0x03, 0xf5, 0xd3, 0x4d,
	0x41, 0x4f, 0xe1, 0xce, 0x51, 0x29, 0xd1, 0x3f, 0x42, 0x82, 0x72, 0xb7, 0x74, 0x90, 0xf4, 0x67,
	0x79, 0xcf, 0x9b, 0xb7, 0xa6, 0x2a, 0xbe, 0xc9, 0x1a, 0x8f, 0xa0, 0x56, 0x02, 0x01, 0x6a, 0x65,
	0x37, 0x52, 0x1d, 0xdd, 0xf8, 0x28, 0xe4, 0xa3, 0xb7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x67, 0xba, 0xf4, 0xfc, 0x00, 0x00, 0x00,
}
