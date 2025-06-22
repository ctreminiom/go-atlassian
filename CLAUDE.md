# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

go-atlassian is a Go client library for Atlassian Cloud services (Jira, Confluence, Admin, Assets/Insight, Bitbucket). It follows interface-driven design with comprehensive testing and error handling.

## Essential Commands

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -race -coverprofile=coverage.out -covermode=atomic ./...

# Run tests for specific package
go test -v ./jira/v3/...

# Run a specific test
go test -v -run TestIssueService_Get ./jira/v3/...

# Lint the code (requires golangci-lint)
golangci-lint run --timeout=10m

# Generate mocks (if needed)
go generate ./...
```

## Architecture

### Package Structure
- `/jira/v2`, `/jira/v3` - Jira Software Cloud API (different versions)
- `/jira/agile` - Jira Agile features
- `/jira/sm` - Jira Service Management
- `/confluence/v2` - Confluence Cloud API
- `/admin` - Atlassian Admin APIs
- `/assets` - Jira Assets (Insight) APIs
- `/bitbucket` - Bitbucket Cloud APIs
- `/service` - Interface definitions for all services
- `/pkg/infra/models` - Shared models and error definitions

### Key Design Patterns

1. **Client-Service Architecture**: Each product has a Client struct containing service instances:
   ```go
   client := &jira.Client{
       HTTP: httpClient,
       Site: siteURL,
       Auth: auth,
   }
   ```

2. **Interface-Based Services**: All functionality exposed through interfaces defined in `/service`:
   - Interface definition: `/service/{product}/{feature}_connector.go`
   - Implementation: `/{product}/internal/{feature}_impl.go`
   - Mock: `/service/mocks/{product}_{feature}.go`

3. **Error Handling**: Use predefined errors from `/pkg/infra/models/errors.go`:
   ```go
   return nil, nil, models.ErrNoProjectIDOrKeyError
   ```

4. **HTTP Request Pattern**: Standard pattern for all API calls:
   ```go
   endpoint := fmt.Sprintf("rest/api/3/issue/%s", issueKey)
   request, err := i.client.NewRequest(ctx, http.MethodGet, endpoint, nil)
   // ... error handling ...
   response, err := i.client.Call(request, &issue)
   ```

## Testing Guidelines

- Uses `testify` for assertions and mocking
- Each implementation file must have corresponding `_test.go`
- Mock services are available in `/service/mocks`
- Tests should cover both success and error cases
- Use table-driven tests where appropriate

## Development Workflow

1. When implementing new features:
   - Define interface in `/service/{product}/`
   - Implement in `/{product}/internal/`
   - Add comprehensive tests
   - Update models in `/pkg/infra/models` if needed

2. When fixing bugs:
   - Add failing test first
   - Fix the issue
   - Ensure all tests pass
   - Run linter before committing

3. Error wrapping convention:
   ```go
   return nil, nil, fmt.Errorf("error getting issue %s: %w", issueKey, err)
   ```

## Important Notes

- Always use the predefined errors from `models` package
- Follow the existing service interface patterns
- Maintain backward compatibility when updating APIs
- Use context.Context for all API methods
- Return (*Response, error) for consistency
- Mock interfaces are generated - don't edit manually