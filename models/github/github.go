package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gobuild/gorelease/models/goutils"
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
	Message     string `json:"message"`
	StatusCode  int    `json:"-"`
	DocumentURL string `json:"documentation_url"`
}

func (e *ErrReturn) Error() string {
	return fmt.Sprintf("code = %d, msg = %s", e.StatusCode, e.Message)
}

func (t *Github) doRequest(method, apiPath string, query url.Values, body io.Reader, v interface{}) error {
	u := &url.URL{
		Scheme: "https",
		Path:   "api.github.com" + apiPath,
	}

	if query != nil {
		u.RawQuery = query.Encode()
	}
	client := &http.Client{}

	req, _ := http.NewRequest(method, u.String(), body)
	req.Header.Set("Authorization", "token "+t.token)

	resp, err := client.Do(req)
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
	if v == nil {
		return nil
	}
	return dec.Decode(v)
}

func (t *Github) doGet(apiPath string, query url.Values, v interface{}) error {
	return t.doRequest("GET", apiPath, query, nil, v)
}

//func (t *Github) loopdoGet(api string, type

func (t *Github) User() (user *User, err error) {
	user = new(User)
	err = t.doGet("/user", nil, user)
	return
}

func (t *Github) Repositories() (repos []*Repository, err error) {
	q := url.Values{}
	q.Set("per_page", "100")
	page := 1
	for {
		var rs []*Repository
		q.Set("page", strconv.Itoa(page))
		err = t.doGet("/user/repos", q, &rs)
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
		err = t.doGet(fmt.Sprintf("/repos/%s/%s/hooks", owner, repo), q, &rs)
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

type CommitFile struct {
	Path    string `json:"-"`
	Message string `json:"message"`
	Content string `json:"content"`
	Branch  string `json:"branch"` // optional, default master
}

func NewCommitFile(path string, message, content, branch string) *CommitFile {
	if branch == "" {
		branch = "master"
	}
	fmt.Println(content)
	return &CommitFile{
		Path:    strings.TrimPrefix(path, "/"),
		Message: message,
		Content: base64.StdEncoding.EncodeToString([]byte(content)),
		Branch:  branch,
	}
}

type FileContent struct {
	Type        string `json:"type"`
	Sha         string `json:"sha"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Size        uint64 `json:"size"`
	Encoding    string `json:"encoding"`
	DownloadURL string `json:"download_url"`
}

// ref
// https://developer.github.com/v3/repos/contents/#update-a-file

func (t *Github) GetFile(owner, repo string, path string) (*FileContent, error) {
	path = strings.TrimPrefix(path, "/")
	fc := new(FileContent)
	err := t.doGet(fmt.Sprintf("/repos/%s/%s/contents/%s",
		owner, repo, path), nil, fc)
	return fc, err
}

func (t *Github) UpdateFile(owner, repo string, file *CommitFile) error {
	fc, err := t.GetFile(owner, repo, file.Path)
	if err != nil {
		return err
	}
	var commitBody = map[string]string{}
	commitBody["sha"] = fc.Sha
	commitBody["message"] = file.Message
	commitBody["content"] = base64.StdEncoding.EncodeToString([]byte(file.Content))
	if file.Branch != "" {
		commitBody["branch"] = file.Branch
	}

	data, _ := json.Marshal(commitBody)
	rd := bytes.NewBuffer(data)
	return t.doRequest("PUT",
		goutils.StrFormat("/repos/{owner}/{repo}/contents/{path}", map[string]interface{}{
			"owner": owner,
			"repo":  repo,
			"path":  file.Path}),
		nil, rd, nil)
}
