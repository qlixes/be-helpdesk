package provider

import (
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/gorm"
)

var connector gorm.Dialector

type Database struct {
	Db *gorm.DB
}

type DatabaseMethod interface {

}

func NewDatabase() *Database {

	db, err := gorm.Open(connector, &gorm.Config{})

	if err != nil {
		log.Fatal("Unable handle connection : %s \n", err.Error())
	}

	return &Database{
		Db: db,
	}
}

func NewMysql() *Database {
	
}