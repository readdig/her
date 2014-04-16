package her

import (
	"database/sql"
	"log"
)

type DB struct{}

var connections = make(map[string]*Connection)

type Connection struct {
	Driver     string
	DataSource string
}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) Connection(key, driver, dataSource string) {
	connections[key] = &Connection{Driver: driver, DataSource: dataSource}
}

func (d *DB) Open(key string) *sql.DB {
	conn := connections[key]
	if conn != nil {
		db, err := sql.Open(conn.Driver, conn.DataSource)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}
	return nil
}
