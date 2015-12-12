package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Unknwon/macaron"
	"github.com/gorelease/gorelease/models"
	"github.com/gorelease/gorelease/models/github"
	"github.com/gorelease/oauth2"
)

var rdx = models.GetRedisClient()

func Applications(ctx *macaron.Context) {
	res, err := rdx.Keys("downloads:*/*").Result()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	var jsondata = []interface{}{}
	for _, key := range res {
		if strings.Count(key, ":") != 1 {
			continue
		}
		repoName := strings.Split(key, ":")[1]
		dcnt, _ := rdx.Get(key).Int64()
		jsondata = append(jsondata, map[string]interface{}{
			"name":           repoName,
			"download_count": dcnt,
		})
	}
	ctx.JSON(200, jsondata)
}

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

func RepoList(user *models.User, ctx *macaron.Context) {
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
		go func() {
			err := models.TriggerTravisBuild(repo.Owner, repo.Repo, branch, user.Email)
			if err != nil {
				log.Println(err)
			}
		}()
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

func UserInfo(user *models.User, ctx *macaron.Context) {
	ctx.JSON(200, user)
}
