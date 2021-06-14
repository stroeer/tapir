// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: stroeer/page/article/v1/article_page.proto

package article

import (
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

// Source of the related article.
type RelatedArticleSource int32

const (
	// Not specified.
	RelatedArticleSource_RELATED_ARTICLE_SOURCE_UNSPECIFIED RelatedArticleSource = 0
	// Article was related manually by the editorial department.
	RelatedArticleSource_RELATED_ARTICLE_SOURCE_EDITORIAL RelatedArticleSource = 1
)

// Enum value maps for RelatedArticleSource.
var (
	RelatedArticleSource_name = map[int32]string{
		0: "RELATED_ARTICLE_SOURCE_UNSPECIFIED",
		1: "RELATED_ARTICLE_SOURCE_EDITORIAL",
	}
	RelatedArticleSource_value = map[string]int32{
		"RELATED_ARTICLE_SOURCE_UNSPECIFIED": 0,
		"RELATED_ARTICLE_SOURCE_EDITORIAL":   1,
	}
)

func (x RelatedArticleSource) Enum() *RelatedArticleSource {
	p := new(RelatedArticleSource)
	*p = x
	return p
}

func (x RelatedArticleSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelatedArticleSource) Descriptor() protoreflect.EnumDescriptor {
	return file_stroeer_page_article_v1_article_page_proto_enumTypes[0].Descriptor()
}

func (RelatedArticleSource) Type() protoreflect.EnumType {
	return &file_stroeer_page_article_v1_article_page_proto_enumTypes[0]
}

func (x RelatedArticleSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelatedArticleSource.Descriptor instead.
func (RelatedArticleSource) EnumDescriptor() ([]byte, []int) {
	return file_stroeer_page_article_v1_article_page_proto_rawDescGZIP(), []int{0}
}

// Article page with all render relevant data for the user and SEO bots.
type ArticlePage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The editorial article with the ID of a `GetArticlePageRequest`.
	//
	// Unpublished elements will be filtered according to their metadata.
	Article *v1.Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	// Additional related editorial articles. Articles which are not published
	// according to their `Metadata` will be filtered.
	RelatedArticles []*RelatedArticle `protobuf:"bytes,2,rep,name=related_articles,json=relatedArticles,proto3" json:"related_articles,omitempty"`
}

func (x *ArticlePage) Reset() {
	*x = ArticlePage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_article_v1_article_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlePage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlePage) ProtoMessage() {}

func (x *ArticlePage) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_article_v1_article_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlePage.ProtoReflect.Descriptor instead.
func (*ArticlePage) Descriptor() ([]byte, []int) {
	return file_stroeer_page_article_v1_article_page_proto_rawDescGZIP(), []int{0}
}

func (x *ArticlePage) GetArticle() *v1.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *ArticlePage) GetRelatedArticles() []*RelatedArticle {
	if x != nil {
		return x.RelatedArticles
	}
	return nil
}

// An editorial article, which is related to another article.
//
// Related articles are specified manually be the editorial department or
// identified by recommendation systems.
type RelatedArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data of the related article. This message will only contain data required
	// to render the article as a _Teaser_ (e.g. with `Article.body` set to `null`
	// thus not containing any data that is only required on detail pages).
	Article *v1.Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	// Source of the related article.
	Source RelatedArticleSource `protobuf:"varint,2,opt,name=source,proto3,enum=stroeer.page.article.v1.RelatedArticleSource" json:"source,omitempty"`
}

func (x *RelatedArticle) Reset() {
	*x = RelatedArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_article_v1_article_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatedArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatedArticle) ProtoMessage() {}

func (x *RelatedArticle) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_article_v1_article_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatedArticle.ProtoReflect.Descriptor instead.
func (*RelatedArticle) Descriptor() ([]byte, []int) {
	return file_stroeer_page_article_v1_article_page_proto_rawDescGZIP(), []int{1}
}

func (x *RelatedArticle) GetArticle() *v1.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *RelatedArticle) GetSource() RelatedArticleSource {
	if x != nil {
		return x.Source
	}
	return RelatedArticleSource_RELATED_ARTICLE_SOURCE_UNSPECIFIED
}

var File_stroeer_page_article_v1_article_page_proto protoreflect.FileDescriptor

var file_stroeer_page_article_v1_article_page_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0x64, 0x0a, 0x14, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x52,
	0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x01,
	0x42, 0x5b, 0x0a, 0x1a, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_page_article_v1_article_page_proto_rawDescOnce sync.Once
	file_stroeer_page_article_v1_article_page_proto_rawDescData = file_stroeer_page_article_v1_article_page_proto_rawDesc
)

func file_stroeer_page_article_v1_article_page_proto_rawDescGZIP() []byte {
	file_stroeer_page_article_v1_article_page_proto_rawDescOnce.Do(func() {
		file_stroeer_page_article_v1_article_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_page_article_v1_article_page_proto_rawDescData)
	})
	return file_stroeer_page_article_v1_article_page_proto_rawDescData
}

var file_stroeer_page_article_v1_article_page_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stroeer_page_article_v1_article_page_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_page_article_v1_article_page_proto_goTypes = []interface{}{
	(RelatedArticleSource)(0), // 0: stroeer.page.article.v1.RelatedArticleSource
	(*ArticlePage)(nil),       // 1: stroeer.page.article.v1.ArticlePage
	(*RelatedArticle)(nil),    // 2: stroeer.page.article.v1.RelatedArticle
	(*v1.Article)(nil),        // 3: stroeer.core.v1.Article
}
var file_stroeer_page_article_v1_article_page_proto_depIdxs = []int32{
	3, // 0: stroeer.page.article.v1.ArticlePage.article:type_name -> stroeer.core.v1.Article
	2, // 1: stroeer.page.article.v1.ArticlePage.related_articles:type_name -> stroeer.page.article.v1.RelatedArticle
	3, // 2: stroeer.page.article.v1.RelatedArticle.article:type_name -> stroeer.core.v1.Article
	0, // 3: stroeer.page.article.v1.RelatedArticle.source:type_name -> stroeer.page.article.v1.RelatedArticleSource
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_stroeer_page_article_v1_article_page_proto_init() }
func file_stroeer_page_article_v1_article_page_proto_init() {
	if File_stroeer_page_article_v1_article_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_page_article_v1_article_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlePage); i {
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
		file_stroeer_page_article_v1_article_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatedArticle); i {
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
			RawDescriptor: file_stroeer_page_article_v1_article_page_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stroeer_page_article_v1_article_page_proto_goTypes,
		DependencyIndexes: file_stroeer_page_article_v1_article_page_proto_depIdxs,
		EnumInfos:         file_stroeer_page_article_v1_article_page_proto_enumTypes,
		MessageInfos:      file_stroeer_page_article_v1_article_page_proto_msgTypes,
	}.Build()
	File_stroeer_page_article_v1_article_page_proto = out.File
	file_stroeer_page_article_v1_article_page_proto_rawDesc = nil
	file_stroeer_page_article_v1_article_page_proto_goTypes = nil
	file_stroeer_page_article_v1_article_page_proto_depIdxs = nil
}
