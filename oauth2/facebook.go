package oauth2

import (
	"encoding/json"
	"errors"
	"net/http"
)

type FacebookAuth struct {
	Authentication
}

func (f FacebookAuth) Authenticated() error {
	req, _ := http.NewRequest("GET", "https://graph.facebook.com/v2.6/me", nil)
	req.URL.Query().Add("access_token", f.Token)
	var c = http.Client{}
	var resp, err = c.Do(req)
	if err != nil {
		return err
	}
	var result = struct {
		ID   string `json:"id"`
		Name string `json:"id"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return errors.New("Access token not found")
	}
	return nil
}
