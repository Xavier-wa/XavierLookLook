package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayActivityLogic {
	return &UpdateHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayActivityLogic) UpdateHomestayActivity(in *pb.UpdateHomestayActivityReq) (*pb.UpdateHomestayActivityResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateHomestayActivityResp{}, nil
}
