package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// send an api request to the discord/users/@me and get user data via accessToken

func GetDiscordUserData(token string) (map[string]interface{}, error) {

	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)

	if err != nil {
		return map[string]interface{}{}, err
	}

	// set reqeuest object's headers
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}

	res, err := RequestClient.Do(req)

	if res.StatusCode != 200 || err != nil {

		// read the http body
		body, _err := io.ReadAll(res.Body)

		if _err != nil {
			return map[string]interface{}{}, _err
		}

		return map[string]interface{}{},
			fmt.Errorf("status: %d, code: %v, body: %s",
				res.StatusCode, err, string(body))

	}

	// store response body (use Fiber to cache/session later)
	var data map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil

}
