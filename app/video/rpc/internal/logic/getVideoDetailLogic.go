package logic

import (
	"context"
	"giligili/app/video/model"
	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoDetailLogic {
	return &GetVideoDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetVideoDetail 获取一条视频详情
func (l *GetVideoDetailLogic) GetVideoDetail(in *pb.GetVideoDetailReq) (*pb.GetVideoDetailResp, error) {
	video, err := l.svcCtx.VideoModel.FindVideoTempById(l.ctx, in.VideoId)
	if err == model.ErrNotFound {
		return &pb.GetVideoDetailResp{Video: nil}, nil
	}
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	return &pb.GetVideoDetailResp{
		Video: &pb.VideoInfo{
			VideoId:      video.Id,
			Title:        video.Title,
			Url:          video.Url,
			UserId:       video.UserId,
			UserNickname: video.Nickname,
			Description:  video.Description.String,
			LikeCount:    video.Like,
			DislikeCount: video.Dislike,
			CreateTime:   video.CreateTime.String(),
			UpdateTime:   video.UpdateTime.String(),
		},
	}, nil
}
