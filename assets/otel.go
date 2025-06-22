package assets

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/ctreminiom/go-atlassian/v2/assets"

func tracer(opts ...trace.TracerOption) trace.Tracer {
	return otel.Tracer(tracerName, opts...)
}