package logic

import (
	"context"
	"giligili/app/video/model"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExistVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistVideoLogic {
	return &IsExistVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IsExistVideo 判断是否存在
func (l *IsExistVideoLogic) IsExistVideo(in *pb.IsExistVideoReq) (*pb.IsExistVideoResp, error) {
	_, err := l.svcCtx.VideoModel.FindOneById(l.ctx, in.VideoId)
	if err == model.ErrNotFound {
		return &pb.IsExistVideoResp{Exists: false}, nil
	}
	if err != nil {
		return nil, err
	}

	return &pb.IsExistVideoResp{Exists: true}, nil
}
