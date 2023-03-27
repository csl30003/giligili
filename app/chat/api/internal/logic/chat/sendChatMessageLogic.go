package chat

import (
	"context"

	"giligili/app/chat/api/internal/svc"
	"giligili/app/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendChatMessageLogic {
	return &SendChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendChatMessageLogic) SendChatMessage(req *types.SendChatMessageReq) (resp *types.SendChatMessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
