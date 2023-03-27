package logic

import (
	"context"
	"giligili/app/user/model"
	"google.golang.org/grpc/status"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOneById(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}

		return nil, status.Error(500, err.Error())
	}

	return &pb.GetUserInfoResp{
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email.String,
		Phone:    user.Phone.String,
		Avatar:   user.Avatar.String,
	}, nil
}
