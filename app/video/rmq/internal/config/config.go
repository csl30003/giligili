package config

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// RabbitMQ 配置，包括消费者
	ListenerConf rabbitmq.RabbitListenerConf
	ExchangeConf rabbitmq.ExchangeConf
	MySQL        struct {
		DataSource string
	}
	Cache       cache.CacheConf
	UserRpcConf zrpc.RpcClientConf
}
