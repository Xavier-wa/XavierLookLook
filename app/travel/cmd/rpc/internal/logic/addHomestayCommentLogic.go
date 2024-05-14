package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayCommentLogic {
	return &AddHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------民宿评价-----------------------
func (l *AddHomestayCommentLogic) AddHomestayComment(in *pb.AddHomestayCommentReq) (*pb.AddHomestayCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddHomestayCommentResp{}, nil
}
