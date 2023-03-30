package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetVideoList 获取视频列表
func (l *GetVideoListLogic) GetVideoList(in *pb.GetVideoListReq) (*pb.GetVideoListResp, error) {
	if in.Page < 1 {
		return nil, status.Error(100, "页面不能小于1")
	}
	if in.PageSize < 1 {
		return nil, status.Error(100, "页面大小不能小于1")
	}

	videoList, err := l.svcCtx.VideoModel.FindMany(l.ctx, in.Page, in.PageSize, in.Keyword)
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	if len(videoList) == 0 {
		return &pb.GetVideoListResp{}, nil
	}

	var videoListResp []*pb.VideoInfo
	for _, v := range videoList {
		videoListResp = append(videoListResp, &pb.VideoInfo{
			VideoId:      v.Id,
			Title:        v.Title,
			Url:          v.Url,
			UserId:       v.UserId,
			UserNickname: v.Nickname,
			Description:  v.Description.String,
			LikeCount:    v.Like,
			DislikeCount: v.Dislike,
			CreateTime:   v.CreateTime.String(),
			UpdateTime:   v.UpdateTime.String(),
		})
	}

	return &pb.GetVideoListResp{VideoList: videoListResp}, nil
}
