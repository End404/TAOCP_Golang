package main

//获取form表单提交的参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//获取表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("password", "***")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
