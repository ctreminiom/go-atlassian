# OpenTelemetry Update Guide for go-atlassian

## Summary

This guide documents the systematic update of all implementation files in `/jira/internal/` to implement proper OpenTelemetry error recording and span attributes.

## Pattern to Apply

### 1. Add imports (if not already present):
```go
"go.opentelemetry.io/otel/attribute"
"go.opentelemetry.io/otel/trace"
```

### 2. Update span creation to include span kind:
```go
ctx, span := tracer().Start(ctx, "spanName", spanWithKind(trace.SpanKindClient))
```

### 3. Add operation attributes after span creation:
```go
addAttributes(span,
    attribute.String("jira.issue.key", issueKeyOrID), // use appropriate attributes
    attribute.String("operation.name", "operation_name"),
)
```

### 4. Add error handling before return:
```go
result, response, err := someCall(...)
if err != nil {
    recordError(span, err)
    return nil, response, err // or appropriate return values
}

setOK(span)
return result, response, nil
```

## Attribute Guidelines

- For IDs: `jira.{resource}.id` (e.g., `jira.issue.id`, `jira.project.id`)
- For Keys: `jira.{resource}.key` (e.g., `jira.issue.key`, `jira.project.key`)
- For operations: `operation.name` with snake_case value
- For API version (in impl methods): `api.version`
- For arrays: use `attribute.StringSlice`
- For booleans: use `attribute.Bool`
- For integers: use `attribute.Int`

## Files Updated

### Completed:
1. ✅ announcement_banner_impl.go
2. ✅ application_role_impl.go
3. ✅ attachment_impl.go
4. ✅ audit_impl.go
5. ✅ issue_impl_adf.go (example provided)

### Remaining Files (58):
- authentication_impl.go (no tracer calls - skip)
- comment_impl.go
- comment_impl_adf.go
- comment_impl_rich_text.go
- dashboard_impl.go
- field_configuration_impl.go
- field_configuration_item_impl.go
- field_configuration_scheme.go
- field_context_impl.go (24 functions)
- field_context_option_impl.go
- field_impl.go
- field_trash_impl.go
- filter_impl.go (16 functions)
- filter_share_impl.go
- group_impl.go
- group_user_picker_impl.go
- issue_archive_impl.go
- issue_impl.go
- issue_impl_rich_text.go
- issue_property_impl.go
- jql_impl.go
- label_impl.go
- link_impl.go
- link_impl_adf.go
- link_impl_rich_text.go
- link_type_impl.go
- metadata_impl.go
- myself_impl.go
- notification_scheme_impl.go (16 functions)
- permission_impl.go
- permission_scheme_grant_impl.go
- permission_scheme_impl.go
- priority_impl.go
- project_category_impl.go
- project_component_impl.go
- project_feature_impl.go
- project_impl.go (20 functions)
- project_permission_scheme_impl.go
- project_property_impl.go
- project_role_actor_impl.go
- project_role_impl.go
- project_type_impl.go
- project_validator_impl.go
- project_version_impl.go (16 functions)
- remote_link_impl.go
- resolution_impl.go
- screen_impl.go
- screen_scheme_impl.go
- screen_tab_field_impl.go
- screen_tab_impl.go
- search_impl.go
- search_impl_adf.go
- search_impl_rich_text.go
- server_impl.go
- task_impl.go
- team_impl.go
- type_impl.go
- type_scheme_impl.go (20 functions)
- type_screen_scheme_impl.go (22 functions)
- user_impl.go
- user_search_impl.go
- vote_impl.go
- watcher_impl.go
- workflow_impl.go (18 functions)
- workflow_scheme_impl.go (14 functions)
- workflow_scheme_issue_type_impl.go
- workflow_status_impl.go (14 functions)
- worklog_impl_adf.go
- worklog_impl_rich_text.go

## Example Transformation

### Before:
```go
func (i *IssueFieldContextService) Gets(ctx context.Context, fieldID string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {
    ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Gets")
    defer span.End()

    return i.internalClient.Gets(ctx, fieldID, options, startAt, maxResults)
}
```

### After:
```go
func (i *IssueFieldContextService) Gets(ctx context.Context, fieldID string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {
    ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Gets", spanWithKind(trace.SpanKindClient))
    defer span.End()

    addAttributes(span,
        attribute.String("operation.name", "get_field_contexts"),
        attribute.String("jira.field.id", fieldID),
        attribute.Int("jira.pagination.start_at", startAt),
        attribute.Int("jira.pagination.max_results", maxResults),
    )

    result, response, err := i.internalClient.Gets(ctx, fieldID, options, startAt, maxResults)
    if err != nil {
        recordError(span, err)
        return nil, response, err
    }

    setOK(span)
    return result, response, nil
}
```

## Notes

- Some files (like authentication_impl.go) don't have tracer calls and can be skipped
- Files with many functions (field_context_impl.go, type_screen_scheme_impl.go, etc.) require careful attention
- Always maintain consistent attribute naming conventions
- Ensure proper error handling and span status setting