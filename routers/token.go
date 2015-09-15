package routers

import (
	"github.com/Unknwon/macaron"
	"github.com/gorelease/gorelease/models/github"
	"github.com/gorelease/gorelease/models/goutils"
	"github.com/gorelease/oauth2"
)

func Token(tokens oauth2.Tokens, ctx *macaron.Context) {
	gh := github.New(tokens)
	user, err := gh.User()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	repos, err := gh.Repositories()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}

	rdx.Set("user:"+user.Login+":github_token", tokens.Access(), 0)
	tokenKey := "user:" + user.Login + ":token"
	if !rdx.Exists(tokenKey).Val() {
		rdx.Set(tokenKey, "gr"+goutils.RandNString(40), 0)
	}
	token := rdx.Get(tokenKey).Val()
	rdx.SAdd("token:"+token+":orgs", user.Login)
	ctx.Data["User"] = user
	ctx.Data["Token"] = token
	ctx.Data["Repos"] = repos
	ctx.HTML(200, "token")

}
