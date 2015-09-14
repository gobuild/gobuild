package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorelease/oauth2"
)

type Github struct {
	token oauth2.Tokens
}

func New(token oauth2.Tokens) *Github {
	return &Github{
		token: token,
	}
}

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

type ErrReturn struct {
	Message    string `json:"message"`
	StatusCode int
}

func (e *ErrReturn) Error() string {
	return fmt.Sprintf("code = %d, msg = %s", e.StatusCode, e.Message)
}

func (t *Github) User() (user *User, err error) {
	user = new(User)
	u := &url.URL{
		Scheme: "https",
		Path:   "api.github.com/user",
	}
	query := u.Query()
	query.Set("access_token", t.token.Access())
	u.RawQuery = query.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusOK {
		err = dec.Decode(user)
		return
	} else {
		er := &ErrReturn{
			StatusCode: resp.StatusCode,
		}
		dec.Decode(er)
		return nil, er
	}
	return
}
