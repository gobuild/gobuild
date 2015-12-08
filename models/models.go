package models

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id          uint64
	Name        string
	Email       string `xorm:"unique"`
	GithubToken string `xorm:"github_token"`
	Admin       bool

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type Repository struct {
	Id        uint64
	Owner     string `xorm:"unique(nn)"`
	Repo      string `xorm:"unique(nn)" unique(offcial)`
	UserId    uint64 `xorm:"'user_id'"`
	Official  bool   `xorm:"unique(offcial)"`
	Download  uint64
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
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
