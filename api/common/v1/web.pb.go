// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: web.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_UNKNOWN Code = 0
	Code_OK      Code = 2000
	// 请求参数错误
	Code_PARAMS_ERROR Code = 4000
	// AUTH_ERROR
	Code_AUTH_ERROR         Code = 4001
	Code_PREMISSION_DENIED  Code = 4003
	Code_NOT_FOUND          Code = 4004
	Code_TOKEN_NOT_FOUND    Code = 4006
	Code_RESOURCE_EXHAUSTED Code = 4008
	Code_INTERNAL_ERROR     Code = 5000
	Code_TIMEOUT_ERROR      Code = 5004
	Code_METADATA_MISSING   Code = 4007
	Code_CONFLICT           Code = 4009
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:    "UNKNOWN",
		2000: "OK",
		4000: "PARAMS_ERROR",
		4001: "AUTH_ERROR",
		4003: "PREMISSION_DENIED",
		4004: "NOT_FOUND",
		4006: "TOKEN_NOT_FOUND",
		4008: "RESOURCE_EXHAUSTED",
		5000: "INTERNAL_ERROR",
		5004: "TIMEOUT_ERROR",
		4007: "METADATA_MISSING",
		4009: "CONFLICT",
	}
	Code_value = map[string]int32{
		"UNKNOWN":            0,
		"OK":                 2000,
		"PARAMS_ERROR":       4000,
		"AUTH_ERROR":         4001,
		"PREMISSION_DENIED":  4003,
		"NOT_FOUND":          4004,
		"TOKEN_NOT_FOUND":    4006,
		"RESOURCE_EXHAUSTED": 4008,
		"INTERNAL_ERROR":     5000,
		"TIMEOUT_ERROR":      5004,
		"METADATA_MISSING":   4007,
		"CONFLICT":           4009,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_web_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_web_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{0}
}

type APIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code         float64 `protobuf:"fixed64,2,opt,name=code,proto3" json:"code,omitempty"`
	ResponseType string  `protobuf:"bytes,3,opt,name=response_type,proto3" json:"response_type,omitempty"`
	Data         []byte  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *APIResponse) Reset() {
	*x = APIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIResponse) ProtoMessage() {}

