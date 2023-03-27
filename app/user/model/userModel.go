package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		FindOneById(ctx context.Context, id int64) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

// FindOneByUsername 通过用户名得到用户
func (c *customUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `username` = ? and delete_time is null limit 1", userRows, c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOneById 通过 id 得到用户，与 FindOne 的差别就是加个 and delete_time is null，因为我的表有这个属性
func (c *customUserModel) FindOneById(ctx context.Context, id int64) (*User, error) {
	giligiliUserIdKey := fmt.Sprintf("%s%v", cacheGiligiliUserIdPrefix, id)
	var resp User
	err := c.QueryRowCtx(ctx, &resp, giligiliUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and delete_time is null limit 1", userRows, c.table)
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
