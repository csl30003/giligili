package video

import (
	"giligili/app/video/api/internal/logic/video"
	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"
	"giligili/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UploadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadVideoReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 这里把 r 传过去
		l := video.NewUploadVideoLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadVideo(&req)
		response.Response(w, resp, err)
	}
}
