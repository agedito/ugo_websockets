package ws_connection

import "log"

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func (connection *WsConnection) Response(message string) error {
	var response WsJsonResponse
	response.Message = message
	err := connection.connection.WriteJSON(response)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
