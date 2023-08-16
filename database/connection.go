package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "notas"
	password = "123"
	dbname   = "notas"
)

func GetConn() *sql.DB {
	return db
}

func Close(db2 *sql.DB) error {
	if err := db2.Close(); err != nil {
		return err
	}
	return nil
}

func Init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")
}
