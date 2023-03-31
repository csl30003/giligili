package logic

import (
	"context"
	"giligili/app/user/rpc/user"
	"google.golang.org/grpc/status"

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

// GetChatHistory 获取聊天信息历史记录
func (l *GetChatHistoryLogic) GetChatHistory(in *pb.GetChatHistoryReq) (*pb.GetChatHistoryResp, error) {
	// 判断两个 id 是否相等
	if in.FromUserId == in.ToUserId {
		return nil, status.Error(100, "不能获取自己的聊天记录")
	}

	// 通过对方 id 查一下对方是否存在
	isExistResp, err := l.svcCtx.UserRpcClient.IsExist(l.ctx, &user.IsExistReq{
		UserId: in.ToUserId,
	})
	if err != nil {
		return nil, status.Error(100, err.Error())
	}
	if !isExistResp.IsExist {
		return nil, status.Error(100, "对方账户不存在")
	}

	// 查询记录
	chatMessages, err := l.svcCtx.ChatModel.FindManyByFromUserIdAndToUserId(l.ctx, in.FromUserId, in.ToUserId)
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	result := &pb.GetChatHistoryResp{
		ChatHistory: make([]*pb.ChatMessage, 0, len(chatMessages)),
	}
	for _, chatMessage := range chatMessages {
		result.ChatHistory = append(result.ChatHistory, &pb.ChatMessage{
			Id:               chatMessage.Id,
			FromUserId:       chatMessage.FromUserId,
			FromUserNickname: chatMessage.FromUserNickname,
			ToUserId:         chatMessage.ToUserId,
			ToUserNickname:   chatMessage.ToUserNickname,
			Content:          chatMessage.Content,
			CreateTime:       chatMessage.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return result, nil
}
