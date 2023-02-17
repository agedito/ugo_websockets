package main

import (
	"fmt"
	"github.com/agedito/ugo_websockets/internal/platform/jet_html_client"
	"github.com/agedito/ugo_websockets/internal/platform/web_server"
	"github.com/agedito/ugo_websockets/internal/platform/ws_server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	config, err := godotenv.Read("env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	port := config["WEB_PORT"]
	fmt.Println("Starting server at port", port, "...")

	htmlClient := jet_html_client.New()
	wsServer, err := ws_server.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	server := web_server.New(wsServer, htmlClient)

	err = server.Run(port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
