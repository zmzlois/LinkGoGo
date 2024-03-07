package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zmzlois/LinkGoGo/pkg/model"
	"github.com/zmzlois/LinkGoGo/pkg/utils"
)

type cookieKey int

var CookieName = fmt.Sprintf("_LinkGoGo-session-token-%s", ScopeIdentify)

const (
	ContextKey cookieKey = iota
)

type Claims struct {
	UserID       int     `json:"user_id"`
	Username     string  `json:"username"`
	State        string  `json:"state"`
	RefreshToken string  `json:"refresh_token"`
	AccessToken  string  `json:"access_token"`
	Exp          float64 `json:"exp"`
	jwt.RegisteredClaims
}

func (dc *Client) CreateToken(userData *model.LoginUserInput, tokenInput *model.TokenInput, state string) (tokenString string, err error) {

	jwtMapClaim := jwt.MapClaims{
		"user_id":       userData.Id,
		"username":      userData.Username,
		"state":         state,
		"refresh_token": tokenInput.RefreshToken,
		"access_token":  tokenInput.AccessToken,
		"exp":           time.Now().Add(168 * time.Hour).Unix(),
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
		Secure:   false,
		HttpOnly: true,
		// SameSite: http.SameSiteStrictMode,
		MaxAge: 168 * 60 * 60,
	}

	http.SetCookie(w, &cookie)
}

func ValidateToken(ctx context.Context, tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.Config("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}
	return false, nil
}

func GetToken(r *http.Request) (string, error) {
	tokenFromCookie, err := r.Cookie(CookieName)

	if err != nil {

		return "", fmt.Errorf("discordClient.GetToken: %w", err)
	}

	return tokenFromCookie.Value, nil
}

func RetrieveTokenValue(field string, r *http.Request) (string, error) {
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

func WithSession(ctx context.Context, CookieValue string) context.Context {
	return context.WithValue(ctx, ContextKey, CookieValue)
}
