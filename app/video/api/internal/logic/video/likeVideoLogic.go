package video

import (
	"context"
	"giligili/app/video/rpc/video"
	"giligili/common/ctxData"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LikeVideo 用户给视频点赞
func (l *LikeVideoLogic) LikeVideo(req *types.LikeVideoReq) (resp *types.LikeVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	likeVideoResp, err := l.svcCtx.VideoRpcClient.LikeVideo(l.ctx, &video.LikeVideoReq{
		UserId:  userId,
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	return &types.LikeVideoResp{Success: likeVideoResp.Success}, nil
}
