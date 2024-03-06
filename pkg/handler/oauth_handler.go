package handler

import (
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/utils"
)

func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	var cred = auth.AuthCred()
	state, err := utils.StateGenerator()

	if err != nil {
		panic("Error generating state")
	}

	var ds *auth.Client = auth.Init(&auth.Client{
		RedirectURI: cred.RedirectURI,
		// RefreshRedirectURI: cred.RefreshRedirectURI,
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		Scopes:       []string{auth.ScopeIdentify},
	})

	ds.RedirectHandler(w, r, state)
}
