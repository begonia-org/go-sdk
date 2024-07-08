// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: sys.proto

package v1

import (
	_ "github.com/begonia-org/go-sdk/begonia/api/v1/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{0}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Commit    string `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	BuildTime string `protobuf:"bytes,3,opt,name=buildTime,proto3" json:"buildTime,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InfoResponse) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *InfoResponse) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

var File_sys_proto protoreflect.FileDescriptor

var file_sys_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x65, 0x67,
	0x6f, 0x6e, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x65,
	0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x0c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x92, 0x01, 0x0a, 0x0d, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x23, 0x88, 0xb7, 0x18, 0x01, 0xb2,
	0xb7, 0x18, 0x1b, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67,
	0x6f, 0x6e, 0x69, 0x61, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sys_proto_rawDescOnce sync.Once
	file_sys_proto_rawDescData = file_sys_proto_rawDesc
)

func file_sys_proto_rawDescGZIP() []byte {
	file_sys_proto_rawDescOnce.Do(func() {
		file_sys_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_proto_rawDescData)
	})
	return file_sys_proto_rawDescData
}

var file_sys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sys_proto_goTypes = []interface{}{
	(*InfoRequest)(nil),  // 0: begonia.org.sdk.InfoRequest
	(*InfoResponse)(nil), // 1: begonia.org.sdk.InfoResponse
}
var file_sys_proto_depIdxs = []int32{
	0, // 0: begonia.org.sdk.SystemService.Get:input_type -> begonia.org.sdk.InfoRequest
	1, // 1: begonia.org.sdk.SystemService.Get:output_type -> begonia.org.sdk.InfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sys_proto_init() }
func file_sys_proto_init() {
	if File_sys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_sys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
			RawDescriptor: file_sys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_proto_goTypes,
		DependencyIndexes: file_sys_proto_depIdxs,
		MessageInfos:      file_sys_proto_msgTypes,
	}.Build()
	File_sys_proto = out.File
	file_sys_proto_rawDesc = nil
	file_sys_proto_goTypes = nil
	file_sys_proto_depIdxs = nil
}
