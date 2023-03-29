package logic

import (
	"context"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UndislikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUndislikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UndislikeVideoLogic {
	return &UndislikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UndislikeVideoLogic) UndislikeVideo(in *pb.UndislikeVideoReq) (*pb.UndislikeVideoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UndislikeVideoResp{}, nil
}
