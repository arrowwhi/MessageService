// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.3
// source: message_service/users/v1/get_status_info.proto

package usersv1

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

type GetStatusInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetStatusInfoRequest) Reset() {
	*x = GetStatusInfoRequest{}
	mi := &file_message_service_users_v1_get_status_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusInfoRequest) ProtoMessage() {}

func (x *GetStatusInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_users_v1_get_status_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusInfoRequest.ProtoReflect.Descriptor instead.
func (*GetStatusInfoRequest) Descriptor() ([]byte, []int) {
	return file_message_service_users_v1_get_status_info_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatusInfoRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetStatusInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users *User `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *GetStatusInfoResponse) Reset() {
	*x = GetStatusInfoResponse{}
	mi := &file_message_service_users_v1_get_status_info_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusInfoResponse) ProtoMessage() {}

func (x *GetStatusInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_users_v1_get_status_info_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusInfoResponse.ProtoReflect.Descriptor instead.
func (*GetStatusInfoResponse) Descriptor() ([]byte, []int) {
	return file_message_service_users_v1_get_status_info_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatusInfoResponse) GetUsers() *User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_message_service_users_v1_get_status_info_proto protoreflect.FileDescriptor

var file_message_service_users_v1_get_status_info_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_service_users_v1_get_status_info_proto_rawDescOnce sync.Once
	file_message_service_users_v1_get_status_info_proto_rawDescData = file_message_service_users_v1_get_status_info_proto_rawDesc
)

func file_message_service_users_v1_get_status_info_proto_rawDescGZIP() []byte {
	file_message_service_users_v1_get_status_info_proto_rawDescOnce.Do(func() {
		file_message_service_users_v1_get_status_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_service_users_v1_get_status_info_proto_rawDescData)
	})
	return file_message_service_users_v1_get_status_info_proto_rawDescData
}

var file_message_service_users_v1_get_status_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_service_users_v1_get_status_info_proto_goTypes = []any{
	(*GetStatusInfoRequest)(nil),  // 0: message_service.users.v1.GetStatusInfoRequest
	(*GetStatusInfoResponse)(nil), // 1: message_service.users.v1.GetStatusInfoResponse
	(*User)(nil),                  // 2: message_service.users.v1.User
}
var file_message_service_users_v1_get_status_info_proto_depIdxs = []int32{
	2, // 0: message_service.users.v1.GetStatusInfoResponse.users:type_name -> message_service.users.v1.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_service_users_v1_get_status_info_proto_init() }
func file_message_service_users_v1_get_status_info_proto_init() {
	if File_message_service_users_v1_get_status_info_proto != nil {
		return
	}
	file_message_service_users_v1_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_service_users_v1_get_status_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_service_users_v1_get_status_info_proto_goTypes,
		DependencyIndexes: file_message_service_users_v1_get_status_info_proto_depIdxs,
		MessageInfos:      file_message_service_users_v1_get_status_info_proto_msgTypes,
	}.Build()
	File_message_service_users_v1_get_status_info_proto = out.File
	file_message_service_users_v1_get_status_info_proto_rawDesc = nil
	file_message_service_users_v1_get_status_info_proto_goTypes = nil
	file_message_service_users_v1_get_status_info_proto_depIdxs = nil
}
