package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html_client/html"),
	jet.InDevelopmentMode(),
)

func renderPage(w http.ResponseWriter, template string, data jet.VarMap) error {
	view, err := views.GetTemplate(template + ".jet.html")
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

func Home(w http.ResponseWriter, _ *http.Request) {
	err := renderPage(w, "home", nil)
	if err != nil {
		log.Println(err)
	}
}
