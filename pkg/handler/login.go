package handler

import (
	"net/http"

	"github.com/zmzlois/LinkGoGo/web/pages"
)

func Login(w http.ResponseWriter, r *http.Request) {

	pages.LogInPage().Render(r.Context(), w)
}
