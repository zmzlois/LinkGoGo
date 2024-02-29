package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/zmzlois/LinkGoGo/auth"
)

func AccessToken(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != state {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("State does not match."))

	}

	// Step 3: We exchange the code we got for an access token
	// Then we can use the access token to do actions, limited to scopes we requested
	token, err := conf.Exchange(context.Background(), r.FormValue("code"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}

	// Step 4: Use the access token, here we use it to get the logged in user's info.
	res, err := conf.Client(context.Background(), token).Get(auth.Endpoint.TokenURL)

	if err != nil || res.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			w.Write([]byte(err.Error()))

		} else {
			w.Write([]byte(res.Status))
		}

	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}
	fmt.Println("Give me the fk body: ", body)

	w.Write(body)

}
