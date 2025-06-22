# Testing Infrastructure

This package provides testing utilities for the go-atlassian library.

## OpenTelemetry Test Configuration

The `otel_test_init.go` file initializes OpenTelemetry with a custom context-preserving tracer for tests. This solves a common issue where:

1. Production code uses OpenTelemetry tracing that wraps contexts with trace information
2. Test mocks expect exactly `context.Background()` 
3. The wrapped contexts don't match mock expectations, causing test failures

## Solution

The custom tracer provider ensures that:
- `tracer().Start(ctx, "operation")` returns the original context unmodified
- All tracing operations are no-ops during tests
- Production tracing behavior is preserved (no changes to production code)
- No changes needed to existing test files

## Usage

Each `internal` package that uses tracing should have an `otel_test_init.go` file that imports this package:

```go
package internal

// Import the testing package to initialize OpenTelemetry with a noop tracer
// that doesn't modify contexts, preventing test failures.
import _ "github.com/ctreminiom/go-atlassian/v2/pkg/infra/testing"
```

This automatically sets up the context-preserving tracer when tests run.

## Benefits

- ✅ Zero changes to production code
- ✅ Zero changes to existing test files  
- ✅ Tracing remains enabled in production
- ✅ Tests pass without context mismatches
- ✅ Clean and maintainable solution