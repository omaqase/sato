// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: catalogue.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CategoryService_CreateCategory_FullMethodName         = "/api.v1.catalogue.CategoryService/CreateCategory"
	CategoryService_UpdateCategory_FullMethodName         = "/api.v1.catalogue.CategoryService/UpdateCategory"
	CategoryService_GetCategoryByID_FullMethodName        = "/api.v1.catalogue.CategoryService/GetCategoryByID"
	CategoryService_GetCategoryBySlug_FullMethodName      = "/api.v1.catalogue.CategoryService/GetCategoryBySlug"
	CategoryService_ListCategory_FullMethodName           = "/api.v1.catalogue.CategoryService/ListCategory"
	CategoryService_DeleteCategoryByID_FullMethodName     = "/api.v1.catalogue.CategoryService/DeleteCategoryByID"
	CategoryService_DeleteCategoryBySlug_FullMethodName   = "/api.v1.catalogue.CategoryService/DeleteCategoryBySlug"
	CategoryService_GetCategoryAncestors_FullMethodName   = "/api.v1.catalogue.CategoryService/GetCategoryAncestors"
	CategoryService_GetCategoryDescendants_FullMethodName = "/api.v1.catalogue.CategoryService/GetCategoryDescendants"
	CategoryService_AddCategoryHierarchy_FullMethodName   = "/api.v1.catalogue.CategoryService/AddCategoryHierarchy"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
	GetCategoryByID(ctx context.Context, in *GetCategoryByIDRequest, opts ...grpc.CallOption) (*GetCategoryByIDResponse, error)
	GetCategoryBySlug(ctx context.Context, in *GetCategoryBySlugRequest, opts ...grpc.CallOption) (*GetCategoryBySlugResponse, error)
	ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryResponse, error)
	DeleteCategoryByID(ctx context.Context, in *DeleteCategoryByIDRequest, opts ...grpc.CallOption) (*DeleteCategoryByIDResponse, error)
	DeleteCategoryBySlug(ctx context.Context, in *DeleteCategoryBySlugRequest, opts ...grpc.CallOption) (*DeleteCategoryBySlugResponse, error)
	GetCategoryAncestors(ctx context.Context, in *GetCategoryAncestorsRequest, opts ...grpc.CallOption) (*GetCategoryAncestorsResponse, error)
	GetCategoryDescendants(ctx context.Context, in *GetCategoryDescendantsRequest, opts ...grpc.CallOption) (*GetCategoryDescendantsResponse, error)
	AddCategoryHierarchy(ctx context.Context, in *AddCategoryHierarchyRequest, opts ...grpc.CallOption) (*AddCategoryHierarchyResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_UpdateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryByID(ctx context.Context, in *GetCategoryByIDRequest, opts ...grpc.CallOption) (*GetCategoryByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryByIDResponse)
	err := c.cc.Invoke(ctx, CategoryService_GetCategoryByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryBySlug(ctx context.Context, in *GetCategoryBySlugRequest, opts ...grpc.CallOption) (*GetCategoryBySlugResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryBySlugResponse)
	err := c.cc.Invoke(ctx, CategoryService_GetCategoryBySlug_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_ListCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategoryByID(ctx context.Context, in *DeleteCategoryByIDRequest, opts ...grpc.CallOption) (*DeleteCategoryByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCategoryByIDResponse)
	err := c.cc.Invoke(ctx, CategoryService_DeleteCategoryByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategoryBySlug(ctx context.Context, in *DeleteCategoryBySlugRequest, opts ...grpc.CallOption) (*DeleteCategoryBySlugResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCategoryBySlugResponse)
	err := c.cc.Invoke(ctx, CategoryService_DeleteCategoryBySlug_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryAncestors(ctx context.Context, in *GetCategoryAncestorsRequest, opts ...grpc.CallOption) (*GetCategoryAncestorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryAncestorsResponse)
	err := c.cc.Invoke(ctx, CategoryService_GetCategoryAncestors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryDescendants(ctx context.Context, in *GetCategoryDescendantsRequest, opts ...grpc.CallOption) (*GetCategoryDescendantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryDescendantsResponse)
	err := c.cc.Invoke(ctx, CategoryService_GetCategoryDescendants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) AddCategoryHierarchy(ctx context.Context, in *AddCategoryHierarchyRequest, opts ...grpc.CallOption) (*AddCategoryHierarchyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCategoryHierarchyResponse)
	err := c.cc.Invoke(ctx, CategoryService_AddCategoryHierarchy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations should embed UnimplementedCategoryServiceServer
// for forward compatibility.
type CategoryServiceServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error)
	GetCategoryByID(context.Context, *GetCategoryByIDRequest) (*GetCategoryByIDResponse, error)
	GetCategoryBySlug(context.Context, *GetCategoryBySlugRequest) (*GetCategoryBySlugResponse, error)
	ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryResponse, error)
	DeleteCategoryByID(context.Context, *DeleteCategoryByIDRequest) (*DeleteCategoryByIDResponse, error)
	DeleteCategoryBySlug(context.Context, *DeleteCategoryBySlugRequest) (*DeleteCategoryBySlugResponse, error)
	GetCategoryAncestors(context.Context, *GetCategoryAncestorsRequest) (*GetCategoryAncestorsResponse, error)
	GetCategoryDescendants(context.Context, *GetCategoryDescendantsRequest) (*GetCategoryDescendantsResponse, error)
	AddCategoryHierarchy(context.Context, *AddCategoryHierarchyRequest) (*AddCategoryHierarchyResponse, error)
}

// UnimplementedCategoryServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoryServiceServer struct{}

func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryByID(context.Context, *GetCategoryByIDRequest) (*GetCategoryByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryByID not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryBySlug(context.Context, *GetCategoryBySlugRequest) (*GetCategoryBySlugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryBySlug not implemented")
}
func (UnimplementedCategoryServiceServer) ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategoryByID(context.Context, *DeleteCategoryByIDRequest) (*DeleteCategoryByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryByID not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategoryBySlug(context.Context, *DeleteCategoryBySlugRequest) (*DeleteCategoryBySlugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryBySlug not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryAncestors(context.Context, *GetCategoryAncestorsRequest) (*GetCategoryAncestorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryAncestors not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryDescendants(context.Context, *GetCategoryDescendantsRequest) (*GetCategoryDescendantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryDescendants not implemented")
}
func (UnimplementedCategoryServiceServer) AddCategoryHierarchy(context.Context, *AddCategoryHierarchyRequest) (*AddCategoryHierarchyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategoryHierarchy not implemented")
}
func (UnimplementedCategoryServiceServer) testEmbeddedByValue() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedCategoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_GetCategoryByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryByID(ctx, req.(*GetCategoryByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryBySlugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_GetCategoryBySlug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryBySlug(ctx, req.(*GetCategoryBySlugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_ListCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).ListCategory(ctx, req.(*ListCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategoryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategoryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_DeleteCategoryByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategoryByID(ctx, req.(*DeleteCategoryByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategoryBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryBySlugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategoryBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_DeleteCategoryBySlug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategoryBySlug(ctx, req.(*DeleteCategoryBySlugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryAncestors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryAncestorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryAncestors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_GetCategoryAncestors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryAncestors(ctx, req.(*GetCategoryAncestorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryDescendants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryDescendantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryDescendants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_GetCategoryDescendants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryDescendants(ctx, req.(*GetCategoryDescendantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_AddCategoryHierarchy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryHierarchyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).AddCategoryHierarchy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_AddCategoryHierarchy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).AddCategoryHierarchy(ctx, req.(*AddCategoryHierarchyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.catalogue.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryService_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CategoryService_UpdateCategory_Handler,
		},
		{
			MethodName: "GetCategoryByID",
			Handler:    _CategoryService_GetCategoryByID_Handler,
		},
		{
			MethodName: "GetCategoryBySlug",
			Handler:    _CategoryService_GetCategoryBySlug_Handler,
		},
		{
			MethodName: "ListCategory",
			Handler:    _CategoryService_ListCategory_Handler,
		},
		{
			MethodName: "DeleteCategoryByID",
			Handler:    _CategoryService_DeleteCategoryByID_Handler,
		},
		{
			MethodName: "DeleteCategoryBySlug",
			Handler:    _CategoryService_DeleteCategoryBySlug_Handler,
		},
		{
			MethodName: "GetCategoryAncestors",
			Handler:    _CategoryService_GetCategoryAncestors_Handler,
		},
		{
			MethodName: "GetCategoryDescendants",
			Handler:    _CategoryService_GetCategoryDescendants_Handler,
		},
		{
			MethodName: "AddCategoryHierarchy",
			Handler:    _CategoryService_AddCategoryHierarchy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalogue.proto",
}

const (
	ProductService_CreateProduct_FullMethodName       = "/api.v1.catalogue.ProductService/CreateProduct"
	ProductService_UpdateProduct_FullMethodName       = "/api.v1.catalogue.ProductService/UpdateProduct"
	ProductService_GetProductByID_FullMethodName      = "/api.v1.catalogue.ProductService/GetProductByID"
	ProductService_ListProduct_FullMethodName         = "/api.v1.catalogue.ProductService/ListProduct"
	ProductService_DeleteProduct_FullMethodName       = "/api.v1.catalogue.ProductService/DeleteProduct"
	ProductService_AddToFavorites_FullMethodName      = "/api.v1.catalogue.ProductService/AddToFavorites"
	ProductService_RemoveFromFavorites_FullMethodName = "/api.v1.catalogue.ProductService/RemoveFromFavorites"
	ProductService_GetUserFavorites_FullMethodName    = "/api.v1.catalogue.ProductService/GetUserFavorites"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDResponse, error)
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	AddToFavorites(ctx context.Context, in *AddToFavoritesRequest, opts ...grpc.CallOption) (*AddToFavoritesResponse, error)
	RemoveFromFavorites(ctx context.Context, in *RemoveFromFavoritesRequest, opts ...grpc.CallOption) (*RemoveFromFavoritesResponse, error)
	GetUserFavorites(ctx context.Context, in *GetUserFavoritesRequest, opts ...grpc.CallOption) (*GetUserFavoritesResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, ProductService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, ProductService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductByIDResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProductByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, ProductService_ListProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, ProductService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddToFavorites(ctx context.Context, in *AddToFavoritesRequest, opts ...grpc.CallOption) (*AddToFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddToFavoritesResponse)
	err := c.cc.Invoke(ctx, ProductService_AddToFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) RemoveFromFavorites(ctx context.Context, in *RemoveFromFavoritesRequest, opts ...grpc.CallOption) (*RemoveFromFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveFromFavoritesResponse)
	err := c.cc.Invoke(ctx, ProductService_RemoveFromFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetUserFavorites(ctx context.Context, in *GetUserFavoritesRequest, opts ...grpc.CallOption) (*GetUserFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserFavoritesResponse)
	err := c.cc.Invoke(ctx, ProductService_GetUserFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility.
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	GetProductByID(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	AddToFavorites(context.Context, *AddToFavoritesRequest) (*AddToFavoritesResponse, error)
	RemoveFromFavorites(context.Context, *RemoveFromFavoritesRequest) (*RemoveFromFavoritesResponse, error)
	GetUserFavorites(context.Context, *GetUserFavoritesRequest) (*GetUserFavoritesResponse, error)
}

// UnimplementedProductServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProductByID(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedProductServiceServer) ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) AddToFavorites(context.Context, *AddToFavoritesRequest) (*AddToFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToFavorites not implemented")
}
func (UnimplementedProductServiceServer) RemoveFromFavorites(context.Context, *RemoveFromFavoritesRequest) (*RemoveFromFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromFavorites not implemented")
}
func (UnimplementedProductServiceServer) GetUserFavorites(context.Context, *GetUserFavoritesRequest) (*GetUserFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavorites not implemented")
}
func (UnimplementedProductServiceServer) testEmbeddedByValue() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductByID(ctx, req.(*GetProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ListProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddToFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddToFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddToFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddToFavorites(ctx, req.(*AddToFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_RemoveFromFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).RemoveFromFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_RemoveFromFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).RemoveFromFavorites(ctx, req.(*RemoveFromFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetUserFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetUserFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetUserFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetUserFavorites(ctx, req.(*GetUserFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.catalogue.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _ProductService_GetProductByID_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _ProductService_ListProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "AddToFavorites",
			Handler:    _ProductService_AddToFavorites_Handler,
		},
		{
			MethodName: "RemoveFromFavorites",
			Handler:    _ProductService_RemoveFromFavorites_Handler,
		},
		{
			MethodName: "GetUserFavorites",
			Handler:    _ProductService_GetUserFavorites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalogue.proto",
}
