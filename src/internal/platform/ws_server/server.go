package ws_server

import "github.com/agedito/ugo_websockets/internal/platform/ws_connection"

type WsServer struct {
	wsChan  chan WsPayload
	clients map[ws_connection.WsConnection]string
}

func New() (*WsServer, error) {
	server := WsServer{
		wsChan:  make(chan WsPayload),
		clients: make(map[ws_connection.WsConnection]string),
	}
	return &server, nil
}
