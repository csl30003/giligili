package svc

import (
	"giligili/app/user/rpc/user"
	"giligili/app/video/model"
	"giligili/app/video/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VideoModel    model.VideoModel
	UserRpcClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		VideoModel:    model.NewVideoModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
		UserRpcClient: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
