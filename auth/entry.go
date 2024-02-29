package auth

import (
	"net/http"
)

var RequestClient *http.Client = &http.Client{}

type Strategy string

const (
	Github  Strategy = "Github"
	Discord Strategy = "Discord"
)

type Client struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string   // click on sign in button and then where?
	RefreshUri   string   // Redirect URI for refresh token
	Scopes       []string // Application scope
	AuthUrl      string   // not required, provide one or not providing one are both fine, if none generate one
	Prompt       string   // Consent prompt parameter for auth reapproval
	Implicit     bool     // Whether using implicit grant for access token
}

// Make sure all the required fields are provided
func (c *Client) DiscordFieldValidation() {

	if len(c.ClientId) == 0 {
		panic("Client ID is required for discord oauth.")
	}

	if len(c.ClientSecret) == 0 {
		panic("Client Secret is required for discord oauth.")
	}

	if len(c.RedirectUri) == 0 {
		panic("Redirect URI is required for discord oauth.")
	}

	if len(c.Scopes) < 1 {
		panic("Not enough scopes[] are provided for discord oauth.")
	}

}

// append scopes to OAuth URL, only run if the number of scope is valid

func (c *Client) AppendScopesToAuthUrl(url []byte) string {

	// use append() and a byte slice(url) is faster than using += to a string
	// "..." is unpacking the string into individual bytes.
	url = append(url, "&scope="...)

	for i := 0; i < len(c.Scopes); i++ {

		url = append(url, c.Scopes[i]...)

		// if there are more scopes to append, add a space before the last one
		if i < len(c.Scopes)-1 {
			url = append(url, "%20"...)
		}
	}
	return string(url)
}

// only run if there is no previously provided OAuth URL
func (c *Client) provideAuthUrl() string {

	// Explicit Grant
	var tempUrl string = c.explicitGrant()

	// Implicit Grant

	if c.Implicit {
		tempUrl = c.implicitGrant()

	}

	if len(c.Scopes) > 0 {
		tempUrl = c.AppendScopesToAuthUrl([]byte(tempUrl))
	}
	return tempUrl

}

func DiscordInitialise(c *Client) *Client {
	c.DiscordFieldValidation()

	// normally the url needs to be 40 characters long
	if len(c.AuthUrl) < 40 {
		c.AuthUrl = c.provideAuthUrl()
	}
	return c
}
