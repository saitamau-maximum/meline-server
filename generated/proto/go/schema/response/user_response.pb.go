// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: schema/response/user_response.proto

package response

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

type UserMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *UserMeResponse) Reset() {
	*x = UserMeResponse{}
	mi := &file_schema_response_user_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeResponse) ProtoMessage() {}

func (x *UserMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_response_user_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeResponse.ProtoReflect.Descriptor instead.
func (*UserMeResponse) Descriptor() ([]byte, []int) {
	return file_schema_response_user_response_proto_rawDescGZIP(), []int{0}
}

func (x *UserMeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserMeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserMeResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_schema_response_user_response_proto protoreflect.FileDescriptor

var file_schema_response_user_response_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x42, 0xcc, 0x01, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x69, 0x74, 0x61, 0x6d, 0x61, 0x75, 0x2d, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x2f, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xa2, 0x02, 0x03,
	0x53, 0x52, 0x58, 0xaa, 0x02, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0xca, 0x02, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xe2, 0x02, 0x1b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_response_user_response_proto_rawDescOnce sync.Once
	file_schema_response_user_response_proto_rawDescData = file_schema_response_user_response_proto_rawDesc
)

func file_schema_response_user_response_proto_rawDescGZIP() []byte {
	file_schema_response_user_response_proto_rawDescOnce.Do(func() {
		file_schema_response_user_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_response_user_response_proto_rawDescData)
	})
	return file_schema_response_user_response_proto_rawDescData
}

var file_schema_response_user_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_response_user_response_proto_goTypes = []any{
	(*UserMeResponse)(nil), // 0: schema.response.UserMeResponse
}
var file_schema_response_user_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_response_user_response_proto_init() }
func file_schema_response_user_response_proto_init() {
	if File_schema_response_user_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_response_user_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_response_user_response_proto_goTypes,
		DependencyIndexes: file_schema_response_user_response_proto_depIdxs,
		MessageInfos:      file_schema_response_user_response_proto_msgTypes,
	}.Build()
	File_schema_response_user_response_proto = out.File
	file_schema_response_user_response_proto_rawDesc = nil
	file_schema_response_user_response_proto_goTypes = nil
	file_schema_response_user_response_proto_depIdxs = nil
}
