// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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

// FilterServiceClient is the client API for FilterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilterServiceClient interface {
	GetFilterValues(ctx context.Context, in *GetFilterValuesReq, opts ...grpc.CallOption) (*GetFilterValuesResp, error)
	GetSearchQuery(ctx context.Context, in *GetSearchQueryReq, opts ...grpc.CallOption) (*GetSearchQueryResp, error)
	GetLemmasByFilterID(ctx context.Context, in *GetLemmasByFilterIDReq, opts ...grpc.CallOption) (*GetLemmasByFilterIDResp, error)
	GetKeywordsByLemmas(ctx context.Context, in *GetKeywordsByLemmasReq, opts ...grpc.CallOption) (*GetKeywordsByLemmasResp, error)
}

type filterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilterServiceClient(cc grpc.ClientConnInterface) FilterServiceClient {
	return &filterServiceClient{cc}
}

func (c *filterServiceClient) GetFilterValues(ctx context.Context, in *GetFilterValuesReq, opts ...grpc.CallOption) (*GetFilterValuesResp, error) {
	out := new(GetFilterValuesResp)
	err := c.cc.Invoke(ctx, "/main.FilterService/GetFilterValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filterServiceClient) GetSearchQuery(ctx context.Context, in *GetSearchQueryReq, opts ...grpc.CallOption) (*GetSearchQueryResp, error) {
	out := new(GetSearchQueryResp)
	err := c.cc.Invoke(ctx, "/main.FilterService/GetSearchQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filterServiceClient) GetLemmasByFilterID(ctx context.Context, in *GetLemmasByFilterIDReq, opts ...grpc.CallOption) (*GetLemmasByFilterIDResp, error) {
	out := new(GetLemmasByFilterIDResp)
	err := c.cc.Invoke(ctx, "/main.FilterService/GetLemmasByFilterID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filterServiceClient) GetKeywordsByLemmas(ctx context.Context, in *GetKeywordsByLemmasReq, opts ...grpc.CallOption) (*GetKeywordsByLemmasResp, error) {
	out := new(GetKeywordsByLemmasResp)
	err := c.cc.Invoke(ctx, "/main.FilterService/GetKeywordsByLemmas", in, out, opts...)
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
	GetSearchQuery(context.Context, *GetSearchQueryReq) (*GetSearchQueryResp, error)
	GetLemmasByFilterID(context.Context, *GetLemmasByFilterIDReq) (*GetLemmasByFilterIDResp, error)
	GetKeywordsByLemmas(context.Context, *GetKeywordsByLemmasReq) (*GetKeywordsByLemmasResp, error)
	mustEmbedUnimplementedFilterServiceServer()
}

// UnimplementedFilterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilterServiceServer struct {
}

func (UnimplementedFilterServiceServer) GetFilterValues(context.Context, *GetFilterValuesReq) (*GetFilterValuesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilterValues not implemented")
}
func (UnimplementedFilterServiceServer) GetSearchQuery(context.Context, *GetSearchQueryReq) (*GetSearchQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchQuery not implemented")
}
func (UnimplementedFilterServiceServer) GetLemmasByFilterID(context.Context, *GetLemmasByFilterIDReq) (*GetLemmasByFilterIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLemmasByFilterID not implemented")
}
func (UnimplementedFilterServiceServer) GetKeywordsByLemmas(context.Context, *GetKeywordsByLemmasReq) (*GetKeywordsByLemmasResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordsByLemmas not implemented")
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
		FullMethod: "/main.FilterService/GetFilterValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilterServiceServer).GetFilterValues(ctx, req.(*GetFilterValuesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilterService_GetSearchQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilterServiceServer).GetSearchQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.FilterService/GetSearchQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilterServiceServer).GetSearchQuery(ctx, req.(*GetSearchQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilterService_GetLemmasByFilterID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLemmasByFilterIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilterServiceServer).GetLemmasByFilterID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.FilterService/GetLemmasByFilterID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilterServiceServer).GetLemmasByFilterID(ctx, req.(*GetLemmasByFilterIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilterService_GetKeywordsByLemmas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordsByLemmasReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilterServiceServer).GetKeywordsByLemmas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.FilterService/GetKeywordsByLemmas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilterServiceServer).GetKeywordsByLemmas(ctx, req.(*GetKeywordsByLemmasReq))
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
		{
			MethodName: "GetSearchQuery",
			Handler:    _FilterService_GetSearchQuery_Handler,
		},
		{
			MethodName: "GetLemmasByFilterID",
			Handler:    _FilterService_GetLemmasByFilterID_Handler,
		},
		{
			MethodName: "GetKeywordsByLemmas",
			Handler:    _FilterService_GetKeywordsByLemmas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filter_service.proto",
}
