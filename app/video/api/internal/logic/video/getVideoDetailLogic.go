package video

import (
	"context"
	"giligili/app/video/rpc/video"

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

// GetVideoDetail 获取一条视频详情，没有视频文件，只有信息
func (l *GetVideoDetailLogic) GetVideoDetail(req *types.GetVideoDetailReq) (resp *types.GetVideoDetailResp, err error) {
	getVideoDetailResp, err := l.svcCtx.VideoRpcClient.GetVideoDetail(l.ctx, &video.GetVideoDetailReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetVideoDetailResp{Video: types.Video{
		VideoId:      getVideoDetailResp.Video.VideoId,
		Title:        getVideoDetailResp.Video.Title,
		Url:          getVideoDetailResp.Video.Url,
		UserId:       getVideoDetailResp.Video.UserId,
		UserNickname: getVideoDetailResp.Video.UserNickname,
		Description:  getVideoDetailResp.Video.Description,
		LikeCount:    getVideoDetailResp.Video.LikeCount,
		DislikeCount: getVideoDetailResp.Video.DislikeCount,
		CreateTime:   getVideoDetailResp.Video.CreateTime,
		UpdateTime:   getVideoDetailResp.Video.UpdateTime,
	}}, nil
}
