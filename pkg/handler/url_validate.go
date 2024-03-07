package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
)

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "param")

	r.ParseForm()

	if param == "url" {
		result, err := url.ParseRequestURI(r.Form.Get("url"))

		if err != nil {
			fmt.Fprintf(w, "Please input valid URL: %s", err)
		}
		fmt.Println("Result", result)
	}
}
