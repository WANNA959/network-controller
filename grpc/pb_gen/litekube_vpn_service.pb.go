// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: litekube_vpn_service.proto

package pb_gen

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

type HelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HelloText string `protobuf:"bytes,1,opt,name=hello_text,json=helloText,proto3" json:"hello_text,omitempty"`
}

func (x *HelloWorldRequest) Reset() {
	*x = HelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_litekube_vpn_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldRequest) ProtoMessage() {}

func (x *HelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_litekube_vpn_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldRequest.ProtoReflect.Descriptor instead.
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_litekube_vpn_service_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldRequest) GetHelloText() string {
	if x != nil {
		return x.HelloText
	}
	return ""
}

type HelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThanksText string `protobuf:"bytes,1,opt,name=thanks_text,json=thanksText,proto3" json:"thanks_text,omitempty"`
}

func (x *HelloWorldResponse) Reset() {
	*x = HelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_litekube_vpn_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldResponse) ProtoMessage() {}

func (x *HelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_litekube_vpn_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldResponse.ProtoReflect.Descriptor instead.
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_litekube_vpn_service_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldResponse) GetThanksText() string {
	if x != nil {
		return x.ThanksText
	}
	return ""
}

type CheckConnStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckConnStateRequest) Reset() {
	*x = CheckConnStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_litekube_vpn_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckConnStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckConnStateRequest) ProtoMessage() {}

func (x *CheckConnStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_litekube_vpn_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckConnStateRequest.ProtoReflect.Descriptor instead.
func (*CheckConnStateRequest) Descriptor() ([]byte, []int) {
	return file_litekube_vpn_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckConnStateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckConnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ConnState int32  `protobuf:"varint,3,opt,name=connState,proto3" json:"connState,omitempty"`
}

func (x *CheckConnResponse) Reset() {
	*x = CheckConnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_litekube_vpn_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckConnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckConnResponse) ProtoMessage() {}

func (x *CheckConnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_litekube_vpn_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckConnResponse.ProtoReflect.Descriptor instead.
func (*CheckConnResponse) Descriptor() ([]byte, []int) {
	return file_litekube_vpn_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckConnResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CheckConnResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CheckConnResponse) GetConnState() int32 {
	if x != nil {
		return x.ConnState
	}
	return 0
}

type UnRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UnRegisterRequest) Reset() {
	*x = UnRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_litekube_vpn_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnRegisterRequest) ProtoMessage() {}

func (x *UnRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_litekube_vpn_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnRegisterRequest.ProtoReflect.Descriptor instead.
func (*UnRegisterRequest) Descriptor() ([]byte, []int) {
	return file_litekube_vpn_service_proto_rawDescGZIP(), []int{4}
}

func (x *UnRegisterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UnRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Result  bool   `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UnRegisterResponse) Reset() {
	*x = UnRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_litekube_vpn_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnRegisterResponse) ProtoMessage() {}

func (x *UnRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_litekube_vpn_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnRegisterResponse.ProtoReflect.Descriptor instead.
func (*UnRegisterResponse) Descriptor() ([]byte, []int) {
	return file_litekube_vpn_service_proto_rawDescGZIP(), []int{5}
}

func (x *UnRegisterResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UnRegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UnRegisterResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_litekube_vpn_service_proto protoreflect.FileDescriptor

var file_litekube_vpn_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6c, 0x69, 0x74, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x70, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x32, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x54, 0x65, 0x78, 0x74, 0x22, 0x35, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68,
	0x61, 0x6e, 0x6b, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x54, 0x65, 0x78, 0x74, 0x22, 0x2d, 0x0a, 0x15, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x55,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x12, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xd8, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x56,
	0x70, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0a, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x3b, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_litekube_vpn_service_proto_rawDescOnce sync.Once
	file_litekube_vpn_service_proto_rawDescData = file_litekube_vpn_service_proto_rawDesc
)

func file_litekube_vpn_service_proto_rawDescGZIP() []byte {
	file_litekube_vpn_service_proto_rawDescOnce.Do(func() {
		file_litekube_vpn_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_litekube_vpn_service_proto_rawDescData)
	})
	return file_litekube_vpn_service_proto_rawDescData
}

var file_litekube_vpn_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_litekube_vpn_service_proto_goTypes = []interface{}{
	(*HelloWorldRequest)(nil),     // 0: pb.HelloWorldRequest
	(*HelloWorldResponse)(nil),    // 1: pb.HelloWorldResponse
	(*CheckConnStateRequest)(nil), // 2: pb.CheckConnStateRequest
	(*CheckConnResponse)(nil),     // 3: pb.CheckConnResponse
	(*UnRegisterRequest)(nil),     // 4: pb.UnRegisterRequest
	(*UnRegisterResponse)(nil),    // 5: pb.UnRegisterResponse
}
var file_litekube_vpn_service_proto_depIdxs = []int32{
	0, // 0: pb.LiteKubeVpnService.HelloWorld:input_type -> pb.HelloWorldRequest
	2, // 1: pb.LiteKubeVpnService.CheckConnState:input_type -> pb.CheckConnStateRequest
	4, // 2: pb.LiteKubeVpnService.UnRegister:input_type -> pb.UnRegisterRequest
	1, // 3: pb.LiteKubeVpnService.HelloWorld:output_type -> pb.HelloWorldResponse
	3, // 4: pb.LiteKubeVpnService.CheckConnState:output_type -> pb.CheckConnResponse
	5, // 5: pb.LiteKubeVpnService.UnRegister:output_type -> pb.UnRegisterResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_litekube_vpn_service_proto_init() }
func file_litekube_vpn_service_proto_init() {
	if File_litekube_vpn_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_litekube_vpn_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldRequest); i {
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
		file_litekube_vpn_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldResponse); i {
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
		file_litekube_vpn_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckConnStateRequest); i {
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
		file_litekube_vpn_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckConnResponse); i {
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
		file_litekube_vpn_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnRegisterRequest); i {
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
		file_litekube_vpn_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnRegisterResponse); i {
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
			RawDescriptor: file_litekube_vpn_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_litekube_vpn_service_proto_goTypes,
		DependencyIndexes: file_litekube_vpn_service_proto_depIdxs,
		MessageInfos:      file_litekube_vpn_service_proto_msgTypes,
	}.Build()
	File_litekube_vpn_service_proto = out.File
	file_litekube_vpn_service_proto_rawDesc = nil
	file_litekube_vpn_service_proto_goTypes = nil
	file_litekube_vpn_service_proto_depIdxs = nil
}
