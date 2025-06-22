// Package testing provides test utilities for the go-atlassian library.
// This file initializes OpenTelemetry with a custom noop tracer that
// doesn't modify contexts, preventing test failures due to context mismatches.
package testing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/attribute"
)

func init() {
	// Set a custom tracer provider that returns tracers which don't modify contexts
	otel.SetTracerProvider(&contextPreservingTracerProvider{})
}

// contextPreservingTracerProvider is a custom TracerProvider that creates tracers
// which don't modify contexts at all.
type contextPreservingTracerProvider struct{
	trace.TracerProvider
}

func (n *contextPreservingTracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return &contextPreservingTracer{}
}

// contextPreservingTracer is a custom Tracer that returns the same context without modification
type contextPreservingTracer struct{
	trace.Tracer
}

func (n *contextPreservingTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	// Return the original context unmodified and a noop span
	return ctx, &noopSpan{}
}

// noopSpan implements the Span interface with no-op methods
type noopSpan struct{
	trace.Span
}

func (n *noopSpan) End(options ...trace.SpanEndOption)                       {}
func (n *noopSpan) AddEvent(name string, options ...trace.EventOption)       {}
func (n *noopSpan) IsRecording() bool                                        { return false }
func (n *noopSpan) RecordError(err error, options ...trace.EventOption)      {}
func (n *noopSpan) SpanContext() trace.SpanContext                           { return trace.SpanContext{} }
func (n *noopSpan) SetStatus(code codes.Code, description string)            {}
func (n *noopSpan) SetName(name string)                                      {}
func (n *noopSpan) SetAttributes(kv ...attribute.KeyValue)                   {}
func (n *noopSpan) TracerProvider() trace.TracerProvider                     { return &contextPreservingTracerProvider{} }