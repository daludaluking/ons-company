// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/validator_pb2/validator.proto

package validator_pb2

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

type Message_MessageType int32

const (
	Message_DEFAULT Message_MessageType = 0
	// Registration request from the transaction processor to the validator
	Message_TP_REGISTER_REQUEST Message_MessageType = 1
	// Registration response from the validator to the
	// transaction processor
	Message_TP_REGISTER_RESPONSE Message_MessageType = 2
	// Tell the validator that the transaction processor
	// won't take any more transactions
	Message_TP_UNREGISTER_REQUEST Message_MessageType = 3
	// Response from the validator to the tp that it won't
	// send any more transactions
	Message_TP_UNREGISTER_RESPONSE Message_MessageType = 4
	// Process Request from the validator/executor to the
	// transaction processor
	Message_TP_PROCESS_REQUEST Message_MessageType = 5
	// Process response from the transaction processor to the validator/executor
	Message_TP_PROCESS_RESPONSE Message_MessageType = 6
	// State get request from the transaction processor to validator/context_manager
	Message_TP_STATE_GET_REQUEST Message_MessageType = 7
	// State get response from the validator/context_manager to the transaction processor
	Message_TP_STATE_GET_RESPONSE Message_MessageType = 8
	// State set request from the transaction processor to the validator/context_manager
	Message_TP_STATE_SET_REQUEST Message_MessageType = 9
	// State set response from the validator/context_manager to the transaction processor
	Message_TP_STATE_SET_RESPONSE Message_MessageType = 10
	// State delete request from the transaction processor to the validator/context_manager
	Message_TP_STATE_DELETE_REQUEST Message_MessageType = 11
	// State delete response from the validator/context_manager to the transaction processor
	Message_TP_STATE_DELETE_RESPONSE Message_MessageType = 12
	// Message to append data to a transaction receipt
	Message_TP_RECEIPT_ADD_DATA_REQUEST Message_MessageType = 13
	// Response from validator to tell transaction processor that data has been appended
	Message_TP_RECEIPT_ADD_DATA_RESPONSE Message_MessageType = 14
	// Message to add event
	Message_TP_EVENT_ADD_REQUEST Message_MessageType = 15
	// Response from validator to tell transaction processor that event has been created
	Message_TP_EVENT_ADD_RESPONSE Message_MessageType = 16
	// Submission of a batchlist from the web api or another client to the validator
	Message_CLIENT_BATCH_SUBMIT_REQUEST Message_MessageType = 100
	// Response from the validator to the web api/client that the submission was accepted
	Message_CLIENT_BATCH_SUBMIT_RESPONSE Message_MessageType = 101
	// A request to list blocks from the web api/client to the validator
	Message_CLIENT_BLOCK_LIST_REQUEST        Message_MessageType = 102
	Message_CLIENT_BLOCK_LIST_RESPONSE       Message_MessageType = 103
	Message_CLIENT_BLOCK_GET_BY_ID_REQUEST   Message_MessageType = 104
	Message_CLIENT_BLOCK_GET_RESPONSE        Message_MessageType = 105
	Message_CLIENT_BATCH_LIST_REQUEST        Message_MessageType = 106
	Message_CLIENT_BATCH_LIST_RESPONSE       Message_MessageType = 107
	Message_CLIENT_BATCH_GET_REQUEST         Message_MessageType = 108
	Message_CLIENT_BATCH_GET_RESPONSE        Message_MessageType = 109
	Message_CLIENT_TRANSACTION_LIST_REQUEST  Message_MessageType = 110
	Message_CLIENT_TRANSACTION_LIST_RESPONSE Message_MessageType = 111
	Message_CLIENT_TRANSACTION_GET_REQUEST   Message_MessageType = 112
	Message_CLIENT_TRANSACTION_GET_RESPONSE  Message_MessageType = 113
	// Client state request of the current state hash to be retrieved from the journal
	Message_CLIENT_STATE_CURRENT_REQUEST Message_MessageType = 114
	// Response with the current state hash
	Message_CLIENT_STATE_CURRENT_RESPONSE Message_MessageType = 115
	// A request of all the addresses under a particular prefix, for a state hash.
	Message_CLIENT_STATE_LIST_REQUEST Message_MessageType = 116
	// The response of the addresses
	Message_CLIENT_STATE_LIST_RESPONSE Message_MessageType = 117
	// Get the address:data entry at a particular address
	Message_CLIENT_STATE_GET_REQUEST Message_MessageType = 118
	// The response with the entry
	Message_CLIENT_STATE_GET_RESPONSE Message_MessageType = 119
	// A request for the status of a batch or batches
	Message_CLIENT_BATCH_STATUS_REQUEST Message_MessageType = 120
	// A response with the batch statuses
	Message_CLIENT_BATCH_STATUS_RESPONSE Message_MessageType = 121
	// A request for one or more transaction receipts
	Message_CLIENT_RECEIPT_GET_REQUEST Message_MessageType = 122
	// A response with the receipts
	Message_CLIENT_RECEIPT_GET_RESPONSE     Message_MessageType = 123
	Message_CLIENT_BLOCK_GET_BY_NUM_REQUEST Message_MessageType = 124
	// A request for a validator's peers
	Message_CLIENT_PEERS_GET_REQUEST Message_MessageType = 125
	// A response with the validator's peers
	Message_CLIENT_PEERS_GET_RESPONSE                  Message_MessageType = 126
	Message_CLIENT_BLOCK_GET_BY_TRANSACTION_ID_REQUEST Message_MessageType = 127
	Message_CLIENT_BLOCK_GET_BY_BATCH_ID_REQUEST       Message_MessageType = 128
	// A request for a validator's status
	Message_CLIENT_STATUS_GET_REQUEST Message_MessageType = 129
	// A response with the validator's status
	Message_CLIENT_STATUS_GET_RESPONSE Message_MessageType = 130
	// Message types for events
	Message_CLIENT_EVENTS_SUBSCRIBE_REQUEST    Message_MessageType = 500
	Message_CLIENT_EVENTS_SUBSCRIBE_RESPONSE   Message_MessageType = 501
	Message_CLIENT_EVENTS_UNSUBSCRIBE_REQUEST  Message_MessageType = 502
	Message_CLIENT_EVENTS_UNSUBSCRIBE_RESPONSE Message_MessageType = 503
	Message_CLIENT_EVENTS                      Message_MessageType = 504
	Message_CLIENT_EVENTS_GET_REQUEST          Message_MessageType = 505
	Message_CLIENT_EVENTS_GET_RESPONSE         Message_MessageType = 506
	// Temp message types until a discussion can be had about gossip msg
	Message_GOSSIP_MESSAGE                         Message_MessageType = 200
	Message_GOSSIP_REGISTER                        Message_MessageType = 201
	Message_GOSSIP_UNREGISTER                      Message_MessageType = 202
	Message_GOSSIP_BLOCK_REQUEST                   Message_MessageType = 205
	Message_GOSSIP_BLOCK_RESPONSE                  Message_MessageType = 206
	Message_GOSSIP_BATCH_BY_BATCH_ID_REQUEST       Message_MessageType = 207
	Message_GOSSIP_BATCH_BY_TRANSACTION_ID_REQUEST Message_MessageType = 208
	Message_GOSSIP_BATCH_RESPONSE                  Message_MessageType = 209
	Message_GOSSIP_GET_PEERS_REQUEST               Message_MessageType = 210
	Message_GOSSIP_GET_PEERS_RESPONSE              Message_MessageType = 211
	Message_NETWORK_ACK                            Message_MessageType = 300
	Message_NETWORK_CONNECT                        Message_MessageType = 301
	Message_NETWORK_DISCONNECT                     Message_MessageType = 302
	// Message types for Authorization Types
	Message_AUTHORIZATION_CONNECTION_RESPONSE Message_MessageType = 600
	Message_AUTHORIZATION_VIOLATION           Message_MessageType = 601
	Message_AUTHORIZATION_TRUST_REQUEST       Message_MessageType = 602
	Message_AUTHORIZATION_TRUST_RESPONSE      Message_MessageType = 603
	Message_AUTHORIZATION_CHALLENGE_REQUEST   Message_MessageType = 604
	Message_AUTHORIZATION_CHALLENGE_RESPONSE  Message_MessageType = 605
	Message_AUTHORIZATION_CHALLENGE_SUBMIT    Message_MessageType = 606
	Message_AUTHORIZATION_CHALLENGE_RESULT    Message_MessageType = 607
	Message_PING_REQUEST                      Message_MessageType = 700
	Message_PING_RESPONSE                     Message_MessageType = 701
	// Consensus service messages
	Message_CONSENSUS_REGISTER_REQUEST          Message_MessageType = 800
	Message_CONSENSUS_REGISTER_RESPONSE         Message_MessageType = 801
	Message_CONSENSUS_SEND_TO_REQUEST           Message_MessageType = 802
	Message_CONSENSUS_SEND_TO_RESPONSE          Message_MessageType = 803
	Message_CONSENSUS_BROADCAST_REQUEST         Message_MessageType = 804
	Message_CONSENSUS_BROADCAST_RESPONSE        Message_MessageType = 805
	Message_CONSENSUS_INITIALIZE_BLOCK_REQUEST  Message_MessageType = 806
	Message_CONSENSUS_INITIALIZE_BLOCK_RESPONSE Message_MessageType = 807
	Message_CONSENSUS_FINALIZE_BLOCK_REQUEST    Message_MessageType = 808
	Message_CONSENSUS_FINALIZE_BLOCK_RESPONSE   Message_MessageType = 809
	Message_CONSENSUS_CANCEL_BLOCK_REQUEST      Message_MessageType = 810
	Message_CONSENSUS_CANCEL_BLOCK_RESPONSE     Message_MessageType = 811
	Message_CONSENSUS_CHECK_BLOCKS_REQUEST      Message_MessageType = 812
	Message_CONSENSUS_CHECK_BLOCKS_RESPONSE     Message_MessageType = 813
	Message_CONSENSUS_COMMIT_BLOCK_REQUEST      Message_MessageType = 814
	Message_CONSENSUS_COMMIT_BLOCK_RESPONSE     Message_MessageType = 815
	Message_CONSENSUS_IGNORE_BLOCK_REQUEST      Message_MessageType = 816
	Message_CONSENSUS_IGNORE_BLOCK_RESPONSE     Message_MessageType = 817
	Message_CONSENSUS_FAIL_BLOCK_REQUEST        Message_MessageType = 818
	Message_CONSENSUS_FAIL_BLOCK_RESPONSE       Message_MessageType = 819
	Message_CONSENSUS_SETTINGS_GET_REQUEST      Message_MessageType = 820
	Message_CONSENSUS_SETTINGS_GET_RESPONSE     Message_MessageType = 821
	Message_CONSENSUS_STATE_GET_REQUEST         Message_MessageType = 822
	Message_CONSENSUS_STATE_GET_RESPONSE        Message_MessageType = 823
	Message_CONSENSUS_BLOCKS_GET_REQUEST        Message_MessageType = 824
	Message_CONSENSUS_BLOCKS_GET_RESPONSE       Message_MessageType = 825
	Message_CONSENSUS_CHAIN_HEAD_GET_REQUEST    Message_MessageType = 826
	Message_CONSENSUS_CHAIN_HEAD_GET_RESPONSE   Message_MessageType = 827
	// Consensus notification messages
	Message_CONSENSUS_NOTIFY_PEER_CONNECTED    Message_MessageType = 900
	Message_CONSENSUS_NOTIFY_PEER_DISCONNECTED Message_MessageType = 901
	Message_CONSENSUS_NOTIFY_PEER_MESSAGE      Message_MessageType = 902
	Message_CONSENSUS_NOTIFY_BLOCK_NEW         Message_MessageType = 903
	Message_CONSENSUS_NOTIFY_BLOCK_VALID       Message_MessageType = 904
	Message_CONSENSUS_NOTIFY_BLOCK_INVALID     Message_MessageType = 905
	Message_CONSENSUS_NOTIFY_BLOCK_COMMIT      Message_MessageType = 906
	Message_CONSENSUS_NOTIFY_ACK               Message_MessageType = 999
)

