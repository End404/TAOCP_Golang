package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//data := map[string]interface{}{ //方法1使用map
		//	"name": "xwz",
		//	"age":  120,
		//}
		data := gin.H{"name": "向前走", "age": 22}
		c.JSON(http.StatusOK, data)
	})
	type msg struct { //方法2 结构体
		Name string `json:"name"`
		Age  int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"洗完澡",
			24,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
