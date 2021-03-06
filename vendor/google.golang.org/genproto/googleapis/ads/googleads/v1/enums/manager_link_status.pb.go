// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/manager_link_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible statuses of a link.
type ManagerLinkStatusEnum_ManagerLinkStatus int32

const (
	// Not specified.
	ManagerLinkStatusEnum_UNSPECIFIED ManagerLinkStatusEnum_ManagerLinkStatus = 0
	// Used for return value only. Represents value unknown in this version.
	ManagerLinkStatusEnum_UNKNOWN ManagerLinkStatusEnum_ManagerLinkStatus = 1
	// Indicates current in-effect relationship
	ManagerLinkStatusEnum_ACTIVE ManagerLinkStatusEnum_ManagerLinkStatus = 2
	// Indicates terminated relationship
	ManagerLinkStatusEnum_INACTIVE ManagerLinkStatusEnum_ManagerLinkStatus = 3
	// Indicates relationship has been requested by manager, but the client
	// hasn't accepted yet.
	ManagerLinkStatusEnum_PENDING ManagerLinkStatusEnum_ManagerLinkStatus = 4
	// Relationship was requested by the manager, but the client has refused.
	ManagerLinkStatusEnum_REFUSED ManagerLinkStatusEnum_ManagerLinkStatus = 5
	// Indicates relationship has been requested by manager, but manager
	// canceled it.
	ManagerLinkStatusEnum_CANCELED ManagerLinkStatusEnum_ManagerLinkStatus = 6
)

var ManagerLinkStatusEnum_ManagerLinkStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ACTIVE",
	3: "INACTIVE",
	4: "PENDING",
	5: "REFUSED",
	6: "CANCELED",
}

var ManagerLinkStatusEnum_ManagerLinkStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ACTIVE":      2,
	"INACTIVE":    3,
	"PENDING":     4,
	"REFUSED":     5,
	"CANCELED":    6,
}

func (x ManagerLinkStatusEnum_ManagerLinkStatus) String() string {
	return proto.EnumName(ManagerLinkStatusEnum_ManagerLinkStatus_name, int32(x))
}

func (ManagerLinkStatusEnum_ManagerLinkStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8f80561365406c5, []int{0, 0}
}

// Container for enum describing possible status of a manager and client link.
type ManagerLinkStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagerLinkStatusEnum) Reset()         { *m = ManagerLinkStatusEnum{} }
func (m *ManagerLinkStatusEnum) String() string { return proto.CompactTextString(m) }
func (*ManagerLinkStatusEnum) ProtoMessage()    {}
func (*ManagerLinkStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8f80561365406c5, []int{0}
}

func (m *ManagerLinkStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagerLinkStatusEnum.Unmarshal(m, b)
}
func (m *ManagerLinkStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagerLinkStatusEnum.Marshal(b, m, deterministic)
}
func (m *ManagerLinkStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagerLinkStatusEnum.Merge(m, src)
}
func (m *ManagerLinkStatusEnum) XXX_Size() int {
	return xxx_messageInfo_ManagerLinkStatusEnum.Size(m)
}
func (m *ManagerLinkStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagerLinkStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ManagerLinkStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.ManagerLinkStatusEnum_ManagerLinkStatus", ManagerLinkStatusEnum_ManagerLinkStatus_name, ManagerLinkStatusEnum_ManagerLinkStatus_value)
	proto.RegisterType((*ManagerLinkStatusEnum)(nil), "google.ads.googleads.v1.enums.ManagerLinkStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/manager_link_status.proto", fileDescriptor_e8f80561365406c5)
}

