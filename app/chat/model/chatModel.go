package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ ChatModel = (*customChatModel)(nil)

// ChatTemp 为了得到连表查询的结果所创建的结构体
type ChatTemp struct {
	Id         int64        `db:"id"`           // 聊天ID
	CreateTime time.Time    `db:"create_time"`  // 创建时间
	DeleteTime sql.NullTime `db:"delete_time"`  // 删除时间
	FromUserId int64        `db:"from_user_id"` // 发送者ID
	ToUserId   int64        `db:"to_user_id"`   // 接收者ID
	Content    string       `db:"content"`      // 聊天内容

	FromUserNickname string `db:"from_user_nickname"` // 发送者昵称
	ToUserNickname   string `db:"to_user_nickname"`   // 接收者昵称
}

type (
	// ChatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatModel.
	ChatModel interface {
		chatModel
		FindManyByFromUserIdAndToUserId(ctx context.Context, fromUserId, toUserId int64) ([]*ChatTemp, error)
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

// FindManyByFromUserIdAndToUserId 通过双方 id 获取聊天历史记录
func (c *customChatModel) FindManyByFromUserIdAndToUserId(ctx context.Context, fromUserId, toUserId int64) ([]*ChatTemp, error) {
	chatList := make([]*ChatTemp, 0)
	query := "select chat.id, chat.create_time, chat.delete_time, chat.from_user_id, chat.to_user_id, chat.content, uf.nickname as from_user_nickname, ut.nickname as to_user_nickname from chat left join user uf on chat.from_user_id=uf.id left join user ut on chat.to_user_id=ut.id where chat.from_user_id = ? and chat.to_user_id = ? and chat.delete_time is null"
	err := c.QueryRowsNoCacheCtx(ctx, &chatList, query, fromUserId, toUserId)
	switch err {
	case nil:
		return chatList, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
