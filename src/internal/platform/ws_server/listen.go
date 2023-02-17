package ws_server

import (
	"fmt"
	"github.com/agedito/ugo_websockets/internal/platform/ws_connection"
	"log"
)

func (server *WsServer) ListenForWs(connection ws_connection.WsConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload
	for {
		err := connection.Read(payload)
		if err == nil {
			payload.Conn = connection
			server.wsChan <- payload
		}
	}
}

func (server *WsServer) ListenToWsChannel() {
	var response WsJsonResponse
	for {
		e := <-server.wsChan
		response.Action = "Got here"
		response.Message = fmt.Sprintf("Some Message, and action ws %s", e.Action)
		server.broadcastToAll(response)
	}
}

func (server *WsServer) broadcastToAll(response WsJsonResponse) {
	for client := range server.clients {
		err := client.Response(response)
		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(server.clients, client)
		}
	}
}
