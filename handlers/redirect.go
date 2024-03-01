package handlers

import (
	"fmt"
	"net/http"

	dsc "github.com/realTristan/disgoauth"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request, ds *dsc.Client) {
	code := r.URL.Query()["code"][0]

	token, _ := ds.GetOnlyAccessToken(code)

	if len(token) > 1 {

	}
	// get user data to store in the database and also to display on the page
	result, _ := dsc.GetUserData(token)

	fmt.Println(result)

	http.Redirect(w, r, "/edit", http.StatusFound)
}
