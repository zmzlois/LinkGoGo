package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var publicRoutes = []string{
	"/",
	"/login",
	"/discord-redirect",
	"/discord-oauth",
}

func ProtectedRoutes(ctx *fiber.Ctx) error {

	// Check if the user is authenticated

	// get the code from cookies
	grantCode := ctx.Cookies("discord_code")

	fmt.Println("Grant Code: ", grantCode)

	url := ctx.Path()

	for _, route := range publicRoutes {
		if url != route {
			if grantCode == "" || len(grantCode) < 5 {
				return ctx.Redirect("/", fiber.StatusTemporaryRedirect)
			}
		}
	}

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
