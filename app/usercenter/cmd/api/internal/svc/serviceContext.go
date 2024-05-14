package svc

import (
	"XavierLookLook/app/usercenter/cmd/api/internal/config"
	"XavierLookLook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	usercenter usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		usercenter: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
