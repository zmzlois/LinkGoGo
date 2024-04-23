package auth

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var RequestClient *http.Client = &http.Client{}

func AuthCred() *Client {

	// Create a new oauth2 config
	var cred *Client = &Client{
		ClientID:           os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret:       os.Getenv("DISCORD_CLIENT_SECRET"),
		RedirectURI:        os.Getenv("DISCORD_REDIRECT_URI"),
		RefreshRedirectURI: os.Getenv("DISCORD_REFRESH_REDIRECT_URI"),
	}

	return cred

}

var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://discord.com/api/oauth2/authorize",
	TokenURL:  "https://discord.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}

type Strategy string

const (
	Github  Strategy = "Github"
	Discord Strategy = "Discord"
)

type Client struct {
	ClientID           string
	ClientSecret       string
	Endpoint           string
	RedirectURI        string   // click on sign in button and then where?
	RefreshRedirectURI string   // Redirect URI for refresh token
	Scopes             []string // Application scope

	OAuthURL string // not required, provide one or not providing one are both fine, if none generate one
	Prompt   string // Consent prompt parameter for auth reapproval
	Implicit bool   // Whether using implicit grant for access token
}

func (dc *Client) checkStructErrors() {
	// Make sure the user has provided
	// a valid client id
	if len(dc.ClientID) < 1 {
		panic("DisGOAuth Error: invalid ClientID in Client (ClientID: string)")
	}
	// Make sure the user has provided
	// a valid client secret
	if len(dc.ClientSecret) < 1 {
		panic("DisGOAuth Error: invalid ClientSecret in Client (ClientSecret: string)")
	}
	// Make sure the user has provided
	// a valid redirect uri
	if len(dc.RedirectURI) < 1 {
		panic("DisGOAuth Error: invalid RedirectURI in Client (RedirectURI: string)")
	}
	// Make sure the user has provided
	// a sufficient number of scopes
	if len(dc.Scopes) == 0 {
		panic("DisGOAuth Error: not enough scopes in Client (Scopes: []string)")
	}
}

// The appendScopes() function is used to append
// the provided scopes to the OAuth URL. This function
// is called from the InitOAuthURL() function and is
// only ran if the number of provided scopes is valid.
//
// Using append() and a byte slice is much faster than
// using += to a string!
func (dc *Client) appendScopes(url []byte) string {
	// Append the initial parameter name (scope)
	url = append(url, "&scope="...)

	// For each of the discord client's scopes
	for i := 0; i < len(dc.Scopes); i++ {
		// Append the scope to the OAuth URL
		url = append(url, dc.Scopes[i]...)

		// If there are multiple scopes and the
		// current index isn't the last scope
		if i != len(dc.Scopes)-1 {
			// Append %20 (space)
			url = append(url, "%20"...)
		}
	}
	return string(url)
}

// The initOAuthURL() function is used to initialize
// a discord OAuth URL. This function is called from
// the Init() function and is only ran if there is
// no previously provided OAuth URL.
func (dc *Client) initOAuthURL() string {
	// Non Implicit OAuth
	var tempUrl string = dc.nonImplicitOAuth() // implicit.go

	// Implicit OAuth
	if dc.Implicit {
		tempUrl = dc.implicitOAuth() // implicit.go
	}
	// If user provided scopes
	if len(dc.Scopes) > 0 {
		// Append the scopes to the OAuth URL
		tempUrl = dc.appendScopes([]byte(tempUrl)) // discord_client.go (this file)
	}
	return tempUrl
}

// The Init() function is used to initalize
// the required data for the discord oauth to work
// It panics if required parameters are missing from
// the provided Client struct
func Init(dc *Client) *Client {
	// Check for Client struct errors
	dc.checkStructErrors() // discord_client.go (this file)

	// Initialize the OAuth URL
	if len(dc.OAuthURL) < 40 {
		dc.OAuthURL = dc.initOAuthURL() // discord_client.go (this file)
	}
	// Return the discord client
	return dc
}

func DiscordInit() *Client {
	var cred = AuthCred()
	return &Client{
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		RedirectURI:  cred.RedirectURI,
		Scopes:       []string{ScopeIdentify},
	}
}
