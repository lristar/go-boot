package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
	"io"
)

func InitJaeger(serverKey, jaegerAddressCollectorEndpoint string) io.Closer {
	// 根据配置初始化Tracer 返回Closer
	var tracer opentracing.Tracer
	var closer io.Closer
	var err error
	tracer, closer, err = (&config.Configuration{
		ServiceName: serverKey,
		Disabled:    false,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: jaegerAddressCollectorEndpoint,
		},
	}).NewTracer()
	if err != nil {
		sentry.LogAndSentry(err)
		return closer
	}

	// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
	if tracer != nil {
		opentracing.SetGlobalTracer(tracer)
	}
	return closer
}
