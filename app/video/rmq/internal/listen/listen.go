package listen

import (
	"context"
	"giligili/app/video/rmq/internal/config"
	"giligili/app/video/rmq/internal/svc"
	"github.com/zeromicro/go-zero/core/service"
)

// Mqs 返回所有消费者
func Mqs(c config.Config) []service.Service {
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	services = append(services, RqMqs(c, ctx, svcContext)...)

	return services
}
