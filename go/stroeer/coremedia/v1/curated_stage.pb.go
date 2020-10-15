// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.13.0
// source: stroeer/coremedia/v1/curated_stage.proto

package coremedia

import (
	proto "github.com/golang/protobuf/proto"
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

type CuratedStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles   []*CuratedArticle `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	StageLabel string            `protobuf:"bytes,2,opt,name=stage_label,json=stageLabel,proto3" json:"stage_label,omitempty"`
	StageId    string            `protobuf:"bytes,3,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	ListId     int64             `protobuf:"varint,4,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
}

func (x *CuratedStage) Reset() {
	*x = CuratedStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_coremedia_v1_curated_stage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CuratedStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CuratedStage) ProtoMessage() {}

func (x *CuratedStage) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_coremedia_v1_curated_stage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CuratedStage.ProtoReflect.Descriptor instead.
func (*CuratedStage) Descriptor() ([]byte, []int) {
	return file_stroeer_coremedia_v1_curated_stage_proto_rawDescGZIP(), []int{0}
}

func (x *CuratedStage) GetArticles() []*CuratedArticle {
	if x != nil {
		return x.Articles
	}
	return nil
}

func (x *CuratedStage) GetStageLabel() string {
	if x != nil {
		return x.StageLabel
	}
	return ""
}

func (x *CuratedStage) GetStageId() string {
	if x != nil {
		return x.StageId
	}
	return ""
}

func (x *CuratedStage) GetListId() int64 {
	if x != nil {
		return x.ListId
	}
	return 0
}

type CuratedArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CuratedHeadline string `protobuf:"bytes,3,opt,name=curated_headline,json=curatedHeadline,proto3" json:"curated_headline,omitempty"`
}

func (x *CuratedArticle) Reset() {
	*x = CuratedArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_coremedia_v1_curated_stage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CuratedArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CuratedArticle) ProtoMessage() {}

func (x *CuratedArticle) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_coremedia_v1_curated_stage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CuratedArticle.ProtoReflect.Descriptor instead.
func (*CuratedArticle) Descriptor() ([]byte, []int) {
	return file_stroeer_coremedia_v1_curated_stage_proto_rawDescGZIP(), []int{1}
}

func (x *CuratedArticle) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CuratedArticle) GetCuratedHeadline() string {
	if x != nil {
		return x.CuratedHeadline
	}
	return ""
}

var File_stroeer_coremedia_v1_curated_stage_proto protoreflect.FileDescriptor

var file_stroeer_coremedia_v1_curated_stage_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31,
	0x22, 0xa5, 0x01, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0e, 0x43, 0x75, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x48, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x57, 0x0a, 0x17, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_coremedia_v1_curated_stage_proto_rawDescOnce sync.Once
	file_stroeer_coremedia_v1_curated_stage_proto_rawDescData = file_stroeer_coremedia_v1_curated_stage_proto_rawDesc
)

func file_stroeer_coremedia_v1_curated_stage_proto_rawDescGZIP() []byte {
	file_stroeer_coremedia_v1_curated_stage_proto_rawDescOnce.Do(func() {
		file_stroeer_coremedia_v1_curated_stage_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_coremedia_v1_curated_stage_proto_rawDescData)
	})
	return file_stroeer_coremedia_v1_curated_stage_proto_rawDescData
}

var file_stroeer_coremedia_v1_curated_stage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_coremedia_v1_curated_stage_proto_goTypes = []interface{}{
	(*CuratedStage)(nil),   // 0: stroeer.coremedia.v1.CuratedStage
	(*CuratedArticle)(nil), // 1: stroeer.coremedia.v1.CuratedArticle
}
var file_stroeer_coremedia_v1_curated_stage_proto_depIdxs = []int32{
	1, // 0: stroeer.coremedia.v1.CuratedStage.articles:type_name -> stroeer.coremedia.v1.CuratedArticle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stroeer_coremedia_v1_curated_stage_proto_init() }
func file_stroeer_coremedia_v1_curated_stage_proto_init() {
	if File_stroeer_coremedia_v1_curated_stage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_coremedia_v1_curated_stage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CuratedStage); i {
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
		file_stroeer_coremedia_v1_curated_stage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CuratedArticle); i {
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
			RawDescriptor: file_stroeer_coremedia_v1_curated_stage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stroeer_coremedia_v1_curated_stage_proto_goTypes,
		DependencyIndexes: file_stroeer_coremedia_v1_curated_stage_proto_depIdxs,
		MessageInfos:      file_stroeer_coremedia_v1_curated_stage_proto_msgTypes,
	}.Build()
	File_stroeer_coremedia_v1_curated_stage_proto = out.File
	file_stroeer_coremedia_v1_curated_stage_proto_rawDesc = nil
	file_stroeer_coremedia_v1_curated_stage_proto_goTypes = nil
	file_stroeer_coremedia_v1_curated_stage_proto_depIdxs = nil
}
