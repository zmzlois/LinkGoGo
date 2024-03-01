package handlers

import (
	"fmt"
	"net/http"

	dsc "github.com/realTristan/disgoauth"
)

func RedirectHandler(ds *dsc.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query()["code"][0]

		token, _ := ds.GetOnlyAccessToken(code)

		if len(token) < 5 {
			authFailed(Realm{authenticate: "Unauthorized"})
			w.Write([]byte("Something went wrong when we are authenticating you, please come back again later"))
		}
		// get user data to store in the database and also to display on the page
		result, _ := dsc.GetUserData(token)

		avatar := result["avatar"].(string)
		global_name := result["global_name"].(string)

		fmt.Printf("Avatar: %s\n Global name: %s \n", avatar, global_name)

		fmt.Println(result)

		http.Redirect(w, r, "/edit", http.StatusFound)
	}
}

// func CreateUser(userData map[string]interface{}) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := context.Background()
// 		db, err := DatabaseClient(ctx)

// 		if err != nil {
// 			log.Printf("failed to create database client: %v", err)
// 		}

// 		task, err := db.Users.Create().
// 			SetUsername(userData["username"].(string)).
// 			SetEmail(userData["email"].(string)).
// 			SetAccessToken(userData["access_token"].(string)).
// 			SetAvatar(userData["avatar"].(string)).
// 			SetGlobalName(userData["global_name"].(string)).
// 			Save(ctx)

// 		if err != nil {
// 			w.Write([]byte("Something went wrong when we are creating your account, please come back again later"))
// 		}

// 		fmt.Println(task)

// 	}
// }
