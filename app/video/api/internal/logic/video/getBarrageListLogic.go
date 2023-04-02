package video

import (
	"context"
	"giligili/app/video/rpc/video"

	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBarrageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBarrageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBarrageListLogic {
	return &GetBarrageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetBarrageList 获取视频弹幕
func (l *GetBarrageListLogic) GetBarrageList(req *types.GetBarrageListReq) (resp *types.GetBarrageListResp, err error) {
	getBarrageListResp, err := l.svcCtx.VideoRpcClient.GetBarrageList(l.ctx, &video.GetBarrageListReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	if len(getBarrageListResp.BarrageList) == 0 {
		return &types.GetBarrageListResp{}, nil
	}

	var barrageList []types.Barrage
	for _, v := range getBarrageListResp.BarrageList {
		barrageList = append(barrageList, types.Barrage{
			BarrageId:    v.BarrageId,
			UserId:       v.UserId,
			UserNickname: v.UserNickname,
			Text:         v.Text,
			Color:        v.Color,
			Type:         v.Type,
			Timestamp:    v.Timestamp,
		})
	}

	return &types.GetBarrageListResp{BarrageList: barrageList}, nil
}
