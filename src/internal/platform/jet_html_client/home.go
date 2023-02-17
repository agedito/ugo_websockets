package jet_html_client

import (
	"log"
	"net/http"
)

func (client *JetHtmlClient) Home(w http.ResponseWriter, _ *http.Request) {
	err := client.renderPage(w, "home", nil)
	if err != nil {
		log.Println(err)
	}
}
