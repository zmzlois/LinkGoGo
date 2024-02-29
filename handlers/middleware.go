package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var publicRoutes = []string{
	"/",
	"/login",
	"/discord-redirect",
	"/discord-oauth",
}

func ProtectedRoutes(ctx *fiber.Ctx) error {

	fmt.Println("Implementing middleware for protected routes")

	grantCode := ctx.Cookies("discord_code")

	fmt.Println("Grant Code: ", grantCode, "Length: ", len(grantCode))

	url := ctx.Path()

	for _, route := range publicRoutes {

		rawUrl := strings.SplitAfterN(url, "/", 2)[1]

		rawRoute := strings.TrimSpace(strings.SplitAfterN(route, "/", 2)[1])

		if rawRoute == "" && rawRoute == rawUrl {
			continue
		}

		if rawRoute != "" {

			if strings.Contains(rawUrl, rawRoute) && grantCode == "" {
				fmt.Println("User is in public route and has no grant code, we do nothing.")
				return ctx.Next()
			}

			fmt.Printf("Does the url contains any public route? %v \n url: |%s| public route:|%s|\n", strings.Contains(rawUrl, rawRoute), rawUrl, rawRoute)

			if !strings.Contains(rawUrl, rawRoute) && grantCode == "" {
				fmt.Println("User is in protected route and has no grant code, redirecting to home page.")

				return ctx.Redirect("/", fiber.StatusFound)

			}
		}

	}

	fmt.Println("User has a grant code, continuing to next middleware")
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
