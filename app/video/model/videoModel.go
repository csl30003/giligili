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

var (
	_                              VideoModel = (*customVideoModel)(nil)
	cacheGiligiliVideoTempIdPrefix            = "cache:giligili:video:Temp:id:"
)

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

	Nickname string `db:"nickname"` // 视频作者昵称
}

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindOneById(ctx context.Context, id int64) (*Video, error)
		FindOneByIdAndIgnoreStatus(ctx context.Context, id int64) (*Video, error)
		FindVideoTempById(ctx context.Context, id int64) (*VideoTemp, error)
		FindManyByKeyword(ctx context.Context, page, pageSize int64, keyword string) ([]*VideoTemp, error)
		FindMany(ctx context.Context) ([]*Video, error)
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

// FindOneById 通过 id 获取视频
func (c *customVideoModel) FindOneById(ctx context.Context, id int64) (*Video, error) {
	giligiliVideoIdKey := fmt.Sprintf("%s%v", cacheGiligiliVideoIdPrefix, id)
	var resp Video
	err := c.QueryRowCtx(ctx, &resp, giligiliVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and delete_time is null and status=1 limit 1", videoRows, c.table)
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

// FindOneByIdAndIgnoreStatus 通过 id 获取视频，并且无论状态是如何
func (c *customVideoModel) FindOneByIdAndIgnoreStatus(ctx context.Context, id int64) (*Video, error) {
	giligiliVideoIdKey := fmt.Sprintf("%s%v", cacheGiligiliVideoIdPrefix, id)
	var resp Video
	err := c.QueryRowCtx(ctx, &resp, giligiliVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and delete_time is null limit 1", videoRows, c.table)
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

// FindVideoTempById 通过 id 得到视频
func (c *customVideoModel) FindVideoTempById(ctx context.Context, id int64) (*VideoTemp, error) {
	giligiliVideoTempIdKey := fmt.Sprintf("%s%v", cacheGiligiliVideoTempIdPrefix, id)
	var resp VideoTemp
	err := c.QueryRowCtx(ctx, &resp, giligiliVideoTempIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := "select v.id, v.create_time, v.update_time, v.title, v.url, v.user_id, v.description, v.like, v.dislike, u.nickname from video v left join user u on v.user_id=u.id where v.id = ? and v.delete_time is null and v.status=1 limit 1"
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

// FindManyByKeyword 通过页面、页面大小、模糊查询值获取视频列表
func (c *customVideoModel) FindManyByKeyword(ctx context.Context, page, pageSize int64, keyword string) ([]*VideoTemp, error) {
	var resp []*VideoTemp
	query := "select v.id, v.create_time, v.update_time, v.title, v.url, v.user_id, u.nickname, v.description, v.like, v.dislike from video v left join user u on v.user_id=u.id where v.delete_time is null and v.status=1"
	if keyword != "" {
		query += " and v.title like '%" + keyword + "%'"
	}
	query += " order by v.id desc limit ?,?"
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*pageSize, pageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindMany 查询所有的视频
func (c *customVideoModel) FindMany(ctx context.Context) ([]*Video, error) {
	var resp []*Video
	query := fmt.Sprintf("select %s from %s where delete_time is null and status=1", videoRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
