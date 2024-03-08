package handler

import (
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     auth.CookieName,
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	auth.CleanUpContext(r.Context())

	http.Redirect(w, r, "/", http.StatusFound)
}
