package logic

import (
	"context"
	"fmt"
	"giligili/app/video/model"
	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"
	"giligili/common/redisConstant"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LikeVideo 用户给视频点赞
func (l *LikeVideoLogic) LikeVideo(in *pb.LikeVideoReq) (*pb.LikeVideoResp, error) {
	// 根据视频id判断视频记录是否存在
	video, err := l.svcCtx.VideoModel.FindOneById(l.ctx, in.VideoId)
	if err != nil || err == model.ErrNotFound {
		return nil, status.Error(100, "没有视频记录")
	}

	// 在Redis中创建一个Hash结构，以视频ID为键，以点赞用户ID为值，用于存储点赞信息
	key := fmt.Sprintf("%s%d", redisConstant.LikeSetKeyPrefix, in.VideoId)

	// 查询Redis中是否存在点赞信息
	exist, err := l.svcCtx.RedisClient.SIsMember(l.ctx, key, in.UserId).Result()
	if err != nil {
		return nil, status.Error(100, "查询Redis中点赞信息失败")
	}
	if exist {
		// 如果点赞记录已存在，返回错误信息
		return &pb.LikeVideoResp{Success: false}, nil
	}

	// 向Redis中添加点赞信息
	err = l.svcCtx.RedisClient.SAdd(l.ctx, key, in.UserId).Err()
	if err != nil {
		return nil, status.Error(100, "添加点赞信息到Redis失败")
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

	return &pb.LikeVideoResp{Success: true}, nil
}
