package middleware

import (
	"fmt"
	"net/http"

	"github.com/aidenwallis/go-write/write"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/model"
	"go.uber.org/zap"
)

var PublicRoutes = []string{
	"/",
	"/login",
	"/discord-redirect",
	"/discord-oauth",
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := auth.GetToken(r)

		if err != nil || tokenString == "" || len(tokenString) < 1 {
			fmt.Printf("AuthMiddleware.GetToken: %v", zap.Error(err))

			http.Redirect(w, r, "/unauthorised", http.StatusFound)
			return
		}

		// Parse and validate JWT token
		ctx := r.Context()
		authenticated, err := auth.ValidateToken(ctx, tokenString)
		if err != nil {

			fmt.Println("failed to validate session token", zap.Error(err))
			_ = write.InternalServerError(w).JSON(model.UnknownError)

			http.Redirect(w, r, "/unauthorised", http.StatusFound)

			return
		}

		if authenticated {
			ctx = auth.WithSession(ctx, "token", tokenString)
		}

		// Token is valid, call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
