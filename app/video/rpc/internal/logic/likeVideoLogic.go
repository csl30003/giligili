package logic

import (
	"context"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeVideoLogic) LikeVideo(in *pb.LikeVideoReq) (*pb.LikeVideoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LikeVideoResp{}, nil
}
