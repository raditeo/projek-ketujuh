package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	db  *gorm.DB
)

func StartDB() {
	config := "xx:xx@tcp(127.0.0.1:3306)/learning-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
