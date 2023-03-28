package logic

import (
	"context"
	"giligili/app/user/model"
	"google.golang.org/grpc/status"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFolloweeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFolloweeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFolloweeListLogic {
	return &GetFolloweeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFolloweeList 获取关注我的用户列表，目前只获取 id
func (l *GetFolloweeListLogic) GetFolloweeList(in *pb.GetFolloweeListReq) (*pb.GetFolloweeListResp, error) {
	followList, err := l.svcCtx.FollowModel.FindManyByFollowerId(l.ctx, in.UserId)
	if err == model.ErrNotFound {
		return nil, status.Error(100, "你没关注任何人")
	}
	if err != nil {
		return nil, status.Error(100, "查询关注错误")
	}

	userIdList := make([]int64, len(followList))
	for i := range followList {
		userIdList[i] = followList[i].FolloweeId
	}

	return &pb.GetFolloweeListResp{
		FolloweeId: userIdList,
	}, nil
}
