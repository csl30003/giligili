package svc

import (
	"giligili/app/chat/api/internal/config"
	"giligili/app/chat/rpc/chat"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ChatRpcClient chat.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ChatRpcClient: chat.NewChat(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}
