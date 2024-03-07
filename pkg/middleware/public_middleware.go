package middleware

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
)

func PublicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("PublicMiddleware")

		user, _, err := auth.RetrieveTokenValue("username", r)

		if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
			if err == nil {
				hub.ConfigureScope(func(scope *sentry.Scope) {
					scope.SetUser(sentry.User{
						ID:        user["user_id"].(string),
						Email:     "",
						IPAddress: r.RemoteAddr,
						Username:  user["username"].(string),
						Name:      user["global_name"].(string),
						Segment:   "",
						Data:      make(map[string]string),
					})
				})
			} else {
				hub.CaptureException(err)
			}
		}

		next.ServeHTTP(w, r)
	})
}
