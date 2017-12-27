package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type GoogleAuth struct {
	Authentication
}

func (g GoogleAuth) Authenticated() error {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/tokeninfo?access_token="+g.Token, nil)
	var c = http.Client{}
	var resp, err = c.Do(req)
	if err != nil {
		return err
	}
	var result = struct {
		Email            string `json:"email"`
		EmailVerified    string `json:"email_verified"`
		AccessType       string `json:"access_type"`
		ErrorDescription string `json:"error_description"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result.ErrorDescription)
	if result.ErrorDescription != "" {
		return errors.New(result.ErrorDescription)
	}
	if err != nil {
		return err
	}
	log.Fatal("dasd")
	return nil
}
