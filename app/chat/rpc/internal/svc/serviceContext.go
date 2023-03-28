package svc

import (
	"giligili/app/chat/model"
	"giligili/app/chat/rpc/internal/config"
	"giligili/app/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	ChatModel     model.ChatModel
	UserRpcClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ChatModel:     model.NewChatModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
		UserRpcClient: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
