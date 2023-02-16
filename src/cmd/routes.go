package main

import (
	htmlHandlers "github.com/agedito/ugo_websockets/html_client/handlers"
	wsHandlers "github.com/agedito/ugo_websockets/ws_server/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(htmlHandlers.Home))
	mux.Get("/ws", http.HandlerFunc(wsHandlers.WsEndpoint))
	return mux
}
