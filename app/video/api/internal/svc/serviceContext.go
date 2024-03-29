package svc

import (
	"giligili/app/video/api/internal/config"
	"giligili/app/video/rpc/video"
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	VideoRpcClient      video.Video
	RqSendBarrageClient rabbitmq.Sender
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		VideoRpcClient:      video.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		RqSendBarrageClient: rabbitmq.MustNewSender(c.RqSendBarrageConf),
	}
}
