package migration

import (
	"be-pemilu/database"
	"be-pemilu/models"
	"fmt"
	"log"
)

func Migrate() {
	database.ConnectDB()
	db := database.DB

	err := db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Party{},
		&models.Paslon{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully migrated")
}
