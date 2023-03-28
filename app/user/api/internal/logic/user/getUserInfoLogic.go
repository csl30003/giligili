package user

import (
	"context"
	"giligili/app/user/rpc/user"
	"giligili/common/ctxData"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUserInfo 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// 用于忽略警告
	_ = req

	userId := ctxData.GetUserIdFromCtx(l.ctx)

	getUserInfoResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		Username: getUserInfoResp.Username,
		Nickname: getUserInfoResp.Nickname,
		Email:    getUserInfoResp.Email,
		Phone:    getUserInfoResp.Phone,
		Avatar:   getUserInfoResp.Avatar,
	}, nil
}
