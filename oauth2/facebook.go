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
	req, _ := http.NewRequest("GET", "https://graph.facebook.com/v2.6/me?access_token="+f.Token, nil)
	var c = http.Client{}
	var resp, err = c.Do(req)
	if err != nil {
		return err
	}
	var result = struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Error *struct {
			Message string `json:"message"`
			Code    string `json:"code"`
		} `json:"error"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	if result.Error != nil {
		return errors.New(result.Error.Message)
	}
	return nil
}
