package Database

import (
	"database/sql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(connString string) (*Database, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) InsertUser(User *UserData) error {
	_, err := d.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		User.Name, User.Email, User.Password)
	if err != nil {
		return err
	}

	return nil
}
