// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: UserAuth_service.proto

package protofiles

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

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserAuth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_UserAuth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_UserAuth_service_proto_rawDescGZIP(), []int{0}
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserAuth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_UserAuth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_UserAuth_service_proto_rawDescGZIP(), []int{1}
}

func (x *Confirmation) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ConfirmationPlus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Jwt string `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *ConfirmationPlus) Reset() {
	*x = ConfirmationPlus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserAuth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationPlus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationPlus) ProtoMessage() {}

func (x *ConfirmationPlus) ProtoReflect() protoreflect.Message {
	mi := &file_UserAuth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationPlus.ProtoReflect.Descriptor instead.
func (*ConfirmationPlus) Descriptor() ([]byte, []int) {
	return file_UserAuth_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmationPlus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ConfirmationPlus) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type Jwt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Jwt) Reset() {
	*x = Jwt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserAuth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jwt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jwt) ProtoMessage() {}

func (x *Jwt) ProtoReflect() protoreflect.Message {
	mi := &file_UserAuth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jwt.ProtoReflect.Descriptor instead.
func (*Jwt) Descriptor() ([]byte, []int) {
	return file_UserAuth_service_proto_rawDescGZIP(), []int{3}
}

func (x *Jwt) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_UserAuth_service_proto protoreflect.FileDescriptor

var file_UserAuth_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x36, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x17, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xc8,
	0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3f, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x75, 0x73, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4a, 0x77, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UserAuth_service_proto_rawDescOnce sync.Once
	file_UserAuth_service_proto_rawDescData = file_UserAuth_service_proto_rawDesc
)

func file_UserAuth_service_proto_rawDescGZIP() []byte {
	file_UserAuth_service_proto_rawDescOnce.Do(func() {
		file_UserAuth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserAuth_service_proto_rawDescData)
	})
	return file_UserAuth_service_proto_rawDescData
}

var file_UserAuth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_UserAuth_service_proto_goTypes = []interface{}{
	(*Credentials)(nil),      // 0: protofiles.Credentials
	(*Confirmation)(nil),     // 1: protofiles.Confirmation
	(*ConfirmationPlus)(nil), // 2: protofiles.ConfirmationPlus
	(*Jwt)(nil),              // 3: protofiles.Jwt
}
var file_UserAuth_service_proto_depIdxs = []int32{
	0, // 0: protofiles.UserAuth.Register:input_type -> protofiles.Credentials
	0, // 1: protofiles.UserAuth.Login:input_type -> protofiles.Credentials
	3, // 2: protofiles.UserAuth.HelloWorld:input_type -> protofiles.Jwt
	1, // 3: protofiles.UserAuth.Register:output_type -> protofiles.Confirmation
	2, // 4: protofiles.UserAuth.Login:output_type -> protofiles.ConfirmationPlus
	1, // 5: protofiles.UserAuth.HelloWorld:output_type -> protofiles.Confirmation
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_UserAuth_service_proto_init() }
func file_UserAuth_service_proto_init() {
	if File_UserAuth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UserAuth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_UserAuth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
		file_UserAuth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationPlus); i {
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
		file_UserAuth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jwt); i {
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
			RawDescriptor: file_UserAuth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_UserAuth_service_proto_goTypes,
		DependencyIndexes: file_UserAuth_service_proto_depIdxs,
		MessageInfos:      file_UserAuth_service_proto_msgTypes,
	}.Build()
	File_UserAuth_service_proto = out.File
	file_UserAuth_service_proto_rawDesc = nil
	file_UserAuth_service_proto_goTypes = nil
	file_UserAuth_service_proto_depIdxs = nil
}