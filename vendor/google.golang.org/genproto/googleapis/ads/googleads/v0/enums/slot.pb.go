// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/slot.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enumerates possible positions of the Ad.
type SlotEnum_Slot int32

const (
	// Not specified.
	SlotEnum_UNSPECIFIED SlotEnum_Slot = 0
	// The value is unknown in this version.
	SlotEnum_UNKNOWN SlotEnum_Slot = 1
	// Google search: Side.
	SlotEnum_SEARCH_SIDE SlotEnum_Slot = 2
	// Google search: Top.
	SlotEnum_SEARCH_TOP SlotEnum_Slot = 3
	// Google search: Other.
	SlotEnum_SEARCH_OTHER SlotEnum_Slot = 4
	// Google Display Network.
	SlotEnum_CONTENT SlotEnum_Slot = 5
	// Search partners: Top.
	SlotEnum_SEARCH_PARTNER_TOP SlotEnum_Slot = 6
	// Search partners: Other.
	SlotEnum_SEARCH_PARTNER_OTHER SlotEnum_Slot = 7
	// Cross-network.
	SlotEnum_MIXED SlotEnum_Slot = 8
)

var SlotEnum_Slot_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH_SIDE",
	3: "SEARCH_TOP",
	4: "SEARCH_OTHER",
	5: "CONTENT",
	6: "SEARCH_PARTNER_TOP",
	7: "SEARCH_PARTNER_OTHER",
	8: "MIXED",
}
var SlotEnum_Slot_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"SEARCH_SIDE":          2,
	"SEARCH_TOP":           3,
	"SEARCH_OTHER":         4,
	"CONTENT":              5,
	"SEARCH_PARTNER_TOP":   6,
	"SEARCH_PARTNER_OTHER": 7,
	"MIXED":                8,
}

func (x SlotEnum_Slot) String() string {
	return proto.EnumName(SlotEnum_Slot_name, int32(x))
}
func (SlotEnum_Slot) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_slot_140eee701bab0100, []int{0, 0}
}

// Container for enumeration of possible positions of the Ad.
type SlotEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlotEnum) Reset()         { *m = SlotEnum{} }
func (m *SlotEnum) String() string { return proto.CompactTextString(m) }
func (*SlotEnum) ProtoMessage()    {}
func (*SlotEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_slot_140eee701bab0100, []int{0}
}
func (m *SlotEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlotEnum.Unmarshal(m, b)
}
func (m *SlotEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlotEnum.Marshal(b, m, deterministic)
}
func (dst *SlotEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlotEnum.Merge(dst, src)
}
func (m *SlotEnum) XXX_Size() int {
	return xxx_messageInfo_SlotEnum.Size(m)
}
func (m *SlotEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SlotEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SlotEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SlotEnum)(nil), "google.ads.googleads.v0.enums.SlotEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.SlotEnum_Slot", SlotEnum_Slot_name, SlotEnum_Slot_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/slot.proto", fileDescriptor_slot_140eee701bab0100)
}

var fileDescriptor_slot_140eee701bab0100 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xa4, 0xff, 0xa7, 0x1f, 0x3a, 0x0c, 0x22, 0x6e, 0xba, 0x68, 0x57, 0xae, 0x26,
	0x01, 0x77, 0xe3, 0x6a, 0xda, 0x8e, 0x6d, 0x10, 0x27, 0x21, 0x49, 0xab, 0x48, 0x40, 0xaa, 0x29,
	0x41, 0x48, 0x33, 0xa5, 0xd3, 0xf6, 0x7d, 0x74, 0xe9, 0xc2, 0x07, 0xf1, 0x21, 0x7c, 0x00, 0x9f,
	0x42, 0x66, 0x26, 0xcd, 0x42, 0xd0, 0x4d, 0x38, 0xf7, 0xfe, 0xce, 0x09, 0x73, 0x0f, 0x38, 0xcf,
	0x84, 0xc8, 0xf2, 0xa5, 0xb3, 0x48, 0xa5, 0x63, 0xa4, 0x52, 0x7b, 0xd7, 0x59, 0x16, 0xbb, 0x95,
	0x74, 0x64, 0x2e, 0xb6, 0x78, 0xbd, 0x11, 0x5b, 0x81, 0x7a, 0x06, 0xe3, 0x45, 0x2a, 0x71, 0xe5,
	0xc4, 0x7b, 0x17, 0x6b, 0xe7, 0xe0, 0xdd, 0x02, 0xed, 0x28, 0x17, 0x5b, 0x56, 0xec, 0x56, 0x83,
	0x17, 0x0b, 0xd4, 0xd5, 0x80, 0x8e, 0x41, 0x77, 0xc6, 0xa3, 0x80, 0x8d, 0xbc, 0x2b, 0x8f, 0x8d,
	0xe1, 0x3f, 0xd4, 0x05, 0xad, 0x19, 0xbf, 0xe6, 0xfe, 0x2d, 0x87, 0x96, 0xa2, 0x11, 0xa3, 0xe1,
	0x68, 0xfa, 0x10, 0x79, 0x63, 0x06, 0x6d, 0x74, 0x04, 0x40, 0xb9, 0x88, 0xfd, 0x00, 0xd6, 0x10,
	0x04, 0xff, 0xcb, 0xd9, 0x8f, 0xa7, 0x2c, 0x84, 0x75, 0x95, 0x1f, 0xf9, 0x3c, 0x66, 0x3c, 0x86,
	0x0d, 0x74, 0x0a, 0x50, 0x89, 0x03, 0x1a, 0xc6, 0x9c, 0x85, 0x3a, 0xd6, 0x44, 0x67, 0xe0, 0xe4,
	0xc7, 0xde, 0xc4, 0x5b, 0xa8, 0x03, 0x1a, 0x37, 0xde, 0x1d, 0x1b, 0xc3, 0xf6, 0xf0, 0xd3, 0x02,
	0xfd, 0x27, 0xb1, 0xc2, 0x7f, 0x9e, 0x35, 0xec, 0xa8, 0x33, 0x02, 0x55, 0x40, 0x60, 0xdd, 0x0f,
	0x4b, 0x6f, 0x26, 0xf2, 0x45, 0x91, 0x61, 0xb1, 0xc9, 0x9c, 0x6c, 0x59, 0xe8, 0x7a, 0x0e, 0xe5,
	0xad, 0x9f, 0xe5, 0x2f, 0x5d, 0x5e, 0xea, 0xef, 0xab, 0x5d, 0x9b, 0x50, 0xfa, 0x66, 0xf7, 0x26,
	0xe6, 0x57, 0x34, 0x95, 0xd8, 0x48, 0xa5, 0xe6, 0x2e, 0x56, 0xfd, 0xc9, 0x8f, 0x03, 0x4f, 0x68,
	0x2a, 0x93, 0x8a, 0x27, 0x73, 0x37, 0xd1, 0xfc, 0xcb, 0xee, 0x9b, 0x25, 0x21, 0x34, 0x95, 0x84,
	0x54, 0x0e, 0x42, 0xe6, 0x2e, 0x21, 0xda, 0xf3, 0xd8, 0xd4, 0x0f, 0xbb, 0xf8, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0x0f, 0xcd, 0x52, 0xe3, 0x01, 0x00, 0x00,
}