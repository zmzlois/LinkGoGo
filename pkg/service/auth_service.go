package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/zmzlois/LinkGoGo/ent"
	"github.com/zmzlois/LinkGoGo/ent/users"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/model"
)

type AuthService struct {
	db            *ent.Client
	discordClient *auth.DiscordService
	State         string
}

func NewAuthService(db *ent.Client, discordClient *auth.DiscordService, state string) *AuthService {
	return &AuthService{
		db:            db,
		discordClient: discordClient,
		State:         state,
	}
}

func (s AuthService) Login(state string) (*ent.Users, error) {
	user, err := s.db.Users.Query().
		Where(users.SessionState(state)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	// update last login time
	s.db.Users.UpdateOneID(user.ID).
		SetUpdatedAt(time.Now()).
		Save(context.Background())

	return user, nil
}

func (s AuthService) Redirect(code string, state string, w http.ResponseWriter) (*model.LoginUserInput, *model.TokenInput, bool, error) {

	var userData = &model.LoginUserInput{}
	var tokenPayload *model.TokenInput

	data, err1 := s.discordClient.DiscordService.GetAccessTokenMap(code)

	if err1 != nil {
		panic(fmt.Sprintf("[AuthService] Failed to get access token. Error: %s\n", err1))
	}

	accessToken := data["token_type"].(string) + " " + data["access_token"].(string)

	userDataMap, _ := s.discordClient.DiscordService.GetUserData(accessToken)
	userData = &model.LoginUserInput{
		Username:   userDataMap["username"].(string),
		Avatar:     fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", userDataMap["id"].(string), userDataMap["avatar"].(string)),
		GlobalName: userDataMap["global_name"].(string),
		Id:         userDataMap["id"].(string),
	}

	tokenPayload = &model.TokenInput{
		AccessToken:  data["access_token"].(string),
		RefreshToken: data["refresh_token"].(string),
		ExpiresIn:    data["expires_in"].(float64),
		Scope:        data["scope"].(string),
		TokenType:    data["token_type"].(string),
	}

	// if user exists, issue them a new session token
	_, err := s.Login(state)

	cookie := http.Cookie{
		Name:  fmt.Sprintf("_LinkGoGo-session-token-%s", s.discordClient.DiscordService.Scopes),
		Value: s.State,
	}

	if err != nil {

		w.Write([]byte("User doesn't exist, creating user..."))

		// user doesn't exists, no user found, move to the next middleware to create the user
		return userData, tokenPayload, false, nil
	}

	// user exists, check the state if it is valid, then
	http.SetCookie(w, &cookie)

	return userData, tokenPayload, true, nil
}

func (s AuthService) CreateSession() {
	// create a session

}
