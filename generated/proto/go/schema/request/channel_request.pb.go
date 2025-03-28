// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: schema/request/channel_request.proto

package request

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

type GetChannelByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetChannelByIDRequest) Reset() {
	*x = GetChannelByIDRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChannelByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelByIDRequest) ProtoMessage() {}

func (x *GetChannelByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelByIDRequest.ProtoReflect.Descriptor instead.
func (*GetChannelByIDRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetChannelByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type JoinChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JoinChannelRequest) Reset() {
	*x = JoinChannelRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChannelRequest) ProtoMessage() {}

func (x *JoinChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChannelRequest.ProtoReflect.Descriptor instead.
func (*JoinChannelRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{1}
}

func (x *JoinChannelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateChannelRequest) Reset() {
	*x = CreateChannelRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelRequest) ProtoMessage() {}

func (x *CreateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateChannelRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{2}
}

func (x *CreateChannelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateChildChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateChildChannelRequest) Reset() {
	*x = CreateChildChannelRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChildChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChildChannelRequest) ProtoMessage() {}

func (x *CreateChildChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChildChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateChildChannelRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{3}
}

func (x *CreateChildChannelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateChildChannelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateChannelRequest) Reset() {
	*x = UpdateChannelRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelRequest) ProtoMessage() {}

func (x *UpdateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelRequest.ProtoReflect.Descriptor instead.
func (*UpdateChannelRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateChannelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateChannelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteChannelRequest) Reset() {
	*x = DeleteChannelRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelRequest) ProtoMessage() {}

func (x *DeleteChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelRequest.ProtoReflect.Descriptor instead.
func (*DeleteChannelRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteChannelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LeaveChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LeaveChannelRequest) Reset() {
	*x = LeaveChannelRequest{}
	mi := &file_schema_request_channel_request_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveChannelRequest) ProtoMessage() {}

func (x *LeaveChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_request_channel_request_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveChannelRequest.ProtoReflect.Descriptor instead.
func (*LeaveChannelRequest) Descriptor() ([]byte, []int) {
	return file_schema_request_channel_request_proto_rawDescGZIP(), []int{6}
}

func (x *LeaveChannelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_schema_request_channel_request_proto protoreflect.FileDescriptor

var file_schema_request_channel_request_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3f, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0xc8, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x69, 0x74, 0x61, 0x6d, 0x61, 0x75,
	0x2d, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2f, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0xa2, 0x02, 0x03, 0x53, 0x52, 0x58, 0xaa, 0x02, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0xca, 0x02, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0xe2, 0x02, 0x1a, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a,
	0x3a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_request_channel_request_proto_rawDescOnce sync.Once
	file_schema_request_channel_request_proto_rawDescData = file_schema_request_channel_request_proto_rawDesc
)

func file_schema_request_channel_request_proto_rawDescGZIP() []byte {
	file_schema_request_channel_request_proto_rawDescOnce.Do(func() {
		file_schema_request_channel_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_request_channel_request_proto_rawDescData)
	})
	return file_schema_request_channel_request_proto_rawDescData
}

var file_schema_request_channel_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_schema_request_channel_request_proto_goTypes = []any{
	(*GetChannelByIDRequest)(nil),     // 0: schema.request.GetChannelByIDRequest
	(*JoinChannelRequest)(nil),        // 1: schema.request.JoinChannelRequest
	(*CreateChannelRequest)(nil),      // 2: schema.request.CreateChannelRequest
	(*CreateChildChannelRequest)(nil), // 3: schema.request.CreateChildChannelRequest
	(*UpdateChannelRequest)(nil),      // 4: schema.request.UpdateChannelRequest
	(*DeleteChannelRequest)(nil),      // 5: schema.request.DeleteChannelRequest
	(*LeaveChannelRequest)(nil),       // 6: schema.request.LeaveChannelRequest
}
var file_schema_request_channel_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_request_channel_request_proto_init() }
func file_schema_request_channel_request_proto_init() {
	if File_schema_request_channel_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_request_channel_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_request_channel_request_proto_goTypes,
		DependencyIndexes: file_schema_request_channel_request_proto_depIdxs,
		MessageInfos:      file_schema_request_channel_request_proto_msgTypes,
	}.Build()
	File_schema_request_channel_request_proto = out.File
	file_schema_request_channel_request_proto_rawDesc = nil
	file_schema_request_channel_request_proto_goTypes = nil
	file_schema_request_channel_request_proto_depIdxs = nil
}
