package video

import (
	"context"
	"giligili/app/video/rpc/video"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotVideoListLogic {
	return &GetHotVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetHotVideoList 获取热门视频列表
func (l *GetHotVideoListLogic) GetHotVideoList() (resp *types.GetHotVideoListResp, err error) {
	getHotVideoListResp, err := l.svcCtx.VideoRpcClient.GetHotVideoList(l.ctx, &video.GetHotVideoListReq{})
	if err != nil {
		return nil, err
	}

	if len(getHotVideoListResp.HotVideoList) == 0 {
		return &types.GetHotVideoListResp{}, nil
	}

	var hotVideoList []types.HotVideo
	for _, v := range getHotVideoListResp.HotVideoList {
		hotVideoList = append(hotVideoList, types.HotVideo{
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
			HotIndex:     v.HotIndex,
		})
	}

	return &types.GetHotVideoListResp{HotVideoList: hotVideoList}, nil
}
