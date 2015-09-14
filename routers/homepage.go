package routers

import "github.com/Unknwon/macaron"

func Homepage(ctx *macaron.Context) {
	ctx.Data["Host"] = ctx.Req.Host
	ctx.HTML(200, "homepage")
}
