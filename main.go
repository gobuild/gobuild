package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Unknwon/macaron"
	"github.com/franela/goreq"
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

func (r *Repo) Pub(goos, arch string) *Publish {
	link := goutils.StrFormat("http://{domain}/gorelease/{org}/{name}/{branch}/{os}-{arch}/{name}", map[string]interface{}{
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

	app.Get("/:org/:name", func(ctx *macaron.Context, r *http.Request) {
		org := ctx.Params(":org")
		name := ctx.Params(":name")
		branch := "master"

		// Here need redis connection
		repoPath := org + "/" + name
		domain := "dn-gobuild5.qbox.me"
		buildJson := fmt.Sprintf("//%s/gorelease/%s/%s/%s/%s", domain, org, name, branch, "builds.json")
		res, err := goreq.Request{
			Method: "HEAD",
			Uri:    "http:" + buildJson,
		}.Do()
		log.Println(res.StatusCode, res.Status)
		if err != nil || res.StatusCode != http.StatusOK {
			ctx.Error(406, "No downloads avaliable now.: +"+strconv.Itoa(res.StatusCode))
			return
		}

		rdx.Incr("pageview:" + repoPath)
		pv, _ := rdx.Get("pageview:" + repoPath).Int64()

		ctx.Data["PageView"] = pv
		ctx.Data["DlCount"], _ = rdx.Get("downloads:" + repoPath).Int64()

		ctx.Data["Name"] = name
		ctx.Data["Branch"] = branch
		ctx.Data["BuildJSON"] = template.URL(buildJson)
		//rels := make([]*Release, 0)
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
	flag.Parse()
	if err := rdx.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	app := InitApp(*debug)

	port := 4000
	fmt.Sscanf(os.Getenv("PORT"), "%d", &port)
	app.Run(port)
}
