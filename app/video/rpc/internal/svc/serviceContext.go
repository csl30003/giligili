package svc

import (
	"giligili/app/user/rpc/user"
	"giligili/app/video/model"
	"giligili/app/video/rpc/internal/config"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VideoModel   model.VideoModel
	BarrageModel model.BarrageModel
	// 这里用的是 go-redis，而不是 go-zero 的 redis，因为参考太少了
	RedisClient   *redis.Client
	UserRpcClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

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
