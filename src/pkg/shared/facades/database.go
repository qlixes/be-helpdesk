package facades

import (
	"database/sql"
)

type Database struct {
	Db *sql.DB
}

type DatabaseManager interface {
}

func NewDatabase() *Database {

	return &Database{}
}
