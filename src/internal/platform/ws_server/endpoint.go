package ws_server

import (
	"github.com/agedito/ugo_websockets/internal/platform/ws_connection"
	"log"
	"net/http"
)

func (server *WsServer) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := ws_connection.New(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	response := `<em><small>Connected to server</em></small>`
	err = server.Response(ws, response)
	if err != nil {
		log.Println(err)
		return
	}
}
