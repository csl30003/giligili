package video

import (
	"context"
	"encoding/json"
	"errors"
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

	if !isColor(req.Color) {
		return nil, errors.New("弹幕颜色值错误")
	}

	if req.Type < 0 || req.Type > 2 {
		return nil, errors.New("弹幕类型错误")
	}

	if req.Timestamp < 0 {
		return nil, errors.New("弹幕时间戳错误")
	}

	data := map[string]interface{}{
		"userId":    userId,
		"videoId":   req.VideoId,
		"text":      req.Text,
		"color":     req.Color,
		"type":      req.Type,
		"timestamp": req.Timestamp,
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

// isColor 判断十进制数字是否代表一个颜色
func isColor(num int64) bool {
	if num >= 0 && num <= 16777215 {
		return true
	}
	return false
}
