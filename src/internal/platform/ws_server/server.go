package ws_server

type WsServer struct {
}

func New() (*WsServer, error) {
	server := WsServer{}
	return &server, nil
}
