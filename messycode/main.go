package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	data := struct {
		Message string `json:"message"`
	}{Message: "Hello world"}
	err := json.NewEncoder(os.Stdout).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}
