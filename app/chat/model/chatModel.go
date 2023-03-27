package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatModel = (*customChatModel)(nil)

type (
	// ChatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatModel.
	ChatModel interface {
		chatModel
	}

	customChatModel struct {
		*defaultChatModel
	}
)

// NewChatModel returns a model for the database table.
func NewChatModel(conn sqlx.SqlConn, c cache.CacheConf) ChatModel {
	return &customChatModel{
		defaultChatModel: newChatModel(conn, c),
	}
}
