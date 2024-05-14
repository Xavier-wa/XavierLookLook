package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayBusinessLogic {
	return &AddHomestayBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------民宿店铺-----------------------
func (l *AddHomestayBusinessLogic) AddHomestayBusiness(in *pb.AddHomestayBusinessReq) (*pb.AddHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddHomestayBusinessResp{}, nil
}
