package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://somelong/json")
	var res interface{}
	err := json.NewDecoder(io.LimitReader(resp.Body, 1000)).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
}
