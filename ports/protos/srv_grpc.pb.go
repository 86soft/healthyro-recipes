// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hproto

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

// RecipeSvcClient is the client API for RecipeSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipeSvcClient interface {
	// Read
	ListRecipe(ctx context.Context, in *ListRecipeRequest, opts ...grpc.CallOption) (*ListRecipeResponse, error)
	FindRecipesByName(ctx context.Context, in *FindRecipesByNameRequest, opts ...grpc.CallOption) (*FindRecipesByNameResponse, error)
	GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error)
	FindRecipesByTags(ctx context.Context, in *FindRecipesByTagsRequest, opts ...grpc.CallOption) (*FindRecipesByTagsResponse, error)
	FindRecipesByNameAndTags(ctx context.Context, in *FindRecipesByNameAndTagsRequest, opts ...grpc.CallOption) (*FindRecipesByNameAndTagsResponse, error)
	// Write
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	UpdateRecipeTitle(ctx context.Context, in *UpdateRecipeTitleRequest, opts ...grpc.CallOption) (*UpdateRecipeTitleResponse, error)
	UpdateRecipeDescription(ctx context.Context, in *UpdateRecipeDescriptionRequest, opts ...grpc.CallOption) (*UpdateRecipeDescriptionResponse, error)
	DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error)
	RemoveResourceFromRecipe(ctx context.Context, in *RemoveResourceFromRecipeRequest, opts ...grpc.CallOption) (*RemoveRecipeFromResourceResponse, error)
	AddRecipeResource(ctx context.Context, in *AddRecipeResourceRequest, opts ...grpc.CallOption) (*AddRecipeResourceResponse, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error)
	AddTagToRecipe(ctx context.Context, in *AddTagToRecipeRequest, opts ...grpc.CallOption) (*AddTagToRecipeResponse, error)
	RemoveTagFromRecipe(ctx context.Context, in *RemoveTagFromRecipeRequest, opts ...grpc.CallOption) (*RemoveTagFromRecipeResponse, error)
}

type recipeSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipeSvcClient(cc grpc.ClientConnInterface) RecipeSvcClient {
	return &recipeSvcClient{cc}
}

