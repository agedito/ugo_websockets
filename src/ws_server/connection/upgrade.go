package connection

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var UpgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

func (ws *WsConnection) Upgrade(w http.ResponseWriter, r *http.Request) error {
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	ws.connection = conn

	return nil
}
