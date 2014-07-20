package her

import (
	"database/sql"
	"log"
)

type DB struct {
	key string
}

var connections = make(map[string]*Connection)

type Connection struct {
	Driver     string
	DataSource string
}

func NewDB(a ...interface{}) *DB {
	if len(a) > 0 {
		if v, ok := a[0].(string); ok {
			return &DB{key: v}
		}
	}
	return &DB{}
}

func (d *DB) Connection(key, driver, dataSource string) {
	connections[key] = &Connection{Driver: driver, DataSource: dataSource}
}

func (d *DB) Open() *sql.DB {
	conn := connections[d.key]
	if conn != nil {
		db, err := sql.Open(conn.Driver, conn.DataSource)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}
	return nil
}
