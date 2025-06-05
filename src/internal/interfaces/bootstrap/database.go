package bootstrap

import "gorm.io/gorm"

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {

	return &Database{}
}
