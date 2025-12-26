package i18n_demo

import (
	"base_util/common"
	"ordersvc/internal/locales"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func GetI18n(r *ghttp.Request) {
	lang := r.Get("lang").String()
	g.Log().Infof(r.GetCtx(), "lang: %s", lang)

	common.Success(r, map[string]string{
		"lang": lang,
		"msg":  locales.GetLocaleMsg(r.GetCtx(), lang, locales.OrderServerError),
	})
}
