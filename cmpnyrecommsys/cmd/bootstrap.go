package main

import (
	"fmt"
	"github.com/afzal442/encore-hacks/pkg/database"
	"github.com/afzal442/encore-hacks/pkg/database/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase(database.Config())

	err = database.DB.AutoMigrate(
		&models.Company{},
		&models.Preset{},
	)
	if err != nil {
		panic("Could not migrate database")
	}

	fmt.Println("Database migrated successfully.")
}
