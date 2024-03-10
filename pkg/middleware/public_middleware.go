package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
)

func PublicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("PublicMiddleware")
		ip := getUserIP(w, r)

		fmt.Println("IP: ", ip)

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

// https://blog.golang.org/context/userip/userip.go
func getUserIP(httpWriter http.ResponseWriter, httpServer *http.Request) string {
	var userIP string
	if len(httpServer.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = httpServer.Header.Get("CF-Connecting-IP")
		fmt.Println(net.ParseIP(userIP))
	} else if len(httpServer.Header.Get("X-Forwarded-For")) > 1 {
		userIP = httpServer.Header.Get("X-Forwarded-For")
		fmt.Println(net.ParseIP(userIP))
	} else if len(httpServer.Header.Get("X-Real-IP")) > 1 {
		userIP = httpServer.Header.Get("X-Real-IP")
		fmt.Println(net.ParseIP(userIP))
	} else {
		userIP = httpServer.RemoteAddr
		if strings.Contains(userIP, ":") {
			fmt.Println(net.ParseIP(strings.Split(userIP, ":")[0]))
		} else {
			fmt.Println(net.ParseIP(userIP))
		}
	}
	return userIP
}
