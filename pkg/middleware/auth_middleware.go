package middleware

import (
	"log"
	"net/http"

	"github.com/aidenwallis/go-write/write"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
)

var PublicRoutes = []string{
	"/",
	"/login",
	"/discord-redirect",
	"/discord-oauth",
}

// Change this to state
// var privateKey = []byte(database.Config("SECRET_KEY"))

// AuthMiddleware to authenticate users

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// jwtKey := utils.Config("JWT_SECRET_KEY")
		// Extract JWT token from Authorization header
		tokenString, err := auth.GetToken(r)

		_, write := write.New()

		log.Printf("AuthMiddleware.GetToken: %s", tokenString)

		if err != nil || tokenString == "" || len(tokenString) < 1 {
			log.Printf("AuthMiddleware.GetToken: %s", err)

			http.Redirect(w, r, "/unauthorised", http.StatusFound)
			return
		}

		// Parse and validate JWT token
		authenticated, err := auth.ValidateToken(tokenString)
		if err != nil || !authenticated {

			http.Redirect(w, r, "/unauthorised", http.StatusFound)

			return
		}

		// Token is valid, call next handler
		next.ServeHTTP(w, r)
	})
}

// func extractTokenFromHeader(r *http.Request) string {
// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" {
// 		return ""
// 	}

// 	parts := strings.Split(authHeader, " ")
// 	if len(parts) != 2 || parts[0] != "Bearer" {
// 		return ""
// 	}

// 	return parts[1]
// }
