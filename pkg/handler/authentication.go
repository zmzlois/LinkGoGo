package handler

import (
	"fmt"
	"net/http"

	"github.com/zmzlois/LinkGoGo/pkg/auth"
)

func AuthenticationHandler(ds *auth.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var dc auth.Client
		var (
			// Get the code from the redirect parameters (&code=...)
			codeFromURLParamaters = r.URL.Query()["code"][0]

			data, err1 = ds.GetAccessTokenMap(codeFromURLParamaters)

			accessToken = data["token_type"].(string) + " " + data["access_token"].(string)

			// Get the authorized user's data using the above accessToken
			userData, _ = dc.GetUserData(accessToken)
		)

		if err1 != nil {
			panic(fmt.Sprintf("Fialed to get access token: %s", err1))
		}

		CreateUser(userData, data).ServeHTTP(w, r)

		http.Redirect(w, r, "/edit", http.StatusFound)
	}
}

func CreateUser(userData map[string]interface{}, data map[string]interface{}) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("We hit the create user middleware ")
		// ctx := context.Background()

		// fmt.Println("User data: ", userData)

		// // Formatting user's data
		// avatarHash := userData["avatar"]
		// fmt.Println("Avatar hash: ", avatarHash)
		// userId := userData["id"]
		// fmt.Println("User ID: ", userId)
		// avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", userId, avatarHash)

		// fmt.Println("Avatar: ", avatar)

		// global_name := userData["global_name"].(string)
		// fmt.Println("Global name: ", global_name)
		// username := userData["username"].(string)

		// fmt.Println("Username: ", username)
		// // Formatting authentication related data
		// accessToken := data["token_type"].(string) + " " + data["access_token"].(string)
		// fmt.Println("Access token: ", accessToken)
		// refreshToken := data["refresh_token"].(string)
		// fmt.Println("Refresh token: ", refreshToken)
		// expiresIn := data["expires_in"].(float64)
		// fmt.Println("Expires in: ", expiresIn)

		// task, err := db.Users.Create().
		// 	SetExternalID(userId.(string)).
		// 	SetUsername(username).
		// 	SetGlobalName(global_name).
		// 	SetSlug(username).
		// 	SetAvatar(avatar).
		// 	SetAccessToken(accessToken).
		// 	SetRefreshToken(refreshToken).
		// 	SetExpiresIn(expiresIn).
		// 	Save(ctx)

		// if err != nil {
		// 	w.Write([]byte("Something went wrong when we are creating your account, please come back again later"))

		// }

		// fmt.Println(task)

	})

}
