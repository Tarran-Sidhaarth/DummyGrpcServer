// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/timesquared/timesquared.proto

package timesquared

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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestValue int32 `protobuf:"varint,1,opt,name=request_value,json=requestValue,proto3" json:"request_value,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesquared_timesquared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesquared_timesquared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_timesquared_timesquared_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckRequest) GetRequestValue() int32 {
	if x != nil {
		return x.RequestValue
	}
	return 0
}

// Response message for the Healthcheck operation
type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HealthStatus bool `protobuf:"varint,1,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesquared_timesquared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesquared_timesquared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_timesquared_timesquared_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetHealthStatus() bool {
	if x != nil {
		return x.HealthStatus
	}
	return false
}

var File_proto_timesquared_timesquared_proto protoreflect.FileDescriptor

var file_proto_timesquared_timesquared_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x64, 0x22, 0x39, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a,
	0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x68, 0x0a, 0x12, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1f,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_timesquared_timesquared_proto_rawDescOnce sync.Once
	file_proto_timesquared_timesquared_proto_rawDescData = file_proto_timesquared_timesquared_proto_rawDesc
)

func file_proto_timesquared_timesquared_proto_rawDescGZIP() []byte {
	file_proto_timesquared_timesquared_proto_rawDescOnce.Do(func() {
		file_proto_timesquared_timesquared_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_timesquared_timesquared_proto_rawDescData)
	})
	return file_proto_timesquared_timesquared_proto_rawDescData
}

var file_proto_timesquared_timesquared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_timesquared_timesquared_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),  // 0: timesquared.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: timesquared.HealthCheckResponse
}
var file_proto_timesquared_timesquared_proto_depIdxs = []int32{
	0, // 0: timesquared.TimeSquaredService.HealthCheck:input_type -> timesquared.HealthCheckRequest
	1, // 1: timesquared.TimeSquaredService.HealthCheck:output_type -> timesquared.HealthCheckResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_timesquared_timesquared_proto_init() }
func file_proto_timesquared_timesquared_proto_init() {
	if File_proto_timesquared_timesquared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_timesquared_timesquared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_proto_timesquared_timesquared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_proto_timesquared_timesquared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_timesquared_timesquared_proto_goTypes,
		DependencyIndexes: file_proto_timesquared_timesquared_proto_depIdxs,
		MessageInfos:      file_proto_timesquared_timesquared_proto_msgTypes,
	}.Build()
	File_proto_timesquared_timesquared_proto = out.File
	file_proto_timesquared_timesquared_proto_rawDesc = nil
	file_proto_timesquared_timesquared_proto_goTypes = nil
	file_proto_timesquared_timesquared_proto_depIdxs = nil
}
