package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/experience-digital/idea-generation/internal/db/models"
)

func Init() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}
