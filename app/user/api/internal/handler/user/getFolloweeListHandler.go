package user

import (
	"giligili/common/response"
	"net/http"

	"giligili/app/user/api/internal/logic/user"
	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFolloweeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFolloweeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetFolloweeListLogic(r.Context(), svcCtx)
		resp, err := l.GetFolloweeList(&req)
		response.Response(w, resp, err)
	}
}
