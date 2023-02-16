package connection

import (
	"github.com/gorilla/websocket"
)

type WsConnection struct {
	upgrade    websocket.Upgrader
	connection *websocket.Conn
}

func New() *WsConnection {
	conn := WsConnection{
		upgrade: UpgradeConnection,
	}

	return &conn
}
