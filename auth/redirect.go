package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// If using a state, base64encode the data beforehand
func (c *Client) RedirectHandler(ctx *fiber.Ctx, w *fiber.Response, r *fiber.Request, state string) {

	// make a copy of authUrl
	var url string = c.AuthUrl

	// if user give a sate make sure is base64 encoded

	if len(state) > 0 {
		url = c.AuthUrl + "&state=" + state
	}

	// Redirect urser to auth url
	ctx.Redirect(url, http.StatusTemporaryRedirect)
}
