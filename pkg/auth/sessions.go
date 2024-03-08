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

var CookieName = fmt.Sprintf("_LinkGoGo-session-token-%s", ScopeIdentify)

type Claims struct {
	UserID       int     `json:"user_id"`
	Username     string  `json:"username"`
	GlobalName   string  `json:"global_name"`
	Avatar       string  `json:"avatar"`
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
		"avatar":        userData.Avatar,
		"global_name":   userData.GlobalName,
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

func ValidateToken(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.Config("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}
	return true, nil
}

func GetToken(r *http.Request) (string, error) {
	tokenFromCookie, err := r.Cookie(CookieName)

	if err != nil {

		return "", fmt.Errorf("discordClient.GetToken: %w failed to get token, loging user out", err)
	}

	return tokenFromCookie.Value, nil
}

func RetrieveTokenValue(field string, r *http.Request) (jwt.MapClaims, interface{}, error) {

	claims := jwt.MapClaims{}
	tokenFromCookie, err := r.Cookie(CookieName)

	if err != nil {
		return claims, "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	tokenString := tokenFromCookie.Value

	// Parse the JWT token string
	// FIXME: https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the key used for signing the token (you need this to validate the token)
		// Here, you would typically return the same key that was used for signing the token
		return []byte(utils.Config("JWT_SECRET_KEY")), nil
	})

	// Check for errors during parsing
	if err != nil {
		fmt.Println("discordClient.RetrieveTokenValue.Error parsing JWT token:", err)
		return claims, "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Println("Invalid JWT token")
		return claims, "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Retrieve claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims format")
		return claims, "", fmt.Errorf("discordClient.RetrieveTokenValue: %w", err)
	}

	// Access specific claims
	result := claims[field]
	// username := claims["username"].(string)

	return claims, result, nil
}

func WithSession(ctx context.Context, key interface{}, ctxValue string) context.Context {
	return context.WithValue(ctx, key, ctxValue)
}

func GetSessionValue(ctx context.Context, key interface{}) string {
	value, ok := ctx.Value(key).(string)

	if !ok {
		return ""
	}

	return value

}

func CleanUpContext(ctx context.Context) context.Context {
	// Create a new context derived from the existing one, without the unwanted values
	newCtx := context.Background()

	return newCtx
}

func IsUser(r *http.Request) (bool, error) {

	token, err := GetToken(r)

	if err != nil {

		return false, err
	}

	result, err := ValidateToken(token)

	if err != nil {

		return false, err
	}

	return result, nil

}
