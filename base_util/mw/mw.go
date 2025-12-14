package mw

import (
	"base_util/common"
	"base_util/uerrors"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		handleError(r, err)
	} else {
		if !(r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0) {
			res := r.GetHandlerResponse()
			r.Response.WriteJson(res)
		}
	}
}

func handleError(r *ghttp.Request, err error) {
	common.Fail2(r, uerrors.NewVideoTranslateErr(err.Error(), 1))
}
