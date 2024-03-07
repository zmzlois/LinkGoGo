package handler

import (
	"fmt"
	"net/http"
)

func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	title := r.Form.Get("title")

	fmt.Fprintf(w, "Create link with title: %s", title)
	link := r.Form.Get("url")
	fmt.Fprintf(w, "Create link with link: %s", link)
}
