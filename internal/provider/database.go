package provider

import (
	"database/sql"
    "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// using database/sql package
type Database struct {
	Db *sql.DB
}

type DatabaseMethod interface {

}

// using https://github.com/lib/pq
func NewPgsql() *Database {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		bootstrap.Config.Db.Host,
		bootstrap.Config.Db.Port,
		bootstrap.Config.Db.User,
		bootstrap.Config.Db.Pass,
		bootstrap.Config.Db.Name,
	)

	db, err := sql.Open("postgres", dsn)

	if err  != nil {
		log.Fatalf("Unable connect db : %s \n", err.Error())
	}

	return &Database{
		Db: db,
	}
}

// using https://github.com/go-sql-driver/mysql
func NewMysql() *Database {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		bootstrap.Config.Db.User,
		bootstrap.Config.Db.Pass,
		bootstrap.Config.Db.Host,
		bootstrap.Config.Db.Port,
		bootstrap.Config.Db.Name,
	)

	db, err := sql.Open("mysql", dsn)

	if err  != nil {
		log.Fatalf("Unable connect db : %s \n", err.Error())
	}

	return &Database{
		Db: db,
	}
}