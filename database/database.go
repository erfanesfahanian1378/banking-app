package database

import (
	"booking-app/helpers"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=erfan dbname=bankapp password=181352 sslmode=disable")
	helpers.HandleErr(err)
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
