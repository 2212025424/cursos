package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		email,
		password,
	}

	data := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(data).Encode(&login)

	if err != nil {
		log.Fatalf("Marshal.Login: %v\n", err)
	}

	resp := httpClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("ERROR: %v  -  response: %s\n", err, string(body))
	}

	dataResponse := LoginResponse{}

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)

	if err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}

	return dataResponse
}
