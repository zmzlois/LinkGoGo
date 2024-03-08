package handler

import (
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/web/pages"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var isUser = false

	isUser, err := auth.IsUser(r)

	if err != nil {
		fmt.Printf("Index.isUser: %v", zap.Error(err))
	}

	pages.HomePage(isUser).Render(r.Context(), w)
}
