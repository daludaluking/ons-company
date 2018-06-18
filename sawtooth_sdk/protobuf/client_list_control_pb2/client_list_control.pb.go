// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/client_list_control_pb2/client_list_control.proto

package client_list_control_pb2

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

// Paging controls to be sent with List requests.
// Attributes:
//     start: The id of a resource to start the page with
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingControls struct {
	Start                string   `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientPagingControls) Reset()         { *m = ClientPagingControls{} }
func (m *ClientPagingControls) String() string { return proto.CompactTextString(m) }
func (*ClientPagingControls) ProtoMessage()    {}
func (*ClientPagingControls) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_list_control_60bfc4d788568132, []int{0}
}
func (m *ClientPagingControls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientPagingControls.Unmarshal(m, b)
}
func (m *ClientPagingControls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientPagingControls.Marshal(b, m, deterministic)
}
func (dst *ClientPagingControls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientPagingControls.Merge(dst, src)
}
func (m *ClientPagingControls) XXX_Size() int {
	return xxx_messageInfo_ClientPagingControls.Size(m)
}
func (m *ClientPagingControls) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientPagingControls.DiscardUnknown(m)
}

var xxx_messageInfo_ClientPagingControls proto.InternalMessageInfo

func (m *ClientPagingControls) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ClientPagingControls) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Information about the pagination used, sent back with List responses.
// Attributes:
//     next: The id of the first resource in the next page
//     start: The id of the first resource in the returned page
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingResponse struct {
	Next                 string   `protobuf:"bytes,1,opt,name=next" json:"next,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientPagingResponse) Reset()         { *m = ClientPagingResponse{} }
func (m *ClientPagingResponse) String() string { return proto.CompactTextString(m) }
func (*ClientPagingResponse) ProtoMessage()    {}
func (*ClientPagingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_list_control_60bfc4d788568132, []int{1}
}
func (m *ClientPagingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientPagingResponse.Unmarshal(m, b)
}
func (m *ClientPagingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientPagingResponse.Marshal(b, m, deterministic)
}
func (dst *ClientPagingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientPagingResponse.Merge(dst, src)
}
func (m *ClientPagingResponse) XXX_Size() int {
	return xxx_messageInfo_ClientPagingResponse.Size(m)
}
func (m *ClientPagingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientPagingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientPagingResponse proto.InternalMessageInfo

func (m *ClientPagingResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *ClientPagingResponse) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ClientPagingResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Sorting controls to be sent with List requests. More than one can be sent.
// If so, the first is used, and additional controls are tie-breakers.
// Attributes:
//     keys: Nested set of keys to sort by (i.e. ['default, block_num'])
//     reverse: Whether or not to reverse the sort (i.e. descending order)
type ClientSortControls struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Reverse              bool     `protobuf:"varint,2,opt,name=reverse" json:"reverse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientSortControls) Reset()         { *m = ClientSortControls{} }
func (m *ClientSortControls) String() string { return proto.CompactTextString(m) }
func (*ClientSortControls) ProtoMessage()    {}
func (*ClientSortControls) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_list_control_60bfc4d788568132, []int{2}
}
func (m *ClientSortControls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSortControls.Unmarshal(m, b)
}
func (m *ClientSortControls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSortControls.Marshal(b, m, deterministic)
}
func (dst *ClientSortControls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSortControls.Merge(dst, src)
}
func (m *ClientSortControls) XXX_Size() int {
	return xxx_messageInfo_ClientSortControls.Size(m)
}
func (m *ClientSortControls) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSortControls.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSortControls proto.InternalMessageInfo

func (m *ClientSortControls) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ClientSortControls) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

func init() {
	proto.RegisterType((*ClientPagingControls)(nil), "ClientPagingControls")
	proto.RegisterType((*ClientPagingResponse)(nil), "ClientPagingResponse")
	proto.RegisterType((*ClientSortControls)(nil), "ClientSortControls")
}

func init() {
	proto.RegisterFile("github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/client_list_control_pb2/client_list_control.proto", fileDescriptor_client_list_control_60bfc4d788568132)
}

var fileDescriptor_client_list_control_60bfc4d788568132 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc9, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0x51, 0x96, 0x2e, 0x73, 0x14, 0x29, 0xf5, 0xd1, 0xe4, 0x94, 0x93, 0x03, 0xe9, 0x1b,
	0x38, 0x87, 0x5e, 0x83, 0x0b, 0x3d, 0xf4, 0x22, 0xec, 0x64, 0x9a, 0x0a, 0xab, 0x1a, 0xa3, 0x99,
	0x6e, 0x6f, 0x5f, 0x2c, 0xa1, 0x2e, 0x90, 0xdc, 0xfe, 0x45, 0xfa, 0x18, 0x7e, 0xb8, 0xe7, 0xf6,
	0x43, 0x88, 0xe4, 0xc5, 0xf0, 0xa1, 0x5f, 0x0f, 0x81, 0x84, 0xba, 0xb7, 0xe7, 0xf5, 0xde, 0x59,
	0xf4, 0x62, 0x9c, 0x65, 0x31, 0x7b, 0xf2, 0x12, 0xc8, 0x99, 0xa1, 0xdb, 0x9c, 0xca, 0xab, 0xf8,
	0x69, 0x59, 0xc3, 0x62, 0x1b, 0xcb, 0x5d, 0x7b, 0xb4, 0xfe, 0xb8, 0x4d, 0x25, 0xeb, 0x05, 0xcc,
	0x59, 0xda, 0x20, 0x85, 0x2a, 0xd5, 0xea, 0xba, 0x49, 0x66, 0x4c, 0x9d, 0x7d, 0xb5, 0x52, 0x4c,
	0x4a, 0xb5, 0x9a, 0x37, 0xc9, 0x2c, 0x1f, 0xff, 0x33, 0x1a, 0xe4, 0x81, 0x3c, 0xa3, 0xd6, 0x30,
	0xf3, 0xf8, 0x99, 0x11, 0x51, 0xff, 0x72, 0x27, 0x27, 0xb9, 0xd3, 0xbf, 0xdc, 0x1a, 0x74, 0xe2,
	0x3e, 0x50, 0x90, 0x9f, 0xcb, 0x34, 0xcc, 0x7a, 0xfc, 0xe2, 0x42, 0x95, 0xd3, 0x91, 0x3a, 0x6a,
	0x5d, 0xc0, 0x65, 0xc0, 0x77, 0x0c, 0x8c, 0x91, 0x7b, 0xd5, 0x64, 0x5b, 0x6f, 0xe0, 0x26, 0x4f,
	0x55, 0xf1, 0xa1, 0xaf, 0xf2, 0x54, 0x3b, 0xf5, 0x74, 0x7b, 0x66, 0xad, 0xee, 0x22, 0x3e, 0xba,
	0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x01, 0x8d, 0xcb, 0xe5, 0x65, 0x01, 0x00, 0x00,
}