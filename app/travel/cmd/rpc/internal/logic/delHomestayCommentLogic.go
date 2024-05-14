package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelHomestayCommentLogic {
	return &DelHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelHomestayCommentLogic) DelHomestayComment(in *pb.DelHomestayCommentReq) (*pb.DelHomestayCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelHomestayCommentResp{}, nil
}
