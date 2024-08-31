package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xinchat_gz/demo/internal/logic/user"
	"xinchat_gz/demo/internal/svc"
	"xinchat_gz/demo/internal/types"
)

func DelUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDelUserLogic(r.Context(), svcCtx)
		err := l.DelUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
