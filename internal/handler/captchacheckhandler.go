package handler

import (
	"net/http"

	"captchaImageServer/internal/logic"
	"captchaImageServer/internal/svc"
	"captchaImageServer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaCheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCaptchaCheckLogic(r.Context(), svcCtx)
		resp, err := l.CaptchaCheck(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
