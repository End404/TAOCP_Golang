// Code generated by goctl. DO NOT EDIT!
// Source: device.proto

package server

import (
	"context"

	"iot-platform/device/internal/logic"
	"iot-platform/device/internal/svc"
	"iot-platform/device/types/device"
)

type DeviceServer struct {
	svcCtx *svc.ServiceContext
	device.UnimplementedDeviceServer
}

func NewDeviceServer(svcCtx *svc.ServiceContext) *DeviceServer {
	return &DeviceServer{
		svcCtx: svcCtx,
	}
}

func (s *DeviceServer) SendMessage(ctx context.Context, in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	l := logic.NewSendMessageLogic(ctx, s.svcCtx)
	return l.SendMessage(in)
}
