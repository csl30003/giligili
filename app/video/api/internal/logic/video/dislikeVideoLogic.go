package video

import (
	"context"
	"giligili/app/video/rpc/video"
	"giligili/common/ctxData"

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

// DislikeVideo 用户给视频踩
func (l *DislikeVideoLogic) DislikeVideo(req *types.DislikeVideoReq) (resp *types.DislikeVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	dislikeVideoResp, err := l.svcCtx.VideoRpcClient.DislikeVideo(l.ctx, &video.DislikeVideoReq{
		UserId:  userId,
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	return &types.DislikeVideoResp{Success: dislikeVideoResp.Success}, nil
}
