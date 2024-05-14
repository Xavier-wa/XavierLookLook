package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHomestayBusinessByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHomestayBusinessByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHomestayBusinessByIdLogic {
	return &GetHomestayBusinessByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHomestayBusinessByIdLogic) GetHomestayBusinessById(in *pb.GetHomestayBusinessByIdReq) (*pb.GetHomestayBusinessByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetHomestayBusinessByIdResp{}, nil
}
