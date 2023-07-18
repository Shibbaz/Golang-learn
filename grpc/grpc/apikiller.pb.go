// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: grpc/apikiller.proto

package grpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type APIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *APIRequest) Reset() {
	*x = APIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_apikiller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIRequest) ProtoMessage() {}

func (x *APIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_apikiller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIRequest.ProtoReflect.Descriptor instead.
func (*APIRequest) Descriptor() ([]byte, []int) {
	return file_grpc_apikiller_proto_rawDescGZIP(), []int{0}
}

func (x *APIRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *APIRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type UserArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *UserArgs) Reset() {
	*x = UserArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_apikiller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserArgs) ProtoMessage() {}

func (x *UserArgs) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_apikiller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserArgs.ProtoReflect.Descriptor instead.
func (*UserArgs) Descriptor() ([]byte, []int) {
	return file_grpc_apikiller_proto_rawDescGZIP(), []int{1}
}

func (x *UserArgs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserArgs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserArgs) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type PostArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AuthorId    string `protobuf:"bytes,4,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
}

func (x *PostArgs) Reset() {
	*x = PostArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_apikiller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostArgs) ProtoMessage() {}

func (x *PostArgs) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_apikiller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostArgs.ProtoReflect.Descriptor instead.
func (*PostArgs) Descriptor() ([]byte, []int) {
	return file_grpc_apikiller_proto_rawDescGZIP(), []int{2}
}

func (x *PostArgs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostArgs) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostArgs) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostArgs) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type APIReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *APIReply) Reset() {
	*x = APIReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_apikiller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIReply) ProtoMessage() {}

func (x *APIReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_apikiller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIReply.ProtoReflect.Descriptor instead.
func (*APIReply) Descriptor() ([]byte, []int) {
	return file_grpc_apikiller_proto_rawDescGZIP(), []int{3}
}

func (x *APIReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_apikiller_proto protoreflect.FileDescriptor

var file_grpc_apikiller_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd9, 0x03, 0x0a, 0x03, 0x41, 0x50,
	0x49, 0x12, 0x3b, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72,
	0x67, 0x73, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73,
	0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x6b,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x55, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0e, 0x41, 0x70, 0x69, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69,
	0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x70, 0x69,
	0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_apikiller_proto_rawDescOnce sync.Once
	file_grpc_apikiller_proto_rawDescData = file_grpc_apikiller_proto_rawDesc
)

func file_grpc_apikiller_proto_rawDescGZIP() []byte {
	file_grpc_apikiller_proto_rawDescOnce.Do(func() {
		file_grpc_apikiller_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_apikiller_proto_rawDescData)
	})
	return file_grpc_apikiller_proto_rawDescData
}

var file_grpc_apikiller_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_apikiller_proto_goTypes = []interface{}{
	(*APIRequest)(nil), // 0: apikiller.APIRequest
	(*UserArgs)(nil),   // 1: apikiller.UserArgs
	(*PostArgs)(nil),   // 2: apikiller.PostArgs
	(*APIReply)(nil),   // 3: apikiller.APIReply
}
var file_grpc_apikiller_proto_depIdxs = []int32{
	0, // 0: apikiller.API.getAllUsers:input_type -> apikiller.APIRequest
	1, // 1: apikiller.API.getUserById:input_type -> apikiller.UserArgs
	1, // 2: apikiller.API.createUser:input_type -> apikiller.UserArgs
	1, // 3: apikiller.API.deleteUser:input_type -> apikiller.UserArgs
	1, // 4: apikiller.API.updateUser:input_type -> apikiller.UserArgs
	2, // 5: apikiller.API.createPost:input_type -> apikiller.PostArgs
	2, // 6: apikiller.API.deletePost:input_type -> apikiller.PostArgs
	2, // 7: apikiller.API.updatePost:input_type -> apikiller.PostArgs
	3, // 8: apikiller.API.getAllUsers:output_type -> apikiller.APIReply
	3, // 9: apikiller.API.getUserById:output_type -> apikiller.APIReply
	3, // 10: apikiller.API.createUser:output_type -> apikiller.APIReply
	3, // 11: apikiller.API.deleteUser:output_type -> apikiller.APIReply
	3, // 12: apikiller.API.updateUser:output_type -> apikiller.APIReply
	3, // 13: apikiller.API.createPost:output_type -> apikiller.APIReply
	3, // 14: apikiller.API.deletePost:output_type -> apikiller.APIReply
	3, // 15: apikiller.API.updatePost:output_type -> apikiller.APIReply
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_apikiller_proto_init() }
func file_grpc_apikiller_proto_init() {
	if File_grpc_apikiller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_apikiller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_apikiller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_apikiller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_apikiller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_apikiller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_apikiller_proto_goTypes,
		DependencyIndexes: file_grpc_apikiller_proto_depIdxs,
		MessageInfos:      file_grpc_apikiller_proto_msgTypes,
	}.Build()
	File_grpc_apikiller_proto = out.File
	file_grpc_apikiller_proto_rawDesc = nil
	file_grpc_apikiller_proto_goTypes = nil
	file_grpc_apikiller_proto_depIdxs = nil
}
