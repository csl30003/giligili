package rq

import (
	"context"
	"fmt"
	"giligili/app/video/rmq/internal/svc"
)

type SendBarrageMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendBarrageMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendBarrageMq {
	return &SendBarrageMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendBarrageMq) Consume(message string) error {
	fmt.Println("消费者消费信息：")
	fmt.Println(message)
	return nil
}
