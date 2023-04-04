package svc

import (
	"giligili/app/comment/api/internal/config"
	"giligili/app/comment/rpc/comment"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	CommentRpcClient comment.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		CommentRpcClient: comment.NewComment(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
