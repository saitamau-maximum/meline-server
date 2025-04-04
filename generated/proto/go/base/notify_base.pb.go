// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: base/notify_base.proto

package base

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotifyMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int64 `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
}

func (x *NotifyMeta) Reset() {
	*x = NotifyMeta{}
	mi := &file_base_notify_base_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMeta) ProtoMessage() {}

func (x *NotifyMeta) ProtoReflect() protoreflect.Message {
	mi := &file_base_notify_base_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMeta.ProtoReflect.Descriptor instead.
func (*NotifyMeta) Descriptor() ([]byte, []int) {
	return file_base_notify_base_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyMeta) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ChannelId string   `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	mi := &file_base_notify_base_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_base_notify_base_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_base_notify_base_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Payload) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

var File_base_notify_base_proto protoreflect.FileDescriptor

var file_base_notify_base_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x17,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x51,
	0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x42, 0x87, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x42, 0x0f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61,
	0x69, 0x74, 0x61, 0x6d, 0x61, 0x75, 0x2d, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2f, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0xa2, 0x02, 0x03,
	0x42, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0xca, 0x02, 0x04, 0x42, 0x61, 0x73,
	0x65, 0xe2, 0x02, 0x10, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_base_notify_base_proto_rawDescOnce sync.Once
	file_base_notify_base_proto_rawDescData = file_base_notify_base_proto_rawDesc
)

func file_base_notify_base_proto_rawDescGZIP() []byte {
	file_base_notify_base_proto_rawDescOnce.Do(func() {
		file_base_notify_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_notify_base_proto_rawDescData)
	})
	return file_base_notify_base_proto_rawDescData
}

var file_base_notify_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_base_notify_base_proto_goTypes = []any{
	(*NotifyMeta)(nil), // 0: base.NotifyMeta
	(*Payload)(nil),    // 1: base.Payload
	(*Message)(nil),    // 2: base.Message
}
var file_base_notify_base_proto_depIdxs = []int32{
	2, // 0: base.Payload.message:type_name -> base.Message
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_base_notify_base_proto_init() }
func file_base_notify_base_proto_init() {
	if File_base_notify_base_proto != nil {
		return
	}
	file_base_message_base_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_notify_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_notify_base_proto_goTypes,
		DependencyIndexes: file_base_notify_base_proto_depIdxs,
		MessageInfos:      file_base_notify_base_proto_msgTypes,
	}.Build()
	File_base_notify_base_proto = out.File
	file_base_notify_base_proto_rawDesc = nil
	file_base_notify_base_proto_goTypes = nil
	file_base_notify_base_proto_depIdxs = nil
}
