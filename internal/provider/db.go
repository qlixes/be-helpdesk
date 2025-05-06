package provider

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

type IDatabase interface {
}

var Db = newDatabase()

func newDatabase() *Database {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Jakarta",
		Cfg.Db.User,
		Cfg.Db.Pass,
		Cfg.Db.Host,
		Cfg.Db.Port,
		Cfg.Db.Name,
	)

	pgx, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalf("Unable connect database : %s \n", err.Error())
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: pgx,
	}))

	if err != nil {
		log.Fatalf("Unable connect existing database : %s \n", err.Error())
	}

	log.Printf("Database connected.")

	return &Database{
		Db: db,
	}
}
