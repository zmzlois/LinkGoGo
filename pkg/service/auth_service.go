package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zmzlois/LinkGoGo/ent"
	"github.com/zmzlois/LinkGoGo/ent/users"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/model"
)

type AuthService struct {
	db    *ent.Client
	dsc   *auth.Client
	State string
}

func NewAuthService(db *ent.Client, discordClient *auth.Client, state string) *AuthService {
	return &AuthService{
		db:    db,
		dsc:   discordClient,
		State: state,
	}
}

func (s *AuthService) Login(userId string, tokenString string) (*ent.Users, error) {
	user, err := s.db.Users.Query().
		Where(users.ExternalID(userId)).
		Only(context.Background())
	if err != nil {
		log.Println("[Login]: ", err)
		return nil, fmt.Errorf("login.user not found, create user.%w", err)
	}

	// update last login time
	_, err = s.db.Users.UpdateOneID(user.ID).
		SetUpdatedAt(time.Now()).
		Save(context.Background())

	if err != nil {
		log.Println("[Login]: ", err)
		return nil, fmt.Errorf("login.failed to update user. %w", err)
	}

	if user != nil {
		s.db.Session.Create().
			SetSessionToken(tokenString).
			SetExpiresAt(time.Now().Add(168 * time.Hour)).
			SetUser(user).
			Save(context.Background())
	}

	return user, nil
}

func (s *AuthService) Redirect(w http.ResponseWriter, r *http.Request) {
	s.dsc.RedirectHandler(w, r, s.State)
}

func (s *AuthService) Callback(code string, state string, w http.ResponseWriter) (*model.LoginUserInput, *model.TokenInput, string, bool, error) {

	var userData = &model.LoginUserInput{}
	var tokenPayload = &model.TokenInput{}

	var dc auth.Client

	fmt.Println("Code: ", code)

	data, err1 := s.dsc.GetAccessTokenMap(code)

	if err1 != nil {
		panic(fmt.Sprintf("[AuthService] Failed to get access token map. Error: %s\n", err1))
	}

	accessToken := data["token_type"].(string) + " " + data["access_token"].(string)

	userDataMap, _ := dc.GetUserData(accessToken)
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

	tokenString, err := s.dsc.CreateToken(userData, tokenPayload, state)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, nil, "", false, err
	}
	// if user exists, issue them a new session token
	// TODO: ambiguous about what this does
	_, err = s.Login(userData.Id, tokenString)

	if err != nil {
		w.Write([]byte("User doesn't exist, creating user..."))

		// user doesn't exists, no user found, move to create the user
		return userData, tokenPayload, tokenString, false, nil
	}
	// getting JWT token

	// set cookie
	s.dsc.SetCookie(tokenString, w)

	return userData, tokenPayload, tokenString, true, nil
}
