package logic

import (
	"context"

	"giligili/app/chat/rpc/internal/svc"
	"giligili/app/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendChatMessageLogic {
	return &SendChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendChatMessageLogic) SendChatMessage(in *pb.SendChatMessageReq) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	return &pb.Empty{}, nil
}
