package handler

import (
	"net/http"

	"captchaImageServer/internal/logic"
	"captchaImageServer/internal/svc"
	"captchaImageServer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaImageServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCaptchaImageServerLogic(r.Context(), svcCtx)
		resp, err := l.CaptchaImageServer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
