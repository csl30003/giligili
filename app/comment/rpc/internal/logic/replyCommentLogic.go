package logic

import (
	"context"
	"database/sql"
	"giligili/app/comment/model"
	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommentLogic {
	return &ReplyCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReplyComment 回复评论
func (l *ReplyCommentLogic) ReplyComment(in *pb.ReplyCommentReq) (*pb.ReplyCommentResp, error) {
	// 判断评论是否存在
	comment, err := l.svcCtx.CommentModel.FindOneById(l.ctx, in.CommentId)
	if err == model.ErrNotFound {
		return nil, status.Error(100, "回复的评论不存在")
	}
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	// 插入记录
	newComment := model.Comment{
		UserId:  in.UserId,
		VideoId: comment.VideoId,
		ReplyId: sql.NullInt64{Int64: in.CommentId, Valid: true},
		Content: in.Content,
	}
	result, err := l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
	if err != nil {
		return nil, status.Error(100, "评论存储数据库失败")
	}
	newComment.Id, err = result.LastInsertId()
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	return &pb.ReplyCommentResp{Success: true}, nil
}
