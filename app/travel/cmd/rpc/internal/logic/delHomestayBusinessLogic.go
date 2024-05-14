package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelHomestayBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelHomestayBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelHomestayBusinessLogic {
	return &DelHomestayBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelHomestayBusinessLogic) DelHomestayBusiness(in *pb.DelHomestayBusinessReq) (*pb.DelHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelHomestayBusinessResp{}, nil
}
