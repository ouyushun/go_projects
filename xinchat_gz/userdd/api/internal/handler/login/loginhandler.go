package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xinchat_gz/userdd/api/internal/logic/login"
	"xinchat_gz/userdd/api/internal/svc"
	"xinchat_gz/userdd/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq

		if err := httpx.Parse(r, &req); err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
