package video

import (
	"context"
	"giligili/app/video/rpc/video"
	"giligili/common/ctxData"

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

// UndislikeVideo 用户取消视频踩
func (l *UndislikeVideoLogic) UndislikeVideo(req *types.UndislikeVideoReq) (resp *types.UndislikeVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	undislikeVideoResp, err := l.svcCtx.VideoRpcClient.UndislikeVideo(l.ctx, &video.UndislikeVideoReq{
		UserId:  userId,
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UndislikeVideoResp{Success: undislikeVideoResp.Success}, nil
}
