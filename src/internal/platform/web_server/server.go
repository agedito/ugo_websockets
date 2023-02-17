package web_server

import (
	"github.com/agedito/ugo_websockets/internal/platform/ports"
	"net/http"
)

type WebServer struct {
	wsServer  ports.WsServer
	webClient ports.WebClient
}

func New(wsServer ports.WsServer, webClient ports.WebClient) *WebServer {
	server := WebServer{wsServer: wsServer, webClient: webClient}
	return &server
}

func (server *WebServer) Run(port string) error {
	mux := server.Routes()
	err := http.ListenAndServe(":"+port, mux)
	return err
}
