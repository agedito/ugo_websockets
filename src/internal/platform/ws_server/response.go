package ws_server

import (
	"github.com/agedito/ugo_websockets/internal/platform/ws_connection"
	"log"
)

type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

func (server *WsServer) Response(wsConn *ws_connection.WsConnection, message string) error {
	var response WsJsonResponse
	response.Message = message
	err := wsConn.Response(response)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
