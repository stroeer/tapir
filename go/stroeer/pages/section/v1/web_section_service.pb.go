// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: stroeer/pages/section/v1/web_section_service.proto

package section

import (
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/stroeer/tapir/go/stroeer/core/v1"
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

type GetSectionPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionPath string `protobuf:"bytes,1,opt,name=section_path,json=sectionPath,proto3" json:"section_path,omitempty"`
}

func (x *GetSectionPageRequest) Reset() {
	*x = GetSectionPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_pages_section_v1_web_section_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSectionPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSectionPageRequest) ProtoMessage() {}

func (x *GetSectionPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_pages_section_v1_web_section_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSectionPageRequest.ProtoReflect.Descriptor instead.
func (*GetSectionPageRequest) Descriptor() ([]byte, []int) {
	return file_stroeer_pages_section_v1_web_section_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSectionPageRequest) GetSectionPath() string {
	if x != nil {
		return x.SectionPath
	}
	return ""
}

type GetSectionPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*v1.Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *GetSectionPageResponse) Reset() {
	*x = GetSectionPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_pages_section_v1_web_section_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSectionPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSectionPageResponse) ProtoMessage() {}

func (x *GetSectionPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_pages_section_v1_web_section_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSectionPageResponse.ProtoReflect.Descriptor instead.
func (*GetSectionPageResponse) Descriptor() ([]byte, []int) {
	return file_stroeer_pages_section_v1_web_section_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSectionPageResponse) GetArticles() []*v1.Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

var File_stroeer_pages_section_v1_web_section_service_proto protoreflect.FileDescriptor

var file_stroeer_pages_section_v1_web_section_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x5f, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1d,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4e, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x32, 0x87, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x12, 0x2f,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x5d, 0x0a, 0x1b, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_pages_section_v1_web_section_service_proto_rawDescOnce sync.Once
	file_stroeer_pages_section_v1_web_section_service_proto_rawDescData = file_stroeer_pages_section_v1_web_section_service_proto_rawDesc
)

func file_stroeer_pages_section_v1_web_section_service_proto_rawDescGZIP() []byte {
	file_stroeer_pages_section_v1_web_section_service_proto_rawDescOnce.Do(func() {
		file_stroeer_pages_section_v1_web_section_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_pages_section_v1_web_section_service_proto_rawDescData)
	})
	return file_stroeer_pages_section_v1_web_section_service_proto_rawDescData
}

var file_stroeer_pages_section_v1_web_section_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_pages_section_v1_web_section_service_proto_goTypes = []interface{}{
	(*GetSectionPageRequest)(nil),  // 0: stroeer.pages.section.v1.GetSectionPageRequest
	(*GetSectionPageResponse)(nil), // 1: stroeer.pages.section.v1.GetSectionPageResponse
	(*v1.Article)(nil),             // 2: stroeer.core.v1.Article
}
var file_stroeer_pages_section_v1_web_section_service_proto_depIdxs = []int32{
	2, // 0: stroeer.pages.section.v1.GetSectionPageResponse.articles:type_name -> stroeer.core.v1.Article
	0, // 1: stroeer.pages.section.v1.SectionService.GetSectionPage:input_type -> stroeer.pages.section.v1.GetSectionPageRequest
	1, // 2: stroeer.pages.section.v1.SectionService.GetSectionPage:output_type -> stroeer.pages.section.v1.GetSectionPageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stroeer_pages_section_v1_web_section_service_proto_init() }
func file_stroeer_pages_section_v1_web_section_service_proto_init() {
	if File_stroeer_pages_section_v1_web_section_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_pages_section_v1_web_section_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSectionPageRequest); i {
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
		file_stroeer_pages_section_v1_web_section_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSectionPageResponse); i {
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
			RawDescriptor: file_stroeer_pages_section_v1_web_section_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stroeer_pages_section_v1_web_section_service_proto_goTypes,
		DependencyIndexes: file_stroeer_pages_section_v1_web_section_service_proto_depIdxs,
		MessageInfos:      file_stroeer_pages_section_v1_web_section_service_proto_msgTypes,
	}.Build()
	File_stroeer_pages_section_v1_web_section_service_proto = out.File
	file_stroeer_pages_section_v1_web_section_service_proto_rawDesc = nil
	file_stroeer_pages_section_v1_web_section_service_proto_goTypes = nil
	file_stroeer_pages_section_v1_web_section_service_proto_depIdxs = nil
}
