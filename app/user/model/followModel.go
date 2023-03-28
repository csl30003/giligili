package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		FindOneByFollowerIdAndFolloweeId(ctx context.Context, followerId, followeeId int64) (*Follow, error)
		FindManyByFolloweeId(ctx context.Context, followeeId int64) ([]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn, c cache.CacheConf) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn, c),
	}
}

// FindOneByFollowerIdAndFolloweeId 通过 FollowerId 和 FolloweeId 获取关注记录
func (c *customFollowModel) FindOneByFollowerIdAndFolloweeId(ctx context.Context, followerId, followeeId int64) (*Follow, error) {
	var resp Follow
	query := fmt.Sprintf("select %s from %s where `follower_id` = ? and `followee_id` = ? and `delete_time` is null limit 1", followRows, c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, followerId, followeeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindManyByFolloweeId 通过被关注者 id 获取关注记录列表
func (c *customFollowModel) FindManyByFolloweeId(ctx context.Context, followeeId int64) ([]*Follow, error) {
	followList := make([]*Follow, 0)
	query := fmt.Sprintf("select %s from %s where `followee_id` = ? and `delete_time` is null", followRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &followList, query, followeeId)
	switch err {
	case nil:
		return followList, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
