package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/learn_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		panic(dbErr)
	}

	return db
}
