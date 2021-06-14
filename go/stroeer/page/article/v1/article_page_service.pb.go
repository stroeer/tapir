// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: stroeer/page/article/v1/article_page_service.proto

package article

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

// Request message to get an article page.
type GetArticlePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the article defined by the content management system (required).
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticlePageRequest) Reset() {
	*x = GetArticlePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_article_v1_article_page_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlePageRequest) ProtoMessage() {}

func (x *GetArticlePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_article_v1_article_page_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlePageRequest.ProtoReflect.Descriptor instead.
func (*GetArticlePageRequest) Descriptor() ([]byte, []int) {
	return file_stroeer_page_article_v1_article_page_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetArticlePageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response message for an article page request.
//
// Status codes:
// * `OK`: article exists and is published
// * `NOT_FOUND`: article doesn't exist or is not published according to it's `Metadata`
type GetArticlePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Article page with all render relevant data for the user and SEO bots.
	ArticlePage *ArticlePage `protobuf:"bytes,1,opt,name=article_page,json=articlePage,proto3" json:"article_page,omitempty"`
}

func (x *GetArticlePageResponse) Reset() {
	*x = GetArticlePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_article_v1_article_page_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlePageResponse) ProtoMessage() {}

func (x *GetArticlePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_article_v1_article_page_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlePageResponse.ProtoReflect.Descriptor instead.
func (*GetArticlePageResponse) Descriptor() ([]byte, []int) {
	return file_stroeer_page_article_v1_article_page_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticlePageResponse) GetArticlePage() *ArticlePage {
	if x != nil {
		return x.ArticlePage
	}
	return nil
}

var File_stroeer_page_article_v1_article_page_service_proto protoreflect.FileDescriptor

var file_stroeer_page_article_v1_article_page_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x61, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x2e,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x5b, 0x0a, 0x1a, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_page_article_v1_article_page_service_proto_rawDescOnce sync.Once
	file_stroeer_page_article_v1_article_page_service_proto_rawDescData = file_stroeer_page_article_v1_article_page_service_proto_rawDesc
)

func file_stroeer_page_article_v1_article_page_service_proto_rawDescGZIP() []byte {
	file_stroeer_page_article_v1_article_page_service_proto_rawDescOnce.Do(func() {
		file_stroeer_page_article_v1_article_page_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_page_article_v1_article_page_service_proto_rawDescData)
	})
	return file_stroeer_page_article_v1_article_page_service_proto_rawDescData
}

var file_stroeer_page_article_v1_article_page_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_page_article_v1_article_page_service_proto_goTypes = []interface{}{
	(*GetArticlePageRequest)(nil),  // 0: stroeer.page.article.v1.GetArticlePageRequest
	(*GetArticlePageResponse)(nil), // 1: stroeer.page.article.v1.GetArticlePageResponse
	(*ArticlePage)(nil),            // 2: stroeer.page.article.v1.ArticlePage
}
var file_stroeer_page_article_v1_article_page_service_proto_depIdxs = []int32{
	2, // 0: stroeer.page.article.v1.GetArticlePageResponse.article_page:type_name -> stroeer.page.article.v1.ArticlePage
	0, // 1: stroeer.page.article.v1.ArticlePageService.GetArticlePage:input_type -> stroeer.page.article.v1.GetArticlePageRequest
	1, // 2: stroeer.page.article.v1.ArticlePageService.GetArticlePage:output_type -> stroeer.page.article.v1.GetArticlePageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stroeer_page_article_v1_article_page_service_proto_init() }
func file_stroeer_page_article_v1_article_page_service_proto_init() {
	if File_stroeer_page_article_v1_article_page_service_proto != nil {
		return
	}
	file_stroeer_page_article_v1_article_page_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stroeer_page_article_v1_article_page_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlePageRequest); i {
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
		file_stroeer_page_article_v1_article_page_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlePageResponse); i {
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
			RawDescriptor: file_stroeer_page_article_v1_article_page_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stroeer_page_article_v1_article_page_service_proto_goTypes,
		DependencyIndexes: file_stroeer_page_article_v1_article_page_service_proto_depIdxs,
		MessageInfos:      file_stroeer_page_article_v1_article_page_service_proto_msgTypes,
	}.Build()
	File_stroeer_page_article_v1_article_page_service_proto = out.File
	file_stroeer_page_article_v1_article_page_service_proto_rawDesc = nil
	file_stroeer_page_article_v1_article_page_service_proto_goTypes = nil
	file_stroeer_page_article_v1_article_page_service_proto_depIdxs = nil
}
