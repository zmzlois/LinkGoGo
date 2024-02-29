package auth

import (
	"bytes"
	"fmt"

	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type RequestBody struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

type Auth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (dc *Client) GetAccessToken(c *fiber.Ctx) error {

	code := c.Cookies("discord_code")

	fmt.Println("Code obtained from redirect: ", code)
	// state := c.Cookies("state")

	// dc := new(auth.Client)

	var Url string = "https://discord.com/api/oauth2/token"

	data := RequestBody{
		GrantType:   "authorization_code",
		Code:        code,
		RedirectUri: dc.RedirectUri,
	}

	// auth := Auth{
	// 	ClientId:     dc.ClientId,
	// 	ClientSecret: dc.ClientSecret,
	// }
	requestBody, err := json.Marshal(data)
	authData := fmt.Sprintf("%s,%s \n", dc.ClientId, dc.ClientSecret)

	agent := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(agent)
	agent.Header.SetMethod("POST")

	agent.SetRequestURI(Url)
	agent.SetBody(requestBody)
	agent.Header.SetContentType("application/x-www-form-urlencoded")
	agent.PostArgs().Add("auth", authData)

	fmt.Printf("Post Args: %v\n", agent.PostArgs())

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	client := fasthttp.Client{}

	// Perform the request

	if err := client.Do(agent, resp); err != nil {
		panic(fmt.Sprintf("[fasthttp client error]: %s \n\n", err))
	}

	if err != nil {
		panic(fmt.Sprintf("Request for access token failed with discord: %s \n", err))
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"message": "Request for access token failed with discord: %s",
		// 	"error":   err,
		// })
	}

	fmt.Printf("Response: %v\n", resp)

	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("[Access Token Discord]: expected status code 200, got %d \n\n", resp.StatusCode())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("expected status code 200, got %d \n", resp.StatusCode()),
		})

	}

	// Verify the content type
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		fmt.Printf("Expected content type application/json but got %s\n", contentType)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Expected content type application/json but got %s", contentType),
		})
	}

	// contentEncoding := resp.Header.Peek("Content-Encoding")
	var body []byte
	body = resp.Body()

	fmt.Printf("Response body is: %s", body)

	// pass status code and body received by the proxy
	return c.Status(resp.StatusCode()).Send(body)
}
