package svc

import (
	"XavierLookLook/app/order/cmd/rpc/order"
	"XavierLookLook/app/payment/cmd/api/internal/config"
	"XavierLookLook/app/payment/cmd/rpc/payment"
	"XavierLookLook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	PaymentRpc    payment.Payment
	OrderRpc      order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		PaymentRpc:    payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
	}
}
