package listen

import (
	"context"
	"giligili/app/video/rmq/internal/config"
	"giligili/app/video/rmq/internal/mqs/rq"
	"giligili/app/video/rmq/internal/svc"
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
)

func RqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		// 目前用个 RabbitMQ 够了
		rabbitmq.MustNewListener(c.ListenerConf, rq.NewSendBarrageMq(ctx, svcContext)),
	}
}
