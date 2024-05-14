package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayBussinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBussinessListLogic {
	return &HomestayBussinessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayBussinessListLogic) HomestayBussinessList(in *pb.HomestayBussinessListReq) (*pb.HomestayBussinessListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.HomestayBussinessListResp{}, nil
}