var Message_MessageType_name = map[int32]string{
	0:   "DEFAULT",
	1:   "TP_REGISTER_REQUEST",
	2:   "TP_REGISTER_RESPONSE",
	3:   "TP_UNREGISTER_REQUEST",
	4:   "TP_UNREGISTER_RESPONSE",
	5:   "TP_PROCESS_REQUEST",
	6:   "TP_PROCESS_RESPONSE",
	7:   "TP_STATE_GET_REQUEST",
	8:   "TP_STATE_GET_RESPONSE",
	9:   "TP_STATE_SET_REQUEST",
	10:  "TP_STATE_SET_RESPONSE",
	11:  "TP_STATE_DELETE_REQUEST",
	12:  "TP_STATE_DELETE_RESPONSE",
	13:  "TP_RECEIPT_ADD_DATA_REQUEST",
	14:  "TP_RECEIPT_ADD_DATA_RESPONSE",
	15:  "TP_EVENT_ADD_REQUEST",
	16:  "TP_EVENT_ADD_RESPONSE",
	100: "CLIENT_BATCH_SUBMIT_REQUEST",
	101: "CLIENT_BATCH_SUBMIT_RESPONSE",
	102: "CLIENT_BLOCK_LIST_REQUEST",
	103: "CLIENT_BLOCK_LIST_RESPONSE",
	104: "CLIENT_BLOCK_GET_BY_ID_REQUEST",
	105: "CLIENT_BLOCK_GET_RESPONSE",
	106: "CLIENT_BATCH_LIST_REQUEST",
	107: "CLIENT_BATCH_LIST_RESPONSE",
	108: "CLIENT_BATCH_GET_REQUEST",
	109: "CLIENT_BATCH_GET_RESPONSE",
	110: "CLIENT_TRANSACTION_LIST_REQUEST",
	111: "CLIENT_TRANSACTION_LIST_RESPONSE",
	112: "CLIENT_TRANSACTION_GET_REQUEST",
	113: "CLIENT_TRANSACTION_GET_RESPONSE",
	114: "CLIENT_STATE_CURRENT_REQUEST",
	115: "CLIENT_STATE_CURRENT_RESPONSE",
	116: "CLIENT_STATE_LIST_REQUEST",
	117: "CLIENT_STATE_LIST_RESPONSE",
	118: "CLIENT_STATE_GET_REQUEST",
	119: "CLIENT_STATE_GET_RESPONSE",
	120: "CLIENT_BATCH_STATUS_REQUEST",
	121: "CLIENT_BATCH_STATUS_RESPONSE",
	122: "CLIENT_RECEIPT_GET_REQUEST",
	123: "CLIENT_RECEIPT_GET_RESPONSE",
	124: "CLIENT_BLOCK_GET_BY_NUM_REQUEST",
	125: "CLIENT_PEERS_GET_REQUEST",
	126: "CLIENT_PEERS_GET_RESPONSE",
	127: "CLIENT_BLOCK_GET_BY_TRANSACTION_ID_REQUEST",
	128: "CLIENT_BLOCK_GET_BY_BATCH_ID_REQUEST",
	129: "CLIENT_STATUS_GET_REQUEST",
	130: "CLIENT_STATUS_GET_RESPONSE",
	500: "CLIENT_EVENTS_SUBSCRIBE_REQUEST",
	501: "CLIENT_EVENTS_SUBSCRIBE_RESPONSE",
	502: "CLIENT_EVENTS_UNSUBSCRIBE_REQUEST",
	503: "CLIENT_EVENTS_UNSUBSCRIBE_RESPONSE",
	504: "CLIENT_EVENTS",
	505: "CLIENT_EVENTS_GET_REQUEST",
	506: "CLIENT_EVENTS_GET_RESPONSE",
	200: "GOSSIP_MESSAGE",
	201: "GOSSIP_REGISTER",
	202: "GOSSIP_UNREGISTER",
	205: "GOSSIP_BLOCK_REQUEST",
	206: "GOSSIP_BLOCK_RESPONSE",
	207: "GOSSIP_BATCH_BY_BATCH_ID_REQUEST",
	208: "GOSSIP_BATCH_BY_TRANSACTION_ID_REQUEST",
	209: "GOSSIP_BATCH_RESPONSE",
	210: "GOSSIP_GET_PEERS_REQUEST",
	211: "GOSSIP_GET_PEERS_RESPONSE",
	300: "NETWORK_ACK",
	301: "NETWORK_CONNECT",
	302: "NETWORK_DISCONNECT",
	600: "AUTHORIZATION_CONNECTION_RESPONSE",
	601: "AUTHORIZATION_VIOLATION",
	602: "AUTHORIZATION_TRUST_REQUEST",
	603: "AUTHORIZATION_TRUST_RESPONSE",
	604: "AUTHORIZATION_CHALLENGE_REQUEST",
	605: "AUTHORIZATION_CHALLENGE_RESPONSE",
	606: "AUTHORIZATION_CHALLENGE_SUBMIT",
	607: "AUTHORIZATION_CHALLENGE_RESULT",
	700: "PING_REQUEST",
	701: "PING_RESPONSE",
	800: "CONSENSUS_REGISTER_REQUEST",
	801: "CONSENSUS_REGISTER_RESPONSE",
	802: "CONSENSUS_SEND_TO_REQUEST",
	803: "CONSENSUS_SEND_TO_RESPONSE",
	804: "CONSENSUS_BROADCAST_REQUEST",
	805: "CONSENSUS_BROADCAST_RESPONSE",
	806: "CONSENSUS_INITIALIZE_BLOCK_REQUEST",
	807: "CONSENSUS_INITIALIZE_BLOCK_RESPONSE",
	808: "CONSENSUS_FINALIZE_BLOCK_REQUEST",
	809: "CONSENSUS_FINALIZE_BLOCK_RESPONSE",
	810: "CONSENSUS_CANCEL_BLOCK_REQUEST",
	811: "CONSENSUS_CANCEL_BLOCK_RESPONSE",
	812: "CONSENSUS_CHECK_BLOCKS_REQUEST",
	813: "CONSENSUS_CHECK_BLOCKS_RESPONSE",
	814: "CONSENSUS_COMMIT_BLOCK_REQUEST",
	815: "CONSENSUS_COMMIT_BLOCK_RESPONSE",
	816: "CONSENSUS_IGNORE_BLOCK_REQUEST",
	817: "CONSENSUS_IGNORE_BLOCK_RESPONSE",
	818: "CONSENSUS_FAIL_BLOCK_REQUEST",
	819: "CONSENSUS_FAIL_BLOCK_RESPONSE",
	820: "CONSENSUS_SETTINGS_GET_REQUEST",
	821: "CONSENSUS_SETTINGS_GET_RESPONSE",
	822: "CONSENSUS_STATE_GET_REQUEST",
	823: "CONSENSUS_STATE_GET_RESPONSE",
	824: "CONSENSUS_BLOCKS_GET_REQUEST",
	825: "CONSENSUS_BLOCKS_GET_RESPONSE",
	826: "CONSENSUS_CHAIN_HEAD_GET_REQUEST",
	827: "CONSENSUS_CHAIN_HEAD_GET_RESPONSE",
	900: "CONSENSUS_NOTIFY_PEER_CONNECTED",
	901: "CONSENSUS_NOTIFY_PEER_DISCONNECTED",
	902: "CONSENSUS_NOTIFY_PEER_MESSAGE",
	903: "CONSENSUS_NOTIFY_BLOCK_NEW",
	904: "CONSENSUS_NOTIFY_BLOCK_VALID",
	905: "CONSENSUS_NOTIFY_BLOCK_INVALID",
	906: "CONSENSUS_NOTIFY_BLOCK_COMMIT",
	999: "CONSENSUS_NOTIFY_ACK",
}
var Message_MessageType_value = map[string]int32{
	"DEFAULT":                                    0,
	"TP_REGISTER_REQUEST":                        1,
	"TP_REGISTER_RESPONSE":                       2,
	"TP_UNREGISTER_REQUEST":                      3,
	"TP_UNREGISTER_RESPONSE":                     4,
	"TP_PROCESS_REQUEST":                         5,
	"TP_PROCESS_RESPONSE":                        6,
	"TP_STATE_GET_REQUEST":                       7,
	"TP_STATE_GET_RESPONSE":                      8,
	"TP_STATE_SET_REQUEST":                       9,
	"TP_STATE_SET_RESPONSE":                      10,
	"TP_STATE_DELETE_REQUEST":                    11,
	"TP_STATE_DELETE_RESPONSE":                   12,
	"TP_RECEIPT_ADD_DATA_REQUEST":                13,
	"TP_RECEIPT_ADD_DATA_RESPONSE":               14,
	"TP_EVENT_ADD_REQUEST":                       15,
	"TP_EVENT_ADD_RESPONSE":                      16,
	"CLIENT_BATCH_SUBMIT_REQUEST":                100,
	"CLIENT_BATCH_SUBMIT_RESPONSE":               101,
	"CLIENT_BLOCK_LIST_REQUEST":                  102,
	"CLIENT_BLOCK_LIST_RESPONSE":                 103,
	"CLIENT_BLOCK_GET_BY_ID_REQUEST":             104,
	"CLIENT_BLOCK_GET_RESPONSE":                  105,
	"CLIENT_BATCH_LIST_REQUEST":                  106,
	"CLIENT_BATCH_LIST_RESPONSE":                 107,
	"CLIENT_BATCH_GET_REQUEST":                   108,
	"CLIENT_BATCH_GET_RESPONSE":                  109,
	"CLIENT_TRANSACTION_LIST_REQUEST":            110,
	"CLIENT_TRANSACTION_LIST_RESPONSE":           111,
	"CLIENT_TRANSACTION_GET_REQUEST":             112,
	"CLIENT_TRANSACTION_GET_RESPONSE":            113,
	"CLIENT_STATE_CURRENT_REQUEST":               114,
	"CLIENT_STATE_CURRENT_RESPONSE":              115,
	"CLIENT_STATE_LIST_REQUEST":                  116,
	"CLIENT_STATE_LIST_RESPONSE":                 117,
	"CLIENT_STATE_GET_REQUEST":                   118,
	"CLIENT_STATE_GET_RESPONSE":                  119,
	"CLIENT_BATCH_STATUS_REQUEST":                120,
	"CLIENT_BATCH_STATUS_RESPONSE":               121,
	"CLIENT_RECEIPT_GET_REQUEST":                 122,
	"CLIENT_RECEIPT_GET_RESPONSE":                123,
	"CLIENT_BLOCK_GET_BY_NUM_REQUEST":            124,
	"CLIENT_PEERS_GET_REQUEST":                   125,
	"CLIENT_PEERS_GET_RESPONSE":                  126,
	"CLIENT_BLOCK_GET_BY_TRANSACTION_ID_REQUEST": 127,
	"CLIENT_BLOCK_GET_BY_BATCH_ID_REQUEST":       128,
	"CLIENT_STATUS_GET_REQUEST":                  129,
	"CLIENT_STATUS_GET_RESPONSE":                 130,
	"CLIENT_EVENTS_SUBSCRIBE_REQUEST":            500,
	"CLIENT_EVENTS_SUBSCRIBE_RESPONSE":           501,
	"CLIENT_EVENTS_UNSUBSCRIBE_REQUEST":          502,
	"CLIENT_EVENTS_UNSUBSCRIBE_RESPONSE":         503,
	"CLIENT_EVENTS":                              504,
	"CLIENT_EVENTS_GET_REQUEST":                  505,
	"CLIENT_EVENTS_GET_RESPONSE":                 506,
	"GOSSIP_MESSAGE":                             200,
	"GOSSIP_REGISTER":                            201,
	"GOSSIP_UNREGISTER":                          202,
	"GOSSIP_BLOCK_REQUEST":                       205,
	"GOSSIP_BLOCK_RESPONSE":                      206,
	"GOSSIP_BATCH_BY_BATCH_ID_REQUEST":           207,
	"GOSSIP_BATCH_BY_TRANSACTION_ID_REQUEST":     208,
	"GOSSIP_BATCH_RESPONSE":                      209,
	"GOSSIP_GET_PEERS_REQUEST":                   210,
	"GOSSIP_GET_PEERS_RESPONSE":                  211,
	"NETWORK_ACK":                                300,
	"NETWORK_CONNECT":                            301,
	"NETWORK_DISCONNECT":                         302,
	"AUTHORIZATION_CONNECTION_RESPONSE":          600,
	"AUTHORIZATION_VIOLATION":                    601,
	"AUTHORIZATION_TRUST_REQUEST":                602,
	"AUTHORIZATION_TRUST_RESPONSE":               603,
	"AUTHORIZATION_CHALLENGE_REQUEST":            604,
	"AUTHORIZATION_CHALLENGE_RESPONSE":           605,
	"AUTHORIZATION_CHALLENGE_SUBMIT":             606,
	"AUTHORIZATION_CHALLENGE_RESULT":             607,
	"PING_REQUEST":                               700,
	"PING_RESPONSE":                              701,
	"CONSENSUS_REGISTER_REQUEST":                 800,
	"CONSENSUS_REGISTER_RESPONSE":                801,
	"CONSENSUS_SEND_TO_REQUEST":                  802,
	"CONSENSUS_SEND_TO_RESPONSE":                 803,
	"CONSENSUS_BROADCAST_REQUEST":                804,
	"CONSENSUS_BROADCAST_RESPONSE":               805,
	"CONSENSUS_INITIALIZE_BLOCK_REQUEST":         806,
	"CONSENSUS_INITIALIZE_BLOCK_RESPONSE":        807,
	"CONSENSUS_FINALIZE_BLOCK_REQUEST":           808,
	"CONSENSUS_FINALIZE_BLOCK_RESPONSE":          809,
	"CONSENSUS_CANCEL_BLOCK_REQUEST":             810,
	"CONSENSUS_CANCEL_BLOCK_RESPONSE":            811,
	"CONSENSUS_CHECK_BLOCKS_REQUEST":             812,
	"CONSENSUS_CHECK_BLOCKS_RESPONSE":            813,
	"CONSENSUS_COMMIT_BLOCK_REQUEST":             814,
	"CONSENSUS_COMMIT_BLOCK_RESPONSE":            815,
	"CONSENSUS_IGNORE_BLOCK_REQUEST":             816,
	"CONSENSUS_IGNORE_BLOCK_RESPONSE":            817,
	"CONSENSUS_FAIL_BLOCK_REQUEST":               818,
	"CONSENSUS_FAIL_BLOCK_RESPONSE":              819,
	"CONSENSUS_SETTINGS_GET_REQUEST":             820,
	"CONSENSUS_SETTINGS_GET_RESPONSE":            821,
	"CONSENSUS_STATE_GET_REQUEST":                822,
	"CONSENSUS_STATE_GET_RESPONSE":               823,
	"CONSENSUS_BLOCKS_GET_REQUEST":               824,
	"CONSENSUS_BLOCKS_GET_RESPONSE":              825,
	"CONSENSUS_CHAIN_HEAD_GET_REQUEST":           826,
	"CONSENSUS_CHAIN_HEAD_GET_RESPONSE":          827,
	"CONSENSUS_NOTIFY_PEER_CONNECTED":            900,
	"CONSENSUS_NOTIFY_PEER_DISCONNECTED":         901,
	"CONSENSUS_NOTIFY_PEER_MESSAGE":              902,
	"CONSENSUS_NOTIFY_BLOCK_NEW":                 903,
	"CONSENSUS_NOTIFY_BLOCK_VALID":               904,
	"CONSENSUS_NOTIFY_BLOCK_INVALID":             905,
	"CONSENSUS_NOTIFY_BLOCK_COMMIT":              906,
	"CONSENSUS_NOTIFY_ACK":                       999,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_validator_52afef5ca12f08ff, []int{1, 0}
}

