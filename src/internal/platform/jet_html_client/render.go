package jet_html_client

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

func (client *JetHtmlClient) renderPage(w http.ResponseWriter, template string, data jet.VarMap) error {
	view, err := client.views.GetTemplate(template + ".jet.html")
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
