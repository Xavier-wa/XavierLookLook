package user

import (
	"context"

	"XavierLookLook/app/usercenter/cmd/api/internal/svc"
	"XavierLookLook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxWiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// wechat mini auth
func NewWxWiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxWiniAuthLogic {
	return &WxWiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxWiniAuthLogic) WxWiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
