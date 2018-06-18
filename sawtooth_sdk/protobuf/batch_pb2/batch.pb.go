// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/batch_pb2/batch.proto

package batch_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import transaction_pb2 "github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/transaction_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BatchHeader struct {
	// Public key for the client that signed the BatchHeader
	SignerPublicKey string `protobuf:"bytes,1,opt,name=signer_public_key,json=signerPublicKey" json:"signer_public_key,omitempty"`
	// List of transaction.header_signatures that match the order of
	// transactions required for the batch
	TransactionIds       []string `protobuf:"bytes,2,rep,name=transaction_ids,json=transactionIds" json:"transaction_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchHeader) Reset()         { *m = BatchHeader{} }
func (m *BatchHeader) String() string { return proto.CompactTextString(m) }
func (*BatchHeader) ProtoMessage()    {}
func (*BatchHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_batch_09f1c4fdde06f07c, []int{0}
}
func (m *BatchHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchHeader.Unmarshal(m, b)
}
func (m *BatchHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchHeader.Marshal(b, m, deterministic)
}
func (dst *BatchHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchHeader.Merge(dst, src)
}
func (m *BatchHeader) XXX_Size() int {
	return xxx_messageInfo_BatchHeader.Size(m)
}
func (m *BatchHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BatchHeader proto.InternalMessageInfo

func (m *BatchHeader) GetSignerPublicKey() string {
	if m != nil {
		return m.SignerPublicKey
	}
	return ""
}

func (m *BatchHeader) GetTransactionIds() []string {
	if m != nil {
		return m.TransactionIds
	}
	return nil
}

type Batch struct {
	// The serialized version of the BatchHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature" json:"header_signature,omitempty"`
	// A list of the transactions that match the list of
	// transaction_ids listed in the batch header
	Transactions []*transaction_pb2.Transaction `protobuf:"bytes,3,rep,name=transactions" json:"transactions,omitempty"`
	// A debugging flag which indicates this batch should be traced through the
	// system, resulting in a higher level of debugging output.
	Trace                bool     `protobuf:"varint,4,opt,name=trace" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_batch_09f1c4fdde06f07c, []int{1}
}
func (m *Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Batch.Unmarshal(m, b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
}
func (dst *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(dst, src)
}
func (m *Batch) XXX_Size() int {
	return xxx_messageInfo_Batch.Size(m)
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Batch) GetHeaderSignature() string {
	if m != nil {
		return m.HeaderSignature
	}
	return ""
}

func (m *Batch) GetTransactions() []*transaction_pb2.Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Batch) GetTrace() bool {
	if m != nil {
		return m.Trace
	}
	return false
}

type BatchList struct {
	Batches              []*Batch `protobuf:"bytes,1,rep,name=batches" json:"batches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchList) Reset()         { *m = BatchList{} }
func (m *BatchList) String() string { return proto.CompactTextString(m) }
func (*BatchList) ProtoMessage()    {}
func (*BatchList) Descriptor() ([]byte, []int) {
	return fileDescriptor_batch_09f1c4fdde06f07c, []int{2}
}
func (m *BatchList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchList.Unmarshal(m, b)
}
func (m *BatchList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchList.Marshal(b, m, deterministic)
}
func (dst *BatchList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchList.Merge(dst, src)
}
func (m *BatchList) XXX_Size() int {
	return xxx_messageInfo_BatchList.Size(m)
}
func (m *BatchList) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchList.DiscardUnknown(m)
}

var xxx_messageInfo_BatchList proto.InternalMessageInfo

func (m *BatchList) GetBatches() []*Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchHeader)(nil), "BatchHeader")
	proto.RegisterType((*Batch)(nil), "Batch")
	proto.RegisterType((*BatchList)(nil), "BatchList")
}

func init() {
	proto.RegisterFile("github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/batch_pb2/batch.proto", fileDescriptor_batch_09f1c4fdde06f07c)
}

var fileDescriptor_batch_09f1c4fdde06f07c = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0xc9, 0xf6, 0xbf, 0xfd, 0xdb, 0xd9, 0xe2, 0x6a, 0x50, 0x09, 0x9e, 0x4a, 0x11, 0xac,
	0x8a, 0x5d, 0x59, 0x0f, 0xde, 0xf7, 0xa4, 0xe8, 0x61, 0x89, 0x9e, 0xbc, 0x94, 0xa4, 0x8d, 0xb6,
	0x54, 0xda, 0x92, 0xa4, 0xc8, 0x7e, 0x0d, 0x3f, 0xb1, 0xec, 0xc4, 0x6a, 0x05, 0x6f, 0x33, 0xbf,
	0x99, 0x79, 0x2f, 0xbc, 0xc0, 0x85, 0x11, 0xef, 0xb6, 0x6d, 0x6d, 0x99, 0x99, 0xa2, 0x5e, 0x74,
	0xba, 0xb5, 0xad, 0xec, 0x5f, 0x16, 0x52, 0xd8, 0xbc, 0xcc, 0x3a, 0xb9, 0x74, 0x55, 0x8a, 0x83,
	0xe3, 0x9b, 0xbf, 0x97, 0xad, 0x16, 0x8d, 0x11, 0xb9, 0xad, 0xda, 0x06, 0x4f, 0x46, 0xbd, 0x3b,
	0x8c, 0x25, 0xcc, 0x56, 0x5b, 0x9d, 0x5b, 0x25, 0x0a, 0xa5, 0xe9, 0x39, 0xec, 0x9b, 0xea, 0xb5,
	0x51, 0x3a, 0xeb, 0x7a, 0xf9, 0x56, 0xe5, 0x59, 0xad, 0x36, 0x8c, 0x44, 0x24, 0x09, 0xf8, 0xdc,
	0x0d, 0xd6, 0xc8, 0xef, 0xd5, 0x86, 0x9e, 0xc2, 0x7c, 0xac, 0x5f, 0x15, 0x86, 0x4d, 0x22, 0x2f,
	0x09, 0xf8, 0xee, 0x08, 0xdf, 0x15, 0x26, 0xfe, 0x20, 0x30, 0x45, 0x13, 0x7a, 0x04, 0x7e, 0x89,
	0x46, 0xa8, 0x19, 0xf2, 0xaf, 0x8e, 0x9e, 0xc1, 0x9e, 0xab, 0xb2, 0xad, 0x89, 0xb0, 0xbd, 0x56,
	0x6c, 0xe2, 0x5c, 0x1d, 0x7f, 0x1c, 0x30, 0xbd, 0x82, 0x70, 0x24, 0x6f, 0x98, 0x17, 0x79, 0xc9,
	0x6c, 0x19, 0xa6, 0x4f, 0x3f, 0x90, 0xff, 0xda, 0xa0, 0x07, 0x30, 0xb5, 0x5a, 0xe4, 0x8a, 0xfd,
	0x8b, 0x48, 0xb2, 0xc3, 0x5d, 0x13, 0x5f, 0x42, 0x80, 0x6f, 0x7a, 0xa8, 0x8c, 0xa5, 0x11, 0xfc,
	0xc7, 0x34, 0x95, 0x61, 0x04, 0xf5, 0xfc, 0x14, 0x87, 0x7c, 0xc0, 0xab, 0x13, 0x38, 0x1c, 0x22,
	0x4e, 0x4d, 0x51, 0xa7, 0x43, 0xc4, 0x6b, 0xf2, 0x1c, 0x7c, 0x7f, 0x89, 0xf4, 0x11, 0x5f, 0x7f,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x84, 0x94, 0xe2, 0x46, 0xbc, 0x01, 0x00, 0x00,
}