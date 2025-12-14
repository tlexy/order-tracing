package trace

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

// 新建一个jaeger的trace客户端
func NewJaegerTracer(serviceName, jaegerHost string) (opentracing.Tracer, io.Closer, error) {
	// 初始化jaeger配置
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            false,
			LocalAgentHostPort:  jaegerHost,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	// 初始化jaeger tracer
	jaegerTracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jaeger.NullLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(jaegerTracer)
	return jaegerTracer, closer, nil
}
