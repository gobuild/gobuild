package models

import (
	"fmt"
	"strconv"

	"github.com/gobuild/gorelease/models/github"
	"github.com/gobuild/gorelease/models/travis"
	yaml "gopkg.in/yaml.v2"
)

var (
	DefaultTriggerOwner = "gorelease"
	DefaultTriggerRepo  = "travis-autobuild"
)

func TriggerTravisBuild(owner, repo, branch, email string) error {
	cfg := travis.DefaultTravisConfig
	cfg.Env.Global = []string{
		"GIT_BRANCH=" + branch,
		"GITHUB_REPO=" + owner + "/" + repo,
	}
	cfg.Notifications.Email.Recipients = []string{email}
	data, _ := yaml.Marshal(cfg)

	gh := github.New(GITHUB_TOKEN)
	return gh.UpdateFile(DefaultTriggerOwner, DefaultTriggerRepo, &github.CommitFile{
		Path:    ".travis.yml",
		Message: fmt.Sprintf("trigger %s by %s", strconv.Quote(owner+"/"+repo), email),
		Content: string(data),
	})
}
