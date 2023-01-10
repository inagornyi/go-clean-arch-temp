package mariadb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDB struct {
	db *sql.DB
}

func NewConnection(user, password, name string) (MariaDB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, name))
	if err != nil {
		return MariaDB{}, err
	}
	return MariaDB{db: db}, nil
}

func (db *MariaDB) DB() *sql.DB {
	return db.db
}

func (db *MariaDB) RowExists(query string, args ...interface{}) (bool, error) {
	var exists bool
	err := db.db.QueryRow(fmt.Sprintf("SELECT EXISTS (%s)", query), args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}

func (db *MariaDB) Close() error {
	return db.db.Close()
}
