package handler

import (
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/web/templates/partials"
)

func (rs *TodosController) CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	partials.LinkForm().Render(r.Context(), w)
}

func (rs *TodosController) CreateButtonHandler(w http.ResponseWriter, r *http.Request) {
	partials.CreateLink().Render(r.Context(), w)
}

func (rs *TodosController) EditLinkHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	title := r.Form.Get("title")

	fmt.Fprintf(w, "Edit link with title: %s", title)
	link := r.Form.Get("url")
	fmt.Fprintf(w, "Edit link with link: %s", link)
}
