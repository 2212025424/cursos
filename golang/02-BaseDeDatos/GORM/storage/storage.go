package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
)

var (
	db   *gorm.DB
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

		dsn := "usercrudgo:$dev/E_c0n_db@tcp(157.245.84.7:3306)/crudgodb?parseTime=true"

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Conectando a MySQL")
	})
}

// newPostgresDB Create a postgresql db connection
func newPostgresDB() {
	once.Do(func() {
		var err error

		dsn := "postgres://usercrudgo:usercrudgo@208.68.37.200:5432/crudgodb?sslmode=disable"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Conectando a postgres")
	})
}

// DB return a unique version of db
func DB() *gorm.DB {
	return db
}
