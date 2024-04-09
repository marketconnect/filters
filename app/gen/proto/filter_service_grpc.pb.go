// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: filter_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FilterService_GetFilterValues_FullMethodName = "/main.FilterService/GetFilterValues"
)

// FilterServiceClient is the client API for FilterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilterServiceClient interface {
	GetFilterValues(ctx context.Context, in *GetFilterValuesReq, opts ...grpc.CallOption) (*GetFilterValuesResp, error)
}

type filterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilterServiceClient(cc grpc.ClientConnInterface) FilterServiceClient {
	return &filterServiceClient{cc}
}

func (c *filterServiceClient) GetFilterValues(ctx context.Context, in *GetFilterValuesReq, opts ...grpc.CallOption) (*GetFilterValuesResp, error) {
	out := new(GetFilterValuesResp)
	err := c.cc.Invoke(ctx, FilterService_GetFilterValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilterServiceServer is the server API for FilterService service.
// All implementations must embed UnimplementedFilterServiceServer
// for forward compatibility
type FilterServiceServer interface {
	GetFilterValues(context.Context, *GetFilterValuesReq) (*GetFilterValuesResp, error)
	mustEmbedUnimplementedFilterServiceServer()
}

// UnimplementedFilterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilterServiceServer struct {
}

func (UnimplementedFilterServiceServer) GetFilterValues(context.Context, *GetFilterValuesReq) (*GetFilterValuesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilterValues not implemented")
}
func (UnimplementedFilterServiceServer) mustEmbedUnimplementedFilterServiceServer() {}

// UnsafeFilterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilterServiceServer will
// result in compilation errors.
type UnsafeFilterServiceServer interface {
	mustEmbedUnimplementedFilterServiceServer()
}

func RegisterFilterServiceServer(s grpc.ServiceRegistrar, srv FilterServiceServer) {
	s.RegisterService(&FilterService_ServiceDesc, srv)
}

func _FilterService_GetFilterValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilterValuesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilterServiceServer).GetFilterValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilterService_GetFilterValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilterServiceServer).GetFilterValues(ctx, req.(*GetFilterValuesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FilterService_ServiceDesc is the grpc.ServiceDesc for FilterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.FilterService",
	HandlerType: (*FilterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFilterValues",
			Handler:    _FilterService_GetFilterValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filter_service.proto",
}

const (
	SearchQueryService_GetSearchQuery_FullMethodName = "/main.SearchQueryService/GetSearchQuery"
)

// SearchQueryServiceClient is the client API for SearchQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchQueryServiceClient interface {
	GetSearchQuery(ctx context.Context, in *GetSearchQueryReq, opts ...grpc.CallOption) (*GetSearchQueryResp, error)
}

type searchQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchQueryServiceClient(cc grpc.ClientConnInterface) SearchQueryServiceClient {
	return &searchQueryServiceClient{cc}
}

func (c *searchQueryServiceClient) GetSearchQuery(ctx context.Context, in *GetSearchQueryReq, opts ...grpc.CallOption) (*GetSearchQueryResp, error) {
	out := new(GetSearchQueryResp)
	err := c.cc.Invoke(ctx, SearchQueryService_GetSearchQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchQueryServiceServer is the server API for SearchQueryService service.
// All implementations must embed UnimplementedSearchQueryServiceServer
// for forward compatibility
type SearchQueryServiceServer interface {
	GetSearchQuery(context.Context, *GetSearchQueryReq) (*GetSearchQueryResp, error)
	mustEmbedUnimplementedSearchQueryServiceServer()
}

// UnimplementedSearchQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchQueryServiceServer struct {
}

func (UnimplementedSearchQueryServiceServer) GetSearchQuery(context.Context, *GetSearchQueryReq) (*GetSearchQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchQuery not implemented")
}
func (UnimplementedSearchQueryServiceServer) mustEmbedUnimplementedSearchQueryServiceServer() {}

// UnsafeSearchQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchQueryServiceServer will
// result in compilation errors.
type UnsafeSearchQueryServiceServer interface {
	mustEmbedUnimplementedSearchQueryServiceServer()
}

func RegisterSearchQueryServiceServer(s grpc.ServiceRegistrar, srv SearchQueryServiceServer) {
	s.RegisterService(&SearchQueryService_ServiceDesc, srv)
}

func _SearchQueryService_GetSearchQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchQueryServiceServer).GetSearchQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchQueryService_GetSearchQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchQueryServiceServer).GetSearchQuery(ctx, req.(*GetSearchQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchQueryService_ServiceDesc is the grpc.ServiceDesc for SearchQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.SearchQueryService",
	HandlerType: (*SearchQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSearchQuery",
			Handler:    _SearchQueryService_GetSearchQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filter_service.proto",
}
