package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHomestayActivityLogic {
	return &SearchHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchHomestayActivityLogic) SearchHomestayActivity(in *pb.SearchHomestayActivityReq) (*pb.SearchHomestayActivityResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchHomestayActivityResp{}, nil
}
