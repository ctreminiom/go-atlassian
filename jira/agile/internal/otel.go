package internal

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/ctreminiom/go-atlassian/v2/jira/agile/internal"

func tracer(opts ...trace.TracerOption) trace.Tracer {
	return otel.Tracer(tracerName, opts...)
}

// recordError records an error on the span and sets the span status to error
func recordError(span trace.Span, err error) {
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

// setOK sets the span status to OK
func setOK(span trace.Span) {
	span.SetStatus(codes.Ok, "")
}

// spanWithKind creates a span with the specified kind
func spanWithKind(kind trace.SpanKind) trace.SpanStartOption {
	return trace.WithSpanKind(kind)
}

// addAttributes adds multiple attributes to a span
func addAttributes(span trace.Span, attrs ...attribute.KeyValue) {
	span.SetAttributes(attrs...)
}
