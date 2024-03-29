// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: comment.proto

package pb

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

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	GetVideoCommentList(ctx context.Context, in *GetVideoCommentListReq, opts ...grpc.CallOption) (*GetVideoCommentListResp, error)
	GetCommentReplyList(ctx context.Context, in *GetCommentReplyListReq, opts ...grpc.CallOption) (*GetCommentReplyListResp, error)
	CommentVideo(ctx context.Context, in *CommentVideoReq, opts ...grpc.CallOption) (*CommentVideoResp, error)
	ReplyComment(ctx context.Context, in *ReplyCommentReq, opts ...grpc.CallOption) (*ReplyCommentResp, error)
	DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) GetVideoCommentList(ctx context.Context, in *GetVideoCommentListReq, opts ...grpc.CallOption) (*GetVideoCommentListResp, error) {
	out := new(GetVideoCommentListResp)
	err := c.cc.Invoke(ctx, "/pb.Comment/GetVideoCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetCommentReplyList(ctx context.Context, in *GetCommentReplyListReq, opts ...grpc.CallOption) (*GetCommentReplyListResp, error) {
	out := new(GetCommentReplyListResp)
	err := c.cc.Invoke(ctx, "/pb.Comment/GetCommentReplyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) CommentVideo(ctx context.Context, in *CommentVideoReq, opts ...grpc.CallOption) (*CommentVideoResp, error) {
	out := new(CommentVideoResp)
	err := c.cc.Invoke(ctx, "/pb.Comment/CommentVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) ReplyComment(ctx context.Context, in *ReplyCommentReq, opts ...grpc.CallOption) (*ReplyCommentResp, error) {
	out := new(ReplyCommentResp)
	err := c.cc.Invoke(ctx, "/pb.Comment/ReplyComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error) {
	out := new(DeleteCommentResp)
	err := c.cc.Invoke(ctx, "/pb.Comment/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	GetVideoCommentList(context.Context, *GetVideoCommentListReq) (*GetVideoCommentListResp, error)
	GetCommentReplyList(context.Context, *GetCommentReplyListReq) (*GetCommentReplyListResp, error)
	CommentVideo(context.Context, *CommentVideoReq) (*CommentVideoResp, error)
	ReplyComment(context.Context, *ReplyCommentReq) (*ReplyCommentResp, error)
	DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentResp, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) GetVideoCommentList(context.Context, *GetVideoCommentListReq) (*GetVideoCommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoCommentList not implemented")
}
func (UnimplementedCommentServer) GetCommentReplyList(context.Context, *GetCommentReplyListReq) (*GetCommentReplyListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentReplyList not implemented")
}
func (UnimplementedCommentServer) CommentVideo(context.Context, *CommentVideoReq) (*CommentVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentVideo not implemented")
}
func (UnimplementedCommentServer) ReplyComment(context.Context, *ReplyCommentReq) (*ReplyCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyComment not implemented")
}
func (UnimplementedCommentServer) DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_GetVideoCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetVideoCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Comment/GetVideoCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetVideoCommentList(ctx, req.(*GetVideoCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetCommentReplyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentReplyListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetCommentReplyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Comment/GetCommentReplyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetCommentReplyList(ctx, req.(*GetCommentReplyListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_CommentVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CommentVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Comment/CommentVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CommentVideo(ctx, req.(*CommentVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_ReplyComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).ReplyComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Comment/ReplyComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).ReplyComment(ctx, req.(*ReplyCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Comment/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DeleteComment(ctx, req.(*DeleteCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideoCommentList",
			Handler:    _Comment_GetVideoCommentList_Handler,
		},
		{
			MethodName: "GetCommentReplyList",
			Handler:    _Comment_GetCommentReplyList_Handler,
		},
		{
			MethodName: "CommentVideo",
			Handler:    _Comment_CommentVideo_Handler,
		},
		{
			MethodName: "ReplyComment",
			Handler:    _Comment_ReplyComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Comment_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
