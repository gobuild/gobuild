package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Unknwon/macaron"
	"github.com/gorelease/gorelease/github"
	"github.com/gorelease/gorelease/public"
	"github.com/gorelease/gorelease/templates"
	"github.com/gorelease/oauth2"
	"github.com/macaron-contrib/bindata"
	"github.com/macaron-contrib/session"
	goauth2 "golang.org/x/oauth2"
	redis "gopkg.in/redis.v3"
)

var debug = flag.Bool("debug", false, "enable debug mode")
var rdx *redis.Client

func init() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	rdx = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

type Release struct {
	Domain string
	OS     string
	Arch   string
	Name   string
	Ext    string
	Branch string
	Link   string
}

//http://{{.Domain}}/gorelease/{{.Branch}}/windows-amd64/{{.Name}}.exe
func NewRelease(qiniuDomain, os, arch, branch, name, ext string, redirect bool) *Release {
	r := &Release{
		Domain: qiniuDomain,
		OS:     os,
		Arch:   arch,
		Branch: branch,
		Name:   name,
		Ext:    ext,
	}
	r.makeLink(redirect)
	return r
}

func (r *Release) makeLink(redirect bool) {
	var link string
	link = fmt.Sprintf("http://%s/gorelease/%s/%s/%s/%s", r.Domain, r.Name, r.Branch, r.OS+"-"+r.Arch, r.Name)
	if redirect {
		link = StrFormat("{name}/downloads/{os}/{arch}", map[string]interface{}{
			"name": r.Name,
			"os":   r.OS,
			"arch": r.Arch,
		})
	}
	if r.Ext != "" {
		r.Link = link + r.Ext
		return
	}
	if r.OS == "windows" && !redirect {
		link += ".exe"
	}
	r.Link = link
}

func StrFormat(format string, kv map[string]interface{}) string {
	for key, val := range kv {
		key = "{" + key + "}"
		format = strings.Replace(format, key, fmt.Sprintf("%v", val), -1)
	}
	return format
}

func InitApp(debug bool) *macaron.Macaron {
	app := macaron.Classic()
	app.Use(session.Sessioner())
	app.Use(oauth2.Github(
		&goauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			Scopes:       []string{"user:email", "public_repo"},
			RedirectURL:  "",
		},
	))
	if debug {
		app.Use(macaron.Renderer())
	} else {
		app.Use(macaron.Static("public",
			macaron.StaticOptions{
				FileSystem: bindata.Static(bindata.Options{
					Asset:      public.Asset,
					AssetDir:   public.AssetDir,
					AssetNames: public.AssetNames,
					Prefix:     "",
				}),
			},
		))

		app.Use(macaron.Renderer(macaron.RenderOptions{
			TemplateFileSystem: bindata.Templates(bindata.Options{
				Asset:      templates.Asset,
				AssetDir:   templates.AssetDir,
				AssetNames: templates.AssetNames,
				Prefix:     "",
			}),
		}))
	}

	app.Get("/token", oauth2.LoginRequired, func(tokens oauth2.Tokens, ctx *macaron.Context) {
		gh := github.New(tokens)
		user, err := gh.User()
		if err != nil {
			ctx.Error(500, err.Error())
			return
		}
		io.WriteString(ctx.Resp, tokens.Access()+" "+user.Name)
	})

	app.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Host"] = ctx.Req.Host
		ctx.HTML(200, "homepage")
	})

	app.Get("/:owner/:name/downloads/:os/:arch", func(ctx *macaron.Context, r *http.Request) {
		owner := ctx.Params(":owner")
		name := ctx.Params(":name")
		branch := r.FormValue("branch")
		if branch == "" {
			branch = "master"
		}
		repo := owner + "/" + name
		domain := rdx.Get(repo + ":domain").Val()
		if domain == "" {
			ctx.Error(405, "repo not registed in gorelease, not open register for now")
			return
		}
		osarch := ctx.Params(":os") + "-" + ctx.Params(":arch")
		rdx.Incr(repo + ":dlcnt")
		rdx.Incr(repo + ":dlcnt:" + osarch)
		realURL := StrFormat("http://{domain}/gorelease/{name}/{branch}/{osarch}/{name}",
			map[string]interface{}{
				"domain": domain,
				"name":   name,
				"branch": branch,
				"osarch": osarch,
			})
		if ctx.Params(":os") == "windows" {
			realURL += ".exe"
		}
		ctx.Redirect(realURL, 302)
	})

	app.Get("/:owner/:name", func(ctx *macaron.Context, r *http.Request) {
		owner := ctx.Params(":owner")
		name := ctx.Params(":name")
		//domain := "dn-gobuild5.qbox.me"
		branch := "master"

		// Here need redis connection
		repo := owner + "/" + name
		domain := rdx.Get(repo + ":domain").Val()
		if domain == "" {
			ctx.Error(405, "repo not registed in gorelease, not open register for now")
			return
		}
		log.Println("Domain:", domain)
		rdx.Incr(repo + ":pageview")
		pv, _ := rdx.Get(repo + ":pageview").Int64()

		ctx.Data["PageView"] = pv
		ctx.Data["DlCount"], _ = rdx.Get(repo + ":dlcnt").Int64()

		ctx.Data["Name"] = name
		ctx.Data["Branch"] = branch
		ctx.Data["BuildJSON"] = template.URL(fmt.Sprintf(
			"//%s/gorelease/%s/%s/%s", domain, name, branch, "builds.json"))
		rels := make([]*Release, 0)
		ext := r.FormValue("ext")

		rels = append(rels, NewRelease(domain, "linux", "amd64", branch, name, ext, true))
		rels = append(rels, NewRelease(domain, "linux", "386", branch, name, ext, true))
		rels = append(rels, NewRelease(domain, "darwin", "amd64", branch, name, ext, true))
		rels = append(rels, NewRelease(domain, "darwin", "386", branch, name, ext, true))
		rels = append(rels, NewRelease(domain, "windows", "amd64", branch, name, ext, true))
		rels = append(rels, NewRelease(domain, "windows", "386", branch, name, ext, true))
		ctx.Data["Releases"] = rels
		ctx.HTML(200, "release")

	})

	app.Get("/:domain/:name/:branch", func(ctx *macaron.Context, r *http.Request) {
		domain := ctx.Params(":domain")
		branch := ctx.Params(":branch")
		name := ctx.Params(":name")
		ctx.Data["Name"] = name
		ctx.Data["Branch"] = branch
		ctx.Data["BuildJSON"] = template.URL(fmt.Sprintf(
			"//%s/gorelease/%s/%s/%s", domain, name, branch, "builds.json"))
		rels := make([]*Release, 0)
		ext := r.FormValue("ext")

		rels = append(rels, NewRelease(domain, "linux", "amd64", branch, name, ext, false))
		rels = append(rels, NewRelease(domain, "linux", "386", branch, name, ext, false))
		rels = append(rels, NewRelease(domain, "darwin", "amd64", branch, name, ext, false))
		rels = append(rels, NewRelease(domain, "darwin", "386", branch, name, ext, false))
		rels = append(rels, NewRelease(domain, "windows", "amd64", branch, name, ext, false))
		rels = append(rels, NewRelease(domain, "windows", "386", branch, name, ext, false))
		ctx.Data["Releases"] = rels
		ctx.HTML(200, "release")
	})
	return app
}

func main() {
	flag.Parse()
	if err := rdx.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	app := InitApp(*debug)

	port := 4000
	fmt.Sscanf(os.Getenv("PORT"), "%d", &port)
	app.Run(port)
}
