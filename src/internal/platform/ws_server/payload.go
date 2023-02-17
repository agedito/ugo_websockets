package ws_server

import (
	"github.com/agedito/ugo_websockets/internal/platform/ws_connection"
)

type WsPayload struct {
	Action   string                     `json:"action"`
	Username string                     `json:"username"`
	Message  string                     `json:"message"`
	Conn     ws_connection.WsConnection `json:"-"`
}
