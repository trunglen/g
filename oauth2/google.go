package oauth2

import (
	"encoding/json"
	"errors"
	"net/http"
)

type GoogleAuth struct {
	Authentication
}

func (g GoogleAuth) Authenticated() error {
	req, _ := http.NewRequest("GET", "https://graph.facebook.com/v2.6/me", nil)
	req.URL.Query().Add("access_token", g.Token)
	var c = http.Client{}
	var resp, err = c.Do(req)
	if err != nil {
		return err
	}
	var result = struct {
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
		AccessType    string `json:"access_type"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return errors.New("Access token not found")
	}
	return nil
}
