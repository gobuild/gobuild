package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Unknwon/macaron"
	"github.com/codeskyblue/gorelease/public"
	"github.com/codeskyblue/gorelease/templates"
	"github.com/macaron-contrib/bindata"
)

var debug = flag.Bool("debug", false, "enable debug mode")

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
	app.Get("/:domain/:name", func(ctx *macaron.Context) {
		ctx.Data["Domain"] = ctx.Params(":domain")
		ctx.Data["Name"] = ctx.Params(":name")
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
