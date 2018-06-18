// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/transaction_receipt_pb2/transaction_receipt.proto

package txn_receipt_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events_pb2 "github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/events_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StateChange_Type int32

const (
	StateChange_TYPE_UNSET StateChange_Type = 0
	StateChange_SET        StateChange_Type = 1
	StateChange_DELETE     StateChange_Type = 2
)

var StateChange_Type_name = map[int32]string{
	0: "TYPE_UNSET",
	1: "SET",
	2: "DELETE",
}
var StateChange_Type_value = map[string]int32{
	"TYPE_UNSET": 0,
	"SET":        1,
	"DELETE":     2,
}

func (x StateChange_Type) String() string {
	return proto.EnumName(StateChange_Type_name, int32(x))
}
func (StateChange_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transaction_receipt_b22653c38208ec7a, []int{1, 0}
}

type TransactionReceipt struct {
	// State changes made by this transaction
	// StateChange is defined in protos/transaction_receipt.proto
	StateChanges []*StateChange `protobuf:"bytes,1,rep,name=state_changes,json=stateChanges" json:"state_changes,omitempty"`
	// Events fired by this transaction
	Events []*events_pb2.Event `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
	// Transaction family defined data
	Data                 [][]byte `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	TransactionId        string   `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionReceipt) Reset()         { *m = TransactionReceipt{} }
func (m *TransactionReceipt) String() string { return proto.CompactTextString(m) }
func (*TransactionReceipt) ProtoMessage()    {}
func (*TransactionReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_receipt_b22653c38208ec7a, []int{0}
}
func (m *TransactionReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionReceipt.Unmarshal(m, b)
}
func (m *TransactionReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionReceipt.Marshal(b, m, deterministic)
}
func (dst *TransactionReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionReceipt.Merge(dst, src)
}
func (m *TransactionReceipt) XXX_Size() int {
	return xxx_messageInfo_TransactionReceipt.Size(m)
}
func (m *TransactionReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionReceipt proto.InternalMessageInfo

func (m *TransactionReceipt) GetStateChanges() []*StateChange {
	if m != nil {
		return m.StateChanges
	}
	return nil
}

func (m *TransactionReceipt) GetEvents() []*events_pb2.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *TransactionReceipt) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TransactionReceipt) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

//  StateChange objects have the type of SET, which is either an insert or
//  update, or DELETE. Items marked as a DELETE will have no byte value.
type StateChange struct {
	Address              string           `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Value                []byte           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type                 StateChange_Type `protobuf:"varint,3,opt,name=type,enum=StateChange_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StateChange) Reset()         { *m = StateChange{} }
func (m *StateChange) String() string { return proto.CompactTextString(m) }
func (*StateChange) ProtoMessage()    {}
func (*StateChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_receipt_b22653c38208ec7a, []int{1}
}
func (m *StateChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateChange.Unmarshal(m, b)
}
func (m *StateChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateChange.Marshal(b, m, deterministic)
}
func (dst *StateChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChange.Merge(dst, src)
}
func (m *StateChange) XXX_Size() int {
	return xxx_messageInfo_StateChange.Size(m)
}
func (m *StateChange) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChange.DiscardUnknown(m)
}

var xxx_messageInfo_StateChange proto.InternalMessageInfo

func (m *StateChange) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StateChange) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StateChange) GetType() StateChange_Type {
	if m != nil {
		return m.Type
	}
	return StateChange_TYPE_UNSET
}

// A collection of state changes.
type StateChangeList struct {
	StateChanges         []*StateChange `protobuf:"bytes,1,rep,name=state_changes,json=stateChanges" json:"state_changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StateChangeList) Reset()         { *m = StateChangeList{} }
func (m *StateChangeList) String() string { return proto.CompactTextString(m) }
func (*StateChangeList) ProtoMessage()    {}
func (*StateChangeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_receipt_b22653c38208ec7a, []int{2}
}
func (m *StateChangeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateChangeList.Unmarshal(m, b)
}
func (m *StateChangeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateChangeList.Marshal(b, m, deterministic)
}
func (dst *StateChangeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChangeList.Merge(dst, src)
}
func (m *StateChangeList) XXX_Size() int {
	return xxx_messageInfo_StateChangeList.Size(m)
}
func (m *StateChangeList) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChangeList.DiscardUnknown(m)
}

var xxx_messageInfo_StateChangeList proto.InternalMessageInfo

