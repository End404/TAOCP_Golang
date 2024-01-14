package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string,
	reg registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {

	registerHandlersFunc()
	ctx = startService(ctx, reg.serviceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host,
	port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s://%s", host, port))
		if err != nil {
			log.Panicln(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v started, Press any key to stop(已启动，按任意键停止). \n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s://%s", host, port))
		if err != nil {
			log.Panicln(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
