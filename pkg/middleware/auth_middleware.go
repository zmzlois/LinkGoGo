package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"go.uber.org/zap"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString, err := auth.GetToken(r)

		if err != nil || tokenString == "" {
			fmt.Printf("AuthMiddleware.GetToken: %v", zap.Error(err))

			http.Redirect(w, r, "/unauthorised", http.StatusFound)
			return
		}

		// Parse and validate JWT token
		ctx := r.Context()
		authenticated, user, userID, err := validateTokenAndGetUser(ctx, tokenString, r)
		if err != nil {
			// Handle token validation error
			fmt.Println("Failed to validate session token:", err)
			http.Error(w, "Failed to validate session token", http.StatusInternalServerError)
			return
		}

		if authenticated {
			ctx = setUserInfoInContext(ctx, user, userID)
		}
		// Token is valid, call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateTokenAndGetUser(ctx context.Context, tokenString string, r *http.Request) (bool, map[string]interface{}, string, error) {
	authenticated, err := auth.ValidateToken(tokenString)
	if err != nil {
		return false, nil, "", err
	}

	user, userID, err := auth.RetrieveTokenValue("user_id", r)
	if err != nil {
		return false, nil, "", err
	}

	return authenticated, user, userID.(string), nil
}

func setUserInfoInContext(ctx context.Context, user map[string]interface{}, userID string) context.Context {

	ctx = auth.WithSession(ctx, "user_id", userID)
	ctx = auth.WithSession(ctx, "username", user["username"].(string))
	ctx = auth.WithSession(ctx, "avatar", user["avatar"].(string))
	ctx = auth.WithSession(ctx, "global_name", user["global_name"].(string))
	return ctx
}
