package main

import (
	"auditlimit/api"
	"auditlimit/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetPort(config.PORT)
	s.BindHandler("/", Index)
	s.BindHandler("/api/oauth", api.OAuth)
	s.Run()
}

func Index(r *ghttp.Request) {
	r.Response.Write("Hello SoruxGPT SaaS, this is the oauth target for SoruxGPT SaaS.")
}

func init() {

}
