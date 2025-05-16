package provider

import (
	"log"

	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

type DatabaseMethod interface {
}

func NewDatabase(connector gorm.Dialector) *Database {

	db, err := gorm.Open(connector, &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable handle connection : %s \n", err.Error())
	}

	return &Database{
		Db: db,
	}
}

