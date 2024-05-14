package logic

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHomestayCommentLogic {
	return &SearchHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchHomestayCommentLogic) SearchHomestayComment(in *pb.SearchHomestayCommentReq) (*pb.SearchHomestayCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchHomestayCommentResp{}, nil
}
