package mw

import (
	"base_util/common"
	"base_util/uerrors"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"golang.org/x/time/rate"
)

const (
	ReqId         = "x-request-id"
	TrafficMark   = "x-qw-traffic-mark"
	BusinessType  = "x-voice-business-type"
	ReqToken      = "x-token"
	ReqVerifyTime = "x-verify-time"
	TenantId      = "tenant-id"
)

var (
	rateLimiter = rate.NewLimiter(20, 50) // 每秒20个请求，最大突发50个请求
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

func InjectDefaultHeaderMW(r *ghttp.Request) {
	ctx := r.Context()
	//ctx = context.WithValue(ctx, ReqId, r.Header.Get(ReqId))
	newctx, err := gtrace.WithUUID(ctx, r.Header.Get(ReqId))
	if err != nil {
		//handleError(r, err)
		//return
		traceId := gctx.CtxId(ctx)
		newctx, _ = gtrace.WithUUID(ctx, traceId)
	}
	r.SetCtx(newctx)
	r.Middleware.Next()
}

func MiddlewareLimiter(r *ghttp.Request) {
	// 获取接口路径
	path := r.URL.Path
	g.Log().Infof(r.Context(), "MiddlewareLimiter path: %s", path)
	hystrix.Do(path, func() error {
		r.Middleware.Next()
		return nil
	}, func(err error) error {
		g.Log().Errorf(r.Context(), "请求过于频繁，请稍后重试: %v", err)
		handleError(r, uerrors.NewVideoTranslateErr("请求过于频繁，请稍后重试", 429))
		return nil
	})
	// TODO 限流中间件
	//r.Middleware.Next()
}

func MiddlewareRateLimiter(r *ghttp.Request) {
	if !rateLimiter.Allow() {
		handleError(r, uerrors.NewVideoTranslateErr("请求过于频繁，请稍后重试", 429))
		return
	}
	r.Middleware.Next()
}

func handleError(r *ghttp.Request, err error) {
	g.Log().Errorf(r.Context(), "handleError: %v", err)
	common.Fail2(r, uerrors.NewVideoTranslateErr(err.Error(), 1))
}
