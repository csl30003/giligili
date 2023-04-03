package video

import (
	"giligili/common/response"
	"net/http"

	"giligili/app/video/api/internal/logic/video"
	"giligili/app/video/api/internal/svc"
)

func GetHotVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := video.NewGetHotVideoListLogic(r.Context(), svcCtx)
		resp, err := l.GetHotVideoList()
		response.Response(w, resp, err)
	}
}
