// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/block_pb2/block.proto

package block_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import batch_pb2 "github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/batch_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	// Block number in the chain
	BlockNum uint64 `protobuf:"varint,1,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
	// The header_signature of the previous block that was added to the chain.
	PreviousBlockId string `protobuf:"bytes,2,opt,name=previous_block_id,json=previousBlockId" json:"previous_block_id,omitempty"`
	// Public key for the component internal to the validator that
	// signed the BlockHeader
	SignerPublicKey string `protobuf:"bytes,3,opt,name=signer_public_key,json=signerPublicKey" json:"signer_public_key,omitempty"`
	// List of batch.header_signatures that match the order of batches
	// required for the block
	BatchIds []string `protobuf:"bytes,4,rep,name=batch_ids,json=batchIds" json:"batch_ids,omitempty"`
	// Bytes that are set and verified by the consensus algorithm used to
	// create and validate the block
	Consensus []byte `protobuf:"bytes,5,opt,name=consensus,proto3" json:"consensus,omitempty"`
	// The state_root_hash should match the final state_root after all
	// transactions in the batches have been applied, otherwise the block
	// is not valid
	StateRootHash        string   `protobuf:"bytes,6,opt,name=state_root_hash,json=stateRootHash" json:"state_root_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_f4726eacc980b068, []int{0}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockId() string {
	if m != nil {
		return m.PreviousBlockId
	}
	return ""
}

func (m *BlockHeader) GetSignerPublicKey() string {
	if m != nil {
		return m.SignerPublicKey
	}
	return ""
}

func (m *BlockHeader) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *BlockHeader) GetConsensus() []byte {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *BlockHeader) GetStateRootHash() string {
	if m != nil {
		return m.StateRootHash
	}
	return ""
}

type Block struct {
	// The serialized version of the BlockHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature" json:"header_signature,omitempty"`
	// A list of batches. The batches may contain new batches that other
	// validators may not have received yet, or they will be all batches needed
	// for block validation when passed to the journal
	Batches              []*batch_pb2.Batch `protobuf:"bytes,3,rep,name=batches" json:"batches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_f4726eacc980b068, []int{1}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetHeaderSignature() string {
	if m != nil {
		return m.HeaderSignature
	}
	return ""
}

func (m *Block) GetBatches() []*batch_pb2.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "BlockHeader")
	proto.RegisterType((*Block)(nil), "Block")
}

func init() {
	proto.RegisterFile("github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/block_pb2/block.proto", fileDescriptor_block_f4726eacc980b068)
}

var fileDescriptor_block_f4726eacc980b068 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xa9, 0xdd, 0xea, 0x9a, 0x4d, 0xa6, 0x01, 0x25, 0xa8, 0x0f, 0x65, 0x88, 0x54, 0x85,
	0x0d, 0xe6, 0x37, 0xd8, 0xd3, 0x86, 0x20, 0x23, 0xbe, 0xf9, 0x12, 0xd2, 0x26, 0xb3, 0x65, 0x5b,
	0x53, 0x7a, 0x89, 0xb2, 0x0f, 0xec, 0xf7, 0x90, 0x5e, 0x56, 0x7d, 0xf2, 0xed, 0xf2, 0xbb, 0x7f,
	0xee, 0xfe, 0xc7, 0x9f, 0x3c, 0x81, 0xfc, 0xb2, 0xc6, 0xd8, 0x42, 0x80, 0xda, 0xce, 0xea, 0xc6,
	0x58, 0x93, 0xb9, 0xcd, 0x2c, 0xdb, 0x99, 0x7c, 0x2b, 0xea, 0x6c, 0xee, 0xab, 0x29, 0x36, 0xae,
	0xff, 0x13, 0x4b, 0x9b, 0x17, 0x5e, 0xdc, 0x56, 0x5e, 0x3c, 0xf9, 0x0e, 0xc8, 0x70, 0xd1, 0x7e,
	0x5e, 0x6a, 0xa9, 0x74, 0x43, 0x6f, 0x48, 0xec, 0xa7, 0x56, 0x6e, 0xcf, 0x82, 0x24, 0x48, 0x7b,
	0x7c, 0x80, 0xe0, 0xd5, 0xed, 0xe9, 0x23, 0xb9, 0xa8, 0x1b, 0xfd, 0x59, 0x1a, 0x07, 0xc2, 0xab,
	0x4a, 0xc5, 0x4e, 0x92, 0x20, 0x8d, 0xf9, 0xb8, 0x6b, 0xe0, 0xb0, 0x95, 0x6a, 0xb5, 0x50, 0x7e,
	0x54, 0xba, 0x11, 0xb5, 0xcb, 0x76, 0x65, 0x2e, 0xb6, 0xfa, 0xc0, 0x42, 0xaf, 0xf5, 0x8d, 0x35,
	0xf2, 0x17, 0x7d, 0xc0, 0xa5, 0xe8, 0xae, 0x54, 0xc0, 0x7a, 0x49, 0x98, 0xc6, 0x7c, 0x80, 0x60,
	0xa5, 0x80, 0xde, 0x92, 0x38, 0x37, 0x15, 0xe8, 0x0a, 0x1c, 0xb0, 0x7e, 0x12, 0xa4, 0x23, 0xfe,
	0x07, 0xe8, 0x3d, 0x19, 0x83, 0x95, 0x56, 0x8b, 0xc6, 0x18, 0x2b, 0x0a, 0x09, 0x05, 0x8b, 0x70,
	0xc9, 0x19, 0x62, 0x6e, 0x8c, 0x5d, 0x4a, 0x28, 0x26, 0x3b, 0xd2, 0x47, 0x67, 0xf4, 0x8a, 0x44,
	0x05, 0x9e, 0x8a, 0xd7, 0x8d, 0xf8, 0xf1, 0x45, 0x1f, 0xc8, 0xb9, 0xaf, 0x44, 0xeb, 0x4e, 0x5a,
	0xd7, 0xe8, 0xee, 0x34, 0xcf, 0xdf, 0x3a, 0x4c, 0x13, 0x72, 0x8a, 0xee, 0x34, 0xb0, 0x30, 0x09,
	0xd3, 0xe1, 0x3c, 0x9a, 0x2e, 0xda, 0x37, 0xef, 0xf0, 0xe2, 0x8e, 0x5c, 0x76, 0x21, 0x4c, 0x41,
	0x1d, 0x83, 0xc9, 0xdc, 0x66, 0x1d, 0xbc, 0xc7, 0xbf, 0xa1, 0x65, 0x11, 0xe2, 0xe7, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x28, 0xac, 0x53, 0x9f, 0xde, 0x01, 0x00, 0x00,
}
