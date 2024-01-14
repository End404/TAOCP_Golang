package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")  //静态文件
	r.SetFuncMap(template.FuncMap{ //模版添加自定义函数
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/index.tmpl") //模版解析
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(200, "posts/index.tmpl", gin.H{ //模版渲染
			"title": "liwenzhou.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //模版渲染
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.Run(":9090") //启动服务
}
