package rq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	user2 "giligili/app/user/rpc/user"
	"giligili/app/video/model"
	"giligili/app/video/rmq/internal/svc"
	"giligili/app/video/rmq/internal/types"
	"giligili/app/video/rpc/pb"
	"giligili/common/redisConstant"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SendBarrageMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendBarrageMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendBarrageMq {
	return &SendBarrageMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Consume 消费弹幕，值得一提的是，go-zero 对 RabbitMQ 的封装，
// 在消费前 message 给转成 string 类型了，而这里又要转成 []byte，实在是...
func (l *SendBarrageMq) Consume(message string) error {
	var barrage types.Barrage
	err := json.Unmarshal([]byte(message), &barrage)
	if err != nil {
		return err
	}

	// 判断视频 id 是否正确
	_, err = l.svcCtx.VideoModel.FindOneById(l.ctx, barrage.VideoId)
	if err != nil || err == sqlx.ErrNotFound {
		return errors.New("查询视频失败")
	}

	// 获取用户信息以得到昵称
	user, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &user2.GetUserInfoReq{
		Id: barrage.UserId,
	})

	// 将弹幕保存至数据库
	newBarrage := model.Barrage{
		UserId:       barrage.UserId,
		UserNickname: user.Nickname,
		VideoId:      barrage.VideoId,
		Text:         barrage.Text,
		Color:        barrage.Color,
		Type:         barrage.Type,
		Timestamp:    barrage.Timestamp,
	}
	result, err := l.svcCtx.BarrageModel.Insert(l.ctx, &newBarrage)
	if err != nil {
		return err
	}
	newBarrage.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	// 如果 redis 里有该视频的弹幕缓存，就把该弹幕添加至缓存
	key := fmt.Sprintf("%s%d", redisConstant.BarrageListKeyPrefix, barrage.VideoId)
	cached, err := l.svcCtx.RedisClient.Exists(l.ctx, key).Result()
	if err != nil {
		return err
	}
	if cached == 1 {
		barragePb := &pb.BarrageInfo{
			BarrageId:    newBarrage.Id,
			UserId:       newBarrage.UserId,
			UserNickname: newBarrage.UserNickname,
			Text:         newBarrage.Text,
			Color:        newBarrage.Color,
			Type:         newBarrage.Type,
			Timestamp:    newBarrage.Timestamp,
		}
		barrageJson, err := json.Marshal(barragePb)

		// 插入
		err = l.svcCtx.RedisClient.ZAdd(l.ctx, key, &redis.Z{
			Score:  float64(barrage.Timestamp),
			Member: barrageJson,
		}).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
