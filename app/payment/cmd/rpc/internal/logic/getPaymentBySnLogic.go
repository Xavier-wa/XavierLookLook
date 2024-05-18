package logic

import (
	"context"

	"XavierLookLook/app/payment/cmd/rpc/internal/svc"
	"XavierLookLook/app/payment/cmd/rpc/pb"
	"XavierLookLook/app/payment/model"
	"XavierLookLook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentBySnLogic {
	return &GetPaymentBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据sn查询流水记录
func (l *GetPaymentBySnLogic) GetPaymentBySn(in *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {
	// todo: add your logic here and delete this line
	thirdPayment, err := l.svcCtx.ThirdPaymentModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetPaymentBySn  FindOneBySn  db err:%v , in : %+v", err, in)
	}

	var resp pb.PaymentDetail
	if thirdPayment != nil {
		_ = copier.Copy(&resp, thirdPayment)

		resp.CreateTime = thirdPayment.CreateTime.Unix()
		resp.UpdateTime = thirdPayment.UpdateTime.Unix()
		resp.PayTime = thirdPayment.PayTime.Unix()
	}

	return &pb.GetPaymentBySnResp{
		PaymentDetail: &resp,
	}, nil
}
