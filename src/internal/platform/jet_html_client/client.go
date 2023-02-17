package jet_html_client

import "github.com/CloudyKit/jet/v6"

type JetHtmlClient struct {
	views *jet.Set
}

func New() *JetHtmlClient {
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader("./internal/platform/jet_html_client/templates/"),
		jet.InDevelopmentMode(),
	)

	client := JetHtmlClient{views: views}

	return &client
}
