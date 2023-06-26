package main

import (
	"fmt" 
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DB_HOST = "208.68.37.200"
	DB_USER = "usertaller"
	DB_PASS = "usertaller"
	DB_NAME = "tallerdb"
	DB_SSLM = "disable"
)

func main () {
	conn := fmt.Sprintf("dbname=%s host=%s user=%s password=%s sslmode=%s", DB_NAME, DB_HOST, DB_USER, DB_PASS, DB_SSLM)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Se ha abierto la conexion")
	}
}
