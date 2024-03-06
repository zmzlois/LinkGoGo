package handler

import (
	"net/http"

	"github.com/zmzlois/LinkGoGo/web/pages"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pages.HomePage().Render(r.Context(), w)
}
