package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHomestayCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHomestayCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHomestayCommentByIdLogic {
	return &GetHomestayCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHomestayCommentByIdLogic) GetHomestayCommentById(in *pb.GetHomestayCommentByIdReq) (*pb.GetHomestayCommentByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetHomestayCommentByIdResp{}, nil
}
