package homestayOrder

import (
	"net/http"

	"XavierLookLook/app/order/cmd/api/internal/logic/homestayOrder"
	"XavierLookLook/app/order/cmd/api/internal/svc"
	"XavierLookLook/app/order/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户订单列表
func UserHomestayOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestayOrder.NewUserHomestayOrderListLogic(r.Context(), svcCtx)
		resp, err := l.UserHomestayOrderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
