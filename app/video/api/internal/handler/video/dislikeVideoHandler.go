package video

import (
	"giligili/common/response"
	"net/http"

	"giligili/app/video/api/internal/logic/video"
	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DislikeVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DislikeVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewDislikeVideoLogic(r.Context(), svcCtx)
		resp, err := l.DislikeVideo(&req)
		response.Response(w, resp, err)
	}
}
