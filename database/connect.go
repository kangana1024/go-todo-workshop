package database

import (
	"gotodo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	connection, err := gorm.Open(postgres.Open("postgresql://demo:123456@127.0.0.1/gotodo?sslmode=disable"), &gorm.Config{})
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
