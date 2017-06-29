// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meshpb.proto

/*
Package meshpb is a generated protocol buffer package.

It is generated from these files:
	meshpb.proto

It has these top-level messages:
	HandshakeRequest
	HandshakeResult
*/
package meshpb

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

type HandshakeType int32

const (
	HandshakeType_DIAL   HandshakeType = 0
	HandshakeType_LISTEN HandshakeType = 1
)

var HandshakeType_name = map[int32]string{
	0: "DIAL",
	1: "LISTEN",
}
var HandshakeType_value = map[string]int32{
	"DIAL":   0,
	"LISTEN": 1,
}

func (x HandshakeType) String() string {
	return proto.EnumName(HandshakeType_name, int32(x))
}
func (HandshakeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HandshakeStatus int32

const (
	HandshakeStatus_CONNECTED HandshakeStatus = 0
	HandshakeStatus_FAILED    HandshakeStatus = 1
)

var HandshakeStatus_name = map[int32]string{
	0: "CONNECTED",
	1: "FAILED",
}
var HandshakeStatus_value = map[string]int32{
	"CONNECTED": 0,
	"FAILED":    1,
}

func (x HandshakeStatus) String() string {
	return proto.EnumName(HandshakeStatus_name, int32(x))
}
func (HandshakeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type HandshakeRequest struct {
	Type     HandshakeType `protobuf:"varint,1,opt,name=type,enum=HandshakeType" json:"type,omitempty"`
	Endpoint string        `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *HandshakeRequest) Reset()                    { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()               {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HandshakeRequest) GetType() HandshakeType {
	if m != nil {
		return m.Type
	}
	return HandshakeType_DIAL
}

func (m *HandshakeRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type HandshakeResult struct {
	Status HandshakeStatus `protobuf:"varint,1,opt,name=status,enum=HandshakeStatus" json:"status,omitempty"`
}

func (m *HandshakeResult) Reset()                    { *m = HandshakeResult{} }
func (m *HandshakeResult) String() string            { return proto.CompactTextString(m) }
func (*HandshakeResult) ProtoMessage()               {}
func (*HandshakeResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HandshakeResult) GetStatus() HandshakeStatus {
	if m != nil {
		return m.Status
	}
	return HandshakeStatus_CONNECTED
}

func init() {
	proto.RegisterType((*HandshakeRequest)(nil), "HandshakeRequest")
	proto.RegisterType((*HandshakeResult)(nil), "HandshakeResult")
	proto.RegisterEnum("HandshakeType", HandshakeType_name, HandshakeType_value)
	proto.RegisterEnum("HandshakeStatus", HandshakeStatus_name, HandshakeStatus_value)
}

func init() { proto.RegisterFile("meshpb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2d, 0xce,
	0x28, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x0a, 0xe2, 0x12, 0xf0, 0x48, 0xcc, 0x4b,
	0x29, 0xce, 0x48, 0xcc, 0x4e, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe2, 0x62,
	0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xe2, 0xd3, 0x83, 0x2b, 0x08,
	0xa9, 0x2c, 0x48, 0x0d, 0x02, 0xcb, 0x09, 0x49, 0x71, 0x71, 0xa4, 0xe6, 0xa5, 0x14, 0xe4, 0x67,
	0xe6, 0x95, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4a, 0xd6, 0x5c, 0xfc, 0x48,
	0x66, 0x16, 0x97, 0xe6, 0x94, 0x08, 0x69, 0x70, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x43,
	0x0d, 0x15, 0x40, 0x18, 0x1a, 0x0c, 0x16, 0x0f, 0x82, 0xca, 0x6b, 0xa9, 0x72, 0xf1, 0xa2, 0xd8,
	0x27, 0xc4, 0xc1, 0xc5, 0xe2, 0xe2, 0xe9, 0xe8, 0x23, 0xc0, 0x20, 0xc4, 0xc5, 0xc5, 0xe6, 0xe3,
	0x19, 0x1c, 0xe2, 0xea, 0x27, 0xc0, 0xa8, 0xa5, 0x83, 0x64, 0x07, 0xc4, 0x04, 0x21, 0x5e, 0x2e,
	0x4e, 0x67, 0x7f, 0x3f, 0x3f, 0x57, 0xe7, 0x10, 0x57, 0x17, 0x88, 0x6a, 0x37, 0x47, 0x4f, 0x1f,
	0x57, 0x17, 0x01, 0xc6, 0x24, 0x36, 0xb0, 0x67, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x04, 0xdc, 0xba, 0xfc, 0x00, 0x00, 0x00,
}
