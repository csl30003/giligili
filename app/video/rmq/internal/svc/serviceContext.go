package svc

import (
	"giligili/app/user/rpc/user"
	"giligili/app/video/model"
	"giligili/app/video/rmq/internal/config"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VideoModel    model.VideoModel
	BarrageModel  model.BarrageModel
	RedisClient   *redis.Client
	UserRpcClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		VideoModel:   model.NewVideoModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
		BarrageModel: model.NewBarrageModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Cache[0].Host,
			Password: c.Cache[0].Pass,
			DB:       0,
		}),
		UserRpcClient: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
