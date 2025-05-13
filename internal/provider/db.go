package provider

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

type DatabaseMethod interface {

}

var Pg = newPgsql()

func newPgsql() *Database {

	pgx, err := sql.Open("pgx", Cfg.GetPgsqlConnection())

	if err != nil {
		log.Fatalf("Unable connect database : %s \n", err.Error())
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: pgx,
	}))

	if err != nil {
		log.Fatalf("Unable connect existing database : %s \n", err.Error())
	}

	return &Database{
		Db: db,
	}
}
