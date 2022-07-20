package main

import (
	"booking-app/api"
	"booking-app/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
