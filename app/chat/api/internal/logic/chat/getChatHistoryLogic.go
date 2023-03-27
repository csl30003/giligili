package chat

import (
	"context"

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

func (l *GetChatHistoryLogic) GetChatHistory(req *types.GetChatHistoryReq) (resp *types.GetChatHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
