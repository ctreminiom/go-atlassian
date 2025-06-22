package v2

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/ctreminiom/go-atlassian/v2/confluence/v2"

func tracer(opts ...trace.TracerOption) trace.Tracer {
	return otel.Tracer(tracerName, opts...)
}