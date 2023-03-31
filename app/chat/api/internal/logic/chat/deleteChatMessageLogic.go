package chat

import (
	"context"
	"giligili/app/chat/rpc/chat"
	"giligili/common/ctxData"

	"giligili/app/chat/api/internal/svc"
	"giligili/app/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChatMessageLogic {
	return &DeleteChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteChatMessage 删除聊天记录
func (l *DeleteChatMessageLogic) DeleteChatMessage(req *types.DeleteChatMessageReq) (resp *types.DeleteChatMessageResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	deleteChatMessageResp, err := l.svcCtx.ChatRpcClient.DeleteChatMessage(l.ctx, &chat.DeleteChatMessageReq{
		UserId: userId,
		ChatId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteChatMessageResp{Success: deleteChatMessageResp.Success}, nil
}
