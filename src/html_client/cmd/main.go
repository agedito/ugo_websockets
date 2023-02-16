package main

import (
	"github.com/agedito/ugo_websockets/html_client/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(handlers.Home))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
