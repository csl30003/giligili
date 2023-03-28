package logic

import (
	"context"
	"giligili/app/chat/model"
	"giligili/app/chat/rpc/internal/svc"
	"giligili/app/chat/rpc/pb"
	"giligili/app/user/rpc/user"
	"google.golang.org/grpc/status"

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

// SendChatMessage 发送聊天信息
func (l *SendChatMessageLogic) SendChatMessage(in *pb.SendChatMessageReq) (*pb.SendChatMessageResp, error) {
	// 判断两个 id 是否相等
	if in.FromUserId == in.ToUserId {
		return nil, status.Error(100, "不能给自己发聊天信息")
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

	// 插入记录
	newChatMessage := model.Chat{
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Content:    in.Content,
	}
	result, err := l.svcCtx.ChatModel.Insert(l.ctx, &newChatMessage)
	if err != nil {
		return nil, status.Error(100, "消息存入失败")
	}

	newChatMessage.Id, err = result.LastInsertId()
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	return &pb.SendChatMessageResp{}, nil
}
