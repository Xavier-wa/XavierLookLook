package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
