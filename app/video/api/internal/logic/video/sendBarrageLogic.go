package video

import (
	"context"
	"encoding/json"
	"giligili/common/ctxData"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendBarrageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendBarrageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendBarrageLogic {
	return &SendBarrageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SendBarrage 发送弹幕
func (l *SendBarrageLogic) SendBarrage(req *types.SendBarrageReq) (resp *types.SendBarrageResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	data := map[string]interface{}{
		"userId":  userId,
		"videoId": req.VideoId,
		"text":    req.Text,
		"color":   req.Color,
		"type":    req.Type,
	}
	msg, _ := json.Marshal(data)

	// 发送至消息队列
	err = l.svcCtx.RqSendBarrageClient.Send(
		l.svcCtx.Config.ExchangeConf.ExchangeName,
		l.svcCtx.Config.ExchangeConf.Queues[0].Name,
		msg,
	)
	if err != nil {
		return nil, err
	}

	return &types.SendBarrageResp{Success: true}, nil
}
