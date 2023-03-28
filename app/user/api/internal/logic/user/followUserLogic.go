package user

import (
	"context"
	"giligili/app/user/rpc/user"
	"giligili/common/ctxData"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FollowUser 关注用户
func (l *FollowUserLogic) FollowUser(req *types.FollowUserReq) (resp *types.FollowUserResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	_, err = l.svcCtx.UserRpcClient.FollowUser(l.ctx, &user.FollowUserReq{
		UserId:     userId,
		FolloweeId: req.FolloweeId,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowUserResp{}, nil
}
