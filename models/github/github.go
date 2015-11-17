package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Github struct {
	token string
}

func New(token string) *Github {
	return &Github{
		token: token,
	}
}

// https://developer.github.com/v3/repos/hooks/#list-hooks
type Hook struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Url    string   `json:"url"`
	Config struct {
		Url         string `json:"url"`
		ContentType string `json:"content_type"`
	} `json:"config"`
}

func (h *Hook) String() string {
	return fmt.Sprintf("Hook %s: %s [%s]", h.Name, h.Config.Url, h.Config.ContentType)
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
	query.Set("access_token", t.token)
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

//func (t *Github) loopDecode(api string, type

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

func (t *Github) Hooks(owner, repo string) (hooks []*Hook, err error) {
	q := url.Values{}
	q.Set("per_page", "100")
	page := 1
	for {
		var rs []*Hook
		q.Set("page", strconv.Itoa(page))
		err = t.decode(fmt.Sprintf("/repos/%s/%s/hooks", owner, repo), q, &rs)
		if err != nil {
			return hooks, err
		}
		if len(rs) == 0 {
			break
		}
		hooks = append(hooks, rs...)
		page += 1
	}
	return
}
