package homestayOrder

import (
	"context"

	"XavierLookLook/app/order/cmd/api/internal/svc"
	"XavierLookLook/app/order/cmd/api/internal/types"
	"XavierLookLook/app/order/cmd/rpc/order"
	"XavierLookLook/app/order/model"
	"XavierLookLook/app/payment/cmd/rpc/payment"
	"XavierLookLook/common/ctxdata"
	"XavierLookLook/common/tool"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户订单明细
func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderDetailLogic {
	return &UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req *types.UserHomestayOrderDetailReq) (resp *types.UserHomestayOrderDetailResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)

	orderRpcResp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: req.Sn,
	})

	var typesOrderDetail types.UserHomestayOrderDetailResp
	if orderRpcResp.HomestayOrder != nil && orderRpcResp.HomestayOrder.UserId == userId {
		copier.Copy(&typesOrderDetail, orderRpcResp.HomestayOrder)
		//重置价格.
		typesOrderDetail.OrderTotalPrice = tool.Fen2Yuan(orderRpcResp.HomestayOrder.OrderTotalPrice)
		typesOrderDetail.FoodTotalPrice = tool.Fen2Yuan(orderRpcResp.HomestayOrder.FoodTotalPrice)
		typesOrderDetail.HomestayTotalPrice = tool.Fen2Yuan(orderRpcResp.HomestayOrder.HomestayTotalPrice)
		typesOrderDetail.HomestayPrice = tool.Fen2Yuan(orderRpcResp.HomestayOrder.HomestayPrice)
		typesOrderDetail.FoodPrice = tool.Fen2Yuan(orderRpcResp.HomestayOrder.FoodPrice)
		typesOrderDetail.MarketHomestayPrice = tool.Fen2Yuan(orderRpcResp.HomestayOrder.MarketHomestayPrice)

		//支付信息.
		if typesOrderDetail.TradeState != model.HomestayOrderTradeStateCancel && typesOrderDetail.TradeState != model.HomestayOrderTradeStateWaitPay {
			paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentSuccessRefundByOrderSn(l.ctx, &payment.GetPaymentSuccessRefundByOrderSnReq{
				OrderSn: orderRpcResp.HomestayOrder.Sn,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Failed to get order payment information err : %v , orderSn:%s", err, orderRpcResp.HomestayOrder.Sn)
			}

			if paymentResp.PaymentDetail != nil {
				typesOrderDetail.PayTime = paymentResp.PaymentDetail.PayTime
				typesOrderDetail.PayType = paymentResp.PaymentDetail.PayMode
			}
		}

		return &typesOrderDetail, nil

	}
	return
}
