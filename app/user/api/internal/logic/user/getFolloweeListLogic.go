package user

import (
	"context"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFolloweeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFolloweeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFolloweeListLogic {
	return &GetFolloweeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFolloweeListLogic) GetFolloweeList(req *types.GetFolloweeListReq) (resp *types.GetFolloweeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
