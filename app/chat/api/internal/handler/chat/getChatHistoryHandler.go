package chat

import (
	"giligili/common/response"
	"net/http"

	"giligili/app/chat/api/internal/logic/chat"
	"giligili/app/chat/api/internal/svc"
	"giligili/app/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetChatHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChatHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewGetChatHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetChatHistory(&req)
		response.Response(w, resp, err)
	}
}
