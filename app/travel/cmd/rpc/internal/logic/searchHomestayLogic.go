package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHomestayLogic {
	return &SearchHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchHomestayLogic) SearchHomestay(in *pb.SearchHomestayReq) (*pb.SearchHomestayResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchHomestayResp{}, nil
}
