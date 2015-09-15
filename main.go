package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Unknwon/macaron"
	"github.com/gorelease/gorelease/models"
	"github.com/gorelease/gorelease/models/goutils"
	"github.com/gorelease/gorelease/public"
	"github.com/gorelease/gorelease/routers"
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
	rdx = models.GetRedisClient()
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
		link = goutils.StrFormat("{name}/downloads/{os}/{arch}", map[string]interface{}{
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

func useBindata(app *macaron.Macaron) {
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
		useBindata(app)
	}

	app.Get("/", routers.Homepage)
	app.Get("/token", oauth2.LoginRequired, routers.Token)
	app.Get("/:owner/:name/downloads/:os/:arch", routers.DownloadRedirect)

	app.Get("/:owner/:name", func(ctx *macaron.Context, r *http.Request) {
		owner := ctx.Params(":owner")
		name := ctx.Params(":name")
		branch := "master"

		// Here need redis connection
		repo := owner + "/" + name
		domain := rdx.Get("domain:" + repo).Val()
		if domain == "" {
			ctx.Error(405, "repo not registed in gorelease, not open register for now")
			return
		}
		log.Println("Domain:", domain)
		rdx.Incr("pageview:" + repo)
		pv, _ := rdx.Get("pageview:" + repo).Int64()

		ctx.Data["PageView"] = pv
		ctx.Data["DlCount"], _ = rdx.Get("downloads:" + repo).Int64()

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
