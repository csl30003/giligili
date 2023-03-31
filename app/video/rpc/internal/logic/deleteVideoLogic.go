package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"time"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteVideo 删除视频
func (l *DeleteVideoLogic) DeleteVideo(in *pb.DeleteVideoReq) (*pb.DeleteVideoResp, error) {
	video, err := l.svcCtx.VideoModel.FindOneByIdAndIgnoreStatus(l.ctx, in.VideoId)
	if err != nil {
		return nil, status.Error(100, "查询视频失败")
	}

	// 无权删除别人的视频
	if video.UserId != in.UserId {
		return &pb.DeleteVideoResp{Success: false, Url: ""}, nil
	}

	video.DeleteTime.Time = time.Now()
	video.DeleteTime.Valid = true
	err = l.svcCtx.VideoModel.Update(l.ctx, video)
	if err != nil {
		return nil, status.Error(100, "更新视频失败")
	}

	return &pb.DeleteVideoResp{Success: true, Url: video.Url}, nil
}
