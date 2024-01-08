// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: pb/svc/firebase/firebase.proto

package firebase

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
	Firebase_GetUser_FullMethodName            = "/pb.svc.firebase.Firebase/GetUser"
	Firebase_CreateCustomToken_FullMethodName  = "/pb.svc.firebase.Firebase/CreateCustomToken"
	Firebase_VerifyIdToken_FullMethodName      = "/pb.svc.firebase.Firebase/VerifyIdToken"
	Firebase_GetUserIdByIdToken_FullMethodName = "/pb.svc.firebase.Firebase/GetUserIdByIdToken"
)

// FirebaseClient is the client API for Firebase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirebaseClient interface {
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRes, error)
	CreateCustomToken(ctx context.Context, in *CreateCustomTokenReq, opts ...grpc.CallOption) (*CreateCustomTokenRes, error)
	VerifyIdToken(ctx context.Context, in *VerifyIdTokenReq, opts ...grpc.CallOption) (*VerifyIdTokenRes, error)
	GetUserIdByIdToken(ctx context.Context, in *GetUserIdByIdTokenReq, opts ...grpc.CallOption) (*GetUserIdByIdTokenRes, error)
}

type firebaseClient struct {
	cc grpc.ClientConnInterface
}

func NewFirebaseClient(cc grpc.ClientConnInterface) FirebaseClient {
	return &firebaseClient{cc}
}

func (c *firebaseClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRes, error) {
	out := new(GetUserRes)
	err := c.cc.Invoke(ctx, Firebase_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firebaseClient) CreateCustomToken(ctx context.Context, in *CreateCustomTokenReq, opts ...grpc.CallOption) (*CreateCustomTokenRes, error) {
	out := new(CreateCustomTokenRes)
	err := c.cc.Invoke(ctx, Firebase_CreateCustomToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firebaseClient) VerifyIdToken(ctx context.Context, in *VerifyIdTokenReq, opts ...grpc.CallOption) (*VerifyIdTokenRes, error) {
	out := new(VerifyIdTokenRes)
	err := c.cc.Invoke(ctx, Firebase_VerifyIdToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firebaseClient) GetUserIdByIdToken(ctx context.Context, in *GetUserIdByIdTokenReq, opts ...grpc.CallOption) (*GetUserIdByIdTokenRes, error) {
	out := new(GetUserIdByIdTokenRes)
	err := c.cc.Invoke(ctx, Firebase_GetUserIdByIdToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirebaseServer is the server API for Firebase service.
// All implementations must embed UnimplementedFirebaseServer
// for forward compatibility
type FirebaseServer interface {
	GetUser(context.Context, *GetUserReq) (*GetUserRes, error)
	CreateCustomToken(context.Context, *CreateCustomTokenReq) (*CreateCustomTokenRes, error)
	VerifyIdToken(context.Context, *VerifyIdTokenReq) (*VerifyIdTokenRes, error)
	GetUserIdByIdToken(context.Context, *GetUserIdByIdTokenReq) (*GetUserIdByIdTokenRes, error)
	mustEmbedUnimplementedFirebaseServer()
}

// UnimplementedFirebaseServer must be embedded to have forward compatible implementations.
type UnimplementedFirebaseServer struct {
}

func (UnimplementedFirebaseServer) GetUser(context.Context, *GetUserReq) (*GetUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedFirebaseServer) CreateCustomToken(context.Context, *CreateCustomTokenReq) (*CreateCustomTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomToken not implemented")
}
func (UnimplementedFirebaseServer) VerifyIdToken(context.Context, *VerifyIdTokenReq) (*VerifyIdTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyIdToken not implemented")
}
func (UnimplementedFirebaseServer) GetUserIdByIdToken(context.Context, *GetUserIdByIdTokenReq) (*GetUserIdByIdTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdByIdToken not implemented")
}
func (UnimplementedFirebaseServer) mustEmbedUnimplementedFirebaseServer() {}

// UnsafeFirebaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirebaseServer will
// result in compilation errors.
type UnsafeFirebaseServer interface {
	mustEmbedUnimplementedFirebaseServer()
}

func RegisterFirebaseServer(s grpc.ServiceRegistrar, srv FirebaseServer) {
	s.RegisterService(&Firebase_ServiceDesc, srv)
}

func _Firebase_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Firebase_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firebase_CreateCustomToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseServer).CreateCustomToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Firebase_CreateCustomToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseServer).CreateCustomToken(ctx, req.(*CreateCustomTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firebase_VerifyIdToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyIdTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseServer).VerifyIdToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Firebase_VerifyIdToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseServer).VerifyIdToken(ctx, req.(*VerifyIdTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firebase_GetUserIdByIdToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdByIdTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseServer).GetUserIdByIdToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Firebase_GetUserIdByIdToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseServer).GetUserIdByIdToken(ctx, req.(*GetUserIdByIdTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Firebase_ServiceDesc is the grpc.ServiceDesc for Firebase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Firebase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.svc.firebase.Firebase",
	HandlerType: (*FirebaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Firebase_GetUser_Handler,
		},
		{
			MethodName: "CreateCustomToken",
			Handler:    _Firebase_CreateCustomToken_Handler,
		},
		{
			MethodName: "VerifyIdToken",
			Handler:    _Firebase_VerifyIdToken_Handler,
		},
		{
			MethodName: "GetUserIdByIdToken",
			Handler:    _Firebase_GetUserIdByIdToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/svc/firebase/firebase.proto",
}
