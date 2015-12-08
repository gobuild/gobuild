package models

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `xorm:"unique" json:"email"`
	GithubToken string `xorm:"github_token" json:"github_token"`
	Admin       bool   `json:"admin"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

type Repository struct {
	Id        int64     `json:"id"`
	Owner     string    `xorm:"unique(nn)" json:"owner"`
	Repo      string    `xorm:"unique(nn) unique(offcial)" json:"repo"`
	UserId    uint64    `xorm:"'user_id'" json:"-"`
	Official  bool      `xorm:"unique(offcial)" json:"official"`
	Download  uint64    `json:"download"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	// OSArch []OSArch
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
