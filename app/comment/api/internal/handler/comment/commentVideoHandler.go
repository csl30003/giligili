package comment

import (
	"giligili/common/response"
	"net/http"

	"giligili/app/comment/api/internal/logic/comment"
	"giligili/app/comment/api/internal/svc"
	"giligili/app/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewCommentVideoLogic(r.Context(), svcCtx)
		resp, err := l.CommentVideo(&req)
		response.Response(w, resp, err)
	}
}
