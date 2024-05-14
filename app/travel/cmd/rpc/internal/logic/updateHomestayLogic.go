package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayLogic {
	return &UpdateHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayLogic) UpdateHomestay(in *pb.UpdateHomestayReq) (*pb.UpdateHomestayResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateHomestayResp{}, nil
}