var fileDescriptor_e8f80561365406c5 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xfe, 0xb5, 0xfb, 0x39, 0x25, 0x13, 0xac, 0x05, 0xbd, 0x10, 0x77, 0xb1, 0x3d, 0x40, 0x4a,
	0xf1, 0x42, 0x88, 0x57, 0x59, 0x9b, 0x8d, 0xe2, 0x8c, 0xc5, 0xb9, 0x0a, 0x52, 0x18, 0xd1, 0x96,
	0x50, 0xb6, 0x26, 0x63, 0xe9, 0xf6, 0x14, 0x3e, 0x85, 0x97, 0x3e, 0x8a, 0x8f, 0xa2, 0x2f, 0x21,
	0x49, 0xb7, 0xde, 0x0c, 0xbd, 0x09, 0xdf, 0xc9, 0xf7, 0x87, 0x73, 0x3e, 0x70, 0xcd, 0xa5, 0xe4,
	0x8b, 0xdc, 0x63, 0x99, 0xf2, 0x6a, 0xa8, 0xd1, 0xc6, 0xf7, 0x72, 0xb1, 0x2e, 0x95, 0x57, 0x32,
	0xc1, 0x78, 0xbe, 0x9a, 0x2d, 0x0a, 0x31, 0x9f, 0xa9, 0x8a, 0x55, 0x6b, 0x05, 0x97, 0x2b, 0x59,
	0x49, 0xb7, 0x5b, 0xab, 0x21, 0xcb, 0x14, 0x6c, 0x8c, 0x70, 0xe3, 0x43, 0x63, 0xbc, 0xb8, 0xdc,
	0xe5, 0x2e, 0x0b, 0x8f, 0x09, 0x21, 0x2b, 0x56, 0x15, 0x52, 0x6c, 0xcd, 0xfd, 0x37, 0x0b, 0x9c,
	0xdd, 0xd5, 0xd1, 0xe3, 0x42, 0xcc, 0x27, 0x26, 0x98, 0x88, 0x75, 0xd9, 0x57, 0xe0, 0x74, 0x8f,
	0x70, 0x4f, 0x40, 0x67, 0x4a, 0x27, 0x31, 0x09, 0xa2, 0x61, 0x44, 0x42, 0xe7, 0x9f, 0xdb, 0x01,
	0x87, 0x53, 0x7a, 0x4b, 0xef, 0x9f, 0xa8, 0x63, 0xb9, 0x00, 0xb4, 0x71, 0xf0, 0x18, 0x25, 0xc4,
	0xb1, 0xdd, 0x63, 0x70, 0x14, 0xd1, 0xed, 0xd4, 0xd2, 0xb2, 0x98, 0xd0, 0x30, 0xa2, 0x23, 0xe7,
	0xbf, 0x1e, 0x1e, 0xc8, 0x70, 0x3a, 0x21, 0xa1, 0x73, 0xa0, 0x75, 0x01, 0xa6, 0x01, 0x19, 0x93,
	0xd0, 0x69, 0x0f, 0xbe, 0x2d, 0xd0, 0x7b, 0x95, 0x25, 0xfc, 0xf3, 0xa4, 0xc1, 0xf9, 0xde, 0x62,
	0xb1, 0x3e, 0x26, 0xb6, 0x9e, 0x07, 0x5b, 0x23, 0x97, 0x0b, 0x26, 0x38, 0x94, 0x2b, 0xee, 0xf1,
	0x5c, 0x98, 0x53, 0x77, 0xa5, 0x2e, 0x0b, 0xf5, 0x4b, 0xc7, 0x37, 0xe6, 0x7d, 0xb7, 0x5b, 0x23,
	0x8c, 0x3f, 0xec, 0xee, 0xa8, 0x8e, 0xc2, 0x99, 0x82, 0x35, 0xd4, 0x28, 0xf1, 0xa1, 0x6e, 0x47,
	0x7d, 0xee, 0xf8, 0x14, 0x67, 0x2a, 0x6d, 0xf8, 0x34, 0xf1, 0x53, 0xc3, 0x7f, 0xd9, 0xbd, 0xfa,
	0x13, 0x21, 0x9c, 0x29, 0x84, 0x1a, 0x05, 0x42, 0x89, 0x8f, 0x90, 0xd1, 0xbc, 0xb4, 0xcd, 0x62,
	0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x76, 0xf8, 0xde, 0xfb, 0x01, 0x00, 0x00,
}
