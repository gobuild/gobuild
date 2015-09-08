package main

import (
	"fmt"
	"os"

	"github.com/Unknwon/macaron"
)

var app = macaron.Classic()

func main() {
	app.Use(macaron.Renderer())
	app.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "homepage")
	})
	app.Get("/:domain/:name", func(ctx *macaron.Context) {
		ctx.Data["Domain"] = ctx.Params(":domain")
		ctx.Data["Name"] = ctx.Params(":name")
		ctx.HTML(200, "release")
	})
	port := 4000
	fmt.Sscanf(os.Getenv("PORT"), "%d", &port)
	app.Run(port)
}
