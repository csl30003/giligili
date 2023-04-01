package admin

import (
	"giligili/app/video/rmq/internal/config"
	"github.com/zeromicro/go-queue/rabbitmq"
	"log"
)

// Admin 目前只做 RabbitMQ 的一个交换机对应一条或多条队列的初始化
func Admin(c config.Config) {
	admin := rabbitmq.MustNewAdmin(c.ListenerConf.RabbitConf)

	// 定义交换机
	err := admin.DeclareExchange(c.ExchangeConf, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 定义队列
	err = admin.DeclareQueue(c.ExchangeConf.Queues[0], nil)
	if err != nil {
		log.Fatal(err)
	}

	// 绑定交换机和队列
	err = admin.Bind(
		c.ExchangeConf.Queues[0].Name,
		c.ExchangeConf.Queues[0].Name,
		c.ExchangeConf.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}
