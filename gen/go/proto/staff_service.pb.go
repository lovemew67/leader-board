// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: staff_service.proto

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

type ListStaffV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStaffV1Request) Reset() {
	*x = ListStaffV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffV1Request) ProtoMessage() {}

func (x *ListStaffV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_staff_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffV1Request.ProtoReflect.Descriptor instead.
func (*ListStaffV1Request) Descriptor() ([]byte, []int) {
	return file_staff_service_proto_rawDescGZIP(), []int{0}
}

type ListStaffV1InnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staff []*StaffV1 `protobuf:"bytes,1,rep,name=staff,proto3" json:"staff,omitempty"`
	Total int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListStaffV1InnerResponse) Reset() {
	*x = ListStaffV1InnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffV1InnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffV1InnerResponse) ProtoMessage() {}

func (x *ListStaffV1InnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffV1InnerResponse.ProtoReflect.Descriptor instead.
func (*ListStaffV1InnerResponse) Descriptor() ([]byte, []int) {
	return file_staff_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListStaffV1InnerResponse) GetStaff() []*StaffV1 {
	if x != nil {
		return x.Staff
	}
	return nil
}

func (x *ListStaffV1InnerResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListStaffV1OutterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ListStaffV1InnerResponse `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ListStaffV1OutterResponse) Reset() {
	*x = ListStaffV1OutterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffV1OutterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffV1OutterResponse) ProtoMessage() {}

func (x *ListStaffV1OutterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffV1OutterResponse.ProtoReflect.Descriptor instead.
func (*ListStaffV1OutterResponse) Descriptor() ([]byte, []int) {
	return file_staff_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListStaffV1OutterResponse) GetResult() *ListStaffV1InnerResponse {
	if x != nil {
		return x.Result
	}
	return nil
}

type CreateStaffV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *StaffV1 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateStaffV1Response) Reset() {
	*x = CreateStaffV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffV1Response) ProtoMessage() {}

func (x *CreateStaffV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_staff_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffV1Response.ProtoReflect.Descriptor instead.
func (*CreateStaffV1Response) Descriptor() ([]byte, []int) {
	return file_staff_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStaffV1Response) GetResult() *StaffV1 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_staff_service_proto protoreflect.FileDescriptor

var file_staff_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x52, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x54, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x56, 0x31, 0x4f, 0x75, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x56, 0x31, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9f, 0x02, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12,
	0x8b, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x4f, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x92,
	0x41, 0x2b, 0x12, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x1d,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x73, 0x74, 0x61, 0x66, 0x66, 0x20, 0x6f,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x12, 0x7f, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x12, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x56, 0x31, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x29, 0x12, 0x0b, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x1a, 0x1a, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x73, 0x74, 0x61, 0x66, 0x66, 0x20,
	0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x26,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x1a, 0x12, 0x05, 0x32, 0x03,
	0x30, 0x2e, 0x31, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38,
	0x39, 0x39, 0x39, 0x2a, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staff_service_proto_rawDescOnce sync.Once
	file_staff_service_proto_rawDescData = file_staff_service_proto_rawDesc
)

func file_staff_service_proto_rawDescGZIP() []byte {
	file_staff_service_proto_rawDescOnce.Do(func() {
		file_staff_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_service_proto_rawDescData)
	})
	return file_staff_service_proto_rawDescData
}

var file_staff_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_staff_service_proto_goTypes = []interface{}{
	(*ListStaffV1Request)(nil),        // 0: proto.ListStaffV1Request
	(*ListStaffV1InnerResponse)(nil),  // 1: proto.ListStaffV1InnerResponse
	(*ListStaffV1OutterResponse)(nil), // 2: proto.ListStaffV1OutterResponse
	(*CreateStaffV1Response)(nil),     // 3: proto.CreateStaffV1Response
	(*StaffV1)(nil),                   // 4: proto.StaffV1
}
var file_staff_service_proto_depIdxs = []int32{
	4, // 0: proto.ListStaffV1InnerResponse.staff:type_name -> proto.StaffV1
	1, // 1: proto.ListStaffV1OutterResponse.result:type_name -> proto.ListStaffV1InnerResponse
	4, // 2: proto.CreateStaffV1Response.result:type_name -> proto.StaffV1
	0, // 3: proto.StaffServiceV1.ListStaffV1:input_type -> proto.ListStaffV1Request
	4, // 4: proto.StaffServiceV1.CreateStaffV1:input_type -> proto.StaffV1
	2, // 5: proto.StaffServiceV1.ListStaffV1:output_type -> proto.ListStaffV1OutterResponse
	3, // 6: proto.StaffServiceV1.CreateStaffV1:output_type -> proto.CreateStaffV1Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_staff_service_proto_init() }
func file_staff_service_proto_init() {
	if File_staff_service_proto != nil {
		return
	}
	file_staff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staff_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffV1Request); i {
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
		file_staff_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffV1InnerResponse); i {
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
		file_staff_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffV1OutterResponse); i {
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
		file_staff_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffV1Response); i {
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
			RawDescriptor: file_staff_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_service_proto_goTypes,
		DependencyIndexes: file_staff_service_proto_depIdxs,
		MessageInfos:      file_staff_service_proto_msgTypes,
	}.Build()
	File_staff_service_proto = out.File
	file_staff_service_proto_rawDesc = nil
	file_staff_service_proto_goTypes = nil
	file_staff_service_proto_depIdxs = nil
}
