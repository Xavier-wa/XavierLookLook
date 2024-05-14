package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayBusinessLogic {
	return &UpdateHomestayBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayBusinessLogic) UpdateHomestayBusiness(in *pb.UpdateHomestayBusinessReq) (*pb.UpdateHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateHomestayBusinessResp{}, nil
}
