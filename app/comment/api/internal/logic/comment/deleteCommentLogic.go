package comment

import (
	"context"
	"giligili/app/comment/rpc/comment"
	"giligili/common/ctxData"

	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteComment 删除评论
func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	deleteCommentResp, err := l.svcCtx.CommentRpcClient.DeleteComment(l.ctx, &comment.DeleteCommentReq{
		UserId:    userId,
		CommentId: req.CommentId,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteCommentResp{Success: deleteCommentResp.Success}, nil
}
