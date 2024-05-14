package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayCommentLogic {
	return &UpdateHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayCommentLogic) UpdateHomestayComment(in *pb.UpdateHomestayCommentReq) (*pb.UpdateHomestayCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateHomestayCommentResp{}, nil
}
