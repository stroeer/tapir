// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package article

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WebArticleServiceClient is the client API for WebArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebArticleServiceClient interface {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage(ctx context.Context, in *GetWebArticlePageRequest, opts ...grpc.CallOption) (*GetWebArticlePageResponse, error)
}

type webArticleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebArticleServiceClient(cc grpc.ClientConnInterface) WebArticleServiceClient {
	return &webArticleServiceClient{cc}
}

var webArticleServiceGetWebArticlePageStreamDesc = &grpc.StreamDesc{
	StreamName: "GetWebArticlePage",
}

func (c *webArticleServiceClient) GetWebArticlePage(ctx context.Context, in *GetWebArticlePageRequest, opts ...grpc.CallOption) (*GetWebArticlePageResponse, error) {
	out := new(GetWebArticlePageResponse)
	err := c.cc.Invoke(ctx, "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebArticleServiceService is the service API for WebArticleService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterWebArticleServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type WebArticleServiceService struct {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage func(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error)
}

func (s *WebArticleServiceService) getWebArticlePage(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetWebArticlePage == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetWebArticlePage not implemented")
	}
	in := new(GetWebArticlePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetWebArticlePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetWebArticlePage(ctx, req.(*GetWebArticlePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterWebArticleServiceService registers a service implementation with a gRPC server.
func RegisterWebArticleServiceService(s grpc.ServiceRegistrar, srv *WebArticleServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "stroeer.web.article.v1.WebArticleService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetWebArticlePage",
				Handler:    srv.getWebArticlePage,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "stroeer/web/article/v1/web_article_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewWebArticleServiceService creates a new WebArticleServiceService containing the
// implemented methods of the WebArticleService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewWebArticleServiceService(s interface{}) *WebArticleServiceService {
	ns := &WebArticleServiceService{}
	if h, ok := s.(interface {
		GetWebArticlePage(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error)
	}); ok {
		ns.GetWebArticlePage = h.GetWebArticlePage
	}
	return ns
}

// UnstableWebArticleServiceService is the service API for WebArticleService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableWebArticleServiceService interface {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error)
}