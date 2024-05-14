package homestay

import (
	"context"

	"XavierLookLook/app/travel/cmd/api/internal/svc"
	"XavierLookLook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GuessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// guess homestay room
func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuessListLogic {
	return &GuessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuessListLogic) GuessList(req *types.GuessListReq) (resp *types.GuessListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
