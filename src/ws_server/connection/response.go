package connection

import "log"

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func (ws *WsConnection) Response(message string) {
	var response WsJsonResponse
	response.Message = message

	err := ws.connection.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
}
