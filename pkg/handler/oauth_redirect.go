package handler

import (
	"log"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/model"
	"github.com/zmzlois/LinkGoGo/pkg/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (h *AuthHandler) OAuthHandler(w http.ResponseWriter, r *http.Request) {
	h.AuthService.Redirect(w, r)
}

func (h *AuthHandler) OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {

	// var dc auth.Client

	// var userService service.UserService

	var dsc = auth.DiscordInit()

	// Get the code from the redirect parameters (&code=...)
	var codeFromURLParamaters = r.URL.Query()["code"][0]

	// get access token map
	data, err1 := dsc.GetAccessTokenMap(codeFromURLParamaters)

	accessToken := data["token_type"].(string) + " " + data["access_token"].(string)

	if err1 != nil {
		log.Printf("OAuthCallbackHandler.Fialed to get access token: %s", err1)
	}
	user, err := dsc.GetUserData(accessToken)

	if err != nil {
		log.Printf("OAuthCallbackHandler.Failed to get user data: %s", err)
	}

	// processing the data coming in
	userData, err := model.ParsingUserInput(user)

	if err != nil {
		log.Printf("[OAuth Redirect.Redirect.ParsingUserInput]: %s", err)
	}

	// check if this user exists?

	userExist, err := h.AuthService.Login(userData.Id)

	if err != nil {
		log.Printf("[OAuth Redirect.Login]: %s", err)
	}

	tokenPayload, err := model.ParsingTokenInput(data)

	if err != nil {
		log.Printf("[OAuth Redirect.Redirect]: %s", err)
	}

	// creating signed jwt token
	tokenString, err := dsc.CreateToken(userData, tokenPayload, h.AuthService.State)

	// if user exists, redirect to edit page, update their session
	if userExist != nil {
		http.Redirect(w, r, "/edit", http.StatusFound)
	}

	if err != nil {
		log.Printf("[OAuth Redirect.Redirect]: %s", err)
	}

	// set token in cookie
	dsc.SetCookie(tokenString, w)

	var userService = &service.UserService{}

	User, Session, err := userService.CreateUser(userData, tokenPayload, tokenString)

	if err != nil {
		log.Printf("[OAuth Redirect.Redirect]: %s", err)
	}

	if User == nil || Session == nil {
		log.Printf("[OAuth Redirect.Redirect]: User or Session is nil")
	}

	if User != nil && Session != nil {
		log.Printf("[OAuth Redirect.Redirect]: User %s and Session created", User.Username)
	}

	http.Redirect(w, r, "/edit", http.StatusFound)

}
