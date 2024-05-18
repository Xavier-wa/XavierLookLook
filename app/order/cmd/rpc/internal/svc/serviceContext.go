package svc

import (
	"XavierLookLook/app/order/cmd/rpc/internal/config"
	"XavierLookLook/app/order/model"
	"XavierLookLook/app/travel/cmd/rpc/travel"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

// create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}

type ServiceContext struct {
	Config             config.Config
	TravelRpc          travel.Travel
	HomestayOrderModel model.HomestayOrderModel
	AsynqClient        *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:             c,
		TravelRpc:          travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		HomestayOrderModel: model.NewHomestayOrderModel(sqlConn, c.Cache),
		AsynqClient:  newAsynqClient(c),
	}
}
