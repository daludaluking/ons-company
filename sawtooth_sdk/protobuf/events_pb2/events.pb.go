// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/events_pb2/events.proto

package events_pb2

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

type EventFilter_FilterType int32

const (
	EventFilter_FILTER_TYPE_UNSET EventFilter_FilterType = 0
	EventFilter_SIMPLE_ANY        EventFilter_FilterType = 1
	EventFilter_SIMPLE_ALL        EventFilter_FilterType = 2
	EventFilter_REGEX_ANY         EventFilter_FilterType = 3
	EventFilter_REGEX_ALL         EventFilter_FilterType = 4
)

var EventFilter_FilterType_name = map[int32]string{
	0: "FILTER_TYPE_UNSET",
	1: "SIMPLE_ANY",
	2: "SIMPLE_ALL",
	3: "REGEX_ANY",
	4: "REGEX_ALL",
}
var EventFilter_FilterType_value = map[string]int32{
	"FILTER_TYPE_UNSET": 0,
	"SIMPLE_ANY":        1,
	"SIMPLE_ALL":        2,
	"REGEX_ANY":         3,
	"REGEX_ALL":         4,
}

func (x EventFilter_FilterType) String() string {
	return proto.EnumName(EventFilter_FilterType_name, int32(x))
}
func (EventFilter_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_ae7aac43f6f15367, []int{2, 0}
}

type Event struct {
	// Used to subscribe to events and servers as a hint for how to deserialize
	// event_data and what pairs to expect in attributes.
	EventType  string             `protobuf:"bytes,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Attributes []*Event_Attribute `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty"`
	// Opaque data defined by the event_type.
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_ae7aac43f6f15367, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetAttributes() []*Event_Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Transparent data defined by the event_type.
type Event_Attribute struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Attribute) Reset()         { *m = Event_Attribute{} }
func (m *Event_Attribute) String() string { return proto.CompactTextString(m) }
func (*Event_Attribute) ProtoMessage()    {}
func (*Event_Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_ae7aac43f6f15367, []int{0, 0}
}
func (m *Event_Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Attribute.Unmarshal(m, b)
}
func (m *Event_Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Attribute.Marshal(b, m, deterministic)
}
func (dst *Event_Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Attribute.Merge(dst, src)
}
func (m *Event_Attribute) XXX_Size() int {
	return xxx_messageInfo_Event_Attribute.Size(m)
}
func (m *Event_Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Attribute proto.InternalMessageInfo

func (m *Event_Attribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event_Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type EventList struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventList) Reset()         { *m = EventList{} }
func (m *EventList) String() string { return proto.CompactTextString(m) }
func (*EventList) ProtoMessage()    {}
func (*EventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_ae7aac43f6f15367, []int{1}
}
func (m *EventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventList.Unmarshal(m, b)
}
func (m *EventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventList.Marshal(b, m, deterministic)
}
func (dst *EventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventList.Merge(dst, src)
}
func (m *EventList) XXX_Size() int {
	return xxx_messageInfo_EventList.Size(m)
}
func (m *EventList) XXX_DiscardUnknown() {
	xxx_messageInfo_EventList.DiscardUnknown(m)
}

var xxx_messageInfo_EventList proto.InternalMessageInfo

func (m *EventList) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventFilter struct {
	// EventFilter is used when subscribing to events to limit the events
	// received within a given event type. See
	// validator/server/events/subscription.py for further explanation.
	Key                  string                 `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	MatchString          string                 `protobuf:"bytes,2,opt,name=match_string,json=matchString" json:"match_string,omitempty"`
	FilterType           EventFilter_FilterType `protobuf:"varint,3,opt,name=filter_type,json=filterType,enum=EventFilter_FilterType" json:"filter_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_ae7aac43f6f15367, []int{2}
}
func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (dst *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(dst, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EventFilter) GetMatchString() string {
	if m != nil {
		return m.MatchString
	}
	return ""
}

func (m *EventFilter) GetFilterType() EventFilter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return EventFilter_FILTER_TYPE_UNSET
}

type EventSubscription struct {
	// EventSubscription is used when subscribing to events to specify the type
	// of events being subscribed to, along with any additional filters. See
	// validator/server/events/subscription.py for further explanation.
	EventType            string         `protobuf:"bytes,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Filters              []*EventFilter `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EventSubscription) Reset()         { *m = EventSubscription{} }
func (m *EventSubscription) String() string { return proto.CompactTextString(m) }
func (*EventSubscription) ProtoMessage()    {}
func (*EventSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_ae7aac43f6f15367, []int{3}
}
func (m *EventSubscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSubscription.Unmarshal(m, b)
}
func (m *EventSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSubscription.Marshal(b, m, deterministic)
}
func (dst *EventSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSubscription.Merge(dst, src)
}
func (m *EventSubscription) XXX_Size() int {
	return xxx_messageInfo_EventSubscription.Size(m)
}
func (m *EventSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_EventSubscription proto.InternalMessageInfo

func (m *EventSubscription) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventSubscription) GetFilters() []*EventFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*Event_Attribute)(nil), "Event.Attribute")
	proto.RegisterType((*EventList)(nil), "EventList")
	proto.RegisterType((*EventFilter)(nil), "EventFilter")
	proto.RegisterType((*EventSubscription)(nil), "EventSubscription")
	proto.RegisterEnum("EventFilter_FilterType", EventFilter_FilterType_name, EventFilter_FilterType_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/events_pb2/events.proto", fileDescriptor_events_ae7aac43f6f15367)
}

