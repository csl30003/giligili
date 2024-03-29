package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"giligili/app/video/model"
	"giligili/common/redisConstant"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"

	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBarrageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBarrageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBarrageListLogic {
	return &GetBarrageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetBarrageList 获取视频弹幕
func (l *GetBarrageListLogic) GetBarrageList(in *pb.GetBarrageListReq) (*pb.GetBarrageListResp, error) {
	// 先查询视频是否存在
	_, err := l.svcCtx.VideoModel.FindOneById(l.ctx, in.VideoId)
	if err != nil && err == sqlx.ErrNotFound {
		return nil, status.Error(100, "视频不存在")
	}

	// 先从 Redis 中获取缓存数据
	key := fmt.Sprintf("%s%d", redisConstant.BarrageListKeyPrefix, in.VideoId)
	rst, err := l.svcCtx.RedisClient.ZRevRange(l.ctx, key, 0, -1).Result()
	if err != nil {
		return nil, status.Error(100, "Redis 查询弹幕列表失败")
	}
	if len(rst) != 0 {
		// 如果 Redis 中有缓存数据，直接返回
		var barrageList []*pb.BarrageInfo
		for _, v := range rst {
			var barrageInfo pb.BarrageInfo
			err := json.Unmarshal([]byte(v), &barrageInfo)
			if err != nil {
				return nil, status.Error(100, "解析 Redis 中的弹幕列表失败")
			}
			barrageList = append(barrageList, &barrageInfo)
		}
		return &pb.GetBarrageListResp{BarrageList: barrageList}, nil
	}

	// 如果 Redis 中没有缓存数据，从 MySQL 中查询
	barrageList, err := l.svcCtx.BarrageModel.FindManyByVideoIdAndLimit(l.ctx, in.VideoId, 3000)
	if err == model.ErrNotFound {
		return &pb.GetBarrageListResp{BarrageList: []*pb.BarrageInfo{}}, nil
	}
	if err != nil {
		return nil, status.Error(100, "查询弹幕列表失败")
	}
	if len(barrageList) == 0 {
		return &pb.GetBarrageListResp{BarrageList: []*pb.BarrageInfo{}}, nil
	}

	// 将查询到的数据缓存到 Redis
	var barragePbList []*pb.BarrageInfo
	for _, barrage := range barrageList {
		key := fmt.Sprintf("%s%d", redisConstant.BarrageListKeyPrefix, in.VideoId)
		barragePb := &pb.BarrageInfo{
			BarrageId:    barrage.Id,
			UserId:       barrage.UserId,
			UserNickname: barrage.UserNickname,
			Text:         barrage.Text,
			Color:        barrage.Color,
			Type:         barrage.Type,
			Timestamp:    barrage.Timestamp,
		}
		barrageJson, err := json.Marshal(barragePb)
		if err != nil {
			return nil, status.Error(100, "弹幕列表转 JSON 失败")
		}
		err = l.svcCtx.RedisClient.ZAdd(l.ctx, key, &redis.Z{Member: barrageJson, Score: float64(barrage.Timestamp)}).Err()
		if err != nil {
			return nil, status.Error(100, "弹幕列表缓存 Redis 失败")
		}
		barragePbList = append(barragePbList, barragePb)
	}
	l.svcCtx.RedisClient.Expire(l.ctx, key, redisConstant.BarrageListExpire)
	return &pb.GetBarrageListResp{BarrageList: barragePbList}, nil
}
