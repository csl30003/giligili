package video

import (
	"context"
	"giligili/app/video/rpc/video"
	"giligili/common/ctxData"
	"os"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteVideo 删除视频
func (l *DeleteVideoLogic) DeleteVideo(req *types.DeleteVideoReq) (resp *types.DeleteVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	deleteVideoResp, err := l.svcCtx.VideoRpcClient.DeleteVideo(l.ctx, &video.DeleteVideoReq{
		UserId:  userId,
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	// 如果 model 成功删除视频记录，在此根据 url 删除视频文件
	if deleteVideoResp.Success {
		err := os.Remove(deleteVideoResp.Url)
		if err != nil {
			logx.Error("删除视频文件失败", err)
			return nil, err
		}
	}

	return &types.DeleteVideoResp{Success: deleteVideoResp.Success}, nil
}