func (c *recipeSvcClient) ListRecipe(ctx context.Context, in *ListRecipeRequest, opts ...grpc.CallOption) (*ListRecipeResponse, error) {
	out := new(ListRecipeResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/ListRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) FindRecipesByName(ctx context.Context, in *FindRecipesByNameRequest, opts ...grpc.CallOption) (*FindRecipesByNameResponse, error) {
	out := new(FindRecipesByNameResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/FindRecipesByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error) {
	out := new(GetRecipeResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/GetRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) FindRecipesByTags(ctx context.Context, in *FindRecipesByTagsRequest, opts ...grpc.CallOption) (*FindRecipesByTagsResponse, error) {
	out := new(FindRecipesByTagsResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/FindRecipesByTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) FindRecipesByNameAndTags(ctx context.Context, in *FindRecipesByNameAndTagsRequest, opts ...grpc.CallOption) (*FindRecipesByNameAndTagsResponse, error) {
	out := new(FindRecipesByNameAndTagsResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/FindRecipesByNameAndTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/CreateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) UpdateRecipeTitle(ctx context.Context, in *UpdateRecipeTitleRequest, opts ...grpc.CallOption) (*UpdateRecipeTitleResponse, error) {
	out := new(UpdateRecipeTitleResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/UpdateRecipeTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) UpdateRecipeDescription(ctx context.Context, in *UpdateRecipeDescriptionRequest, opts ...grpc.CallOption) (*UpdateRecipeDescriptionResponse, error) {
	out := new(UpdateRecipeDescriptionResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/UpdateRecipeDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error) {
	out := new(DeleteRecipeResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/DeleteRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) RemoveResourceFromRecipe(ctx context.Context, in *RemoveResourceFromRecipeRequest, opts ...grpc.CallOption) (*RemoveRecipeFromResourceResponse, error) {
	out := new(RemoveRecipeFromResourceResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/RemoveResourceFromRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) AddRecipeResource(ctx context.Context, in *AddRecipeResourceRequest, opts ...grpc.CallOption) (*AddRecipeResourceResponse, error) {
	out := new(AddRecipeResourceResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/AddRecipeResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error) {
	out := new(CreateTagResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) AddTagToRecipe(ctx context.Context, in *AddTagToRecipeRequest, opts ...grpc.CallOption) (*AddTagToRecipeResponse, error) {
	out := new(AddTagToRecipeResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/AddTagToRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeSvcClient) RemoveTagFromRecipe(ctx context.Context, in *RemoveTagFromRecipeRequest, opts ...grpc.CallOption) (*RemoveTagFromRecipeResponse, error) {
	out := new(RemoveTagFromRecipeResponse)
	err := c.cc.Invoke(ctx, "/RecipeSvc/RemoveTagFromRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipeSvcServer is the server API for RecipeSvc service.
// All implementations must embed UnimplementedRecipeSvcServer
// for forward compatibility
type RecipeSvcServer interface {
	// Read
	ListRecipe(context.Context, *ListRecipeRequest) (*ListRecipeResponse, error)
	FindRecipesByName(context.Context, *FindRecipesByNameRequest) (*FindRecipesByNameResponse, error)
	GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error)
	FindRecipesByTags(context.Context, *FindRecipesByTagsRequest) (*FindRecipesByTagsResponse, error)
	FindRecipesByNameAndTags(context.Context, *FindRecipesByNameAndTagsRequest) (*FindRecipesByNameAndTagsResponse, error)
	// Write
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	UpdateRecipeTitle(context.Context, *UpdateRecipeTitleRequest) (*UpdateRecipeTitleResponse, error)
	UpdateRecipeDescription(context.Context, *UpdateRecipeDescriptionRequest) (*UpdateRecipeDescriptionResponse, error)
	DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error)
	RemoveResourceFromRecipe(context.Context, *RemoveResourceFromRecipeRequest) (*RemoveRecipeFromResourceResponse, error)
	AddRecipeResource(context.Context, *AddRecipeResourceRequest) (*AddRecipeResourceResponse, error)
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error)
	AddTagToRecipe(context.Context, *AddTagToRecipeRequest) (*AddTagToRecipeResponse, error)
	RemoveTagFromRecipe(context.Context, *RemoveTagFromRecipeRequest) (*RemoveTagFromRecipeResponse, error)
	mustEmbedUnimplementedRecipeSvcServer()
}

// UnimplementedRecipeSvcServer must be embedded to have forward compatible implementations.
type UnimplementedRecipeSvcServer struct {
}

func (UnimplementedRecipeSvcServer) ListRecipe(context.Context, *ListRecipeRequest) (*ListRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) FindRecipesByName(context.Context, *FindRecipesByNameRequest) (*FindRecipesByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRecipesByName not implemented")
}
func (UnimplementedRecipeSvcServer) GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) FindRecipesByTags(context.Context, *FindRecipesByTagsRequest) (*FindRecipesByTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRecipesByTags not implemented")
}
func (UnimplementedRecipeSvcServer) FindRecipesByNameAndTags(context.Context, *FindRecipesByNameAndTagsRequest) (*FindRecipesByNameAndTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRecipesByNameAndTags not implemented")
}
func (UnimplementedRecipeSvcServer) CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) UpdateRecipeTitle(context.Context, *UpdateRecipeTitleRequest) (*UpdateRecipeTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipeTitle not implemented")
}
func (UnimplementedRecipeSvcServer) UpdateRecipeDescription(context.Context, *UpdateRecipeDescriptionRequest) (*UpdateRecipeDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipeDescription not implemented")
}
func (UnimplementedRecipeSvcServer) DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) RemoveResourceFromRecipe(context.Context, *RemoveResourceFromRecipeRequest) (*RemoveRecipeFromResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveResourceFromRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) AddRecipeResource(context.Context, *AddRecipeResourceRequest) (*AddRecipeResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecipeResource not implemented")
}
func (UnimplementedRecipeSvcServer) CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedRecipeSvcServer) AddTagToRecipe(context.Context, *AddTagToRecipeRequest) (*AddTagToRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTagToRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) RemoveTagFromRecipe(context.Context, *RemoveTagFromRecipeRequest) (*RemoveTagFromRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTagFromRecipe not implemented")
}
func (UnimplementedRecipeSvcServer) mustEmbedUnimplementedRecipeSvcServer() {}

// UnsafeRecipeSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipeSvcServer will
// result in compilation errors.
type UnsafeRecipeSvcServer interface {
	mustEmbedUnimplementedRecipeSvcServer()
}

func RegisterRecipeSvcServer(s grpc.ServiceRegistrar, srv RecipeSvcServer) {
	s.RegisterService(&RecipeSvc_ServiceDesc, srv)
}

func _RecipeSvc_ListRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).ListRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/ListRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).ListRecipe(ctx, req.(*ListRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_FindRecipesByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRecipesByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).FindRecipesByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/FindRecipesByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).FindRecipesByName(ctx, req.(*FindRecipesByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/GetRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).GetRecipe(ctx, req.(*GetRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_FindRecipesByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRecipesByTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).FindRecipesByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/FindRecipesByTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).FindRecipesByTags(ctx, req.(*FindRecipesByTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_FindRecipesByNameAndTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRecipesByNameAndTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).FindRecipesByNameAndTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/FindRecipesByNameAndTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).FindRecipesByNameAndTags(ctx, req.(*FindRecipesByNameAndTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/CreateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_UpdateRecipeTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecipeTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).UpdateRecipeTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/UpdateRecipeTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).UpdateRecipeTitle(ctx, req.(*UpdateRecipeTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_UpdateRecipeDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecipeDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).UpdateRecipeDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/UpdateRecipeDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).UpdateRecipeDescription(ctx, req.(*UpdateRecipeDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/DeleteRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).DeleteRecipe(ctx, req.(*DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_RemoveResourceFromRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveResourceFromRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).RemoveResourceFromRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/RemoveResourceFromRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).RemoveResourceFromRecipe(ctx, req.(*RemoveResourceFromRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_AddRecipeResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRecipeResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).AddRecipeResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/AddRecipeResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).AddRecipeResource(ctx, req.(*AddRecipeResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_AddTagToRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagToRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).AddTagToRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/AddTagToRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).AddTagToRecipe(ctx, req.(*AddTagToRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeSvc_RemoveTagFromRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTagFromRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeSvcServer).RemoveTagFromRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecipeSvc/RemoveTagFromRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeSvcServer).RemoveTagFromRecipe(ctx, req.(*RemoveTagFromRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecipeSvc_ServiceDesc is the grpc.ServiceDesc for RecipeSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecipeSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RecipeSvc",
	HandlerType: (*RecipeSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRecipe",
			Handler:    _RecipeSvc_ListRecipe_Handler,
		},
		{
			MethodName: "FindRecipesByName",
			Handler:    _RecipeSvc_FindRecipesByName_Handler,
		},
		{
			MethodName: "GetRecipe",
			Handler:    _RecipeSvc_GetRecipe_Handler,
		},
		{
			MethodName: "FindRecipesByTags",
			Handler:    _RecipeSvc_FindRecipesByTags_Handler,
		},
		{
			MethodName: "FindRecipesByNameAndTags",
			Handler:    _RecipeSvc_FindRecipesByNameAndTags_Handler,
		},
		{
			MethodName: "CreateRecipe",
			Handler:    _RecipeSvc_CreateRecipe_Handler,
		},
		{
			MethodName: "UpdateRecipeTitle",
			Handler:    _RecipeSvc_UpdateRecipeTitle_Handler,
		},
		{
			MethodName: "UpdateRecipeDescription",
			Handler:    _RecipeSvc_UpdateRecipeDescription_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _RecipeSvc_DeleteRecipe_Handler,
		},
		{
			MethodName: "RemoveResourceFromRecipe",
			Handler:    _RecipeSvc_RemoveResourceFromRecipe_Handler,
		},
		{
			MethodName: "AddRecipeResource",
			Handler:    _RecipeSvc_AddRecipeResource_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _RecipeSvc_CreateTag_Handler,
		},
		{
			MethodName: "AddTagToRecipe",
			Handler:    _RecipeSvc_AddTagToRecipe_Handler,
		},
		{
			MethodName: "RemoveTagFromRecipe",
			Handler:    _RecipeSvc_RemoveTagFromRecipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports/protos/srv.proto",
}
