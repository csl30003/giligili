package video

import (
	"context"
	"giligili/app/video/rpc/video"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetVideoList 获取视频列表
func (l *GetVideoListLogic) GetVideoList(req *types.GetVideoListReq) (resp *types.GetVideoListResp, err error) {
	var videoList []types.Video
	getVideoListResp, err := l.svcCtx.VideoRpcClient.GetVideoList(l.ctx, &video.GetVideoListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	if len(getVideoListResp.VideoList) == 0 {
		return &types.GetVideoListResp{}, nil
	}

	for _, v := range getVideoListResp.VideoList {
		videoList = append(videoList, types.Video{
			VideoId:      v.VideoId,
			Title:        v.Title,
			Url:          v.Url,
			UserId:       v.UserId,
			UserNickname: v.UserNickname,
			Description:  v.Description,
			LikeCount:    v.LikeCount,
			DislikeCount: v.DislikeCount,
			CreateTime:   v.CreateTime,
			UpdateTime:   v.UpdateTime,
		})
	}

	return &types.GetVideoListResp{VideoList: videoList}, nil
}
