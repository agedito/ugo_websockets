package ports

import (
	"github.com/agedito/ugo_websockets/internal/platform/ws_connection"
	"net/http"
)

type WsServer interface {
	WsEndpoint(w http.ResponseWriter, r *http.Request)
	ListenForWs(connection ws_connection.WsConnection)
	ListenToWsChannel()
}
