// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: access.proto

package access_v1

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

// AntiBruteforceClient is the client API for AntiBruteforce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AntiBruteforceClient interface {
	TryAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	ResetBucket(ctx context.Context, in *ResetBucketRequest, opts ...grpc.CallOption) (*ResetBucketResponse, error)
	AddToBlacklist(ctx context.Context, in *AddToBlacklistRequest, opts ...grpc.CallOption) (*AddToBlacklistResponse, error)
	RemoveFromBlacklist(ctx context.Context, in *RemoveFromBlacklistRequest, opts ...grpc.CallOption) (*RemoveFromBlacklistResponse, error)
	AddToWhitelist(ctx context.Context, in *AddToWhitelistRequest, opts ...grpc.CallOption) (*AddToWhitelistResponse, error)
	RemoveFromWhitelist(ctx context.Context, in *RemoveFromWhitelistRequest, opts ...grpc.CallOption) (*RemoveFromWhitelistResponse, error)
}

type antiBruteforceClient struct {
	cc grpc.ClientConnInterface
}

func NewAntiBruteforceClient(cc grpc.ClientConnInterface) AntiBruteforceClient {
	return &antiBruteforceClient{cc}
}

func (c *antiBruteforceClient) TryAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/access.AntiBruteforce/TryAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) ResetBucket(ctx context.Context, in *ResetBucketRequest, opts ...grpc.CallOption) (*ResetBucketResponse, error) {
	out := new(ResetBucketResponse)
	err := c.cc.Invoke(ctx, "/access.AntiBruteforce/ResetBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) AddToBlacklist(ctx context.Context, in *AddToBlacklistRequest, opts ...grpc.CallOption) (*AddToBlacklistResponse, error) {
	out := new(AddToBlacklistResponse)
	err := c.cc.Invoke(ctx, "/access.AntiBruteforce/AddToBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) RemoveFromBlacklist(ctx context.Context, in *RemoveFromBlacklistRequest, opts ...grpc.CallOption) (*RemoveFromBlacklistResponse, error) {
	out := new(RemoveFromBlacklistResponse)
	err := c.cc.Invoke(ctx, "/access.AntiBruteforce/RemoveFromBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) AddToWhitelist(ctx context.Context, in *AddToWhitelistRequest, opts ...grpc.CallOption) (*AddToWhitelistResponse, error) {
	out := new(AddToWhitelistResponse)
	err := c.cc.Invoke(ctx, "/access.AntiBruteforce/AddToWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) RemoveFromWhitelist(ctx context.Context, in *RemoveFromWhitelistRequest, opts ...grpc.CallOption) (*RemoveFromWhitelistResponse, error) {
	out := new(RemoveFromWhitelistResponse)
	err := c.cc.Invoke(ctx, "/access.AntiBruteforce/RemoveFromWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AntiBruteforceServer is the server API for AntiBruteforce service.
// All implementations must embed UnimplementedAntiBruteforceServer
// for forward compatibility
type AntiBruteforceServer interface {
	TryAuth(context.Context, *AuthRequest) (*AuthResponse, error)
	ResetBucket(context.Context, *ResetBucketRequest) (*ResetBucketResponse, error)
	AddToBlacklist(context.Context, *AddToBlacklistRequest) (*AddToBlacklistResponse, error)
	RemoveFromBlacklist(context.Context, *RemoveFromBlacklistRequest) (*RemoveFromBlacklistResponse, error)
	AddToWhitelist(context.Context, *AddToWhitelistRequest) (*AddToWhitelistResponse, error)
	RemoveFromWhitelist(context.Context, *RemoveFromWhitelistRequest) (*RemoveFromWhitelistResponse, error)
	mustEmbedUnimplementedAntiBruteforceServer()
}

// UnimplementedAntiBruteforceServer must be embedded to have forward compatible implementations.
type UnimplementedAntiBruteforceServer struct {
}

func (UnimplementedAntiBruteforceServer) TryAuth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TryAuth not implemented")
}
func (UnimplementedAntiBruteforceServer) ResetBucket(context.Context, *ResetBucketRequest) (*ResetBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetBucket not implemented")
}
func (UnimplementedAntiBruteforceServer) AddToBlacklist(context.Context, *AddToBlacklistRequest) (*AddToBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBlacklist not implemented")
}
func (UnimplementedAntiBruteforceServer) RemoveFromBlacklist(context.Context, *RemoveFromBlacklistRequest) (*RemoveFromBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromBlacklist not implemented")
}
func (UnimplementedAntiBruteforceServer) AddToWhitelist(context.Context, *AddToWhitelistRequest) (*AddToWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWhitelist not implemented")
}
func (UnimplementedAntiBruteforceServer) RemoveFromWhitelist(context.Context, *RemoveFromWhitelistRequest) (*RemoveFromWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWhitelist not implemented")
}
func (UnimplementedAntiBruteforceServer) mustEmbedUnimplementedAntiBruteforceServer() {}

// UnsafeAntiBruteforceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AntiBruteforceServer will
// result in compilation errors.
type UnsafeAntiBruteforceServer interface {
	mustEmbedUnimplementedAntiBruteforceServer()
}

func RegisterAntiBruteforceServer(s grpc.ServiceRegistrar, srv AntiBruteforceServer) {
	s.RegisterService(&AntiBruteforce_ServiceDesc, srv)
}

func _AntiBruteforce_TryAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).TryAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.AntiBruteforce/TryAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).TryAuth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_ResetBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).ResetBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.AntiBruteforce/ResetBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).ResetBucket(ctx, req.(*ResetBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_AddToBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).AddToBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.AntiBruteforce/AddToBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).AddToBlacklist(ctx, req.(*AddToBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_RemoveFromBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).RemoveFromBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.AntiBruteforce/RemoveFromBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).RemoveFromBlacklist(ctx, req.(*RemoveFromBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_AddToWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToWhitelistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).AddToWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.AntiBruteforce/AddToWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).AddToWhitelist(ctx, req.(*AddToWhitelistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_RemoveFromWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromWhitelistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).RemoveFromWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.AntiBruteforce/RemoveFromWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).RemoveFromWhitelist(ctx, req.(*RemoveFromWhitelistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AntiBruteforce_ServiceDesc is the grpc.ServiceDesc for AntiBruteforce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AntiBruteforce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "access.AntiBruteforce",
	HandlerType: (*AntiBruteforceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TryAuth",
			Handler:    _AntiBruteforce_TryAuth_Handler,
		},
		{
			MethodName: "ResetBucket",
			Handler:    _AntiBruteforce_ResetBucket_Handler,
		},
		{
			MethodName: "AddToBlacklist",
			Handler:    _AntiBruteforce_AddToBlacklist_Handler,
		},
		{
			MethodName: "RemoveFromBlacklist",
			Handler:    _AntiBruteforce_RemoveFromBlacklist_Handler,
		},
		{
			MethodName: "AddToWhitelist",
			Handler:    _AntiBruteforce_AddToWhitelist_Handler,
		},
		{
			MethodName: "RemoveFromWhitelist",
			Handler:    _AntiBruteforce_RemoveFromWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access.proto",
}
