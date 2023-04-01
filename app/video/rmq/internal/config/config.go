package config

import "github.com/zeromicro/go-queue/rabbitmq"

type Config struct {
	// RabbitMQ 配置，包括消费者
	ListenerConf rabbitmq.RabbitListenerConf
	ExchangeConf rabbitmq.ExchangeConf
}
