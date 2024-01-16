package svc

import (
	"iot-platform/device/device"
	"iot-platform/open/internal/config"
	"iot-platform/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RpcDevice device.Device
	RpcUser   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: device.NewDevice(zrpc.MustNewClient(c.RpcDevice)),
		RpcUser:   user.NewUser(zrpc.MustNewClient(c.RpcUser)),
	}
}
