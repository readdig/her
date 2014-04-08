package her

import (
	"database/sql"
	"log"
)

type DB struct{}

var Connections = make(map[string]*Connection)

type Connection struct {
	Driver     string
	DataSource string
}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) Connection(key, driver, dataSource string) {
	Connections[key] = &Connection{Driver: driver, DataSource: dataSource}
}

func (d *DB) Open(key string) *sql.DB {
	conn := Connections[key]
	if conn != nil {
		DB, err := sql.Open(conn.Driver, conn.DataSource)
		if err != nil {
			log.Fatal(err)
		}
		return DB
	}
	return nil
}
