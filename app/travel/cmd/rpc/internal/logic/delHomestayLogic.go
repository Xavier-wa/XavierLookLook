package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelHomestayLogic {
	return &DelHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelHomestayLogic) DelHomestay(in *pb.DelHomestayReq) (*pb.DelHomestayResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelHomestayResp{}, nil
}
