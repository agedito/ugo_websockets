package handlers

import (
	"github.com/agedito/ugo_websockets/ws_server/connection"
	"log"
	"net/http"
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws := connection.New()
	err := ws.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client connected to endpoint")
	response := `<em><small>Connected to server</em></small>`
	ws.Response(response)
}
