package test

import (
	"fmt"
	"gin-gorm-oj/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T)  {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin-gorm-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.Problem,0)
	err = db.Debug().Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _,datam := range data {
		fmt.Printf("Problem ==> %v \n",datam)
	}
}

