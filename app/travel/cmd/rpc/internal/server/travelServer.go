// Code generated by goctl. DO NOT EDIT.
// Source: travel.proto

package server

import (
	"context"

	"XavierLookLook/app/travel/cmd/rpc/internal/logic"
	"XavierLookLook/app/travel/cmd/rpc/internal/svc"
	"XavierLookLook/app/travel/cmd/rpc/pb"
)

type TravelServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTravelServer
}

func NewTravelServer(svcCtx *svc.ServiceContext) *TravelServer {
	return &TravelServer{
		svcCtx: svcCtx,
	}
}

func (s *TravelServer) HomestayList(ctx context.Context, in *pb.HomestayListReq) (*pb.HomestayListResp, error) {
	l := logic.NewHomestayListLogic(ctx, s.svcCtx)
	return l.HomestayList(in)
}

func (s *TravelServer) BusinessList(ctx context.Context, in *pb.BusinessListReq) (*pb.BusinessListResp, error) {
	l := logic.NewBusinessListLogic(ctx, s.svcCtx)
	return l.BusinessList(in)
}

func (s *TravelServer) GuessList(ctx context.Context, in *pb.GuessListReq) (*pb.GoodBossResp, error) {
	l := logic.NewGuessListLogic(ctx, s.svcCtx)
	return l.GuessList(in)
}

func (s *TravelServer) HomestayDetail(ctx context.Context, in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	l := logic.NewHomestayDetailLogic(ctx, s.svcCtx)
	return l.HomestayDetail(in)
}

func (s *TravelServer) GoodBoss(ctx context.Context, in *pb.GoodBossReq) (*pb.GoodBossResp, error) {
	l := logic.NewGoodBossLogic(ctx, s.svcCtx)
	return l.GoodBoss(in)
}

func (s *TravelServer) HomestayBussinessList(ctx context.Context, in *pb.HomestayBussinessListReq) (*pb.HomestayBussinessListResp, error) {
	l := logic.NewHomestayBussinessListLogic(ctx, s.svcCtx)
	return l.HomestayBussinessList(in)
}

func (s *TravelServer) HomestayBussinessDetail(ctx context.Context, in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	l := logic.NewHomestayBussinessDetailLogic(ctx, s.svcCtx)
	return l.HomestayBussinessDetail(in)
}

func (s *TravelServer) CommentList(ctx context.Context, in *pb.CommentListReq) (*pb.CommentListResp, error) {
	l := logic.NewCommentListLogic(ctx, s.svcCtx)
	return l.CommentList(in)
}
