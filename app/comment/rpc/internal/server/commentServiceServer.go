// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"giligili/app/comment/rpc/internal/logic"
	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"
)

type CommentServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCommentServiceServer
}

func NewCommentServiceServer(svcCtx *svc.ServiceContext) *CommentServiceServer {
	return &CommentServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServiceServer) GetVideoCommentList(ctx context.Context, in *pb.GetVideoCommentListReq) (*pb.GetVideoCommentListResp, error) {
	l := logic.NewGetVideoCommentListLogic(ctx, s.svcCtx)
	return l.GetVideoCommentList(in)
}

func (s *CommentServiceServer) GetCommentReplyList(ctx context.Context, in *pb.GetCommentReplyListReq) (*pb.GetCommentReplyListResp, error) {
	l := logic.NewGetCommentReplyListLogic(ctx, s.svcCtx)
	return l.GetCommentReplyList(in)
}

func (s *CommentServiceServer) CommentVideo(ctx context.Context, in *pb.CommentVideoReq) (*pb.CommentVideoResp, error) {
	l := logic.NewCommentVideoLogic(ctx, s.svcCtx)
	return l.CommentVideo(in)
}

func (s *CommentServiceServer) ReplyComment(ctx context.Context, in *pb.ReplyCommentReq) (*pb.ReplyCommentResp, error) {
	l := logic.NewReplyCommentLogic(ctx, s.svcCtx)
	return l.ReplyComment(in)
}

func (s *CommentServiceServer) DeleteComment(ctx context.Context, in *pb.DeleteCommentReq) (*pb.DeleteCommentResp, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}