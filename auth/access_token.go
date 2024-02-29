package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (dc *Client) GetAccessToken(c *fiber.Ctx) (err error) {

	code := c.Cookies("discord_code")

	// dc := new(auth.Client)

	var (
		Url  string = "https://discordapp.com/api/oauth2/token"
		Body []byte = dc.AccessTokenBody(code)
	)

	agent := fiber.Post(Url)

	agent.Body(c.Body()) // set body received by request
	statusCode, body, errs := agent.Bytes()

	fmt.Println("Status Code: ", statusCode)
	fmt.Println("Body: ", string(body))
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	result := c.Status(statusCode).Send(Body)

	fmt.Println("[GetAccessToken] Result:", result)

	// pass status code and body received by the proxy
	return result
}
