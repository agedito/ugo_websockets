package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	config, err := godotenv.Read("env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := config["WEB_PORT"]
	fmt.Println("Starting server at port", port, "...")
	mux := Routes()
	err = http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
