package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xinchat_gz/demo/internal/logic/login"
	"xinchat_gz/demo/internal/svc"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
