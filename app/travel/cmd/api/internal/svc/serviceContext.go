package svc

import (
	"XavierLookLook/app/travel/cmd/api/internal/config"
	"XavierLookLook/app/travel/cmd/rpc/travel"
	"XavierLookLook/app/travel/model"
	"XavierLookLook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	//rpc
	UsercenterRpc usercenter.Usercenter
	TravelRpc     travel.Travel

	//model
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
