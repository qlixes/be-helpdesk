package bootstrap

import (
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type DatabaseManager interface {
}

func NewDatabase() *Database {

	db, err := gorm.Open()

	return &Database{
		DB: db,
	}
}
