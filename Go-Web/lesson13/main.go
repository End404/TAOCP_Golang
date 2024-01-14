package main

//gin框架中间件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件m1：统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("计时：", cost)
	fmt.Println("m1 out...")
}

func m2(ctx *gin.Context) {
	fmt.Println("m2 in...")
	ctx.Next()
	fmt.Println("m2 out...")
}

func main() {
	r := gin.Default()

	r.Use(m1, m2) //全局注册中间件函数m1
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {

	})
	r.GET("/user", func(c *gin.Context) {

	})

	r.Run(":9090")
}
