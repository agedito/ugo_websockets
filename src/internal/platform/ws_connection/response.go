package ws_connection

import "log"

func (connection *WsConnection) Response(response interface{}) error {
	err := connection.connection.WriteJSON(response)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
