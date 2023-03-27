package video

import (
	"context"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoDetailLogic {
	return &GetVideoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoDetailLogic) GetVideoDetail(req *types.GetVideoDetailReq) (resp *types.GetVideoDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
