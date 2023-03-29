package logic

import (
	"context"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoDetailLogic {
	return &GetVideoDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoDetailLogic) GetVideoDetail(in *pb.GetVideoDetailReq) (*pb.GetVideoDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVideoDetailResp{}, nil
}
