package user

import (
	"context"
	"giligili/app/user/rpc/user"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.UserRpcClient.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		UserId: registerResp.UserId,
	}, nil
}
