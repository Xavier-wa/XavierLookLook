package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayBussinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBussinessDetailLogic {
	return &HomestayBussinessDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayBussinessDetailLogic) HomestayBussinessDetail(in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.HomestayDetailResp{}, nil
}
