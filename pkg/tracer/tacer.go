package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {

	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //固定采样，对所有数据都进行采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  agentHostPort,
		},
	}

	tracer, closer, err := cfg.NewTracer() //初始化tracer 对象
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
