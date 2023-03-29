package logic

import (
	"context"
	"database/sql"
	"giligili/app/video/model"
	"giligili/app/video/rpc/internal/svc"
	"giligili/app/video/rpc/pb"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadVideo 上传视频
func (l *UploadVideoLogic) UploadVideo(in *pb.UploadVideoReq) (*pb.UploadVideoResp, error) {
	newVideo := model.Video{
		Title:  in.Title,
		Url:    in.Url,
		UserId: in.UserId,
		Description: sql.NullString{
			String: in.Description,
			Valid:  true,
		},
	}
	result, err := l.svcCtx.VideoModel.Insert(l.ctx, &newVideo)
	if err != nil {
		return nil, status.Error(100, "视频存入数据库失败")
	}

	newVideo.Id, err = result.LastInsertId()
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	return &pb.UploadVideoResp{Success: true}, nil
}
