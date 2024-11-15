package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	PORT        = 8080
	ShareUrl    = ""
	AdminApiKey = ""
)

func init() {
	ctx := gctx.GetInitCtx()
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)

	shareUrl := g.Cfg().MustGetWithEnv(ctx, "SHARE_URL").String()
	if shareUrl != "" {
		ShareUrl = shareUrl
	}
	g.Log().Info(ctx, "SHARE_URL:", ShareUrl)

	adminApiKey := g.Cfg().MustGetWithEnv(ctx, "ADMIN_API_KEY").String()
	if adminApiKey != "" {
		AdminApiKey = adminApiKey
	}
	g.Log().Info(ctx, "ADMIN_API_KEY:", AdminApiKey)
}
