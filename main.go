package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Unknwon/macaron"
	"github.com/gobuild/gorelease/models"
	"github.com/gobuild/gorelease/models/goutils"
	"github.com/gobuild/gorelease/routers"
	"github.com/gobuild/gorelease/routers/api"
	"github.com/gobuild/gorelease/routers/middleware"
	"github.com/gorelease/oauth2"
	"github.com/macaron-contrib/session"
	goauth2 "golang.org/x/oauth2"
	redis "gopkg.in/redis.v3"
)

var rdx *redis.Client

func init() {
	rdx = models.GetRedisClient()
}

type Repo struct {
	Domain string
	Org    string
	Name   string
	Branch string
	Ext    string
}

type Publish struct {
	OS     string
	Arch   string
	Branch string
	Link   string
}

// generate each os-arch address
func (r *Repo) Pub(goos, arch string) *Publish {
	link := goutils.StrFormat("http://{domain}/gorelease/{org}/{name}/{branch}/{os}-{arch}/{name}",
		map[string]interface{}{
			"domain": r.Domain,
			"org":    r.Org,
			"name":   r.Name,
			"branch": r.Branch,
			"os":     goos,
			"arch":   arch,
		})
	if goos == "windows" {
		link += ".exe"
	}
	return &Publish{
		OS:     goos,
		Arch:   arch,
		Branch: r.Branch,
		Link:   link,
	}
}

func InitApp() *macaron.Macaron {
	app := macaron.Classic()
	app.Use(macaron.Static("public"))
	app.Use(session.Sessioner())
	app.Use(oauth2.Github(
		&goauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			// https://developer.github.com/v3/oauth/#scopes
			Scopes:      []string{"user:email", "public_repo", "write:repo_hook"},
			RedirectURL: "",
		},
	))

	app.Use(macaron.Renderer(macaron.RenderOptions{
		Delims: macaron.Delims{"{[", "]}"},
	}))

	app.Get("/", routers.Homepage)
	app.Get("/build", oauth2.LoginRequired, routers.Build)
	app.Post("/stats/:org/:name/:branch/:os/:arch", routers.DownloadStats)

	app.Get("/:org/:name", func(ctx *macaron.Context, r *http.Request) {
		org := ctx.Params(":org")
		//name := ctx.Params(":name")
		if org == "js" {
			ctx.Next()
			return
		}
		ctx.Redirect(r.RequestURI+"/"+"master", 302)
	})

	// api
	app.Group("/api", func() {
		app.Get("/repos", api.RepoList)
		app.Post("/repos", api.AnonymousTriggerBuild)
		app.Any("/user/repos", oauth2.LoginRequired, middleware.UserNeeded, api.UserRepoList)
		app.Get("/recent/repos", api.RecentBuild)
		app.Post("/builds", oauth2.LoginRequired, api.TriggerBuild)

		// accept PUT(callback), POST(trigger build)
		app.Any("/repos/:id/build", oauth2.LoginRequired, middleware.UserNeeded, api.RepoBuild)
		app.Get("/user", oauth2.LoginRequired, middleware.UserNeeded, api.UserInfo)
	})

	app.Get("/explore", func(ctx *macaron.Context) {
		ctx.HTML(200, "explore", nil)
	})
	app.Get("/repos", func(ctx *macaron.Context) {
		ctx.HTML(200, "repos", nil)
	})

	app.Get("/:org/:name/:branch", func(ctx *macaron.Context, r *http.Request) {
		org := ctx.Params(":org")
		name := ctx.Params(":name")
		branch := ctx.Params(":branch")

		// Here need redis connection
		repoPath := org + "/" + name
		domain := "dn-gobuild5.qbox.me"
		buildJson := fmt.Sprintf("//%s/gorelease/%s/%s/%s/%s", domain, org, name, branch, "builds.json")

		ctx.Data["DlCount"], _ = rdx.Get("downloads:" + repoPath).Int64()
		ctx.Data["Org"] = org
		ctx.Data["Name"] = name
		ctx.Data["Branch"] = branch
		ctx.Data["BuildJSON"] = template.URL(buildJson)
		prepo := &Repo{
			Domain: domain,
			Org:    org,
			Name:   name,
			Branch: branch,
			Ext:    "",
		}
		pubs := make([]*Publish, 0)
		pubs = append(pubs, prepo.Pub("darwin", "amd64"))
		pubs = append(pubs, prepo.Pub("linux", "amd64"))
		pubs = append(pubs, prepo.Pub("linux", "386"))
		pubs = append(pubs, prepo.Pub("linux", "arm"))
		pubs = append(pubs, prepo.Pub("windows", "amd64"))
		pubs = append(pubs, prepo.Pub("windows", "386"))
		ctx.Data["Pubs"] = pubs
		ctx.HTML(200, "release")
	})
	return app
}

func main() {
	if err := rdx.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	app := InitApp()

	port := 4000
	fmt.Sscanf(os.Getenv("PORT"), "%d", &port)
	app.Run(port)
}
