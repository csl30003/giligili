package logic

import (
	"context"
	"giligili/app/comment/model"
	"google.golang.org/grpc/status"

	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentListLogic {
	return &GetVideoCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetVideoCommentList 获取视频评论列表
func (l *GetVideoCommentListLogic) GetVideoCommentList(in *pb.GetVideoCommentListReq) (*pb.GetVideoCommentListResp, error) {
	videoCommentList, err := l.svcCtx.CommentModel.FindManyByVideoId(l.ctx, in.VideoId)
	if err == model.ErrNotFound {
		return &pb.GetVideoCommentListResp{}, nil
	}
	if err != nil {
		return nil, status.Error(100, "获取评论失败")
	}

	var videoComments []*pb.VideoComment
	for _, videoComment := range videoCommentList {
		videoComments = append(videoComments, &pb.VideoComment{
			CommentId:    videoComment.Id,
			UserId:       videoComment.UserId,
			Content:      videoComment.Content,
			LikeCount:    videoComment.Like,
			DislikeCount: videoComment.Dislike,
			ReplyCount:   videoComment.ReplyCount,
			CreateTime:   videoComment.CreateTime.String(),
			UpdateTime:   videoComment.UpdateTime.String(),
		})
	}

	return &pb.GetVideoCommentListResp{VideoCommentList: videoComments}, nil
}
