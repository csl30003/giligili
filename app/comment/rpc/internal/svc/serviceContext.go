package svc

import (
	"giligili/app/comment/model"
	"giligili/app/comment/rpc/internal/config"
	"giligili/app/video/rpc/video"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommentModel   model.CommentModel
	VideoRpcClient video.Video
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CommentModel:   model.NewCommentModel(sqlx.NewMysql(c.MySQL.DataSource), c.Cache),
		VideoRpcClient: video.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}
