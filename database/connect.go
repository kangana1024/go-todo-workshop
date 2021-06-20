package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() error {
	connection, err := gorm.Open(postgres.Open("postgresql://demo:123456@127.0.0.1/gotodo?sslmode=disable"), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Printf("Connecion : %+v", connection)

	return nil
}
