package gateway

import (
	"log"

	"github.com/micro/micro/v2/plugin"
	"github.com/opentracing/opentracing-go"
	"github.com/simontuo/merp-go/lib/tracer"
)

const name = "tracer"

func NewRegistry() plugin.Plugin {
	SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	return plugin.NewPlugin(
		plugin.WithName(name),
		plugin.WithHandler(TracerWrapper),
	)
}
