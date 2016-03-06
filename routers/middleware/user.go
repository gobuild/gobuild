package middleware

import (
	"github.com/Unknwon/macaron"
	"github.com/gobuild/gobuild/models"
	"github.com/gobuild/gobuild/models/github"
	"github.com/gobuild/oauth2"
)

func UserNeeded(tokens oauth2.Tokens, ctx *macaron.Context) {
	user := &models.User{GithubToken: tokens.Access()}
	has, err := models.DB.Get(user)
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	if !has {
		gh := github.New(tokens.Access())
		ghuser, err := gh.User()
		if err != nil {
			ctx.Error(500, err.Error())
			return
		}
		user.Name = ghuser.Name
		user.Email = ghuser.Email
		user.Admin = false
		models.DB.Insert(user)
	}
	ctx.Map(user)
}
