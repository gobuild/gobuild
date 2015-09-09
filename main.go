package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Unknwon/macaron"
	"github.com/codeskyblue/gorelease/public"
	"github.com/codeskyblue/gorelease/templates"
	"github.com/macaron-contrib/bindata"
)

var debug = flag.Bool("debug", false, "enable debug mode")

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
func NewRelease(qiniuDomain, os, arch, branch, name, ext string) *Release {
	r := &Release{
		Domain: qiniuDomain,
		OS:     os,
		Arch:   arch,
		Branch: branch,
		Name:   name,
		Ext:    ext,
	}
	r.makeLink()
	return r
}

func (r *Release) makeLink() {
	link := fmt.Sprintf("http://%s/gorelease/%s/%s/%s", r.Domain, r.Branch, r.OS+"-"+r.Arch, r.Name)
	if r.Ext != "" {
		r.Link = link + r.Ext
		return
	}
	if r.OS == "windows" {
		link += ".exe"
	}
	r.Link = link
}

func InitApp(debug bool) *macaron.Macaron {
	app := macaron.Classic()
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

	app.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "homepage")
	})
	app.Get("/:domain/:name/:branch", func(ctx *macaron.Context, r *http.Request) {
		domain := ctx.Params(":domain")
		branch := ctx.Params(":branch")
		name := ctx.Params(":name")
		ctx.Data["Name"] = name
		ctx.Data["Branch"] = branch
		rels := make([]*Release, 0)
		ext := r.FormValue("ext")

		rels = append(rels, NewRelease(domain, "linux", "amd64", branch, name, ext))
		rels = append(rels, NewRelease(domain, "linux", "386", branch, name, ext))
		rels = append(rels, NewRelease(domain, "darwin", "amd64", branch, name, ext))
		rels = append(rels, NewRelease(domain, "darwin", "386", branch, name, ext))
		rels = append(rels, NewRelease(domain, "windows", "amd64", branch, name, ext))
		rels = append(rels, NewRelease(domain, "windows", "386", branch, name, ext))
		ctx.Data["Releases"] = rels
		ctx.HTML(200, "release")
	})
	return app
}

func main() {
	flag.Parse()
	app := InitApp(*debug)

	port := 4000
	fmt.Sscanf(os.Getenv("PORT"), "%d", &port)
	app.Run(port)
}
