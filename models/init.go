package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:zdzdsrs20001010@tcp(127.0.0.1:3306)/iot-platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	err = db.AutoMigrate(&ProductBasic{}, &DeviceBasic{}, &UserBasic{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	DB = db
}
