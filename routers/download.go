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

	domain := rdx.Get("domain:" + repo).Val()
	if domain == "" {
		ctx.Error(405, "repo not registed in gorelease, not open register for now")
		return
	}

	osarch := ctx.Params(":os") + "-" + ctx.Params(":arch")
	rdx.Incr("downloads:" + repo)
	rdx.Incr("downloads:" + repo + ":" + osarch)
	ctx.JSON(200, "update success")
	/*
		realURL := goutils.StrFormat("http://{domain}/gorelease/{name}/{branch}/{osarch}/{name}",
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
	*/
}
