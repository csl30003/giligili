package logic

import (
	"context"
	"encoding/json"
	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"
	"giligili/common/redisConstant"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/status"
	"math"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHotVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotVideoListLogic {
	return &GetHotVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetHotVideoList 获取热门视频列表
func (l *GetHotVideoListLogic) GetHotVideoList(in *pb.GetHotVideoListReq) (*pb.GetHotVideoListResp, error) {
	// 先看缓存有没有，没有就去数据库拿出来，存缓存，5分钟或者多少分钟，返回
	_ = in
	// 先从 Redis 中获取缓存数据
	key := redisConstant.HotVideoListKey
	rst, err := l.svcCtx.RedisClient.ZRevRangeWithScores(l.ctx, key, 0, 29).Result()
	if err != nil {
		return nil, status.Error(100, "Redis 查询热门视频列表失败")
	}

	if len(rst) != 0 {
		// 如果 Redis 中有缓存数据，直接返回
		var hotVideoList []*pb.HotVideoInfo
		for _, v := range rst {
			var hotVideoInfo pb.HotVideoInfo
			err := json.Unmarshal([]byte(v.Member.(string)), &hotVideoInfo)
			if err != nil {
				return nil, status.Error(100, "解析 Redis 中的热门视频列表失败")
			}
			hotVideoList = append(hotVideoList, &hotVideoInfo)
		}
		return &pb.GetHotVideoListResp{HotVideoList: hotVideoList}, nil
	}

	// 如果 Redis 中没有缓存数据，从 MySQL 中查询
	videoList, err := l.svcCtx.VideoModel.FindManyByLimitDesc(l.ctx, 100)
	if err != nil {
		return nil, status.Error(100, "查询视频列表失败")
	}
	if len(videoList) == 0 {
		return &pb.GetHotVideoListResp{HotVideoList: []*pb.HotVideoInfo{}}, nil
	}

	// 计算热门指数
	var hotVideoInfos []*pb.HotVideoInfo
	for _, video := range videoList {
		timestamp := float64(video.CreateTime.Unix()) / 1000000000.0
		hotIndex := video.Like + int64(math.Floor(timestamp))
		hotVideoInfo := &pb.HotVideoInfo{
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
			HotIndex:     hotIndex,
		}
		hotVideoInfos = append(hotVideoInfos, hotVideoInfo)
	}

	// 按热门指数排序取前30个
	sort.Slice(hotVideoInfos, func(i, j int) bool {
		return hotVideoInfos[i].HotIndex > hotVideoInfos[j].HotIndex
	})
	if len(hotVideoInfos) > 30 {
		hotVideoInfos = hotVideoInfos[:30]
	}

	// 将查询到的数据缓存到 Redis
	for _, hotVideo := range hotVideoInfos {
		key := redisConstant.HotVideoListKey
		hotVideoJson, err := json.Marshal(hotVideo)
		if err != nil {
			return nil, status.Error(100, "热门视频列表转 JSON 失败")
		}
		err = l.svcCtx.RedisClient.ZAdd(l.ctx, key, &redis.Z{Member: hotVideoJson, Score: float64(hotVideo.HotIndex)}).Err()
		if err != nil {
			return nil, status.Error(100, "热门视频列表缓存 Redis 失败")
		}
	}
	l.svcCtx.RedisClient.Expire(l.ctx, key, redisConstant.HotVideoListExpire)
	return &pb.GetHotVideoListResp{HotVideoList: hotVideoInfos}, nil
}
