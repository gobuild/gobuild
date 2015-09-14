package routers

import (
	"net/http"

	"github.com/Unknwon/macaron"
	"github.com/gorelease/gorelease/models/goutils"
)

func DownloadRedirect(ctx *macaron.Context, r *http.Request) {
	owner := ctx.Params(":owner")
	name := ctx.Params(":name")
	branch := r.FormValue("branch")
	if branch == "" {
		branch = "master"
	}
	repo := owner + "/" + name
	domain := rdx.Get("domain:" + repo).Val()
	if domain == "" {
		ctx.Error(405, "repo not registed in gorelease, not open register for now")
		return
	}
	osarch := ctx.Params(":os") + "-" + ctx.Params(":arch")
	rdx.Incr("downloads:" + repo)
	rdx.Incr("downloads:" + repo + ":" + osarch)
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
}
