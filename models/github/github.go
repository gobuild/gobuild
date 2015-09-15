package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

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
	Login   string `json:"login"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

type Repository struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Fullname string `json:"full_name"`
	Private  bool   `json:"private"`
	Fork     bool   `json:"fork"`
	HtmlURL  string `json:"html_url"`
}

type ErrReturn struct {
	Message    string `json:"message"`
	StatusCode int
}

func (e *ErrReturn) Error() string {
	return fmt.Sprintf("code = %d, msg = %s", e.StatusCode, e.Message)
}

func (t *Github) decode(apiPath string, query url.Values, v interface{}) error {
	u := &url.URL{
		Scheme: "https",
		Path:   "api.github.com" + apiPath,
	}
	if query == nil {
		query = u.Query()
	}
	query.Set("access_token", t.token.Access())
	u.RawQuery = query.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := &ErrReturn{
			StatusCode: resp.StatusCode,
		}
		dec.Decode(er)
		return er
	}
	return dec.Decode(v)
}

func (t *Github) User() (user *User, err error) {
	user = new(User)
	err = t.decode("/user", nil, user)
	return
}

func (t *Github) Repositories() (repos []*Repository, err error) {
	q := url.Values{}
	q.Set("per_page", "100")
	page := 1
	for {
		var rs []*Repository
		q.Set("page", strconv.Itoa(page))
		err = t.decode("/user/repos", q, &rs)
		if err != nil {
			return repos, err
		}
		if len(rs) == 0 {
			break
		}
		repos = append(repos, rs...)
		page += 1
	}
	return
}
