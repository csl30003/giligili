package logic

import (
	"context"
	"giligili/app/user/model"
	"google.golang.org/grpc/status"
	"time"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowUserLogic {
	return &UnfollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UnfollowUser 取关用户
func (l *UnfollowUserLogic) UnfollowUser(in *pb.UnfollowUserReq) (*pb.Empty, error) {
	// 判断是否相等
	if in.UserId == in.FolloweeId {
		return nil, status.Error(100, "不能取关自己")
	}

	// 判断关注记录是否存在
	follow, err := l.svcCtx.FollowModel.FindOneByFollowerIdAndFolloweeId(l.ctx, in.UserId, in.FolloweeId)
	if err != nil || err == model.ErrNotFound {
		// 找不到记录就返回
		return nil, status.Error(100, "找不到记录")
	}

	// 更新记录，delete_time
	follow.DeleteTime.Time = time.Now()
	follow.DeleteTime.Valid = true
	err = l.svcCtx.FollowModel.Update(l.ctx, follow)
	if err != nil {
		return nil, status.Error(100, "取关失败")
	}

	return &pb.Empty{}, nil
}
