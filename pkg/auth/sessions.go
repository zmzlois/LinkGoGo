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

type Claims struct {
	UserID   int     `json:"user_id"`
	Username string  `json:"username"`
	State    string  `json:"state"`
	Exp      float64 `json:"exp"`
	jwt.RegisteredClaims
}

func (dc *Client) CreateToken(userData *model.LoginUserInput, tokenInput *model.TokenInput, state string) (tokenString string, err error) {

	jwtMapClaim := jwt.MapClaims{
		"user_id":  userData.Id,
		"username": userData.Username,
		"state":    state,
		"exp":      tokenInput.ExpiresIn,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMapClaim)

	tokenString, err = token.SignedString([]byte(utils.Config("JWT_SECRET_KEY")))

	if err != nil {
		fmt.Printf("[Create JWT Token]: %s", err)
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

func (dc *Client) RetrieveTokenValue(field string, r *http.Request) (string, error) {
	tokenFromCookie, err := r.Cookie(CookieName)

	tokenString := tokenFromCookie.Value

	if err != nil {
		return "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Parse the JWT token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the key used for signing the token (you need this to validate the token)
		// Here, you would typically return the same key that was used for signing the token
		return []byte(utils.Config("JWT_SECRET_KEY")), nil
	})

	// Check for errors during parsing
	if err != nil {
		fmt.Println("discordClient.RetrieveTokenValue.Error parsing JWT token:", err)
		return "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Println("Invalid JWT token")
		return "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Retrieve claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims format")
		return "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Access specific claims
	result := claims[field].(string)
	// username := claims["username"].(string)

	return result, nil

}
