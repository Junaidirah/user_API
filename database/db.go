package database

import (
	"fmt"
	"log"
	"user_API/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = ""
	dbPort   = "5432"
	dbname   = "user_API"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Connection success to database")
	db.Debug().AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return db
}