// A list of messages to be transmitted together.
type MessageList struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MessageList) Reset()         { *m = MessageList{} }
func (m *MessageList) String() string { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()    {}
func (*MessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_validator_52afef5ca12f08ff, []int{0}
}
func (m *MessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList.Unmarshal(m, b)
}
func (m *MessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList.Marshal(b, m, deterministic)
}
func (dst *MessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList.Merge(dst, src)
}
func (m *MessageList) XXX_Size() int {
	return xxx_messageInfo_MessageList.Size(m)
}
func (m *MessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList proto.InternalMessageInfo

func (m *MessageList) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// The message passed between the validator and client, containing the
// header fields and content.
type Message struct {
	// The type of message, used to determine how to 'route' the message
	// to the appropriate handler as well as how to deserialize the
	// content.
	MessageType Message_MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,enum=Message_MessageType" json:"message_type,omitempty"`
	// The identifier used to correlate response messages to their related
	// request messages.  correlation_id should be set to a random string
	// for messages which are not responses to previously sent messages.  For
	// response messages, correlation_id should be set to the same string as
	// contained in the request message.
	CorrelationId string `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId" json:"correlation_id,omitempty"`
	// The content of the message, defined by message_type.  In many
	// cases, this data has been serialized with Protocol Buffers or
	// CBOR.
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_validator_52afef5ca12f08ff, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_DEFAULT
}

func (m *Message) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageList)(nil), "MessageList")
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/validator_pb2/validator.proto", fileDescriptor_validator_52afef5ca12f08ff)
}

