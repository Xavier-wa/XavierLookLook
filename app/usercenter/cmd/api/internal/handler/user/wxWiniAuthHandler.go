package user

import (
	"net/http"

	"XavierLookLook/app/usercenter/cmd/api/internal/logic/user"
	"XavierLookLook/app/usercenter/cmd/api/internal/svc"
	"XavierLookLook/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// wechat mini auth
func WxWiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewWxWiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxWiniAuth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
