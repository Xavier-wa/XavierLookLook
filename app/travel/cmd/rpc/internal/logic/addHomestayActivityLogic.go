package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayActivityLogic {
	return &AddHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------每一间民宿-----------------------
func (l *AddHomestayActivityLogic) AddHomestayActivity(in *pb.AddHomestayActivityReq) (*pb.AddHomestayActivityResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddHomestayActivityResp{}, nil
}
