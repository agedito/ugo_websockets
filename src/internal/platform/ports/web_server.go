package ports

type WebServer interface {
	Run(port string) error
}
