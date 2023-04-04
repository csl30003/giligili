package logic

import (
	"context"

	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"

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

func (l *ReplyCommentLogic) ReplyComment(in *pb.ReplyCommentReq) (*pb.ReplyCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ReplyCommentResp{}, nil
}
