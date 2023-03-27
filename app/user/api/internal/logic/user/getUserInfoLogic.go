package user

import (
	"context"
	"encoding/json"
	"errors"
	"giligili/app/user/rpc/user"

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

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	if userId, ok := l.ctx.Value("userId").(json.Number); ok {
		if int64UserId, err := userId.Int64(); err == nil {
			getUserInfoResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &user.GetUserInfoReq{
				Id: int64UserId,
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
		} else {
			return nil, err
		}
	}

	return nil, errors.New("无法识别用户")
}
