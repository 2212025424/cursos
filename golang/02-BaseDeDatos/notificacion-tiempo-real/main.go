package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
	"github.com/lib/pq"
)

const (
	DB_HOST = "208.68.37.200"
	DB_USER = "usertaller"
	DB_PASS = "usertaller"
	DB_NAME = "tallerdb"
	DB_SSLM = "disable"
)

func waitForNotification (l *pq.Listener) {
	for {
		select {
			case n := <- l.Notify:
				var j bytes.Buffer
				fmt.Println("Datos recibidos del canal llamado: ", n.Channel)
				if err := json.Indent(&j, []byte(n.Extra), "", "\t"); err != nil {
					fmt.Println("Error al procesar el json", err)
					return
				}
				fmt.Println(string(j.Bytes()))
				return

			case <-time.After(60 * time.Second):
				fmt.Println("No se ha resibido informacion, revisa conexion")
				go func () {
					l.Ping()
				} ()
				return
		}
	}
}

func main () {
	conn := fmt.Sprintf("dbname=%s host=%s user=%s password=%s sslmode=%s", DB_NAME, DB_HOST, DB_USER, DB_PASS, DB_SSLM)

	if _, err := sql.Open("postgres", conn); err != nil {
		panic(err)
	}

	reporte := func (et pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	listener := pq.NewListener(conn, 10*time.Second, time.Minute, reporte)

	if err := listener.Listen("canal"); err != nil {
		panic(err)
	}

	fmt.Println("Inicia proceso de monitoreo")

	for {
		waitForNotification(listener)
	}
}