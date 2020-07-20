// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package coremedia

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurationServiceClient is the client API for CurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurationServiceClient interface {
	GetCuratedList(ctx context.Context, in *GetCuratedListRequest, opts ...grpc.CallOption) (*GetCuratedListResponse, error)
}

type curationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurationServiceClient(cc grpc.ClientConnInterface) CurationServiceClient {
	return &curationServiceClient{cc}
}

func (c *curationServiceClient) GetCuratedList(ctx context.Context, in *GetCuratedListRequest, opts ...grpc.CallOption) (*GetCuratedListResponse, error) {
	out := new(GetCuratedListResponse)
	err := c.cc.Invoke(ctx, "/stroeer.coremedia.v1.CurationService/GetCuratedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurationServiceServer is the server API for CurationService service.
// All implementations must embed UnimplementedCurationServiceServer
// for forward compatibility
type CurationServiceServer interface {
	GetCuratedList(context.Context, *GetCuratedListRequest) (*GetCuratedListResponse, error)
	mustEmbedUnimplementedCurationServiceServer()
}

// UnimplementedCurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurationServiceServer struct {
}

func (*UnimplementedCurationServiceServer) GetCuratedList(context.Context, *GetCuratedListRequest) (*GetCuratedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCuratedList not implemented")
}
func (*UnimplementedCurationServiceServer) mustEmbedUnimplementedCurationServiceServer() {}

func RegisterCurationServiceServer(s *grpc.Server, srv CurationServiceServer) {
	s.RegisterService(&_CurationService_serviceDesc, srv)
}

func _CurationService_GetCuratedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCuratedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurationServiceServer).GetCuratedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stroeer.coremedia.v1.CurationService/GetCuratedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurationServiceServer).GetCuratedList(ctx, req.(*GetCuratedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stroeer.coremedia.v1.CurationService",
	HandlerType: (*CurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCuratedList",
			Handler:    _CurationService_GetCuratedList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stroeer/coremedia/v1/curation_service.proto",
}
