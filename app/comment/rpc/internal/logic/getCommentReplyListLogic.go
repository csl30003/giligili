package logic

import (
	"context"

	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyListLogic {
	return &GetCommentReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentReplyListLogic) GetCommentReplyList(in *pb.GetCommentReplyListReq) (*pb.GetCommentReplyListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCommentReplyListResp{}, nil
}
