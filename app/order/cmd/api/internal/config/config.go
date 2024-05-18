package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	OrderRpcConf   zrpc.RpcClientConf
	PaymentRpcConf zrpc.RpcClientConf
	TravelRpcConf  zrpc.RpcClientConf
}
