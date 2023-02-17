package ports

import "net/http"

type WebClient interface {
	Home(w http.ResponseWriter, r *http.Request)
}
