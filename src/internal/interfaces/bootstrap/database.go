package bootstrap

import (
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

type DatabaseManager interface {
}

func NewDatabase(db *gorm.DB) *Database {

	return &Database{
		Db: db,
	}
}
