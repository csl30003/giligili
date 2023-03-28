// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "giligili/app/user/api/internal/handler/user"
	"giligili/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userInfo",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/followerList",
				Handler: user.GetFollowerListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/followeeList",
				Handler: user.GetFolloweeListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/followUser",
				Handler: user.FollowUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/unfollowUser",
				Handler: user.UnfollowUserHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user/v1"),
	)
}
