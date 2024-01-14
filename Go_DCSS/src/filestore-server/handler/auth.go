package handler

import (
	"filestore-server/util"
	"go/doc/comment"
	"net/http"

	"github.com/gin-gonic/gin"
)

// http请求拦截器
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		r.ParseForm()
		username := c.Request.FormValue("username")
		toekn := c.Request.FormValue("token")

		//验证登录token是否有效
		if len(username) < 3 || !IsTokenValid(toekn) {
			//失败则跳转
			c.Abort()
			resp := util.NewRespMsg(
				int(comment.StatuslidToken),
				"token无效",
				nil,
			)
			c.JSONBytes(http.StatusOK, resp)
			return
		}
		c.Next()
	}

}
