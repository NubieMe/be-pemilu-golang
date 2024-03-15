package database

import (
	"fmt"
	// "time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=localhost user=postgres password=apahayo dbname=go-backend port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")
}
