// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user.proto

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
	UserService_SendOTP_FullMethodName                       = "/api.v1.user.UserService/SendOTP"
	UserService_VerifyOTP_FullMethodName                     = "/api.v1.user.UserService/VerifyOTP"
	UserService_LogOut_FullMethodName                        = "/api.v1.user.UserService/LogOut"
	UserService_RefreshToken_FullMethodName                  = "/api.v1.user.UserService/RefreshToken"
	UserService_Update_FullMethodName                        = "/api.v1.user.UserService/Update"
	UserService_GetSubscribedToPromotionUsers_FullMethodName = "/api.v1.user.UserService/GetSubscribedToPromotionUsers"
	UserService_ParseJWT_FullMethodName                      = "/api.v1.user.UserService/ParseJWT"
	UserService_MakeSeller_FullMethodName                    = "/api.v1.user.UserService/MakeSeller"
	UserService_AddToFavorites_FullMethodName                = "/api.v1.user.UserService/AddToFavorites"
	UserService_RemoveFromFavorites_FullMethodName           = "/api.v1.user.UserService/RemoveFromFavorites"
	UserService_GetFavorites_FullMethodName                  = "/api.v1.user.UserService/GetFavorites"
	UserService_GetCart_FullMethodName                       = "/api.v1.user.UserService/GetCart"
	UserService_AddToCart_FullMethodName                     = "/api.v1.user.UserService/AddToCart"
	UserService_UpdateCartItem_FullMethodName                = "/api.v1.user.UserService/UpdateCartItem"
	UserService_RemoveFromCart_FullMethodName                = "/api.v1.user.UserService/RemoveFromCart"
	UserService_ClearCart_FullMethodName                     = "/api.v1.user.UserService/ClearCart"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*SendOTPResponse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
	LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	GetSubscribedToPromotionUsers(ctx context.Context, in *GetSubscribedToPromotionUsersRequest, opts ...grpc.CallOption) (*GetSubscribedToPromotionUsersResponse, error)
	ParseJWT(ctx context.Context, in *ParseJWTRequest, opts ...grpc.CallOption) (*ParseJWTResponse, error)
	MakeSeller(ctx context.Context, in *MakeSellerRequest, opts ...grpc.CallOption) (*MakeSellerResponse, error)
	AddToFavorites(ctx context.Context, in *AddToFavoritesRequest, opts ...grpc.CallOption) (*AddToFavoritesResponse, error)
	RemoveFromFavorites(ctx context.Context, in *RemoveFromFavoritesRequest, opts ...grpc.CallOption) (*RemoveFromFavoritesResponse, error)
	GetFavorites(ctx context.Context, in *GetFavoritesRequest, opts ...grpc.CallOption) (*GetFavoritesResponse, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
	AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error)
	UpdateCartItem(ctx context.Context, in *UpdateCartItemRequest, opts ...grpc.CallOption) (*UpdateCartItemResponse, error)
	RemoveFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*RemoveFromCartResponse, error)
	ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*SendOTPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendOTPResponse)
	err := c.cc.Invoke(ctx, UserService_SendOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, UserService_VerifyOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogOutResponse)
	err := c.cc.Invoke(ctx, UserService_LogOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, UserService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSubscribedToPromotionUsers(ctx context.Context, in *GetSubscribedToPromotionUsersRequest, opts ...grpc.CallOption) (*GetSubscribedToPromotionUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubscribedToPromotionUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetSubscribedToPromotionUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ParseJWT(ctx context.Context, in *ParseJWTRequest, opts ...grpc.CallOption) (*ParseJWTResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParseJWTResponse)
	err := c.cc.Invoke(ctx, UserService_ParseJWT_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) MakeSeller(ctx context.Context, in *MakeSellerRequest, opts ...grpc.CallOption) (*MakeSellerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MakeSellerResponse)
	err := c.cc.Invoke(ctx, UserService_MakeSeller_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddToFavorites(ctx context.Context, in *AddToFavoritesRequest, opts ...grpc.CallOption) (*AddToFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddToFavoritesResponse)
	err := c.cc.Invoke(ctx, UserService_AddToFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveFromFavorites(ctx context.Context, in *RemoveFromFavoritesRequest, opts ...grpc.CallOption) (*RemoveFromFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveFromFavoritesResponse)
	err := c.cc.Invoke(ctx, UserService_RemoveFromFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFavorites(ctx context.Context, in *GetFavoritesRequest, opts ...grpc.CallOption) (*GetFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFavoritesResponse)
	err := c.cc.Invoke(ctx, UserService_GetFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCartResponse)
	err := c.cc.Invoke(ctx, UserService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddToCartResponse)
	err := c.cc.Invoke(ctx, UserService_AddToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCartItem(ctx context.Context, in *UpdateCartItemRequest, opts ...grpc.CallOption) (*UpdateCartItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCartItemResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateCartItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*RemoveFromCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveFromCartResponse)
	err := c.cc.Invoke(ctx, UserService_RemoveFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearCartResponse)
	err := c.cc.Invoke(ctx, UserService_ClearCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	SendOTP(context.Context, *SendOTPRequest) (*SendOTPResponse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	GetSubscribedToPromotionUsers(context.Context, *GetSubscribedToPromotionUsersRequest) (*GetSubscribedToPromotionUsersResponse, error)
	ParseJWT(context.Context, *ParseJWTRequest) (*ParseJWTResponse, error)
	MakeSeller(context.Context, *MakeSellerRequest) (*MakeSellerResponse, error)
	AddToFavorites(context.Context, *AddToFavoritesRequest) (*AddToFavoritesResponse, error)
	RemoveFromFavorites(context.Context, *RemoveFromFavoritesRequest) (*RemoveFromFavoritesResponse, error)
	GetFavorites(context.Context, *GetFavoritesRequest) (*GetFavoritesResponse, error)
	GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error)
	AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error)
	UpdateCartItem(context.Context, *UpdateCartItemRequest) (*UpdateCartItemResponse, error)
	RemoveFromCart(context.Context, *RemoveFromCartRequest) (*RemoveFromCartResponse, error)
	ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) SendOTP(context.Context, *SendOTPRequest) (*SendOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTP not implemented")
}
func (UnimplementedUserServiceServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedUserServiceServer) LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedUserServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) GetSubscribedToPromotionUsers(context.Context, *GetSubscribedToPromotionUsersRequest) (*GetSubscribedToPromotionUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribedToPromotionUsers not implemented")
}
func (UnimplementedUserServiceServer) ParseJWT(context.Context, *ParseJWTRequest) (*ParseJWTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseJWT not implemented")
}
func (UnimplementedUserServiceServer) MakeSeller(context.Context, *MakeSellerRequest) (*MakeSellerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeSeller not implemented")
}
func (UnimplementedUserServiceServer) AddToFavorites(context.Context, *AddToFavoritesRequest) (*AddToFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToFavorites not implemented")
}
func (UnimplementedUserServiceServer) RemoveFromFavorites(context.Context, *RemoveFromFavoritesRequest) (*RemoveFromFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromFavorites not implemented")
}
func (UnimplementedUserServiceServer) GetFavorites(context.Context, *GetFavoritesRequest) (*GetFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorites not implemented")
}
func (UnimplementedUserServiceServer) GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedUserServiceServer) AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedUserServiceServer) UpdateCartItem(context.Context, *UpdateCartItemRequest) (*UpdateCartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItem not implemented")
}
func (UnimplementedUserServiceServer) RemoveFromCart(context.Context, *RemoveFromCartRequest) (*RemoveFromCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedUserServiceServer) ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedUserServiceServer) testEmbeddedByValue() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_SendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SendOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendOTP(ctx, req.(*SendOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LogOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogOut(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSubscribedToPromotionUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscribedToPromotionUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSubscribedToPromotionUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetSubscribedToPromotionUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSubscribedToPromotionUsers(ctx, req.(*GetSubscribedToPromotionUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ParseJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ParseJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ParseJWT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ParseJWT(ctx, req.(*ParseJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_MakeSeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeSellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).MakeSeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_MakeSeller_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).MakeSeller(ctx, req.(*MakeSellerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddToFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddToFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddToFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddToFavorites(ctx, req.(*AddToFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveFromFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveFromFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveFromFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveFromFavorites(ctx, req.(*RemoveFromFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFavorites(ctx, req.(*GetFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCartItem(ctx, req.(*UpdateCartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveFromCart(ctx, req.(*RemoveFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ClearCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClearCart(ctx, req.(*ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOTP",
			Handler:    _UserService_SendOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _UserService_VerifyOTP_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _UserService_LogOut_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _UserService_RefreshToken_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "GetSubscribedToPromotionUsers",
			Handler:    _UserService_GetSubscribedToPromotionUsers_Handler,
		},
		{
			MethodName: "ParseJWT",
			Handler:    _UserService_ParseJWT_Handler,
		},
		{
			MethodName: "MakeSeller",
			Handler:    _UserService_MakeSeller_Handler,
		},
		{
			MethodName: "AddToFavorites",
			Handler:    _UserService_AddToFavorites_Handler,
		},
		{
			MethodName: "RemoveFromFavorites",
			Handler:    _UserService_RemoveFromFavorites_Handler,
		},
		{
			MethodName: "GetFavorites",
			Handler:    _UserService_GetFavorites_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _UserService_GetCart_Handler,
		},
		{
			MethodName: "AddToCart",
			Handler:    _UserService_AddToCart_Handler,
		},
		{
			MethodName: "UpdateCartItem",
			Handler:    _UserService_UpdateCartItem_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _UserService_RemoveFromCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _UserService_ClearCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
