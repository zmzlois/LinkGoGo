package handler

import (
	"fmt"
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

	user, global_name, _ := auth.RetrieveTokenValue("global_name", r)

	userId := user["user_id"].(string)

	ctx := r.Context()

	links, err := rs.LinkService.GetLinks(ctx, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// don't return here, still render the page
	}

	fmt.Println("All links", links)

	pages.EditPage(user["avatar"].(string), global_name.(string), "Edit your bio here?", links).Render(r.Context(), w)
}
