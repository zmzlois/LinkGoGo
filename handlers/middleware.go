package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
)

var publicRoutes = []string{
	"/",
	"/login",
	"/discord-redirect",
	"/discord-oauth",
}

func ProtectedRoutes(r chi.Router) error {

	fmt.Println("Implementing middleware for protected routes")

	grantCode := w.Cookies("discord_code")

	url := ctx.Path()

	for _, route := range publicRoutes {

		// parsing current url in browser
		rawUrl := strings.SplitAfterN(url, "/", 2)[1]

		// parsing the public routes
		rawRoute := strings.TrimSpace(strings.SplitAfterN(route, "/", 2)[1])

		// if the rawRoute is empty(home page) and the current url equal to public routes we continue and do nothing
		if rawRoute == "" && rawRoute == rawUrl {
			continue
		}

		// if it is not home page
		if rawRoute != "" {

			// and the path match public routes + grant code is empty, we move to the next and let it render what it should render
			if strings.Contains(rawUrl, rawRoute) && grantCode == "" {

				return ctx.Next()
			}

			// if the path is a private route and grant code is empty, we redirect to home page
			if !strings.Contains(rawUrl, rawRoute) && grantCode == "" {

				return ctx.Redirect("/", fiber.StatusFound)
			}
		}

	}

	// if none of the scenario above is met, we move to the next middleware
	return ctx.Next()
}

func ProtectedMiddleware(c *fiber.Ctx) http.Handler {

	var next http.Handler

	code := c.Cookies("discord_code")

	if code != "" {
		return next
	}

	if len(code) < 5 {
		return http.RedirectHandler("/", http.StatusBadRequest)
	}
	return http.RedirectHandler("/", http.StatusBadRequest)
}
