package api

import (
	"log"
	"strings"

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

func RepoList(tokens oauth2.Tokens, ctx *macaron.Context) {
	gh := github.New(tokens.Access())
	// user, err := gh.User()
	// if err != nil {
	// 	ctx.Error(500, err.Error())
	// 	return
	// }

	repos, err := gh.Repositories()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	ctx.JSON(200, repos)
}
