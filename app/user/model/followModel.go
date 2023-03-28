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

func (c *customFollowModel) FindOneByFollowerIdAndFolloweeId(ctx context.Context, followerId, followeeId int64) (*Follow, error) {
	var resp Follow
	query := fmt.Sprintf("select %s from %s where `follower_id` = ? and `followee_id` = ? and delete_time is null limit 1", followRows, c.table)
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
