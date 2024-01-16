package models

import (
	"iot-platform/define"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := define.MysqlDSN + "/iot-platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR]: ", err)
	}
	err = db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})
	if err != nil {
		log.Fatalln("[DB ERROR]: ", err)
	}
	DB = db
}
