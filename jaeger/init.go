package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"gitlab.gf.com.cn/hk-common/go-boot/lib/atomic"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
	"io"
)

var (
	enabled atomic.AtomicBool
	closer  io.Closer
)

type JaegerConfig struct {
	Endpoint string `json:"endpoint"`
}

func (j *JaegerConfig) Enable() bool {
	return enabled.True()
}

func (j *JaegerConfig) Start(serverKey string) (io.Closer, error) {
	var tracer opentracing.Tracer
	var err error
	if j.Endpoint != "" {
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
				CollectorEndpoint: j.Endpoint,
			},
		}).NewTracer()
		if err != nil {
			sentry.LogAndSentry(err)
			return nil, err
		}
		// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
		if tracer != nil {
			opentracing.SetGlobalTracer(tracer)
		}
	}
	enabled.Set(true)
	return closer, nil
}

func (j *JaegerConfig) Close() error {
	if closer != nil {
		return closer.Close()
	}
	return nil
}

func InitJaeger(serverKey, jaegerAddressCollectorEndpoint string) (io.Closer, error) {
	// 根据配置初始化Tracer 返回Closer
	var tracer opentracing.Tracer
	var closer io.Closer
	var err error
	if jaegerAddressCollectorEndpoint != "" {
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
			return closer, err
		}

		// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
		if tracer != nil {
			opentracing.SetGlobalTracer(tracer)
		}
	}
	return closer, nil
}
