package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	DB struct {
		DataSource string
	}
	Cache cache.CacheConf

	UsercenterRpcConf zrpc.RpcClientConf
	TravelRpcConf     zrpc.RpcClientConf
}
