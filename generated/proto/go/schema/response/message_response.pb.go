// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: schema/response/message_response.proto

package response

import (
	base "github.com/saitamau-maximum/meline/generated/proto/go/base"
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

type GetByChannelIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*base.Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetByChannelIDResponse) Reset() {
	*x = GetByChannelIDResponse{}
	mi := &file_schema_response_message_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByChannelIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByChannelIDResponse) ProtoMessage() {}

func (x *GetByChannelIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_response_message_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByChannelIDResponse.ProtoReflect.Descriptor instead.
func (*GetByChannelIDResponse) Descriptor() ([]byte, []int) {
	return file_schema_response_message_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetByChannelIDResponse) GetMessages() []*base.Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_schema_response_message_response_proto protoreflect.FileDescriptor

var file_schema_response_message_response_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x17, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0xcf, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x69, 0x74, 0x61, 0x6d, 0x61, 0x75, 0x2d, 0x6d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x2f, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xa2, 0x02,
	0x03, 0x53, 0x52, 0x58, 0xaa, 0x02, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xca, 0x02, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xe2, 0x02, 0x1b, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a,
	0x3a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_schema_response_message_response_proto_rawDescOnce sync.Once
	file_schema_response_message_response_proto_rawDescData = file_schema_response_message_response_proto_rawDesc
)

func file_schema_response_message_response_proto_rawDescGZIP() []byte {
	file_schema_response_message_response_proto_rawDescOnce.Do(func() {
		file_schema_response_message_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_response_message_response_proto_rawDescData)
	})
	return file_schema_response_message_response_proto_rawDescData
}

var file_schema_response_message_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_response_message_response_proto_goTypes = []any{
	(*GetByChannelIDResponse)(nil), // 0: schema.response.GetByChannelIDResponse
	(*base.Message)(nil),           // 1: base.Message
}
var file_schema_response_message_response_proto_depIdxs = []int32{
	1, // 0: schema.response.GetByChannelIDResponse.messages:type_name -> base.Message
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schema_response_message_response_proto_init() }
func file_schema_response_message_response_proto_init() {
	if File_schema_response_message_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_response_message_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_response_message_response_proto_goTypes,
		DependencyIndexes: file_schema_response_message_response_proto_depIdxs,
		MessageInfos:      file_schema_response_message_response_proto_msgTypes,
	}.Build()
	File_schema_response_message_response_proto = out.File
	file_schema_response_message_response_proto_rawDesc = nil
	file_schema_response_message_response_proto_goTypes = nil
	file_schema_response_message_response_proto_depIdxs = nil
}
