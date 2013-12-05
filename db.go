package web

import (
	"database/sql"
	"log"
)

type db struct{}

var (
	DB = &db{}
)

func (d *db) Open() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
