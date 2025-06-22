package internal

// Import the testing package to initialize OpenTelemetry with a noop tracer
// that doesn't modify contexts, preventing test failures.
import _ "github.com/ctreminiom/go-atlassian/v2/pkg/infra/testing"