package user

import (
	"context"
	"giligili/app/user/rpc/user"
	"giligili/common/ctxData"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnfollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowUserLogic {
	return &UnfollowUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfollowUserLogic) UnfollowUser(req *types.UnfollowUserReq) (resp *types.UnfollowUserResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	_, err = l.svcCtx.UserRpcClient.UnfollowUser(l.ctx, &user.UnfollowUserReq{
		UserId:     userId,
		FolloweeId: req.FolloweeId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UnfollowUserResp{}, nil
}
