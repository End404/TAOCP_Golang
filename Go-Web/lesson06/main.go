package main

//querystring参数
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		name := c.Query("query")
		age := c.Query("age")

		//name := c.DefaultQuery("query", "somebody")

		//name, ok := c.GetQuery("query")

		//if !ok {
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}
