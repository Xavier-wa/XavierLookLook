package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayLogic {
	return &AddHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------每一间民宿-----------------------
func (l *AddHomestayLogic) AddHomestay(in *pb.AddHomestayReq) (*pb.AddHomestayResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddHomestayResp{}, nil
}
