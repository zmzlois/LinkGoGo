package auth

import "fmt"

// use the implicit/less-safe response type for getting the users access token
func (c *Client) implicitGrant() string {

	// return the Auth URL to a formatted string that contains the client id, redirect uri and response type
	return fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=token",
		c.ClientID,
		c.RedirectURI,
	)
}

// use the safer/default response type for getting the users access token
func (c *Client) explicitGrant() string {

	// establish the prompt parameter
	var prompt string = c.Prompt

	if len(c.Prompt) > 0 {
		prompt = "&prompt=" + c.Prompt
	}

	// return the OAuth URL to a formatted string containing the client id, redirect uri, response type and prompt parameter
	return fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code%s",
		c.ClientID,
		c.RefreshRedirectURI,
		prompt,
	)

}
