package handler

import (
	"log"
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

	var isUser bool
	var bio string = "Edit your bio here?"

	avatar := auth.GetSessionValue(ctx, "avatar")
	global_name := auth.GetSessionValue(ctx, "global_name")
	id := auth.GetSessionValue(ctx, "user_id")

	if user, ok := rs.LinkService.GetUser(ctx); ok != nil {
		log.Println("Error getting user")
	} else {
		bio = user.Description
	}

	links, err := rs.LinkService.GetLinks(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// don't return here, still render the page
	}

	pages.EditPage(avatar, global_name, bio, id, links, isUser).Render(r.Context(), w)
}
