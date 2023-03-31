package logic

import (
	"context"
	"fmt"
	"giligili/app/video/rpc/internal/svc"
	"giligili/common/redisConstant"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type UpdateRankingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRankingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRankingListLogic {
	return &UpdateRankingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRankingList 更新排行榜，未完成！！！
func (l *UpdateRankingListLogic) UpdateRankingList() error {
	now := time.Now()

	// 从数据库中读取视频的创建时间和点赞数
	videos, err := l.svcCtx.VideoModel.FindMany(l.ctx)
	if err != nil {
		logx.Error("在更新排行榜时，获取视频列表失败，时间：" + now.String())
		return err
	}

	rankingListKey := redisConstant.RankingListKeyPrefix + "rank"
	// 计算加权值
	for _, video := range videos {
		key := fmt.Sprintf("%s%d", redisConstant.LikeCountKeyPrefix, video.Id)
		count, err := l.svcCtx.RedisClient.Get(l.ctx, key).Int64()
		if err != nil {
			logx.Error("在更新排行榜时，从 redis 获取视频的点赞数失败，时间：" + now.String())
		}

		weight := float64(count)*0.8 + now.Sub(video.CreateTime).Hours()*0.2
		// 将加权值写入Redis
		err = l.svcCtx.RedisClient.ZAdd(l.ctx, rankingListKey, &redis.Z{Member: video.Id, Score: weight}).Err()
		if err != nil {
			logx.Error("在更新排行榜时，将加权值写入 Redis 失败，时间：" + now.String() + " 视频id为：" + strconv.FormatInt(video.Id, 10))
		}
	}

	return nil
}
