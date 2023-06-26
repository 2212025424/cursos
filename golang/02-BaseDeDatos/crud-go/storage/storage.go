package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"time"

	"github.com/2212025424/go-db/pkg/product"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New Create the db connection
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newMySQLDB Create a mysql db connection
func newMySQLDB() {
	once.Do(func() {
		var err error

		db, err = sql.Open("mysql", "usercrudgo:$dev/E_c0n_db@tcp(157.245.84.7:3306)/crudgodb?parseTime=true")

		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Conectando a MySQL")
	})
}

// newPostgresDB Create a postgresql db connection
func newPostgresDB() {
	once.Do(func() {
		var err error

		db, err = sql.Open("postgres", "postgres://usercrudgo:usercrudgo@208.68.37.200:5432/crudgodb?sslmode=disable")

		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Conectando a postgres")
	})
}

// Pool return a unique version of db
func Pool() *sql.DB {
	return db
}

//
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}

	if null.String != "" {
		null.Valid = true
	}

	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}

	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}

// DAOProduct factory of product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}
