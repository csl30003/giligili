package logic

import (
	"context"
	"giligili/app/user/model"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/app/user/rpc/user"
	"giligili/common/cryptx"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 判断用户名是否已经存在，这里有个坑，go-zero查询会先查 redis，没有再去 MySQL，
	// 新用户第一次注册，用用户名查，查不到数据才能注册，而这里 go-zero 查不到会到 redis 设置一个键值对：username:xxx : *，
	// 这是为了防止缓存穿透，这就导致了这 1 分钟内再次用该用户名注册还能成功！因为这用户名在 redis 查不到数据，查到的是 *，
	// 接着下面就符合 err == model.ErrNotFound 了，可恶。
	// 解决方法可以是不走缓存。
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Username: in.Username,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Nickname: in.Nickname,
		}

		// 插入
		result, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		// 获取该用户的 id
		newUser.Id, err = result.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		logx.Info(newUser.Username + "注册成功")
		return &user.RegisterResp{
			UserId: newUser.Id,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
