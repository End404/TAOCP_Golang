package handler

import (
	"context"
	"filestore-server/service/account/proto"

	"filestore-server/common"
	cfg "filestore-server/config"
	dblayer "filestore-server/db"
	"filestore-server/service/account/proto"
	"filestore-server/util"
)

type User struct{}

// 处理用户注册请求
func Signup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password

	//参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		res.Code = common.StatusParamInvalid
		res.Mesaage = "注册参数无效"
		return nil
	}
	//加密
	enc_passwd := util.Sha1([]byte(passwd + cfg.PasswordSalt))
	//注册到用户表
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		res.Code = common.StatusOK
		res.Mesaage = "注册成功"
	} else {
		res.Code = common.StatusRegisterFailed
		res.Mesaage = "注册失败"
	}
	return nil
}
