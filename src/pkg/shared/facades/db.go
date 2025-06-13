package facades

import (
	"database/sql"
)

type Db struct {
	DB *sql.DB
}

type DbManager interface {
	UseEngine(engine, dsn string) *Db
}

func NewDb() *Db {

	return &Db{}
}

func (d *Db) UseMysql() {

}