var fileDescriptor_events_ae7aac43f6f15367 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x6a, 0xea, 0x40,
	0x10, 0x86, 0xcf, 0x1a, 0xf5, 0x90, 0x89, 0x47, 0xe2, 0x72, 0xa4, 0xa1, 0xd0, 0x62, 0x03, 0x2d,
	0x42, 0x69, 0x2c, 0x7a, 0xd3, 0x5b, 0x0b, 0xb1, 0x08, 0xa9, 0xc8, 0x26, 0x85, 0xea, 0x4d, 0x48,
	0x74, 0xad, 0x41, 0x6b, 0x42, 0x76, 0x63, 0xf1, 0x71, 0xfa, 0x64, 0x7d, 0x95, 0x92, 0x4d, 0xa2,
	0xb9, 0x28, 0xf4, 0x2a, 0x33, 0xdf, 0xbf, 0x9b, 0xf9, 0x67, 0x67, 0xe0, 0x8e, 0x79, 0x1f, 0x3c,
	0x0c, 0xf9, 0xda, 0x65, 0xcb, 0x4d, 0x2f, 0x8a, 0x43, 0x1e, 0xfa, 0xc9, 0xaa, 0x47, 0xf7, 0x74,
	0xc7, 0x99, 0x1b, 0xf9, 0xfd, 0x3c, 0x34, 0x84, 0xa4, 0x7f, 0x22, 0xa8, 0x99, 0x29, 0xc0, 0x17,
	0x00, 0x42, 0x71, 0xf9, 0x21, 0xa2, 0x1a, 0xea, 0xa0, 0xae, 0x4c, 0x64, 0x41, 0x9c, 0x43, 0x44,
	0xf1, 0x3d, 0x80, 0xc7, 0x79, 0x1c, 0xf8, 0x09, 0xa7, 0x4c, 0xab, 0x74, 0xa4, 0xae, 0xd2, 0x57,
	0x0d, 0x71, 0xd5, 0x18, 0x16, 0x02, 0x29, 0x9d, 0xc1, 0x18, 0xaa, 0x4b, 0x8f, 0x7b, 0x9a, 0xd4,
	0x41, 0xdd, 0x06, 0x11, 0xf1, 0xf9, 0x00, 0xe4, 0xe3, 0x61, 0xac, 0x82, 0xb4, 0xa1, 0x87, 0xbc,
	0x54, 0x1a, 0xe2, 0xff, 0x50, 0xdb, 0x7b, 0xdb, 0x84, 0x6a, 0x15, 0xc1, 0xb2, 0x44, 0xbf, 0x05,
	0x59, 0xd4, 0xb1, 0x02, 0xc6, 0xf1, 0x25, 0xd4, 0xb3, 0x06, 0x34, 0x24, 0x3c, 0xd4, 0x33, 0x0f,
	0x24, 0xa7, 0xfa, 0x17, 0x02, 0x45, 0x90, 0x51, 0xb0, 0xe5, 0x34, 0xfe, 0xa1, 0xc8, 0x15, 0x34,
	0xde, 0x3d, 0xbe, 0x58, 0xbb, 0x8c, 0xc7, 0xc1, 0xee, 0x2d, 0xaf, 0xa5, 0x08, 0x66, 0x0b, 0x84,
	0x1f, 0x40, 0x59, 0x89, 0xeb, 0xd9, 0x63, 0xa4, 0x1d, 0x34, 0xfb, 0x67, 0x46, 0xe9, 0xbf, 0x46,
	0xf6, 0x49, 0x9f, 0x86, 0xc0, 0xea, 0x18, 0xeb, 0x1e, 0xc0, 0x49, 0xc1, 0x6d, 0x68, 0x8d, 0xc6,
	0x96, 0x63, 0x12, 0xd7, 0x99, 0x4d, 0x4d, 0xf7, 0x65, 0x62, 0x9b, 0x8e, 0xfa, 0x07, 0x37, 0x01,
	0xec, 0xf1, 0xf3, 0xd4, 0x32, 0xdd, 0xe1, 0x64, 0xa6, 0xa2, 0x72, 0x6e, 0x59, 0x6a, 0x05, 0xff,
	0x03, 0x99, 0x98, 0x4f, 0xe6, 0xab, 0x90, 0xa5, 0x52, 0x6a, 0x59, 0x6a, 0x55, 0x9f, 0x43, 0x4b,
	0x18, 0xb1, 0x13, 0x9f, 0x2d, 0xe2, 0x20, 0xe2, 0x41, 0xb8, 0xfb, 0x6d, 0x7a, 0x37, 0xf0, 0x37,
	0x33, 0x59, 0x8c, 0xae, 0x51, 0x6e, 0x86, 0x14, 0xe2, 0xe3, 0x35, 0xb4, 0x8b, 0xfd, 0x31, 0xd8,
	0x72, 0x63, 0x14, 0xfb, 0x33, 0x45, 0x73, 0x38, 0xad, 0x90, 0x5f, 0x17, 0x7c, 0xf0, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x30, 0x48, 0x75, 0x07, 0x6d, 0x02, 0x00, 0x00,
}
