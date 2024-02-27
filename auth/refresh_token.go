package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) accessTokenBody(code string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s&scope=%s",
		c.ClientId, c.ClientSecret, c.RedirectUri, code, ScopeIdentify)))
}

func (c *Client) refreshAccessTokenBody(refreshToken string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=refresh_token&redirect_uri=%s&refresh_token=%s",
		c.ClientId, c.ClientSecret, c.RedirectUri, refreshToken,
	)))
}

func (c *Client) accessTokenRequestObject(body *bytes.Buffer, creds bool) (*http.Request, error) {

	var req, err = http.NewRequest("POST", "https://discordapp.com/api/oauth2/token", body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	}

	// if using the credentials grant, set the request object's headers

	if creds {
		// Base64 encode the client id and client secret
		var auth string = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ClientId, c.ClientSecret)))

		// Set the request object's headers
		req.Header["Authorization"] = []string{"Basic " + auth}
	}

	return req, nil
}

func (c *Client) accessTokenRequest(req *http.Request) (map[string]interface{}, error) {

	// Send the request
	var res, err = RequestClient.Do(req)

	// Handle the response status code
	if res.StatusCode != 200 || err != nil {

		// read the http body
		body, _err := io.ReadAll(res.Body)

		// Handle the response body error
		if _err != nil {
			return map[string]interface{}{}, _err
		}

		// Handle http response error
		return map[string]interface{}{},
			fmt.Errorf("status: %d, code: %v, body: %s",
				res.StatusCode, err, string(body))
	}

	// Store the response body
	var data map[string]interface{}

	// Decode the response body
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return map[string]interface{}{}, err

	}

	// Return the token data
	return data, nil
}

// get the users bearer access token.
func (c *Client) GetOnlyAccessToken(code string) (string, error) {

	// define variables

	var (

		// The access token request body
		tokenBody *bytes.Buffer = c.accessTokenBody(code)
		//Access token request object
		tokenReq, err = c.accessTokenRequestObject(tokenBody, false)
	)

	// Handle the token request object error
	if err != nil {
		return "", err
	}
	// Get the token data map
	var data, _err = c.accessTokenRequest(tokenReq)

	// Handle the token request error
	if _err != nil {
		return "", _err
	}
	// The Bearer access token
	var accessToken string = data["token_type"].(string) + " " + data["access_token"].(string)

	// Return the bearer token from said data
	return accessToken, nil
}

func (c *Client) RefreshAccessToken(refreshToken string) (map[string]interface{}, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = c.refreshAccessTokenBody(refreshToken)
		// The Access Token Request Object
		tokenReq, err = c.accessTokenRequestObject(tokenBody, false)
	)
	// Handle the token request object error
	if err != nil {
		return map[string]interface{}{}, err
	}
	return c.accessTokenRequest(tokenReq)
}
