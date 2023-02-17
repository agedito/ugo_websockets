package web_server

import (
	"fmt"
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
	fmt.Println("Starting channel listener")
	go server.wsServer.ListenToWsChannel()
	err := http.ListenAndServe(":"+port, mux)
	return err
}
