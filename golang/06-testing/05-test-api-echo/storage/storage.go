package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver of storage
type Driver string

// newPostgresDB Create a postgresql db connection
func NewPostgresDB() {
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
