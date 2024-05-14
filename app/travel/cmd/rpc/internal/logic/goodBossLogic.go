package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodBossLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodBossLogic) GoodBoss(in *pb.GoodBossReq) (*pb.GoodBossResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodBossResp{}, nil
}
