// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package chat

import (
	"context"

	"giligili/app/chat/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChatMessage           = pb.ChatMessage
	DeleteChatMessageReq  = pb.DeleteChatMessageReq
	DeleteChatMessageResp = pb.DeleteChatMessageResp
	GetChatHistoryReq     = pb.GetChatHistoryReq
	GetChatHistoryResp    = pb.GetChatHistoryResp
	SendChatMessageReq    = pb.SendChatMessageReq
	SendChatMessageResp   = pb.SendChatMessageResp

	Chat interface {
		SendChatMessage(ctx context.Context, in *SendChatMessageReq, opts ...grpc.CallOption) (*SendChatMessageResp, error)
		GetChatHistory(ctx context.Context, in *GetChatHistoryReq, opts ...grpc.CallOption) (*GetChatHistoryResp, error)
		DeleteChatMessage(ctx context.Context, in *DeleteChatMessageReq, opts ...grpc.CallOption) (*DeleteChatMessageResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) SendChatMessage(ctx context.Context, in *SendChatMessageReq, opts ...grpc.CallOption) (*SendChatMessageResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.SendChatMessage(ctx, in, opts...)
}

func (m *defaultChat) GetChatHistory(ctx context.Context, in *GetChatHistoryReq, opts ...grpc.CallOption) (*GetChatHistoryResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.GetChatHistory(ctx, in, opts...)
}

func (m *defaultChat) DeleteChatMessage(ctx context.Context, in *DeleteChatMessageReq, opts ...grpc.CallOption) (*DeleteChatMessageResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.DeleteChatMessage(ctx, in, opts...)
}
