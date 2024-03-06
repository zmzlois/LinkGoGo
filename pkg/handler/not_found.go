package handler

import (
	"net/http"

	"github.com/zmzlois/LinkGoGo/web/pages"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	pages.NotFound().Render(r.Context(), w)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed"))
}
