package logic

import (
	"context"

	"giligili/app/chat/rpc/internal/svc"
	"giligili/app/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatHistoryLogic {
	return &GetChatHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChatHistoryLogic) GetChatHistory(in *pb.GetChatHistoryReq) (*pb.GetChatHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetChatHistoryResp{}, nil
}
