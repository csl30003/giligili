package logic

import (
	"context"
	"giligili/app/chat/model"
	"google.golang.org/grpc/status"
	"time"

	"giligili/app/chat/rpc/internal/svc"
	"giligili/app/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChatMessageLogic {
	return &DeleteChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteChatMessage 删除聊天记录
func (l *DeleteChatMessageLogic) DeleteChatMessage(in *pb.DeleteChatMessageReq) (*pb.DeleteChatMessageResp, error) {
	chat, err := l.svcCtx.ChatModel.FindOneById(l.ctx, in.ChatId)
	if err == model.ErrNotFound {
		return nil, status.Error(100, "聊天记录不存在")
	}
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	// 如果聊天信息的发起人不是自己,不能删除
	if chat.FromUserId != in.UserId {
		return nil, status.Error(100, "你没有权限删除此聊天记录")
	}

	// 时差大于2分钟不允许删除聊天信息
	now := time.Now()
	if now.Sub(chat.CreateTime).Minutes() > 2 {
		return &pb.DeleteChatMessageResp{Success: false}, nil
	}

	chat.DeleteTime.Time = now
	chat.DeleteTime.Valid = true
	err = l.svcCtx.ChatModel.Update(l.ctx, chat)
	if err != nil {
		return nil, status.Error(100, "删除聊天失败")
	}

	return &pb.DeleteChatMessageResp{Success: true}, nil
}
