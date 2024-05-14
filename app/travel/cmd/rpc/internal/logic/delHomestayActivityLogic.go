package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelHomestayActivityLogic {
	return &DelHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelHomestayActivityLogic) DelHomestayActivity(in *pb.DelHomestayActivityReq) (*pb.DelHomestayActivityResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelHomestayActivityResp{}, nil
}
