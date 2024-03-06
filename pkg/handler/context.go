package handler

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Realm struct {
	authenticate string `json:"authenticate"`
}

type UserContext struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`

	// Add other relevant user data fields here
}

func AuthMiddlewarer(token string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// if !ok {
			// 	authFailed(Realm{authenticate: "Unauthorized"})
			// 	return
			// }

			// If the user is authenticated, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func authFailed(authenticate Realm) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm=%s`, authenticate))
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/failed", http.StatusUnauthorized)
		})
	}

}

func HashValue(v string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CompareToHash(v, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(v))
	return err == nil
}
