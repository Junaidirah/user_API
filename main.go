package main

import (
	"user_API/database"
	"user_API/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
