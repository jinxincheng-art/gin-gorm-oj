package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin-gorm-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error:",err)
	}
	return db
}
