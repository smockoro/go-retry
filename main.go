package main

import (
	"log"
	"net/http"
)

func main() {
	c := &http.Client{}
	rc := NewRetryClient(c)
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	rc.Do(req, nil)
}