func (x *APIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIResponse.ProtoReflect.Descriptor instead.
func (*APIResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{0}
}

func (x *APIResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *APIResponse) GetCode() float64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *APIResponse) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *APIResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type HttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32            `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data    *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *HttpResponse) Reset() {
	*x = HttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpResponse) ProtoMessage() {}

func (x *HttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpResponse.ProtoReflect.Descriptor instead.
func (*HttpResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{1}
}

func (x *HttpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HttpResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HttpResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Id    int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Retry int32  `protobuf:"varint,4,opt,name=retry,proto3" json:"retry,omitempty"`
}

func (x *EventStreamResponse) Reset() {
	*x = EventStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStreamResponse) ProtoMessage() {}

func (x *EventStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStreamResponse.ProtoReflect.Descriptor instead.
func (*EventStreamResponse) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{2}
}

func (x *EventStreamResponse) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *EventStreamResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *EventStreamResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventStreamResponse) GetRetry() int32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

// }
type Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message         string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Action          string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	File            string `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Line            int32  `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
	Fn              string `protobuf:"bytes,6,opt,name=fn,proto3" json:"fn,omitempty"`
	Stack           string `protobuf:"bytes,7,opt,name=stack,proto3" json:"stack,omitempty"`
	ToClientMessage string `protobuf:"bytes,8,opt,name=to_client_message,json=toClientMessage,proto3" json:"to_client_message,omitempty"`
}

func (x *Errors) Reset() {
	*x = Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors) ProtoMessage() {}

func (x *Errors) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors.ProtoReflect.Descriptor instead.
func (*Errors) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{3}
}

func (x *Errors) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Errors) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Errors) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Errors) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Errors) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Errors) GetFn() string {
	if x != nil {
		return x.Fn
	}
	return ""
}

func (x *Errors) GetStack() string {
	if x != nil {
		return x.Stack
	}
	return ""
}

func (x *Errors) GetToClientMessage() string {
	if x != nil {
		return x.ToClientMessage
	}
	return ""
}

type Headers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Authentication string `protobuf:"bytes,2,opt,name=authentication,proto3" json:"authentication,omitempty"`
	Filename       string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Token          string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Headers) Reset() {
	*x = Headers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Headers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Headers) ProtoMessage() {}

func (x *Headers) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Headers.ProtoReflect.Descriptor instead.
func (*Headers) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{4}
}

func (x *Headers) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Headers) GetAuthentication() string {
	if x != nil {
		return x.Authentication
	}
	return ""
}

func (x *Headers) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Headers) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type EventStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Id    int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Retry int32  `protobuf:"varint,4,opt,name=retry,proto3" json:"retry,omitempty"`
}

func (x *EventStream) Reset() {
	*x = EventStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStream) ProtoMessage() {}

func (x *EventStream) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStream.ProtoReflect.Descriptor instead.
func (*EventStream) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{5}
}

func (x *EventStream) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *EventStream) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *EventStream) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventStream) GetRetry() int32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

var file_web_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50031,
		Name:          "begonia.api.v1.description",
		Tag:           "bytes,50031,opt,name=description",
		Filename:      "web.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50032,
		Name:          "begonia.api.v1.http_code",
		Tag:           "varint,50032,opt,name=http_code",
		Filename:      "web.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string description = 50031;
	E_Description = &file_web_proto_extTypes[0]
	// optional int32 http_code = 50032;
	E_HttpCode = &file_web_proto_extTypes[1]
)

var File_web_proto protoreflect.FileDescriptor

var file_web_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x65, 0x67,
	0x6f, 0x6e, 0x69, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0b, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x69, 0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x65, 0x0a,
	0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x66, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x66, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x74, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x75, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x2a, 0x8f, 0x04, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x18, 0xfa, 0xb6, 0x18,
	0x0f, 0xe6, 0x9c, 0xaa, 0xe7, 0x9f, 0xa5, 0xe7, 0x9a, 0x84, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf,
	0x80, 0xb7, 0x18, 0xf4, 0x03, 0x12, 0x24, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0xd0, 0x0f, 0x1a, 0x1b,
	0xfa, 0xb6, 0x18, 0x12, 0xe8, 0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0xe5, 0xa4, 0x84, 0xe7, 0x90, 0x86,
	0xe6, 0x88, 0x90, 0xe5, 0x8a, 0x9f, 0x80, 0xb7, 0x18, 0xc8, 0x01, 0x12, 0x28, 0x0a, 0x0c, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa0, 0x1f, 0x1a, 0x15,
	0xfa, 0xb6, 0x18, 0x0c, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf,
	0x80, 0xb7, 0x18, 0x90, 0x03, 0x12, 0x26, 0x0a, 0x0a, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xa1, 0x1f, 0x1a, 0x15, 0xfa, 0xb6, 0x18, 0x0c, 0xe8, 0xae, 0xa4, 0xe8,
	0xaf, 0x81, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x80, 0xb7, 0x18, 0x91, 0x03, 0x12, 0x2d, 0x0a,
	0x11, 0x50, 0x52, 0x45, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49,
	0x45, 0x44, 0x10, 0xa3, 0x1f, 0x1a, 0x15, 0xfa, 0xb6, 0x18, 0x0c, 0xe6, 0x9d, 0x83, 0xe9, 0x99,
	0x90, 0xe4, 0xb8, 0x8d, 0xe8, 0xb6, 0xb3, 0x80, 0xb7, 0x18, 0x93, 0x03, 0x12, 0x28, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa4, 0x1f, 0x1a, 0x18, 0xfa, 0xb6,
	0x18, 0x0f, 0xe6, 0x89, 0xbe, 0xe4, 0xb8, 0x8d, 0xe5, 0x88, 0xb0, 0xe8, 0xb5, 0x84, 0xe6, 0xba,
	0x90, 0x80, 0xb7, 0x18, 0x94, 0x03, 0x12, 0x2b, 0x0a, 0x0f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa6, 0x1f, 0x1a, 0x15, 0xfa, 0xb6,
	0x18, 0x0c, 0xe5, 0x86, 0x85, 0xe9, 0x83, 0xa8, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x80, 0xb7,
	0x18, 0x94, 0x03, 0x12, 0x2e, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0xa8, 0x1f, 0x1a, 0x15, 0xfa, 0xb6,
	0x18, 0x0c, 0xe8, 0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe8, 0x80, 0x97, 0xe5, 0xb0, 0xbd, 0x80, 0xb7,
	0x18, 0xad, 0x03, 0x12, 0x2a, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x88, 0x27, 0x1a, 0x15, 0xfa, 0xb6, 0x18, 0x0c, 0xe5, 0x86,
	0x85, 0xe9, 0x83, 0xa8, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x80, 0xb7, 0x18, 0xf4, 0x03, 0x12,
	0x29, 0x0a, 0x0d, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x8c, 0x27, 0x1a, 0x15, 0xfa, 0xb6, 0x18, 0x0c, 0xe8, 0xb6, 0x85, 0xe6, 0x97, 0xb6, 0xe9,
	0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x80, 0xb7, 0x18, 0xf8, 0x03, 0x12, 0x2f, 0x0a, 0x10, 0x4d, 0x45,
	0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0xa7,
	0x1f, 0x1a, 0x18, 0xfa, 0xb6, 0x18, 0x0f, 0xe7, 0xbc, 0xba, 0xe5, 0xb0, 0x91, 0xe5, 0x85, 0x83,
	0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0x80, 0xb7, 0x18, 0x90, 0x03, 0x12, 0x2a, 0x0a, 0x08, 0x43,
	0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x10, 0xa9, 0x1f, 0x1a, 0x1b, 0xfa, 0xb6, 0x18, 0x12,
	0xe8, 0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe5, 0xb7, 0xb2, 0xe7, 0xbb, 0x8f, 0xe5, 0xad, 0x98, 0xe5,
	0x9c, 0xa8, 0x80, 0xb7, 0x18, 0x99, 0x03, 0x3a, 0x48, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xef, 0x86, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x3a, 0x43, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xf0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2d, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_proto_rawDescOnce sync.Once
	file_web_proto_rawDescData = file_web_proto_rawDesc
)

func file_web_proto_rawDescGZIP() []byte {
	file_web_proto_rawDescOnce.Do(func() {
		file_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_proto_rawDescData)
	})
	return file_web_proto_rawDescData
}

var file_web_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_web_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_web_proto_goTypes = []interface{}{
	(Code)(0),                             // 0: begonia.api.v1.Code
	(*APIResponse)(nil),                   // 1: begonia.api.v1.APIResponse
	(*HttpResponse)(nil),                  // 2: begonia.api.v1.HttpResponse
	(*EventStreamResponse)(nil),           // 3: begonia.api.v1.EventStreamResponse
	(*Errors)(nil),                        // 4: begonia.api.v1.Errors
	(*Headers)(nil),                       // 5: begonia.api.v1.Headers
	(*EventStream)(nil),                   // 6: begonia.api.v1.EventStream
	(*structpb.Struct)(nil),               // 7: google.protobuf.Struct
	(*descriptorpb.EnumValueOptions)(nil), // 8: google.protobuf.EnumValueOptions
}
var file_web_proto_depIdxs = []int32{
	7, // 0: begonia.api.v1.HttpResponse.data:type_name -> google.protobuf.Struct
	8, // 1: begonia.api.v1.description:extendee -> google.protobuf.EnumValueOptions
	8, // 2: begonia.api.v1.http_code:extendee -> google.protobuf.EnumValueOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_proto_init() }
func file_web_proto_init() {
	if File_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIResponse); i {
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
		file_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpResponse); i {
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
		file_web_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStreamResponse); i {
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
		file_web_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errors); i {
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
		file_web_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Headers); i {
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
		file_web_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStream); i {
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
			RawDescriptor: file_web_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_web_proto_goTypes,
		DependencyIndexes: file_web_proto_depIdxs,
		EnumInfos:         file_web_proto_enumTypes,
		MessageInfos:      file_web_proto_msgTypes,
		ExtensionInfos:    file_web_proto_extTypes,
	}.Build()
	File_web_proto = out.File
	file_web_proto_rawDesc = nil
	file_web_proto_goTypes = nil
	file_web_proto_depIdxs = nil
}
