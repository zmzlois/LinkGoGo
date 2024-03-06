package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zmzlois/LinkGoGo/pkg/model"
	"github.com/zmzlois/LinkGoGo/pkg/utils"
)

var CookieName = fmt.Sprintf("_LinkGoGo-session-token-%s", ScopeIdentify)

func (dc *Client) CreateToken(userData *model.LoginUserInput, tokenInput *model.TokenInput, state string) (tokenString string, err error) {

	jwtMapClaim := &jwt.MapClaims{
		"user_id":  userData.Id,
		"username": userData.Username,
		"state":    state,
		"exp":      tokenInput.ExpiresIn,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtMapClaim)

	tokenString, err = token.SignedString(utils.Config("JWT_SECRET_KEY"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (dc *Client) SetCookie(tokenString string, w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:     CookieName,
		Value:    tokenString,
		Expires:  time.Now().Add(168 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   168 * 60 * 60,
		Raw:      fmt.Sprintf("_LinkGoGo-session-token-%s=%s; Secure; HttpOnly; SameSite=None; Max-Age=604800", ScopeIdentify, tokenString),
	}

	http.SetCookie(w, &cookie)
}

func (dc *Client) ValidateToken(tokenString string) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.Config("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
