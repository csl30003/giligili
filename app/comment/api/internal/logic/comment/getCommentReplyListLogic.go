package comment

import (
	"context"
	"giligili/app/comment/rpc/comment"

	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyListLogic {
	return &GetCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetCommentReplyList 获取评论评论列表
func (l *GetCommentReplyListLogic) GetCommentReplyList(req *types.GetCommentReplyListReq) (resp *types.GetCommentReplyListResp, err error) {
	getCommentReplyListResp, err := l.svcCtx.CommentRpcClient.GetCommentReplyList(l.ctx, &comment.GetCommentReplyListReq{
		CommentId: req.CommentId,
	})
	if err != nil {
		return nil, err
	}

	var commentReplyList []types.CommentReply
	for _, commentReply := range getCommentReplyListResp.CommentReplyList {
		commentReplyList = append(commentReplyList, types.CommentReply{
			CommentId:    commentReply.CommentId,
			UserId:       commentReply.UserId,
			Content:      commentReply.Content,
			LikeCount:    commentReply.LikeCount,
			DislikeCount: commentReply.DislikeCount,
			CreateTime:   commentReply.CreateTime,
			UpdateTime:   commentReply.UpdateTime,
		})
	}

	return &types.GetCommentReplyListResp{CommentReplyList: commentReplyList}, nil
}
