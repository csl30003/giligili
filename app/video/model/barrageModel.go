package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BarrageModel = (*customBarrageModel)(nil)

type (
	// BarrageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBarrageModel.
	BarrageModel interface {
		barrageModel
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
