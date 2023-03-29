package logic

import (
	"context"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DislikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDislikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DislikeVideoLogic {
	return &DislikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DislikeVideoLogic) DislikeVideo(in *pb.DislikeVideoReq) (*pb.DislikeVideoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DislikeVideoResp{}, nil
}
