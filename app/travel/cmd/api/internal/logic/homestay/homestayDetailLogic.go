package homestay

import (
	"context"

	"XavierLookLook/app/travel/cmd/api/internal/svc"
	"XavierLookLook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room detail
func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
