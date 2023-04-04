package logic

import (
	"context"
	"giligili/app/comment/model"
	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"
	"giligili/app/video/rpc/video"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentVideoLogic {
	return &CommentVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommentVideo 评论视频
func (l *CommentVideoLogic) CommentVideo(in *pb.CommentVideoReq) (*pb.CommentVideoResp, error) {
	// 判断视频是否存在
	isExistResp, err := l.svcCtx.VideoRpcClient.IsExistVideo(l.ctx, &video.IsExistVideoReq{
		VideoId: in.VideoId,
	})
	if err != nil {
		return nil, status.Error(100, "查询视频失败")
	}
	if !isExistResp.Exists {
		return nil, status.Error(100, "视频不存在")
	}

	// 插入记录
	newComment := model.Comment{
		UserId:  in.UserId,
		VideoId: in.VideoId,
		Content: in.Content,
	}
	result, err := l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
	if err != nil {
		return nil, status.Error(100, "评论存储数据库失败")
	}
	newComment.Id, err = result.LastInsertId()
	if err != nil {
		return nil, status.Error(100, err.Error())
	}

	return &pb.CommentVideoResp{Success: true}, nil
}
