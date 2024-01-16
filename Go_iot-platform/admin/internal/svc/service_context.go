package svc

import (
	"iot-platform/admin/internal/config"
	"iot-platform/models"
	user2 "iot-platform/user/rpc/types/user"
	"iot-platform/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user.User
	AuthUser *user2.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