var fileDescriptor_validator_52afef5ca12f08ff = []byte{
	// 1368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x57, 0xd9, 0x73, 0xd3, 0x46,
	0x18, 0xaf, 0x62, 0x23, 0xc3, 0xe6, 0x40, 0x2c, 0x81, 0x18, 0x48, 0x82, 0x09, 0x81, 0xa6, 0xed,
	0x4c, 0x98, 0x81, 0x87, 0x3e, 0x2b, 0xf2, 0xc6, 0xd9, 0x89, 0x22, 0xb9, 0xda, 0x15, 0x0c, 0xbc,
	0x68, 0x02, 0x36, 0x90, 0x02, 0x71, 0x1a, 0x1b, 0x28, 0xbd, 0xdb, 0xe9, 0xfd, 0x57, 0xf4, 0x3e,
	0x39, 0x7a, 0xd1, 0xbb, 0xef, 0x6d, 0x67, 0x7a, 0xbf, 0xf4, 0xee, 0x63, 0xff, 0x81, 0xde, 0x4f,
	0x1d, 0x49, 0x9f, 0xa4, 0xdd, 0x95, 0xcc, 0x93, 0xad, 0xfd, 0xfd, 0xbe, 0xdf, 0x77, 0xee, 0x6a,
	0x85, 0x0e, 0x75, 0x97, 0x2f, 0xf6, 0x3a, 0x9d, 0xde, 0xe9, 0xa0, 0xdb, 0x3a, 0x73, 0x60, 0x6d,
	0xbd, 0xd3, 0xeb, 0x1c, 0x3f, 0x7f, 0xf2, 0xc0, 0x85, 0xe5, 0xb3, 0x2b, 0xad, 0xe5, 0x5e, 0x67,
	0x3d, 0x58, 0x3b, 0x7e, 0x30, 0x7b, 0x9a, 0x8d, 0x08, 0x53, 0x87, 0xd0, 0xe0, 0x52, 0xbb, 0xdb,
	0x5d, 0x3e, 0xd5, 0xb6, 0x57, 0xba, 0x3d, 0x3c, 0x8d, 0x36, 0x9e, 0x8b, 0x1f, 0xbb, 0x55, 0xad,
	0x56, 0x9a, 0x19, 0x3c, 0xb8, 0x71, 0x16, 0x70, 0x2f, 0x45, 0xa6, 0xae, 0x4f, 0xa2, 0x0a, 0xac,
	0xe2, 0xdb, 0xd1, 0x10, 0xac, 0x07, 0xbd, 0x4b, 0x6b, 0xed, 0xaa, 0x56, 0xd3, 0x66, 0x46, 0x0e,
	0x8e, 0x26, 0x56, 0xc9, 0x2f, 0xbf, 0xb4, 0xd6, 0xf6, 0x06, 0xcf, 0x65, 0x0f, 0x78, 0x1f, 0x1a,
	0x39, 0xd1, 0x59, 0x5f, 0x6f, 0x9f, 0x5d, 0xee, 0xad, 0x74, 0x56, 0x83, 0x95, 0x56, 0x75, 0xa0,
	0xa6, 0xcd, 0x6c, 0xf2, 0x86, 0x85, 0x55, 0xda, 0xc2, 0x55, 0x54, 0x39, 0xd1, 0x59, 0xed, 0xb5,
	0x57, 0x7b, 0xd5, 0x52, 0x4d, 0x9b, 0x19, 0xf2, 0x92, 0xc7, 0xa9, 0x6b, 0x13, 0x69, 0xec, 0x91,
	0xe0, 0x20, 0xaa, 0xd4, 0xc9, 0xbc, 0xe9, 0xdb, 0xdc, 0xb8, 0x09, 0x8f, 0xa1, 0xad, 0xbc, 0x19,
	0x78, 0xa4, 0x41, 0x19, 0x27, 0x5e, 0xe0, 0x91, 0x3b, 0x7c, 0xc2, 0xb8, 0xa1, 0xe1, 0x2a, 0x1a,
	0x95, 0x01, 0xd6, 0x74, 0x1d, 0x46, 0x8c, 0x01, 0xbc, 0x03, 0x6d, 0xe3, 0xcd, 0xc0, 0x77, 0x72,
	0x46, 0x25, 0xbc, 0x13, 0x6d, 0x57, 0x21, 0x30, 0x2b, 0xe3, 0xed, 0x08, 0xf3, 0x66, 0xd0, 0xf4,
	0x5c, 0x8b, 0x30, 0x96, 0xda, 0x6c, 0x80, 0x08, 0xb2, 0x75, 0x30, 0xd0, 0x21, 0x02, 0xc6, 0x4d,
	0x4e, 0x82, 0x06, 0xe1, 0xa9, 0x49, 0x05, 0x22, 0x10, 0x11, 0x30, 0xda, 0x28, 0x19, 0x31, 0xc1,
	0x68, 0x93, 0x64, 0xc4, 0x44, 0x23, 0x84, 0x77, 0xa1, 0xb1, 0x14, 0xaa, 0x13, 0x9b, 0x70, 0x92,
	0xda, 0x0d, 0xe2, 0x71, 0x54, 0xcd, 0x83, 0x60, 0x3a, 0x84, 0x77, 0xa3, 0x5d, 0x51, 0x99, 0x2c,
	0x42, 0x9b, 0x3c, 0x30, 0xeb, 0xf5, 0xa0, 0x6e, 0x72, 0x33, 0x35, 0x1f, 0xc6, 0x35, 0x34, 0x5e,
	0x4c, 0x00, 0x89, 0x11, 0x08, 0x99, 0x1c, 0x26, 0x4e, 0x8c, 0x27, 0xb6, 0x9b, 0x21, 0x64, 0x11,
	0x01, 0x23, 0x23, 0xf4, 0x6b, 0xd9, 0x34, 0x04, 0xe6, 0x4c, 0x6e, 0x2d, 0x04, 0xcc, 0x9f, 0x5b,
	0xa2, 0x59, 0xba, 0xad, 0xd0, 0x6f, 0x31, 0x01, 0x24, 0xda, 0x78, 0x02, 0xed, 0x48, 0x18, 0xb6,
	0x6b, 0x2d, 0x06, 0x36, 0x65, 0x99, 0xc0, 0x49, 0x3c, 0x89, 0x76, 0x16, 0xc1, 0x60, 0x7e, 0x0a,
	0x4f, 0xa1, 0x49, 0x09, 0x0f, 0x1b, 0x31, 0x77, 0x34, 0xa0, 0x59, 0x02, 0xa7, 0x73, 0x2e, 0xa4,
	0x66, 0xad, 0x88, 0x70, 0x14, 0xa3, 0x14, 0xc1, 0x9d, 0x62, 0x04, 0x22, 0x0c, 0xe6, 0x67, 0xc2,
	0xce, 0x48, 0xb8, 0x38, 0x24, 0x67, 0x73, 0xe2, 0x92, 0xef, 0x73, 0x78, 0x2f, 0xda, 0x0d, 0x30,
	0xf7, 0x4c, 0x87, 0x99, 0x16, 0xa7, 0xae, 0x23, 0x47, 0xb0, 0x8a, 0xa7, 0x51, 0xad, 0x3f, 0x09,
	0xa4, 0x3a, 0x42, 0x25, 0x44, 0x96, 0x18, 0xcd, 0x5a, 0x1f, 0x77, 0x52, 0x4c, 0x77, 0x09, 0x3d,
	0x8b, 0xc7, 0xcd, 0xf2, 0x3d, 0x2f, 0x7c, 0x4a, 0x64, 0xd6, 0xf1, 0x1e, 0x34, 0xd1, 0x87, 0x01,
	0x22, 0x5d, 0x21, 0xef, 0x98, 0x22, 0xa5, 0xd4, 0x13, 0x8a, 0x2a, 0xc1, 0x60, 0x7e, 0x5e, 0x28,
	0x6a, 0x7e, 0xe7, 0x5d, 0xc8, 0x89, 0x4b, 0x09, 0x5c, 0xcc, 0x4f, 0x25, 0x37, 0xb9, 0x9f, 0x6d,
	0xf6, 0xbb, 0xf3, 0x53, 0x99, 0x10, 0x40, 0xe2, 0x92, 0x10, 0x5f, 0xb2, 0x67, 0xc4, 0x08, 0xee,
	0x11, 0x5c, 0xc8, 0x38, 0x08, 0xdc, 0x2b, 0x54, 0x5a, 0x9a, 0x4b, 0xc7, 0x5f, 0x4a, 0x55, 0xee,
	0x13, 0xb2, 0x6c, 0x12, 0xe2, 0x31, 0xc9, 0xc7, 0xfd, 0x42, 0x96, 0x22, 0x0a, 0x1e, 0x1e, 0xc0,
	0xb3, 0xe8, 0xd6, 0x22, 0x0f, 0x62, 0x63, 0x85, 0x5d, 0xf0, 0x20, 0xbe, 0x05, 0x4d, 0x17, 0xf1,
	0xe3, 0x0a, 0x08, 0xcc, 0x87, 0x34, 0x3c, 0x29, 0xd5, 0xd7, 0x97, 0x03, 0x7b, 0x58, 0xc3, 0xbb,
	0xa5, 0xee, 0xf9, 0x4a, 0x68, 0x8f, 0x68, 0x78, 0x3a, 0xcd, 0x3e, 0x3a, 0x36, 0x58, 0xb8, 0xef,
	0x99, 0xe5, 0xd1, 0xb9, 0xec, 0x48, 0xfb, 0xa3, 0x84, 0xf7, 0xa5, 0x73, 0x5d, 0xc0, 0x02, 0xb1,
	0x3f, 0x4b, 0x78, 0x3f, 0xda, 0x23, 0xd3, 0x7c, 0x27, 0x2f, 0xf7, 0x57, 0x09, 0xdf, 0x8c, 0xa6,
	0x6e, 0xc4, 0x03, 0xc1, 0xbf, 0x4b, 0x18, 0xa3, 0x61, 0x89, 0x68, 0xfc, 0x53, 0x12, 0x52, 0x06,
	0x63, 0x31, 0xe5, 0x7f, 0x4b, 0x42, 0xca, 0x12, 0x0e, 0xa2, 0xff, 0x95, 0xf0, 0x56, 0x34, 0xd2,
	0x70, 0x19, 0xa3, 0xcd, 0x60, 0x89, 0x30, 0x66, 0x36, 0x88, 0xf1, 0xa9, 0x86, 0x47, 0xd1, 0x66,
	0x58, 0x4c, 0xde, 0x45, 0xc6, 0x67, 0x1a, 0xde, 0x8e, 0xb6, 0xc0, 0x6a, 0xf6, 0x8e, 0x32, 0x3e,
	0xd7, 0xf0, 0x0e, 0x34, 0x0a, 0xeb, 0x71, 0x87, 0x12, 0xf7, 0x5f, 0x68, 0x78, 0x27, 0xda, 0xa6,
	0x40, 0xe0, 0xf9, 0x4b, 0x2d, 0x2c, 0x63, 0x82, 0x45, 0xbd, 0x2c, 0x6a, 0xea, 0x57, 0x1a, 0xbe,
	0x0d, 0xed, 0x57, 0x69, 0x7d, 0x66, 0xe5, 0x6b, 0xc9, 0x5f, 0x44, 0x4e, 0xfd, 0x7d, 0xa3, 0xe1,
	0x09, 0x54, 0x05, 0x2c, 0xac, 0x41, 0x3c, 0x9b, 0x89, 0xe9, 0xb7, 0xd1, 0xf0, 0x14, 0xc0, 0x60,
	0xfe, 0x9d, 0x86, 0x0d, 0x34, 0xe8, 0x10, 0x7e, 0xc4, 0xf5, 0x16, 0x03, 0xd3, 0x5a, 0x34, 0x2e,
	0x0f, 0x84, 0x55, 0x4a, 0x56, 0x2c, 0xd7, 0x71, 0x88, 0xc5, 0x8d, 0x2b, 0x03, 0x78, 0x0c, 0xe1,
	0x64, 0xb5, 0x4e, 0x59, 0x02, 0x5c, 0x1d, 0x08, 0xe7, 0xc1, 0xf4, 0xf9, 0x82, 0xeb, 0xd1, 0x63,
	0x66, 0x14, 0x3a, 0x60, 0xe1, 0xdf, 0xd4, 0xd1, 0xf7, 0x65, 0x3c, 0x8e, 0xc6, 0x64, 0xde, 0x61,
	0xea, 0xda, 0xd1, 0x3f, 0xe3, 0x87, 0x32, 0xae, 0xa1, 0x5d, 0x32, 0xca, 0x3d, 0x5f, 0x38, 0xa2,
	0x7e, 0x2c, 0xe3, 0x3d, 0x68, 0xbc, 0x98, 0x01, 0x2e, 0x7e, 0x2a, 0x87, 0x73, 0xae, 0x84, 0xb2,
	0x60, 0xda, 0x36, 0x71, 0x1a, 0xd9, 0x60, 0xfe, 0x5c, 0x0e, 0x1b, 0xd4, 0x9f, 0x05, 0x62, 0xbf,
	0x94, 0xf1, 0x5e, 0x34, 0xd9, 0x8f, 0x16, 0xbf, 0x36, 0x8d, 0x5f, 0x6f, 0x48, 0xf2, 0x08, 0x0b,
	0x6f, 0x53, 0xbf, 0x95, 0xf1, 0x16, 0x34, 0xd4, 0xa4, 0x4e, 0x23, 0x8d, 0xe1, 0xe3, 0x0d, 0xe1,
	0xcc, 0xc3, 0x12, 0x38, 0xfc, 0x64, 0x43, 0x34, 0xd3, 0xe1, 0x7f, 0x87, 0x45, 0x87, 0x9f, 0x72,
	0x8f, 0x7a, 0x46, 0x0f, 0x6b, 0x54, 0x48, 0x00, 0x89, 0x67, 0xf5, 0x68, 0xdb, 0xa4, 0x0c, 0x46,
	0x9c, 0x7a, 0xc0, 0xdd, 0x54, 0xe1, 0x39, 0x5d, 0x76, 0x91, 0xe1, 0x20, 0xf0, 0xbc, 0xe2, 0x62,
	0xce, 0x73, 0xcd, 0xba, 0x65, 0x0a, 0x6d, 0x78, 0x41, 0x0f, 0xdb, 0x50, 0xcc, 0x00, 0x91, 0x17,
	0xf5, 0x68, 0xe7, 0xa7, 0x14, 0xea, 0x50, 0x4e, 0x4d, 0x9b, 0x1e, 0x23, 0xca, 0x36, 0x7a, 0x49,
	0xc7, 0x33, 0x68, 0xef, 0x0d, 0x89, 0x20, 0xf9, 0xb2, 0x1e, 0x9d, 0x4d, 0x29, 0x73, 0x9e, 0x3a,
	0x45, 0x82, 0xaf, 0xe8, 0xd1, 0xd9, 0xd4, 0x9f, 0x06, 0x72, 0xaf, 0xea, 0x61, 0xdb, 0x32, 0x9e,
	0x65, 0x3a, 0x16, 0xb1, 0x15, 0xb1, 0xd7, 0xf4, 0xe8, 0xd4, 0xec, 0x47, 0x02, 0xa9, 0xd7, 0x55,
	0xa9, 0x05, 0x62, 0x2d, 0xc6, 0xa4, 0x6c, 0x13, 0x5e, 0x56, 0xa5, 0x64, 0x12, 0x48, 0x5d, 0x51,
	0xa5, 0xdc, 0xa5, 0xf0, 0x6a, 0x26, 0x47, 0x75, 0x55, 0x95, 0x92, 0x49, 0x20, 0x75, 0x4d, 0x91,
	0xa2, 0x0d, 0xc7, 0xf5, 0xd4, 0x6a, 0xbd, 0xa1, 0x48, 0x29, 0x24, 0x90, 0x7a, 0x53, 0x69, 0xf8,
	0xbc, 0x49, 0xd5, 0x4a, 0xbd, 0xa5, 0xe3, 0x29, 0x34, 0xd1, 0x87, 0x02, 0x32, 0x6f, 0x2b, 0x11,
	0x31, 0xc2, 0x39, 0x75, 0x1a, 0xf2, 0xb1, 0xfe, 0x8e, 0x12, 0x91, 0x42, 0x02, 0xa9, 0xeb, 0xca,
	0x90, 0xe6, 0x2f, 0x24, 0xef, 0x2a, 0x31, 0x17, 0x5c, 0x4a, 0xde, 0x53, 0xe7, 0x38, 0x6e, 0x86,
	0xa8, 0xf2, 0xbe, 0x92, 0x96, 0x44, 0x01, 0x99, 0x0f, 0x94, 0xc1, 0xb4, 0x16, 0x4c, 0xea, 0x04,
	0x0b, 0xc4, 0xac, 0x4b, 0x52, 0x1f, 0x2a, 0x83, 0x99, 0xa3, 0x81, 0xdc, 0x47, 0x4a, 0x01, 0x1c,
	0x97, 0xd3, 0xf9, 0xa3, 0xd1, 0x99, 0x9d, 0x1c, 0xaa, 0xa4, 0x6e, 0x3c, 0x5a, 0x91, 0x37, 0x98,
	0xc8, 0xca, 0x4e, 0x66, 0x52, 0x37, 0x1e, 0xab, 0xc8, 0x19, 0x88, 0xc4, 0xe4, 0xa5, 0xf8, 0x78,
	0x45, 0x3e, 0x13, 0x80, 0x13, 0xb7, 0xcf, 0x21, 0x47, 0x8c, 0x27, 0x2a, 0x72, 0xa5, 0x24, 0xc2,
	0x61, 0xd3, 0xa6, 0x75, 0xe3, 0xc9, 0x8a, 0xdc, 0x5c, 0x89, 0x42, 0x9d, 0x98, 0xf4, 0x54, 0x71,
	0x30, 0x31, 0x29, 0x1e, 0x63, 0xe3, 0xe9, 0xf0, 0x23, 0x6e, 0x34, 0xc7, 0x09, 0x5f, 0x4b, 0xbf,
	0x57, 0xe6, 0x66, 0xd0, 0xb6, 0xe4, 0x1b, 0x7d, 0xb6, 0xdb, 0x3a, 0x33, 0x9b, 0x7c, 0xa3, 0x37,
	0xb5, 0x63, 0xc3, 0xd2, 0x67, 0xfa, 0x71, 0x3d, 0x82, 0x0e, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0xc1, 0x84, 0xcf, 0xd4, 0x0f, 0x00, 0x00,
}