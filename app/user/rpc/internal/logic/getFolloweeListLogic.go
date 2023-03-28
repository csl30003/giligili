package logic

import (
	"context"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFolloweeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFolloweeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFolloweeListLogic {
	return &GetFolloweeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFolloweeListLogic) GetFolloweeList(in *pb.GetFolloweeListReq) (*pb.GetFolloweeListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFolloweeListResp{}, nil
}
