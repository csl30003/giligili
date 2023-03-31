package chat

import (
	"giligili/common/response"
	"net/http"

	"giligili/app/chat/api/internal/logic/chat"
	"giligili/app/chat/api/internal/svc"
	"giligili/app/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteChatMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteChatMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewDeleteChatMessageLogic(r.Context(), svcCtx)
		resp, err := l.DeleteChatMessage(&req)
		response.Response(w, resp, err)
	}
}
