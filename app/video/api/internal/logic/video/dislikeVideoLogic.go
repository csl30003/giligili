package video

import (
	"context"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DislikeVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDislikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DislikeVideoLogic {
	return &DislikeVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DislikeVideoLogic) DislikeVideo(req *types.DislikeVideoReq) (resp *types.DislikeVideoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
