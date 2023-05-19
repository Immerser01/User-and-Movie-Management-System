package Database

import (
	"database/sql"
)

type Database struct {
	Db *sql.DB
}

func NewStorage(database *sql.DB) *Database {
	return &Database{Db: database}
}
