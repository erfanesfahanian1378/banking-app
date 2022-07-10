package main

import (
	"booking-app/migrations"
	"fmt"
)

func main() {
	migrations.Migrate()
	fmt.Println("Hello world")
}
