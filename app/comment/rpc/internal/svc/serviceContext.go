package svc

import (
	"giligili/app/comment/model"
	"giligili/app/comment/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
	}
}
