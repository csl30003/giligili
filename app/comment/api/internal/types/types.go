// Code generated by goctl. DO NOT EDIT.
package types

type GetVideoCommentListReq struct {
	VideoId int64 `json:"videoId"`
}

type VideoComment struct {
	CommentId    int64  `json:"commentId"`
	UserId       int64  `json:"userId"`
	Content      string `json:"content"`
	LikeCount    int64  `json:"likeCount"`
	DislikeCount int64  `json:"dislikeCount"`
	ReplyCount   int64  `json:"replyCount"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type GetVideoCommentListResp struct {
	VideoCommentList []VideoComment `json:"commentList"`
}

type GetCommentReplyListReq struct {
	CommentId int64 `json:"commentId"`
}

type CommentReply struct {
	CommentId    int64  `json:"commentId"`
	UserId       int64  `json:"userId"`
	Content      string `json:"content"`
	LikeCount    int64  `json:"likeCount"`
	DislikeCount int64  `json:"dislikeCount"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type GetCommentReplyListResp struct {
	CommentReplyList []CommentReply `json:"commentReplyList"`
}

type CommentVideoReq struct {
	VideoId int64  `json:"videoId"`
	Content string `json:"content"`
}

type CommentVideoResp struct {
	Success bool `json:"success"`
}

type ReplyCommentReq struct {
	CommentId int64  `json:"commentId"`
	Content   string `json:"content"`
}

type ReplyCommentResp struct {
	Success bool `json:"success"`
}

type DeleteCommentReq struct {
	CommentId int64 `json:"commentId"`
}

type DeleteCommentResp struct {
	Success bool `json:"success"`
}
