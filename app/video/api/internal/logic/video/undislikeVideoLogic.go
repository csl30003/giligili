package video

import (
	"context"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UndislikeVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUndislikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UndislikeVideoLogic {
	return &UndislikeVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UndislikeVideoLogic) UndislikeVideo(req *types.UndislikeVideoReq) (resp *types.UndislikeVideoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
