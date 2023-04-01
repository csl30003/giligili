package svc

import (
	"giligili/app/video/rmq/internal/config"
)

type ServiceContext struct {
	Config config.Config

	// 可以加 rpc 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
