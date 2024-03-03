package handlers

import (
	"context"
	"fmt"
	"net/http"

	dsc "github.com/zmzlois/LinkGoGo/auth"
)

func AuthenticationHandler(ds *dsc.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var dc dsc.Client
		var (
			// Get the code from the redirect parameters (&code=...)
			codeFromURLParamaters = r.URL.Query()["code"][0]

			// Get the access token using the above codeFromURLParamaters
			// accessTokenMap, _ = ds.GetAccessTokenMap(codeFromURLParamaters)

			data, err1 = ds.GetAccessTokenMap(codeFromURLParamaters)

			accessToken = data["token_type"].(string) + " " + data["access_token"].(string)

			// Get the authorized user's data using the above accessToken
			userData, _ = dc.GetUserData(accessToken)
		)

		if err1 != nil {
			panic(fmt.Sprintf("Fialed to get access token: %s", err1))
		}

		var next http.Handler
		// make this a promise
		CreateUser(userData, data)(next)

		http.Redirect(w, r, "/edit", http.StatusFound)
	}
}

func CreateUser(userData map[string]interface{}, data map[string]interface{}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			db := DatabaseClient()

			// Formatting user's data
			avatarHash := userData["avatar"].(string)
			userId := userData["id"].(string)
			avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", userId, avatarHash)
			global_name := userData["global_name"].(string)
			username := userData["username"].(string)

			// Formatting authentication related data
			accessToken := data["token_type"].(string) + " " + data["access_token"].(string)
			refreshToken := data["refresh_token"].(string)
			expiresIn := data["expires_in"].(float64)

			task, err := db.Users.Create().
				SetExternalID(userId).
				SetUsername(username).
				SetGlobalName(global_name).
				SetSlug(username).
				SetAvatar(avatar).
				SetAccessToken(accessToken).
				SetRefreshToken(refreshToken).
				SetExpiresIn(expiresIn).
				Save(ctx)

			if err != nil {
				w.Write([]byte("Something went wrong when we are creating your account, please come back again later"))
			}

			fmt.Println(task)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
