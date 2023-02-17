package ports

import "net/http"

type WsServer interface {
	WsEndpoint(w http.ResponseWriter, r *http.Request)
}
