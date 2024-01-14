package handler

import (
	dblayer "filestore-server/db"
	"filestore-server/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	pwd_salt = "*#890"
)

// 处理用户注册请求/响应注册页面
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/sigunp.html")

}

// 处理注册post请求
func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if len(username) < 3 || len(passwd) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Invalid parameter",
			"code": -1,
		})
		return
	}
	//加密
	enc_passwd := util.Sha1([]byte(passwd + pwd_slat))
	//注册到用户表
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册成功",
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册失败",
			"code": -2,
		})
	}
}

// 登录接口/响应登录页面
func SignInHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// 处理登录post请求
func DoSignlnHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	encPasswd := util.Sha1([]byte(password + pwd_slat))

	// 1.校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "密码错误",
			"code": -1,
		})
		return
	}

	// 2.生成访问凭证（token）
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "生成凭证失败",
			"code": -2,
		})
		return
	}

	// 3.登录成功后重定向到首页
	// w.Write([]byte("http://" + r.Host + "/static/view/home.hmtl"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.hmtl",
			Username: username,
			Token:    token,
		},
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}

// 查询用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//1.解析请求参数
	r.ParseForm()
	useranme := r.Form.Get("username")
	token := r.Form.Get("token")

	// // 2.验证token是否有效
	// isValidToken := IsTokenValid(token)
	// if !isValidToken {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }

	// 3.查询用户信息
	user, err := dblayer.GetUserInfo(useranme)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 4.组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// 生产token
func GenToken(username string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// token是否有效
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	//判断token的时效性，是否过期
	// 从数据库表tbl_user_token查询username对应的toeken信息
	// 对比两个toeken是否一致
	return true
}
