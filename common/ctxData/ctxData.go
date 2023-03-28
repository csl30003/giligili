package ctxData

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "userId"

func GetUserIdFromCtx(ctx context.Context) int64 {
	var userId int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			userId = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("获取用户 id 失败: %+v", err)
		}
	}

	return userId
}
