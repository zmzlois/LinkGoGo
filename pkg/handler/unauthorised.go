package handler

import (
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/utils"
	"github.com/zmzlois/LinkGoGo/web/pages"
)

func UnauthorisedHandler(w http.ResponseWriter, r *http.Request) {

	// how this string should looks like based on environment
	// 	<meta http-equiv="refresh" content="5; url=http://localhost:3000/login"/>
	var loginURL string

	URL := utils.Config("URL")

	loginURL = fmt.Sprintf("5; url=%s/login", URL)

	pages.UnauthorisedPage(loginURL).Render(r.Context(), w)
}
