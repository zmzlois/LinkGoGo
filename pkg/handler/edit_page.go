package handler

import (
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/model"
	"github.com/zmzlois/LinkGoGo/pkg/service"
	"github.com/zmzlois/LinkGoGo/web/pages"
)

type TodosController struct {
	LinkService *service.LinkService
}

func (rs *TodosController) Edit(w http.ResponseWriter, r *http.Request) {
	var links = []model.NewLinkInput{}

	ctx := r.Context()

	avatar := auth.GetSessionValue(ctx, "avatar")
	global_name := auth.GetSessionValue(ctx, "global_name")

	links, err := rs.LinkService.GetLinks(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// don't return here, still render the page
	}

	pages.EditPage(avatar, global_name, "Edit your bio here?", links).Render(r.Context(), w)
}
