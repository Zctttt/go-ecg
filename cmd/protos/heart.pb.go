// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: heart.proto

package protos

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

type HeartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID     string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	TIME     string `protobuf:"bytes,2,opt,name=TIME,proto3" json:"TIME,omitempty"`
	BPM      string `protobuf:"bytes,3,opt,name=BPM,proto3" json:"BPM,omitempty"`
	HeartADC string `protobuf:"bytes,4,opt,name=HeartADC,proto3" json:"HeartADC,omitempty"`
}

func (x *HeartResponse) Reset() {
	*x = HeartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartResponse) ProtoMessage() {}

func (x *HeartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartResponse.ProtoReflect.Descriptor instead.
func (*HeartResponse) Descriptor() ([]byte, []int) {
	return file_heart_proto_rawDescGZIP(), []int{0}
}

func (x *HeartResponse) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *HeartResponse) GetTIME() string {
	if x != nil {
		return x.TIME
	}
	return ""
}

func (x *HeartResponse) GetBPM() string {
	if x != nil {
		return x.BPM
	}
	return ""
}

func (x *HeartResponse) GetHeartADC() string {
	if x != nil {
		return x.HeartADC
	}
	return ""
}

type HeartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	TIME string `protobuf:"bytes,2,opt,name=TIME,proto3" json:"TIME,omitempty"`
}

func (x *HeartRequest) Reset() {
	*x = HeartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartRequest) ProtoMessage() {}

func (x *HeartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartRequest.ProtoReflect.Descriptor instead.
func (*HeartRequest) Descriptor() ([]byte, []int) {
	return file_heart_proto_rawDescGZIP(), []int{1}
}

func (x *HeartRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *HeartRequest) GetTIME() string {
	if x != nil {
		return x.TIME
	}
	return ""
}

var File_heart_proto protoreflect.FileDescriptor

var file_heart_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a,
	0x0d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x50, 0x4d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x42, 0x50, 0x4d, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x41, 0x44, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x41, 0x44, 0x43, 0x22, 0x36, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x49, 0x4d, 0x45,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x32, 0x3e, 0x0a, 0x0c,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x48, 0x65, 0x61, 0x72, 0x74, 0x12, 0x0d, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c,
	0x67, 0x6f, 0x5f, 0x73, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heart_proto_rawDescOnce sync.Once
	file_heart_proto_rawDescData = file_heart_proto_rawDesc
)

func file_heart_proto_rawDescGZIP() []byte {
	file_heart_proto_rawDescOnce.Do(func() {
		file_heart_proto_rawDescData = protoimpl.X.CompressGZIP(file_heart_proto_rawDescData)
	})
	return file_heart_proto_rawDescData
}

var file_heart_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_heart_proto_goTypes = []interface{}{
	(*HeartResponse)(nil), // 0: HeartResponse
	(*HeartRequest)(nil),  // 1: HeartRequest
}
var file_heart_proto_depIdxs = []int32{
	1, // 0: HeartService.GetUnaryHeart:input_type -> HeartRequest
	0, // 1: HeartService.GetUnaryHeart:output_type -> HeartResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heart_proto_init() }
func file_heart_proto_init() {
	if File_heart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartResponse); i {
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
		file_heart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartRequest); i {
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
			RawDescriptor: file_heart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heart_proto_goTypes,
		DependencyIndexes: file_heart_proto_depIdxs,
		MessageInfos:      file_heart_proto_msgTypes,
	}.Build()
	File_heart_proto = out.File
	file_heart_proto_rawDesc = nil
	file_heart_proto_goTypes = nil
	file_heart_proto_depIdxs = nil
}
