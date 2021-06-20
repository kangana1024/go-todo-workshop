package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Todo   string `gorm:"type:VARCHAR(150)"`
	Detail sql.NullString
	Status string `gorm:"type:VARCHAR(1)"`
}
