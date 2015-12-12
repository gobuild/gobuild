package models

import (
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gorelease/gorelease/models/github"
)

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `xorm:"unique" json:"email"`
	GithubToken string `xorm:"github_token" json:"github_token"`
	Admin       bool   `json:"admin"`

	CreatedAt     time.Time `xorm:"created" json:"created_at"`
	UpdatedAt     time.Time `xorm:"updated" json:"updated_at"`
	RepoUpdatedAt time.Time `json:"repo_updated_at"`
}

type Repository struct {
	Id        int64     `json:"id"`
	Owner     string    `xorm:"unique(nn)" json:"owner"`
	Repo      string    `xorm:"unique(nn)" json:"repo"`
	UserId    int64     `xorm:"'user_id'" json:"-"`
	Refs      []string  `json:"refs"` // can be branch or tag
	Valid     bool      `json:"valid"`
	Official  bool      `json:"official"`
	Download  uint64    `json:"download"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	TriggerAt time.Time `json:"trigger_at"`
}

var DB *xorm.Engine

func init() {
	var err error
	DB, err = xorm.NewEngine("mysql", MYSQL_URI)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Sync(new(User), new(Repository)); err != nil {
		log.Fatal(err)
	}
}

func (user *User) SyncGithub() error {
	gh := github.New(user.GithubToken)
	repos, err := gh.Repositories()
	if err != nil {
		return err
	}

	for _, ghRepo := range repos {
		parts := strings.Split(ghRepo.Fullname, "/")
		if len(parts) != 2 {
			continue
		}
		var repo = &Repository{
			Owner: parts[0],
			Repo:  parts[1],
		}
		exists, err := DB.Get(repo)
		if err != nil {
			return err
		}
		repo.UserId = user.Id
		if exists {
			repoId := repo.Id
			_, err := DB.Update(repo, &Repository{Id: repoId})
			if err != nil {
				return err
			}
		} else {
			_, err := DB.Insert(repo)
			if err != nil {
				return err
			}
		}
	}
	user.RepoUpdatedAt = time.Now()
	_, err = DB.Id(user.Id).Cols("repo_updated_at").Update(user)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (user *User) Repositories() ([]Repository, error) {
	var repos []Repository
	err := DB.Find(&repos, &Repository{UserId: user.Id})
	return repos, err
}

func (r *Repository) AddBranch(name string) error {
	// CDN_URL_BASE + "/"
	// r.Branches
	return nil
}
