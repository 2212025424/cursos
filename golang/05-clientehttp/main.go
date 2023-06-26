package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const host = "http://localhost:8080"

func main() {
	lc := loginClient(host+"/v1/login", "contacto@ed.team", "123456")
	fmt.Println(lc)

	persona := Person{
		Name: "Lupe",
		Age:  20,
		Communities: []Community{
			{Name: "sdsdvsdvs"},
			{Name: "sdfsdfb sbfs"},
		},
	}

	gr := createPerson(host+"/v1/person", lc.Data.Token, &persona)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Fatalf("Request: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}

	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("Request: %v\n", err)
	}

	return response
}
