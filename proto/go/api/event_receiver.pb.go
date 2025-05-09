// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/event_receiver.proto

package eventv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReceiveEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiveEventRequest) Reset() {
	*x = ReceiveEventRequest{}
	mi := &file_api_event_receiver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiveEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventRequest) ProtoMessage() {}

func (x *ReceiveEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_event_receiver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventRequest.ProtoReflect.Descriptor instead.
func (*ReceiveEventRequest) Descriptor() ([]byte, []int) {
	return file_api_event_receiver_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveEventRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReceiveEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiveEventResponse) Reset() {
	*x = ReceiveEventResponse{}
	mi := &file_api_event_receiver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiveEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventResponse) ProtoMessage() {}

func (x *ReceiveEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_event_receiver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventResponse.ProtoReflect.Descriptor instead.
func (*ReceiveEventResponse) Descriptor() ([]byte, []int) {
	return file_api_event_receiver_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_event_receiver_proto protoreflect.FileDescriptor

const file_api_event_receiver_proto_rawDesc = "" +
	"\n" +
	"\x18api/event_receiver.proto\x12\bevent.v1\"/\n" +
	"\x13ReceiveEventRequest\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"0\n" +
	"\x14ReceiveEventResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2e\n" +
	"\x14EventReceiverService\x12M\n" +
	"\fReceiveEvent\x12\x1d.event.v1.ReceiveEventRequest\x1a\x1e.event.v1.ReceiveEventResponseB\x16Z\x14proto/go/api;eventv1b\x06proto3"

var (
	file_api_event_receiver_proto_rawDescOnce sync.Once
	file_api_event_receiver_proto_rawDescData []byte
)

func file_api_event_receiver_proto_rawDescGZIP() []byte {
	file_api_event_receiver_proto_rawDescOnce.Do(func() {
		file_api_event_receiver_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_event_receiver_proto_rawDesc), len(file_api_event_receiver_proto_rawDesc)))
	})
	return file_api_event_receiver_proto_rawDescData
}

var file_api_event_receiver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_event_receiver_proto_goTypes = []any{
	(*ReceiveEventRequest)(nil),  // 0: event.v1.ReceiveEventRequest
	(*ReceiveEventResponse)(nil), // 1: event.v1.ReceiveEventResponse
}
var file_api_event_receiver_proto_depIdxs = []int32{
	0, // 0: event.v1.EventReceiverService.ReceiveEvent:input_type -> event.v1.ReceiveEventRequest
	1, // 1: event.v1.EventReceiverService.ReceiveEvent:output_type -> event.v1.ReceiveEventResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_event_receiver_proto_init() }
func file_api_event_receiver_proto_init() {
	if File_api_event_receiver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_event_receiver_proto_rawDesc), len(file_api_event_receiver_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_event_receiver_proto_goTypes,
		DependencyIndexes: file_api_event_receiver_proto_depIdxs,
		MessageInfos:      file_api_event_receiver_proto_msgTypes,
	}.Build()
	File_api_event_receiver_proto = out.File
	file_api_event_receiver_proto_goTypes = nil
	file_api_event_receiver_proto_depIdxs = nil
}
