package svc

import (
	"giligili/app/comment/api/internal/config"
	"giligili/app/comment/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// model
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// model
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
	}
}
