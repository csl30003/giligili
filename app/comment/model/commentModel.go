package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ CommentModel = (*customCommentModel)(nil)

type VideoCommentTemp struct {
	Id         int64     `db:"id"`          // 评论ID
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	UserId     int64     `db:"user_id"`     // 评论用户ID
	Content    string    `db:"content"`     // 评论内容
	Like       int64     `db:"like"`        // 点赞数
	Dislike    int64     `db:"dislike"`     // 踩数

	ReplyCount int64 `db:"reply_count"` // 被评论数量
}

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindManyByVideoId(ctx context.Context, videoId int64) ([]*VideoCommentTemp, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}

// FindManyByVideoId 根据视频 id 获取评论列表
func (c *customCommentModel) FindManyByVideoId(ctx context.Context, videoId int64) ([]*VideoCommentTemp, error) {
	var resp []*VideoCommentTemp
	query := "select c.id, c.user_id, c.content, c.like, c.dislike, COUNT(reply.id) as reply_count, c.create_time, c.update_time from comment as c left join comment as reply on c.id = reply.reply_id where c.video_id = ? and c.delete_time is null group by c.id"
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, videoId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
