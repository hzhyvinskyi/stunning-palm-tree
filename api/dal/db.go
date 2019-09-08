package dal

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	DB_USER		= "postgres"
	DB_PASSWORD	= "postgres"
	DB_NAME		= "spt_db"
)

var once sync.Once

// Connect connects to the PostgreSQL ORDBMS
func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error

	once.Do(func() {
		dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
		sql.Open("postgres", dbinfo)
	})
	return db, err
}

// LogAndQuery logs and makes query
func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query)
	return db.Query(query, args...)
}

// MustExec executes query or panics
func MustExec(db *sql.DB, query string, args ...interface{}) {
	if _, err := db.Exec(query, args...); err != nil {
		panic(err)
	}
}
