package main

import(
	micro "github.com/micro/go-micro"

	"filestore-server/service/account/proto"
)

func main(){
	//创建一个service
	service:= micro.NesService(
		micro.Name("go.micro.service.user")
	)
	service.Init()

	 proto.RegisterUserServiceHandler(service.Server(),new())
}