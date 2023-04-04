package comment

import (
	"context"
	"giligili/app/comment/rpc/comment"
	"giligili/common/ctxData"

	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentVideoLogic {
	return &CommentVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CommentVideo 评论视频
func (l *CommentVideoLogic) CommentVideo(req *types.CommentVideoReq) (resp *types.CommentVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	commentVideoResp, err := l.svcCtx.CommentRpcClient.CommentVideo(l.ctx, &comment.CommentVideoReq{
		UserId:  userId,
		VideoId: req.VideoId,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.CommentVideoResp{Success: commentVideoResp.Success}, nil
}
