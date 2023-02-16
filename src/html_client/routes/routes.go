package routes

import (
	"github.com/agedito/ugo_websockets/html_client/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	return mux
}
