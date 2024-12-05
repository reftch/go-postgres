package db

import (
	"fmt"
	"log"
	"os"

	"github.com/reftch/go-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUser, err := getEnv("POSTGRES_USER")
	if err != nil {
		log.Fatalln(err)
	}
	dbPassword, err := getEnv("POSTGRES_PASSWORD")
	if err != nil {
		log.Fatalln(err)
	}
	dbName, err := getEnv("POSTGRES_DB")
	if err != nil {
		log.Fatalln(err)
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}

func getEnv(valueKey string) (string, error) {
	value := os.Getenv(valueKey)
	if value == "" {
		return "", fmt.Errorf("%s environment variable not set", valueKey)
	}
	return value, nil
}
