package middleware

import (
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"go.uber.org/zap"
)

// If user has a valid token, once they request for /login page they should be redirected to /edit
func RedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user has a valid token, redirect to /edit
		// if user doesn't have a valid token, redirect to /login

		tokenString, err := auth.GetToken(r)

		if err != nil || tokenString == "" || len(tokenString) < 1 {
			fmt.Printf("AuthMiddleware.GetToken: %v", zap.Error(err))

			return
		}

		ctx := r.Context()
		authenticated, err := auth.ValidateToken(tokenString)

		if err != nil {
			fmt.Printf("AuthMiddleware.ValidateToken: %v", zap.Error(err))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if authenticated {
			http.Redirect(w, r, "/edit", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
