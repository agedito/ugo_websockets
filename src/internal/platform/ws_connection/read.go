package ws_connection

import (
	"log"
)

func (connection *WsConnection) Read(payload interface{}) error {
	err := connection.connection.ReadJSON(&payload)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
