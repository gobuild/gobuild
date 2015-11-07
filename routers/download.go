package routers

import (
	"net/http"

	"github.com/Unknwon/macaron"
)

func DownloadStats(ctx *macaron.Context, r *http.Request) {
	org := ctx.Params(":org")
	name := ctx.Params(":name")
	branch := ctx.Params(":branch")
	_ = branch
	repo := org + "/" + name

	osarch := ctx.Params(":os") + "-" + ctx.Params(":arch")
	rdx.Incr("downloads:" + repo)
	rdx.Incr("downloads:" + repo + ":" + osarch)
	ctx.JSON(200, "update success")
}
