package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	WxMiniConf WxMiniConf
	WxPayConf WxPayConf
	OrderRpcConf zrpc.RpcClientConf
	PaymentRpcConf zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf

}
