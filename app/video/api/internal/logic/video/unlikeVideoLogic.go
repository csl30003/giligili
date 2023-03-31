package video

import (
	"context"
	"giligili/app/video/rpc/video"
	"giligili/common/ctxData"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnlikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeVideoLogic {
	return &UnlikeVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UnlikeVideo 用户取消视频点赞
func (l *UnlikeVideoLogic) UnlikeVideo(req *types.UnlikeVideoReq) (resp *types.UnlikeVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	unlikeVideoResp, err := l.svcCtx.VideoRpcClient.UnlikeVideo(l.ctx, &video.UnlikeVideoReq{
		UserId:  userId,
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UnlikeVideoResp{Success: unlikeVideoResp.Success}, nil
}
