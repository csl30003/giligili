package logic

import (
	"context"
	"giligili/app/comment/model"
	"google.golang.org/grpc/status"
	"time"

	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteComment 删除评论
func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentReq) (*pb.DeleteCommentResp, error) {
	comment, err := l.svcCtx.CommentModel.FindOneById(l.ctx, in.CommentId)
	if err == model.ErrNotFound {
		return nil, status.Error(100, "评论不存在")
	}
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	if comment.UserId != in.UserId {
		return nil, status.Error(100, "你没有权限删除此评论记录")
	}

	// 删除
	comment.DeleteTime.Time = time.Now()
	comment.DeleteTime.Valid = true
	err = l.svcCtx.CommentModel.Update(l.ctx, comment)
	if err != nil {
		return nil, status.Error(100, "删除评论失败")
	}

	return &pb.DeleteCommentResp{Success: true}, nil
}
