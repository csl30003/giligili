package logic

import (
	"context"
	"giligili/app/user/model"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/app/user/rpc/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FollowUser 关注用户
func (l *FollowUserLogic) FollowUser(in *pb.FollowUserReq) (*pb.FollowUserResp, error) {
	// 判断是否相等
	if in.UserId == in.FolloweeId {
		return nil, status.Error(100, "不能关注自己")
	}

	// 判断用户是否存在
	_, err := l.svcCtx.UserModel.FindOneById(l.ctx, in.FolloweeId)
	if err != nil || err == model.ErrNotFound {
		return nil, status.Error(100, "用户不存在")
	}

	// 判断关注记录是否存在
	_, err = l.svcCtx.FollowModel.FindOneByFollowerIdAndFolloweeId(l.ctx, in.UserId, in.FolloweeId)
	if err == nil {
		// 找得到记录
		return nil, status.Error(100, "已关注")
	}

	if err == model.ErrNotFound {
		newFollow := model.Follow{
			FollowerId: in.UserId,
			FolloweeId: in.FolloweeId,
		}

		// 插入
		result, err := l.svcCtx.FollowModel.Insert(l.ctx, &newFollow)
		if err != nil {
			return nil, status.Error(100, "关注失败")
		}

		newFollow.Id, err = result.LastInsertId()
		if err != nil {
			return nil, status.Error(100, err.Error())
		}

		return &user.FollowUserResp{}, nil
	}

	return nil, status.Error(100, err.Error())
}
