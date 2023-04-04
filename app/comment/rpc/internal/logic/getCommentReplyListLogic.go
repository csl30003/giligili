package logic

import (
	"context"
	"giligili/app/comment/model"
	"google.golang.org/grpc/status"

	"giligili/app/comment/rpc/internal/svc"
	"giligili/app/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyListLogic {
	return &GetCommentReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCommentReplyList 获取评论评论列表
func (l *GetCommentReplyListLogic) GetCommentReplyList(in *pb.GetCommentReplyListReq) (*pb.GetCommentReplyListResp, error) {
	commentReplyList, err := l.svcCtx.CommentModel.FindManyByCommentId(l.ctx, in.CommentId)
	if err == model.ErrNotFound {
		return &pb.GetCommentReplyListResp{}, nil
	}
	if err != nil {
		return nil, status.Error(100, "获取评论失败")
	}

	var commentReplies []*pb.CommentReply
	for _, commentReply := range commentReplyList {
		commentReplies = append(commentReplies, &pb.CommentReply{
			CommentId:    commentReply.Id,
			UserId:       commentReply.UserId,
			Content:      commentReply.Content,
			LikeCount:    commentReply.Like,
			DislikeCount: commentReply.Dislike,
			CreateTime:   commentReply.CreateTime.String(),
			UpdateTime:   commentReply.UpdateTime.String(),
		})
	}

	return &pb.GetCommentReplyListResp{CommentReplyList: commentReplies}, nil
}
