package config

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	VideoRpcConf      zrpc.RpcClientConf
	RqSendBarrageConf rabbitmq.RabbitSenderConf
	ExchangeConf      rabbitmq.ExchangeConf
}
