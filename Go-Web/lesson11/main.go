package main

//gin框架实现重定向

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		//跳转到搜狗页面
		c.Redirect(http.StatusMovedPermanently, "https://www.sogou.com/")
	})
	r.GET("/a", func(c *gin.Context) {
		//跳转到 b 对应的路由处理函数
		c.Request.URL.Path = "/b" //把请求的uri修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":8080")
}
