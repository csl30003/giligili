package chat

import (
	"context"
	"giligili/app/chat/rpc/chat"
	"giligili/common/ctxData"

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

// SendChatMessage 发送聊天信息
func (l *SendChatMessageLogic) SendChatMessage(req *types.SendChatMessageReq) (resp *types.SendChatMessageResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	_, err = l.svcCtx.ChatRpcClient.SendChatMessage(l.ctx, &chat.SendChatMessageReq{
		FromUserId: userId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.SendChatMessageResp{}, nil
}
