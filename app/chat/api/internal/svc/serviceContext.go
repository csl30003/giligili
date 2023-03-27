package svc

import (
	"giligili/app/chat/api/internal/config"
	"giligili/app/chat/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// model
	ChatModel model.ChatModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// model
		ChatModel: model.NewChatModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
	}
}
