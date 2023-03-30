package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ VideoModel = (*customVideoModel)(nil)

type VideoTemp struct {
	Id          int64          `db:"id"`          // 视频ID
	CreateTime  time.Time      `db:"create_time"` // 创建时间
	UpdateTime  time.Time      `db:"update_time"` // 更新时间
	Title       string         `db:"title"`       // 视频标题
	Url         string         `db:"url"`         // 视频URL
	UserId      int64          `db:"user_id"`     // 上传用户ID
	Description sql.NullString `db:"description"` // 视频描述
	Like        int64          `db:"like"`        // 点赞数
	Dislike     int64          `db:"dislike"`     // 踩数
	Status      int64          `db:"status"`      // 审核状态，0表示未审核，1表示已审核

	Nickname string `db:"nickname"` // 视频作者昵称
}

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindOneById(ctx context.Context, id int64) (*VideoTemp, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c),
	}
}

// FindOneById 通过 id 得到视频
func (c *customVideoModel) FindOneById(ctx context.Context, id int64) (*VideoTemp, error) {
	giligiliVideoIdKey := fmt.Sprintf("%s%v", cacheGiligiliVideoIdPrefix, id)
	var resp VideoTemp
	err := c.QueryRowCtx(ctx, &resp, giligiliVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := "select v.id, v.create_time, v.update_time, v.title, v.url, v.user_id, v.description, v.like, v.dislike, v.status, u.nickname from video v left join user u on v.user_id=u.id where v.id = ? and v.delete_time is null limit 1"
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
