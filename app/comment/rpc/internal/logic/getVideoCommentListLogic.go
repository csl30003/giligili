package logic

import (
	"context"

	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentListLogic {
	return &GetVideoCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoCommentListLogic) GetVideoCommentList(in *pb.GetVideoCommentListReq) (*pb.GetVideoCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVideoCommentListResp{}, nil
}
