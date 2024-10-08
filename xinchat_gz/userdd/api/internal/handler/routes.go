// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	login "xinchat_gz/userdd/api/internal/handler/login"
	user "xinchat_gz/userdd/api/internal/handler/user"
	"xinchat_gz/userdd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/userdd/login",
				Handler: login.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Loginverification},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/userdd/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/userdd/info/update",
					Handler: user.UpdateUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
	)
}
