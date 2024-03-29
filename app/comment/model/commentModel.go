package model

import (
	"context"
	"fmt"
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
		FindOneById(ctx context.Context, id int64) (*Comment, error)
		FindManyByVideoId(ctx context.Context, videoId int64) ([]*VideoCommentTemp, error)
		FindManyByCommentId(ctx context.Context, commentId int64) ([]*Comment, error)
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

// FindOneById 通过 id 获取评论记录
func (c *customCommentModel) FindOneById(ctx context.Context, id int64) (*Comment, error) {
	giligiliCommentIdKey := fmt.Sprintf("%s%v", cacheGiligiliCommentIdPrefix, id)
	var resp Comment
	err := c.QueryRowCtx(ctx, &resp, giligiliCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and delete_time is null limit 1", commentRows, c.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindManyByVideoId 根据视频 id 获取评论列表
func (c *customCommentModel) FindManyByVideoId(ctx context.Context, videoId int64) ([]*VideoCommentTemp, error) {
	var resp []*VideoCommentTemp
	query := "select c.id, c.user_id, c.content, c.like, c.dislike, COUNT(reply.id) as reply_count, c.create_time, c.update_time from comment as c left join comment as reply on c.id = reply.reply_id where c.video_id = ? and c.reply_id is null and c.delete_time is null group by c.id"
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

// FindManyByCommentId 通过评论 id 获取评论列表
func (c *customCommentModel) FindManyByCommentId(ctx context.Context, commentId int64) ([]*Comment, error) {
	var resp []*Comment
	query := fmt.Sprintf("select %s from %s where `reply_id` = ? and delete_time is null", commentRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, commentId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
