package main

import (
	"log"
	"net/http"
	"api"
)

func main() {
	http.HandleFunc("/api/cmd", api.cmdHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
