package jaeger

import (
	"fmt"
	"io"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	conf "gpi/libriries/config"
)

func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	jaegerConf := conf.Conf.GetStringMapString("jaeger")
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type: jaegerConf["type"],
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: jaegerConf["host"]+":"+jaegerConf["port"],
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}