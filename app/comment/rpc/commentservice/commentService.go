// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentservice

import (
	"context"

	"giligili/app/comment/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommentReply            = pb.CommentReply
	CommentVideoReq         = pb.CommentVideoReq
	CommentVideoResp        = pb.CommentVideoResp
	DeleteCommentReq        = pb.DeleteCommentReq
	DeleteCommentResp       = pb.DeleteCommentResp
	GetCommentReplyListReq  = pb.GetCommentReplyListReq
	GetCommentReplyListResp = pb.GetCommentReplyListResp
	GetVideoCommentListReq  = pb.GetVideoCommentListReq
	GetVideoCommentListResp = pb.GetVideoCommentListResp
	ReplyCommentReq         = pb.ReplyCommentReq
	ReplyCommentResp        = pb.ReplyCommentResp
	VideoComment            = pb.VideoComment

	CommentService interface {
		GetVideoCommentList(ctx context.Context, in *GetVideoCommentListReq, opts ...grpc.CallOption) (*GetVideoCommentListResp, error)
		GetCommentReplyList(ctx context.Context, in *GetCommentReplyListReq, opts ...grpc.CallOption) (*GetCommentReplyListResp, error)
		CommentVideo(ctx context.Context, in *CommentVideoReq, opts ...grpc.CallOption) (*CommentVideoResp, error)
		ReplyComment(ctx context.Context, in *ReplyCommentReq, opts ...grpc.CallOption) (*ReplyCommentResp, error)
		DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error)
	}

	defaultCommentService struct {
		cli zrpc.Client
	}
)

func NewCommentService(cli zrpc.Client) CommentService {
	return &defaultCommentService{
		cli: cli,
	}
}

func (m *defaultCommentService) GetVideoCommentList(ctx context.Context, in *GetVideoCommentListReq, opts ...grpc.CallOption) (*GetVideoCommentListResp, error) {
	client := pb.NewCommentServiceClient(m.cli.Conn())
	return client.GetVideoCommentList(ctx, in, opts...)
}

func (m *defaultCommentService) GetCommentReplyList(ctx context.Context, in *GetCommentReplyListReq, opts ...grpc.CallOption) (*GetCommentReplyListResp, error) {
	client := pb.NewCommentServiceClient(m.cli.Conn())
	return client.GetCommentReplyList(ctx, in, opts...)
}

func (m *defaultCommentService) CommentVideo(ctx context.Context, in *CommentVideoReq, opts ...grpc.CallOption) (*CommentVideoResp, error) {
	client := pb.NewCommentServiceClient(m.cli.Conn())
	return client.CommentVideo(ctx, in, opts...)
}

func (m *defaultCommentService) ReplyComment(ctx context.Context, in *ReplyCommentReq, opts ...grpc.CallOption) (*ReplyCommentResp, error) {
	client := pb.NewCommentServiceClient(m.cli.Conn())
	return client.ReplyComment(ctx, in, opts...)
}

func (m *defaultCommentService) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error) {
	client := pb.NewCommentServiceClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}
