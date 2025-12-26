package main

import (
	"context"
	i18n_demo "ordersvc/internal/controller/i18n-demo"
	"ordersvc/internal/controller/order"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"

	"base_util/consts"
	"base_util/limits"
	"base_util/mw"
	"base_util/trace"

	_ "ordersvc/internal/locales"
)

func main() {
	// 初始化jaeger tracer
	_, closer, err := trace.NewJaegerTracer("order-svc", "localhost:6831")
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	// 初始化限流器
	limits.InitLimiter()

	cmd.Run(gctx.GetInitCtx())
}

var cmd = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "start http server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		//s.Use(mw.MiddlewareRateLimiter)
		s.Use(mw.MiddlewareLimiter)
		s.Use(mw.MiddlewareHandlerResponse)
		s.Use(mw.InjectDefaultHeaderMW)
		s.Group("/v1", func(group *ghttp.RouterGroup) {
			group.Middleware(middlewareCORS)
			group.POST("/order", order.CreateOrder)
			group.POST("/i18n/:lang", i18n_demo.GetI18n)
		})
		s.BindHandler("/health", func(r *ghttp.Request) {
			r.Response.Write("ok")
		})
		s.SetAddr(":8080")
		//s.SetOpenApiPath(cfg.Server.OpenapiPath)
		//s.SetSwaggerPath(cfg.Server.SwaggerPath)
		s.SetSwaggerUITemplate(consts.SwaggerUITemplate)
		s.SetClientMaxBodySize(10 * 1024 * 1024)
		s.SetReadTimeout(time.Hour * 3)
		s.SetWriteTimeout(time.Hour * 3)
		s.Run()
		return nil
	},
}

func middlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
