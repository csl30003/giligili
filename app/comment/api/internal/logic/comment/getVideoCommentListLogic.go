package comment

import (
	"context"
	"giligili/app/comment/rpc/comment"

	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentListLogic {
	return &GetVideoCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetVideoCommentList 获取视频评论列表
func (l *GetVideoCommentListLogic) GetVideoCommentList(req *types.GetVideoCommentListReq) (resp *types.GetVideoCommentListResp, err error) {
	getVideoCommentListResp, err := l.svcCtx.CommentRpcClient.GetVideoCommentList(l.ctx, &comment.GetVideoCommentListReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	var videoCommentList []types.VideoComment
	for _, videoComment := range getVideoCommentListResp.VideoCommentList {
		videoCommentList = append(videoCommentList, types.VideoComment{
			CommentId:    videoComment.CommentId,
			UserId:       videoComment.UserId,
			Content:      videoComment.Content,
			LikeCount:    videoComment.LikeCount,
			DislikeCount: videoComment.DislikeCount,
			ReplyCount:   videoComment.ReplyCount,
			CreateTime:   videoComment.CreateTime,
			UpdateTime:   videoComment.UpdateTime,
		})
	}

	return &types.GetVideoCommentListResp{VideoCommentList: videoCommentList}, nil
}
