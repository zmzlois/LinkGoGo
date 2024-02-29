package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (c *Client) AccessTokenBody(code string) []byte {

	if len(c.ClientId) < 1 || len(c.ClientSecret) < 1 {
		panic("Client ID and Client Secret are required for oauth strategy:")
	}

	return []byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s&scope=%s",
		c.ClientId, c.ClientSecret, c.RedirectUri, code, ScopeIdentify))
}

func (c *Client) RefreshAccessTokenBody(refreshToken string) []byte {
	return []byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=refresh_token&redirect_uri=%s&refresh_token=%s",
		c.ClientId, c.ClientSecret, c.RedirectUri, refreshToken,
	))
}

func (c *Client) GetAccessTokenV1(body []byte, creds bool) (err error) {

	// req := &fiber.Request{}
	ctx := fiber.Ctx{}

	var (
		// Method string = "POST"
		Url  string = "https://discordapp.com/api/oauth2/token"
		Body []byte = body
	)

	agent := fiber.Post(Url).Body(Body)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errs,
		})
	}

	return ctx.Status(statusCode).Send(body)
}

// func (dc *Client) AccessTokenRequest(req *fiber.Request) (map[string]interface{}, error) {

// 	// Send the request
// 	var res, err = client(req)

// 	// Handle the response status code
// 	if res.StatusCode != 200 || err != nil {

// 		// read the http body
// 		body, _err := io.ReadAll(res.Body)

// 		// Handle the response body error
// 		if _err != nil {
// 			return map[string]interface{}{}, _err
// 		}

// 		// Handle http response error
// 		return map[string]interface{}{},
// 			fmt.Errorf("status: %d, code: %v, body: %s",
// 				res.StatusCode, err, string(body))
// 	}

// 	// Store the response body
// 	var data map[string]interface{}

// 	// Decode the response body
// 	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
// 		return map[string]interface{}{}, err

// 	}

// 	// Return the token data
// 	return data, nil
// }

// func (c *Client) RefreshAccessToken(refreshToken string) (map[string]interface{}, error) {
// 	// Define Variables
// 	var (
// 		// The Access Token Request Body
// 		tokenBody *bytes.Buffer = c.RefreshAccessTokenBody(refreshToken)
// 		// The Access Token Request Object
// 		tokenReq, err = c.AccessTokenRequestObject(tokenBody, false)
// 	)
// 	// Handle the token request object error
// 	if err != nil {
// 		return map[string]interface{}{}, err
// 	}
// 	return c.AccessTokenRequest(tokenReq)
// }
