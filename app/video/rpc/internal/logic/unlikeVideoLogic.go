package logic

import (
	"context"
	"fmt"
	"giligili/app/video/model"
	"giligili/common/redisConstant"
	"google.golang.org/grpc/status"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnlikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeVideoLogic {
	return &UnlikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UnlikeVideo 用户取消视频点赞
func (l *UnlikeVideoLogic) UnlikeVideo(in *pb.UnlikeVideoReq) (*pb.UnlikeVideoResp, error) {
	// 根据视频id判断视频记录是否存在
	video, err := l.svcCtx.VideoModel.FindOneById(l.ctx, in.VideoId)
	if err != nil || err == model.ErrNotFound {
		return nil, status.Error(100, "没有视频记录")
	}

	// 在Redis中创建一个Hash结构，以视频ID为键，以点赞用户ID为值，用于存储点赞信息
	key := fmt.Sprintf("%s%d", redisConstant.LikeSetKeyPrefix, in.VideoId)

	// 判断点赞信息是否存在
	exists, err := l.svcCtx.RedisClient.SIsMember(l.ctx, key, in.UserId).Result()
	if err != nil {
		return nil, status.Error(100, "从Redis中查询点赞信息失败")
	}
	if !exists {
		return &pb.UnlikeVideoResp{Success: false}, nil
	}

	// 从Redis中删除点赞信息
	err = l.svcCtx.RedisClient.SRem(l.ctx, key, in.UserId).Err()
	if err != nil {
		return nil, status.Error(100, "从Redis中删除点赞信息失败")
	}

	// 从Redis中获取点赞数，并更新缓存
	count, err := l.svcCtx.RedisClient.SCard(l.ctx, key).Result()
	if err != nil {
		return nil, status.Error(100, "从Redis中获取点赞数失败")
	}
	countKey := fmt.Sprintf("%s%d", redisConstant.LikeCountKeyPrefix, in.VideoId)
	err = l.svcCtx.RedisClient.Set(l.ctx, countKey, count, 0).Err()
	if err != nil {
		return nil, status.Error(100, "更新点赞数缓存失败")
	}

	// 将点赞信息写入数据库
	video.Like = count
	err = l.svcCtx.VideoModel.Update(l.ctx, video)
	if err != nil {
		return nil, status.Error(100, "数据库更新视频记录失败")
	}
	return &pb.UnlikeVideoResp{Success: true}, nil
}
