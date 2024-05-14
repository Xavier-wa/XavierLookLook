package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHomestayBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchHomestayBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHomestayBusinessLogic {
	return &SearchHomestayBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchHomestayBusinessLogic) SearchHomestayBusiness(in *pb.SearchHomestayBusinessReq) (*pb.SearchHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchHomestayBusinessResp{}, nil
}
