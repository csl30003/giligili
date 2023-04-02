package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BarrageModel = (*customBarrageModel)(nil)

type (
	// BarrageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBarrageModel.
	BarrageModel interface {
		barrageModel
		FindManyByVideoIdAndLimit(ctx context.Context, videoId, limit int64) ([]*Barrage, error)
	}

	customBarrageModel struct {
		*defaultBarrageModel
	}
)

// NewBarrageModel returns a model for the database table.
func NewBarrageModel(conn sqlx.SqlConn, c cache.CacheConf) BarrageModel {
	return &customBarrageModel{
		defaultBarrageModel: newBarrageModel(conn, c),
	}
}

// FindManyByVideoIdAndLimit 根据视频 id 和限制条数查询弹幕列表
func (c *customBarrageModel) FindManyByVideoIdAndLimit(ctx context.Context, videoId, limit int64) ([]*Barrage, error) {
	var resp []*Barrage
	query := fmt.Sprintf("select %s from %s where video_id = ? and delete_time is null order by timestamp desc limit ?", barrageRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, videoId, limit)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
