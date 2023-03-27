package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 自定义组件
	MySQL struct {
		DataSource string
	}
	Cache cache.CacheConf
	Salt  string
}
