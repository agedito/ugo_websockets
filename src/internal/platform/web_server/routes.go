package web_server

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (server *WebServer) Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(server.webClient.Home))
	mux.Get("/ws", http.HandlerFunc(server.wsServer.WsEndpoint))
	return mux
}
