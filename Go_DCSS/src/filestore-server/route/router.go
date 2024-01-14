package route

import (
	"filestore-server/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin framework,包括Logger,Recovery
	router := gin.Default()

	//处理静态资源
	router.Static("/static/", "./static/")

	//不需要经过验证就能访问的接口
	router.GET("/user/signup", handler.SignupHandler)
	router.POST("/user/signup", handler.DoSignupHandler)

	router.GET("/user/signin", handler.SignupHandler)
	router.POST("/user/signin", handler.DoSignlnHandler)

	//加入中间件，用于校验toekn的拦截器
	router.Use(handler.HTTPInterceptor())

	//拦截器token校验

}
