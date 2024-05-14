package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHomestayByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHomestayByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHomestayByIdLogic {
	return &GetHomestayByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHomestayByIdLogic) GetHomestayById(in *pb.GetHomestayByIdReq) (*pb.GetHomestayByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetHomestayByIdResp{}, nil
}
