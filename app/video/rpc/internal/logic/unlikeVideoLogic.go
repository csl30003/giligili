package logic

import (
	"context"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnlikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeVideoLogic {
	return &UnlikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnlikeVideoLogic) UnlikeVideo(in *pb.UnlikeVideoReq) (*pb.UnlikeVideoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UnlikeVideoResp{}, nil
}
