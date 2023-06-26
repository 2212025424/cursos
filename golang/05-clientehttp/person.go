package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func createPerson(url, token string, persona *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(data).Encode(&persona)

	if err != nil {
		log.Fatalf("Marshal.Login: %v\n", err)
	}

	resp := httpClient(http.MethodPost, url, token, data)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("ERROR: %v  -  response: %s\n", err, string(body))
	}

	dataResponse := GeneralResponse{}

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)

	if err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}

	return dataResponse

}
