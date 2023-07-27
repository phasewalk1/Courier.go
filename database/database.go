package database

import (
	"fmt"
	"os"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"github.com/phasewalk1/courier-go/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Los_Angeles",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config {
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil  {
		log.Fatal("Failed to connect to database. Error: ", err)
		os.Exit(2)
	}

	log.Println("Connected to database.")
	log.Println("Running migrations...")

	db.AutoMigrate(&models.Message{})

	DB = DbInstance {
		Db: db,
	}
}
