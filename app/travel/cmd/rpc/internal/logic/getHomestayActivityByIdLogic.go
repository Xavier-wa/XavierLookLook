package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHomestayActivityByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHomestayActivityByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHomestayActivityByIdLogic {
	return &GetHomestayActivityByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHomestayActivityByIdLogic) GetHomestayActivityById(in *pb.GetHomestayActivityByIdReq) (*pb.GetHomestayActivityByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetHomestayActivityByIdResp{}, nil
}
