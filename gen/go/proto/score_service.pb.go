// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: score_service.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListTopKScoresV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTopKScoresV1Request) Reset() {
	*x = ListTopKScoresV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopKScoresV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopKScoresV1Request) ProtoMessage() {}

func (x *ListTopKScoresV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopKScoresV1Request.ProtoReflect.Descriptor instead.
func (*ListTopKScoresV1Request) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{0}
}

type ListTopKScoresV1InnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*ScoreV1 `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *ListTopKScoresV1InnerResponse) Reset() {
	*x = ListTopKScoresV1InnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopKScoresV1InnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopKScoresV1InnerResponse) ProtoMessage() {}

func (x *ListTopKScoresV1InnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopKScoresV1InnerResponse.ProtoReflect.Descriptor instead.
func (*ListTopKScoresV1InnerResponse) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListTopKScoresV1InnerResponse) GetScores() []*ScoreV1 {
	if x != nil {
		return x.Scores
	}
	return nil
}

type ListTopKScoresV1OutterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ListTopKScoresV1InnerResponse `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ListTopKScoresV1OutterResponse) Reset() {
	*x = ListTopKScoresV1OutterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopKScoresV1OutterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopKScoresV1OutterResponse) ProtoMessage() {}

func (x *ListTopKScoresV1OutterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopKScoresV1OutterResponse.ProtoReflect.Descriptor instead.
func (*ListTopKScoresV1OutterResponse) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTopKScoresV1OutterResponse) GetResult() *ListTopKScoresV1InnerResponse {
	if x != nil {
		return x.Result
	}
	return nil
}

type InsertScoreV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ScoreV1 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InsertScoreV1Response) Reset() {
	*x = InsertScoreV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertScoreV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertScoreV1Response) ProtoMessage() {}

func (x *InsertScoreV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertScoreV1Response.ProtoReflect.Descriptor instead.
func (*InsertScoreV1Response) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{3}
}

func (x *InsertScoreV1Response) GetResult() *ScoreV1 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_score_service_proto protoreflect.FileDescriptor

var file_score_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x70, 0x4b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x47, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x56, 0x31, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x56, 0x31, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x1e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x56, 0x31, 0x4f,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x56, 0x31, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x56, 0x31, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xe8, 0x02, 0x0a,
	0x0e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12,
	0xbe, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x70, 0x4b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x70, 0x4b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x56, 0x31, 0x4f, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x92, 0x41, 0x50,
	0x12, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x70, 0x20, 0x6b, 0x20, 0x68, 0x69, 0x67,
	0x68, 0x65, 0x73, 0x74, 0x20, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x1a, 0x2d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x70, 0x20, 0x6b, 0x20, 0x68, 0x69,
	0x67, 0x68, 0x65, 0x73, 0x74, 0x20, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x94, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x56, 0x31, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x56, 0x31, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x55, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3d, 0x12, 0x15, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x20, 0x61, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x1a, 0x24, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x20, 0x61, 0x20, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x20, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x26, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x92, 0x41, 0x1a, 0x12, 0x05, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x1a, 0x0e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x39, 0x39, 0x39, 0x2a, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_score_service_proto_rawDescOnce sync.Once
	file_score_service_proto_rawDescData = file_score_service_proto_rawDesc
)

func file_score_service_proto_rawDescGZIP() []byte {
	file_score_service_proto_rawDescOnce.Do(func() {
		file_score_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_score_service_proto_rawDescData)
	})
	return file_score_service_proto_rawDescData
}

var file_score_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_score_service_proto_goTypes = []interface{}{
	(*ListTopKScoresV1Request)(nil),        // 0: proto.ListTopKScoresV1Request
	(*ListTopKScoresV1InnerResponse)(nil),  // 1: proto.ListTopKScoresV1InnerResponse
	(*ListTopKScoresV1OutterResponse)(nil), // 2: proto.ListTopKScoresV1OutterResponse
	(*InsertScoreV1Response)(nil),          // 3: proto.InsertScoreV1Response
	(*ScoreV1)(nil),                        // 4: proto.ScoreV1
}
var file_score_service_proto_depIdxs = []int32{
	4, // 0: proto.ListTopKScoresV1InnerResponse.scores:type_name -> proto.ScoreV1
	1, // 1: proto.ListTopKScoresV1OutterResponse.result:type_name -> proto.ListTopKScoresV1InnerResponse
	4, // 2: proto.InsertScoreV1Response.result:type_name -> proto.ScoreV1
	0, // 3: proto.ScoreServiceV1.ListTopKScores:input_type -> proto.ListTopKScoresV1Request
	4, // 4: proto.ScoreServiceV1.InsertScoreV1:input_type -> proto.ScoreV1
	2, // 5: proto.ScoreServiceV1.ListTopKScores:output_type -> proto.ListTopKScoresV1OutterResponse
	3, // 6: proto.ScoreServiceV1.InsertScoreV1:output_type -> proto.InsertScoreV1Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_score_service_proto_init() }
func file_score_service_proto_init() {
	if File_score_service_proto != nil {
		return
	}
	file_score_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_score_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopKScoresV1Request); i {
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
		file_score_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopKScoresV1InnerResponse); i {
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
		file_score_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopKScoresV1OutterResponse); i {
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
		file_score_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertScoreV1Response); i {
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
			RawDescriptor: file_score_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_score_service_proto_goTypes,
		DependencyIndexes: file_score_service_proto_depIdxs,
		MessageInfos:      file_score_service_proto_msgTypes,
	}.Build()
	File_score_service_proto = out.File
	file_score_service_proto_rawDesc = nil
	file_score_service_proto_goTypes = nil
	file_score_service_proto_depIdxs = nil
}
