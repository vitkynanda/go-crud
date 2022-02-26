package connection

import (
	"fmt"
	"go-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/posts?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection success")

	DB.AutoMigrate(&models.Article{})
	
}