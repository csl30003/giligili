package svc

import (
	"giligili/app/video/api/internal/config"
	"giligili/app/video/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// model
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// model
		VideoModel: model.NewVideoModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
	}
}
