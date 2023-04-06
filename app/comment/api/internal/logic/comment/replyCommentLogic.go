package comment

import (
	"context"
	"giligili/app/comment/rpc/comment"
	"giligili/common/ctxData"

	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommentLogic {
	return &ReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReplyComment 回复评论
func (l *ReplyCommentLogic) ReplyComment(req *types.ReplyCommentReq) (resp *types.ReplyCommentResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	replyCommentResp, err := l.svcCtx.CommentRpcClient.ReplyComment(l.ctx, &comment.ReplyCommentReq{
		UserId:    userId,
		CommentId: req.CommentId,
		Content:   req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.ReplyCommentResp{Success: replyCommentResp.Success}, nil
}
