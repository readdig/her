package her

import (
	"database/sql"
	"log"
)

type DB struct{}

func (d *DB) Open() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
