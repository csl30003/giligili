package logic

import (
	"context"
	"giligili/app/user/model"
	"google.golang.org/grpc/status"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFollowerList 获取关注我的用户列表，目前只获取 id
func (l *GetFollowerListLogic) GetFollowerList(in *pb.GetFollowerListReq) (*pb.GetFollowerListResp, error) {
	followList, err := l.svcCtx.FollowModel.FindManyByFolloweeId(l.ctx, in.UserId)
	if err == model.ErrNotFound {
		return nil, status.Error(100, "没有人关注你")
	}
	if err != nil {
		return nil, status.Error(100, "查询关注错误")
	}

	userIdList := make([]int64, len(followList))
	for i := range followList {
		userIdList[i] = followList[i].FollowerId
	}

	return &pb.GetFollowerListResp{
		FollowerId: userIdList,
	}, nil
}
