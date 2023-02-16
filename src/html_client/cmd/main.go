package main

import (
	"github.com/agedito/ugo_websockets/html_client/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.Routes()
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
