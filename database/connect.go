package database

import (
	"gotodo/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	connection, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = connection
	err = connection.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}

	return nil
}