func (m *StateChangeList) GetStateChanges() []*StateChange {
	if m != nil {
		return m.StateChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionReceipt)(nil), "TransactionReceipt")
	proto.RegisterType((*StateChange)(nil), "StateChange")
	proto.RegisterType((*StateChangeList)(nil), "StateChangeList")
	proto.RegisterEnum("StateChange_Type", StateChange_Type_name, StateChange_Type_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/transaction_receipt_pb2/transaction_receipt.proto", fileDescriptor_transaction_receipt_b22653c38208ec7a)
}

var fileDescriptor_transaction_receipt_b22653c38208ec7a = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xdd, 0x4a, 0xc3, 0x30,
	0x18, 0x35, 0x6b, 0xed, 0xf0, 0x5b, 0xf7, 0x63, 0x50, 0x08, 0x5e, 0x48, 0x29, 0x0c, 0x8a, 0x62,
	0x87, 0xf3, 0x0d, 0x74, 0x45, 0x84, 0x21, 0x23, 0xab, 0x17, 0x7a, 0x53, 0xb2, 0x25, 0xba, 0x32,
	0x69, 0x4b, 0x93, 0x4d, 0xf7, 0x18, 0xbe, 0x82, 0x4f, 0x2a, 0x49, 0x1d, 0xab, 0xb0, 0x2b, 0xef,
	0xce, 0x39, 0xdf, 0xf9, 0x4e, 0xfb, 0xe5, 0xc0, 0xbd, 0x64, 0x1f, 0x2a, 0xcf, 0xd5, 0x22, 0x91,
	0x7c, 0x39, 0x28, 0xca, 0x5c, 0xe5, 0xb3, 0xd5, 0xeb, 0x40, 0x95, 0x2c, 0x93, 0x6c, 0xae, 0xd2,
	0x3c, 0x4b, 0x4a, 0x31, 0x17, 0x69, 0xa1, 0x92, 0x62, 0x36, 0xdc, 0xa7, 0x87, 0x66, 0xe9, 0xec,
	0x6a, 0x7f, 0x90, 0x58, 0x8b, 0x4c, 0x49, 0xb3, 0x5b, 0xc1, 0xca, 0xee, 0x7f, 0x23, 0xc0, 0xf1,
	0x2e, 0x8c, 0x56, 0x59, 0xf8, 0x1a, 0xda, 0x52, 0x31, 0x25, 0x92, 0xf9, 0x82, 0x65, 0x6f, 0x42,
	0x12, 0xe4, 0x59, 0x41, 0x6b, 0xe8, 0x86, 0x53, 0xad, 0xde, 0x19, 0x91, 0xba, 0x72, 0x47, 0x24,
	0x3e, 0x07, 0xa7, 0x4a, 0x26, 0x0d, 0xe3, 0x75, 0xc2, 0x48, 0x53, 0xfa, 0xab, 0x62, 0x0c, 0x36,
	0x67, 0x8a, 0x11, 0xcb, 0xb3, 0x02, 0x97, 0x1a, 0x8c, 0xfb, 0xd0, 0xa9, 0x5f, 0x92, 0x72, 0x62,
	0x7b, 0x28, 0x38, 0xa2, 0xed, 0x9a, 0xfa, 0xc0, 0xfd, 0x2f, 0x04, 0xad, 0xda, 0x87, 0x31, 0x81,
	0x26, 0xe3, 0xbc, 0x14, 0x52, 0xff, 0x97, 0xf6, 0x6f, 0x29, 0x3e, 0x81, 0xc3, 0x35, 0x7b, 0x5f,
	0x09, 0xd2, 0xf0, 0x50, 0xe0, 0xd2, 0x8a, 0xe0, 0x3e, 0xd8, 0x6a, 0x53, 0x08, 0x62, 0x79, 0x28,
	0xe8, 0x0c, 0x8f, 0xeb, 0x47, 0x84, 0xf1, 0xa6, 0x10, 0xd4, 0x8c, 0xfd, 0x4b, 0xb0, 0x35, 0xc3,
	0x1d, 0x80, 0xf8, 0x79, 0x12, 0x25, 0x4f, 0x8f, 0xd3, 0x28, 0xee, 0x1d, 0xe0, 0x26, 0x58, 0x1a,
	0x20, 0x0c, 0xe0, 0x8c, 0xa2, 0x71, 0x14, 0x47, 0xbd, 0x86, 0x3f, 0x82, 0x6e, 0x2d, 0x66, 0x9c,
	0xca, 0xff, 0x3c, 0xda, 0xed, 0x05, 0x9c, 0x6e, 0xfb, 0x0a, 0x25, 0x5f, 0x86, 0xdb, 0xbe, 0x26,
	0xe8, 0xa5, 0xab, 0x3e, 0xff, 0x74, 0x3e, 0x73, 0xcc, 0xf0, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x3b, 0xb5, 0x67, 0xd8, 0x2b, 0x02, 0x00, 0x00,
}
