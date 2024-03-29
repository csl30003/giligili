package user

import (
	"context"
	"giligili/app/user/rpc/user"
	"giligili/common/ctxData"

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

// GetFolloweeList 获取我关注的用户列表，目前只获取 id
func (l *GetFolloweeListLogic) GetFolloweeList(req *types.GetFolloweeListReq) (resp *types.GetFolloweeListResp, err error) {
	// 用于忽略警告
	_ = req

	userId := ctxData.GetUserIdFromCtx(l.ctx)

	getFolloweeListResp, err := l.svcCtx.UserRpcClient.GetFolloweeList(l.ctx, &user.GetFolloweeListReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetFolloweeListResp{
		FolloweeList: getFolloweeListResp.FolloweeId,
	}, nil
}
