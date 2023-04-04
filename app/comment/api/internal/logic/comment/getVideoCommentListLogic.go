package comment

import (
	"context"

	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentListLogic {
	return &GetVideoCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoCommentListLogic) GetVideoCommentList(req *types.GetVideoCommentListReq) (resp *types.GetVideoCommentListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
