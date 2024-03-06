package model

import (
	"fmt"

	"github.com/zmzlois/LinkGoGo/ent"
)

type LoginUserInput struct {
	Id               string  `json:"id"`
	AccentColor      float64 `json:"accent_color"`
	Avatar           string  `json:"avatar"`
	AvatarDecoration string  `json:"avatar_decoratio_data"`
	Banner           string  `json:"banner"`
	BannerColor      string  `json:"banner_color"`
	Flags            float64 `json:"flags"`
	MfaEnabled       bool    `json:"mfa_enabled"`
	PremiumType      float64 `json:"premium_type"`
	PublicFlags      float64 `json:"public_flags"`
	Discriminator    string  `json:"discriminator"`
	GlobalName       string  `json:"global_name"`
	Locale           string  `json:"locale"`
	Username         string  `json:"username"`
}

func ParsingUserInput(input map[string]interface{}) (*LoginUserInput, error) {

	var user = &LoginUserInput{
		Id:         input["id"].(string),
		Avatar:     fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", input["id"].(string), input["avatar"].(string)),
		GlobalName: input["global_name"].(string),
		Locale:     input["locale"].(string),
		Username:   input["username"].(string),
	}

	return user, nil

}

type CreateUserInput struct {
	CreateUserInput *ent.Users
}

type TokenInput struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	ExpiresIn    float64 `json:"expires_in"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
}

func ParsingTokenInput(input map[string]interface{}) (*TokenInput, error) {

	var token = &TokenInput{
		AccessToken:  input["access_token"].(string),
		RefreshToken: input["refresh_token"].(string),
		ExpiresIn:    input["expires_in"].(float64),
		Scope:        input["scope"].(string),
		TokenType:    input["token_type"].(string),
	}

	return token, nil
}

type NewLinkInput struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Image  string `json:"image"`
	Author string `json:"author"`
}
