package main

//GORM连接MySQL基本示例
// GORM使用之模型定义

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct { //数据表
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

type User struct { //定义模型
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root123@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/*	基本示例
		//创建表 自动迁移
		db.AutoMigrate(&UserInfo{})
		//创建数据行
		//u1 := UserInfo{1, "qimi", "男", "游泳"}
		//db.Create(&u1)
		//查询
		var u UserInfo
		db.First(&u)
		fmt.Println("查询：", u)
		//更新
		db.Model(&u).Update("hobby", "球")
		//删除
		db.Delete(&u)
	*/
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

}
