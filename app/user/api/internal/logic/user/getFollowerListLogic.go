package user

import (
	"context"
	"giligili/app/user/rpc/user"
	"giligili/common/ctxData"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetFollowerList 获取关注我的用户列表，目前只获取 id
func (l *GetFollowerListLogic) GetFollowerList(req *types.GetFollowerListReq) (resp *types.GetFollowerListResp, err error) {
	// 用于忽略警告
	_ = req

	userId := ctxData.GetUserIdFromCtx(l.ctx)

	getFollowerListresp, err := l.svcCtx.UserRpcClient.GetFollowerList(l.ctx, &user.GetFollowerListReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetFollowerListResp{
		FollowerList: getFollowerListresp.FollowerId,
	}, nil
}
