package handler

import (
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (h *AuthHandler) OAuthRedirectHandler(w http.ResponseWriter, r *http.Request) {

	// var dc auth.Client

	var userService service.UserService

	// Get the code from the redirect parameters (&code=...)
	var codeFromURLParamaters = r.URL.Query()["code"][0]

	userData, tokenPayload, tokenString, result, err := h.AuthService.Redirect(codeFromURLParamaters, h.AuthService.State, w)

	if err != nil {
		panic(fmt.Sprintf("[OAuth Redirect]: %s", err))
	}

	if !result {
		user, session, err := userService.CreateUser(userData, tokenPayload, tokenString)

		if err != nil {
			panic(fmt.Sprintf("[OAuth Redirect]: %s", err))
		}

		if user == nil || session == nil {
			panic("[OAuth Redirect]: Failed to create user or session")
		}
	}

	http.Redirect(w, r, "/edit", http.StatusFound)

}
