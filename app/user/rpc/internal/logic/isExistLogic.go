package logic

import (
	"context"
	"giligili/app/user/model"
	"google.golang.org/grpc/status"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistLogic {
	return &IsExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IsExist 判断用户是否存在
func (l *IsExistLogic) IsExist(in *pb.IsExistReq) (*pb.IsExistResp, error) {
	_, err := l.svcCtx.UserModel.FindOneById(l.ctx, in.UserId)
	if err != nil {
		return nil, status.Error(100, "判断用户是否存在失败")
	}
	if err == model.ErrNotFound {
		return &pb.IsExistResp{IsExist: false}, nil
	}

	return &pb.IsExistResp{IsExist: true}, nil
}
