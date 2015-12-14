package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Unknwon/macaron"
	"github.com/gobuild/gorelease/models"
	"github.com/gobuild/gorelease/models/github"
	"github.com/gorelease/oauth2"
)

var rdx = models.GetRedisClient()

func TriggerBuild(tokens oauth2.Tokens, ctx *macaron.Context) {
	r := ctx.Req
	owner, repo := r.FormValue("owner"), r.FormValue("repo")
	branch := r.FormValue("branch")

	gh := github.New(tokens.Access())
	user, err := gh.User()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}

	go func() {
		err := models.TriggerTravisBuild(owner, repo, branch, user.Email)
		if err != nil {
			log.Println(err)
		}
	}()
	ctx.JSON(200, map[string]string{
		"message": "build is triggered",
	})
}

type RepositoryInfo struct {
	IsCgo bool
	IsCmd bool
}

var ErrURL404 = errors.New("URL 404 error")

func checkURL(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return ErrURL404
	}
	return nil
}

// FIXME(ssx): not finished
func GetRepoInfo(owner, repo string) (*RepositoryInfo, error) {
	url1 := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/master/main.go",
		owner, repo)
	url2 := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/master/%s.go",
		owner, repo, repo)
	if checkURL(url1) != nil && checkURL(url2) != nil {
		return nil, errors.New("Repo maybe not a golang cli repo")
	}
	return &RepositoryInfo{
		IsCgo: true,
		IsCmd: true,
	}, nil
}

var LimitBuildInterval = time.Minute * 10

func triggerBuild(owner, repo, ref string, email string) error {
	var mr = &models.Repository{
		Owner: owner,
		Repo:  repo,
	}
	has, err := models.DB.Get(mr)
	if err != nil {
		return err
	}

	var dur time.Duration
	if !has {
		// not exists in db, create one
		ri, err := GetRepoInfo(owner, repo)
		if err != nil {
			return err
		}
		if !ri.IsCmd {
			return errors.New("repo is not a go cli repo")
		}

		mr.Valid = true
		mr.TriggerAt = time.Now()
		if _, err = models.DB.Insert(mr); err != nil {
			return err
		}
		dur = LimitBuildInterval + time.Minute

	} else {
		dur = time.Since(mr.TriggerAt)
	}

	if dur < LimitBuildInterval { // too offen is not good
		return fmt.Errorf("Too offen, retry after %v", LimitBuildInterval-dur)
	}

	// update trigger time
	if has {
		mr.TriggerAt = time.Now()
		models.DB.Id(mr.Id).Update(mr)
	}

	go func() {
		err := models.TriggerTravisBuild(owner, repo, ref, email)
		if err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func AnonymousTriggerBuild(ctx *macaron.Context) {
	r := ctx.Req
	owner, repo := r.FormValue("owner"), r.FormValue("repo")
	branch := r.FormValue("branch")

	if owner == "" || repo == "" {
		ctx.Error(500, "owner or repo should not be empty")
		return
	}

	var mrepo = &models.Repository{
		Owner: owner,
		Repo:  repo,
	}
	has, err := models.DB.Get(mrepo)
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	if has && mrepo.UserId != 0 {
		ctx.Error(500, "This repo is limited to its author to build") // TODO: show who is owned
		return
	}

	if err := triggerBuild(owner, repo, branch, "codeskyblue@gmail.com"); err != nil {
		ctx.Error(500, err.Error())
		return
	}

	ctx.JSON(200, map[string]string{
		"message": "build is triggered",
	})
}

func RepoList(ctx *macaron.Context) {
	var repos []models.Repository

	// TODO: change limit to paginate
	err := models.DB.Limit(100).Where("valid=?", true).Desc("updated_at").Find(&repos)
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	ctx.JSON(200, repos)
}

func UserRepoList(user *models.User, ctx *macaron.Context) {
	if ctx.Req.Method == "POST" {
		if time.Since(user.RepoUpdatedAt) > time.Minute {
			if err := user.SyncGithub(); err != nil {
				ctx.JSON(500, err.Error())
			} else {
				ctx.JSON(200, map[string]string{
					"message": "sync github success",
				})
			}
		} else {
			ctx.JSON(200, map[string]string{
				"message": "try after a minute",
			})
		}
		return
	}

	if ctx.Req.Method != "GET" {
		ctx.JSON(500, map[string]string{
			"message": fmt.Sprintf("Method %s not supported", ctx.Req.Method),
		})
		return
	}

	repos, err := user.Repositories()

	if err != nil {
		ctx.Error(500, err.Error())
		return
	}

	if len(repos) == 0 && time.Since(user.RepoUpdatedAt) > time.Hour*24 {
		user.SyncGithub()
		repos, _ = user.Repositories()
	}
	ctx.JSON(200, repos)
}

func getRepository(id int64) (*models.Repository, error) {
	var repo = &models.Repository{Id: id}
	exists, err := models.DB.Get(repo)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("Repo id:%d not found", id)
	}
	return repo, nil
}

func RepoBuild(user *models.User, ctx *macaron.Context) {
	repo, err := getRepository(ctx.ParamsInt64(":id"))
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	branch := ctx.Req.FormValue("branch")
	if ctx.Req.Method == "POST" {
		err := triggerBuild(repo.Owner, repo.Repo, "", user.Email)
		if err != nil {
			ctx.Error(500, err.Error())
			return
		}
		// dur := time.Since(repo.TriggerAt)
		// var limitTime = time.Minute * 10
		// if dur < limitTime { // too offen is not good
		// 	ctx.Error(500, fmt.Sprintf("Too offen, retry after %v", limitTime-dur))
		// 	return
		// }

		// repo.TriggerAt = time.Now()
		// repo.Valid = true
		// log.Println("%v", repo)
		// models.DB.Id(repo.Id).Cols("trigger_at", "valid").Update(repo)

		// go func() {
		// 	err := models.TriggerTravisBuild(repo.Owner, repo.Repo, branch, user.Email)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// }()
		ctx.JSON(200, map[string]string{
			"message": fmt.Sprintf("build for repo: %s is triggered", strconv.Quote(repo.Owner+"/"+repo.Repo)),
		})
		return
	}
	if ctx.Req.Method == "PUT" {
		if err := repo.AddBranch(branch); err != nil {
			ctx.Error(500, err.Error())
			return
		}

		ctx.JSON(200, map[string]string{
			"message": "change to valid",
		})
		return
	}

	// other methods
	ctx.JSON(500, map[string]string{
		"message": fmt.Sprintf("Method %s not supported", ctx.Req.Method),
	})
}

func RecentBuild(ctx *macaron.Context) {
	var repos []models.Repository
	err := models.DB.Limit(10).Desc("trigger_at").Where("valid=?", true).Find(&repos)
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	ctx.JSON(200, repos)
}

func UserInfo(user *models.User, ctx *macaron.Context) {
	ctx.JSON(200, user)
}
