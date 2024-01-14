package main

//小清单项目练习

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct { //模型
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySql() (err error) {
	dsn := "root:root1234@tcp(1270.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()

}

func main() {
	//创建数据库
	//连接数据库
	err := initMySql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "static") //找引用的静态文件
	r.LoadHTMLGlob("templates/*") //加载模版文件
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{ //待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项 点击提交 会发送请求到这里
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			//3.返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//查看
		v1Group.GET("/todo", func(c *gin.Context) { //查看所有
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		v1Group.GET("/tode/:id", func(c *gin.Context) { //查看某一个

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) { //修改某一个待办事项
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) { //删除某一个
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run()

}
