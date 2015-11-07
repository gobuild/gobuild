package routers

import (
	"strings"

	"github.com/Unknwon/macaron"
)

func ApiApplications(ctx *macaron.Context) {
	res, err := rdx.Keys("downloads:*/*").Result()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	var jsondata = []interface{}{}
	for _, key := range res {
		if strings.Count(key, ":") != 1 {
			continue
		}
		repoName := strings.Split(key, ":")[1]
		dcnt, _ := rdx.Get(key).Int64()
		jsondata = append(jsondata, map[string]interface{}{
			"name":           repoName,
			"download_count": dcnt,
		})
	}
	ctx.JSON(200, jsondata)
}
