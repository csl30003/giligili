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

type UndislikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUndislikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UndislikeVideoLogic {
	return &UndislikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UndislikeVideo 用户取消视频踩
func (l *UndislikeVideoLogic) UndislikeVideo(in *pb.UndislikeVideoReq) (*pb.UndislikeVideoResp, error) {
	// 根据视频id判断视频记录是否存在
	video, err := l.svcCtx.VideoModel.FindOneById(l.ctx, in.VideoId)
	if err != nil || err == model.ErrNotFound {
		return nil, status.Error(100, "没有视频记录")
	}

	// 在Redis中创建一个Hash结构，以视频ID为键，以踩用户ID为值，用于存储踩信息
	key := fmt.Sprintf("%s%d", redisConstant.DislikeSetKeyPrefix, in.VideoId)

	// 判断踩信息是否存在
	exists, err := l.svcCtx.RedisClient.SIsMember(l.ctx, key, in.UserId).Result()
	if err != nil {
		return nil, status.Error(100, "从Redis中查询踩信息失败")
	}
	if !exists {
		return &pb.UndislikeVideoResp{Success: false}, nil
	}

	// 从Redis中删除踩信息
	err = l.svcCtx.RedisClient.SRem(l.ctx, key, in.UserId).Err()
	if err != nil {
		return nil, status.Error(100, "从Redis中删除踩信息失败")
	}

	// 从Redis中获取踩数，并更新缓存
	count, err := l.svcCtx.RedisClient.SCard(l.ctx, key).Result()
	if err != nil {
		return nil, status.Error(100, "从Redis中获取踩数失败")
	}
	countKey := fmt.Sprintf("%s%d", redisConstant.DislikeCountKeyPrefix, in.VideoId)
	err = l.svcCtx.RedisClient.Set(l.ctx, countKey, count, 0).Err()
	if err != nil {
		return nil, status.Error(100, "更新踩数缓存失败")
	}

	// 将踩信息写入数据库
	video.Dislike = count
	err = l.svcCtx.VideoModel.Update(l.ctx, video)
	if err != nil {
		return nil, status.Error(100, "数据库更新视频记录失败")
	}
	return &pb.UndislikeVideoResp{Success: true}, nil
}
