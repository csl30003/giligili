package chat

import (
	"context"
	"giligili/app/chat/rpc/chat"
	"giligili/common/ctxData"

	"giligili/app/chat/api/internal/svc"
	"giligili/app/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatHistoryLogic {
	return &GetChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetChatHistory 获取聊天信息历史记录
func (l *GetChatHistoryLogic) GetChatHistory(req *types.GetChatHistoryReq) (resp *types.GetChatHistoryResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	chatHistoryResp, err := l.svcCtx.ChatRpcClient.GetChatHistory(l.ctx, &chat.GetChatHistoryReq{
		FromUserId: userId,
		ToUserId:   req.ToUserId,
	})
	if err != nil {
		return nil, err
	}

	result := &types.GetChatHistoryResp{
		ChatHistory: make([]types.ChatMessage, 0, len(chatHistoryResp.ChatHistory)),
	}
	for _, chatMessage := range chatHistoryResp.ChatHistory {
		result.ChatHistory = append(result.ChatHistory, types.ChatMessage{
			FromUserId:       chatMessage.FromUserId,
			FromUserNickname: chatMessage.FromUserNickname,
			ToUserId:         chatMessage.ToUserId,
			ToUserNickname:   chatMessage.ToUserNickname,
			Content:          chatMessage.Content,
			CreateTime:       chatMessage.CreateTime,
		})
	}

	return result, nil
}
