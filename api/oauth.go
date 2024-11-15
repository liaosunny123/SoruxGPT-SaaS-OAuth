package api

import (
	"auditlimit/config"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func OAuth(r *ghttp.Request) {
	// 下面开始判断用户状态
	ctx := r.GetCtx()
	token := r.GetForm("usertoken").String()

	resp, err := g.Client().SetHeaderMap(g.MapStrStr{
		"apiauth":      config.AdminApiKey,
		"Content-Type": "application/json",
	}).Post(ctx, config.ShareUrl+"/adminapi/chatgpt/user/page", g.Map{
		"keyWord": token,
		"page":    1,
		"size":    20,
	})
	if err != nil {
		g.Log().Error(ctx, "GetJson", err)
		r.Response.Status = 500
		r.Response.WriteJson(g.Map{
			"detail": g.Map{
				"message": "Internal Server Error.",
			},
		})
		return
	}
	respJson := gjson.New(resp.ReadAllString())
	if len(respJson.Get("data.list").Array()) == 0 {
		r.Response.Status = 200
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "无效的激活码",
		})
		return
	}

	r.Response.Status = 200
	r.Response.WriteJson(g.Map{
		"code": 1,
		"msg":  "登入成功",
	})
	return
}
