package ws_connection

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

func (connection *WsConnection) Upgrade(w http.ResponseWriter, r *http.Request) error {
	upgradedConnection, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	connection.connection = upgradedConnection
	return nil
}
