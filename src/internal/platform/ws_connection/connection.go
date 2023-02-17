package ws_connection

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WsConnection struct {
	connection *websocket.Conn
}

func New(w http.ResponseWriter, r *http.Request) (*WsConnection, error) {
	connection := WsConnection{}
	err := connection.Upgrade(w, r)
	if err != nil {
		return &connection, err
	}

	return &connection, nil
}

func (connection *WsConnection) Close() error {
	return connection.connection.Close()
}
