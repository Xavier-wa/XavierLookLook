package homestay

import (
	"context"

	"XavierLookLook/app/travel/cmd/api/internal/svc"
	"XavierLookLook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room list
func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (resp *types.HomestayListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
