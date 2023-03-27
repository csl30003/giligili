package svc

import (
	"giligili/app/user/api/internal/config"
	"giligili/app/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// model
	UserModel   model.UserModel
	FollowModel model.FollowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// model
		UserModel:   model.NewUserModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
		FollowModel: model.NewFollowModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
	}
}
