package models

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

var RDB = RdbInit()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin-gorm-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error:", err)
	}
	return db
}

func RdbInit() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
