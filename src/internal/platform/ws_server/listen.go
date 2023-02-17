package ws_server

import (
	"fmt"
	"github.com/agedito/ugo_websockets/internal/platform/ws_connection"
	"log"
	"sort"
)

func (server *WsServer) ListenForWs(connection ws_connection.WsConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload
	for {
		err := connection.Read(&payload)
		if err == nil {
			payload.Conn = connection
			server.wsChan <- payload
			log.Println("Received", payload)
		}
	}
}

func (server *WsServer) ListenToWsChannel() {
	var response WsJsonResponse
	for {
		e := <-server.wsChan

		switch e.Action {
		case "username":
			server.clients[e.Conn] = e.Username
			users := server.getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			server.broadcastToAll(response)

		case "left":
			response.Action = "list_users"
			delete(server.clients, e.Conn)
			users := server.getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			server.broadcastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
			server.broadcastToAll(response)
		}

	}
}

func (server *WsServer) getUserList() []string {
	var userList []string
	for _, x := range server.clients {
		if x != "" {
			userList = append(userList, x)
		}
	}
	sort.Strings(userList)
	return userList
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
