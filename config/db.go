package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var (
		user   = "root"
		pass   = ""
		host   = "localhost"
		dbname = "orders_by"
	)
	dsn := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error when connect")
	}

	db.AutoMigrate(&Order{}, &Item{})
}

func GetDB() *gorm.DB {
	return db
}
