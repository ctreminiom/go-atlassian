package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
)

// IssueServices groups various services related to issue management in Jira Service Management.
type IssueServices struct {
	// Attachment is the service for managing issue attachments.
	Attachment *IssueAttachmentService
	// CommentRT is the service for managing rich text comments.
	CommentRT *CommentRichTextService
	// CommentADF is the service for managing ADF comments.
	CommentADF *CommentADFService
	// Field is the service for managing issue fields.
	Field *IssueFieldService
	// Label is the service for managing issue labels.
	Label *LabelService
	// LinkRT is the service for managing rich text issue links.
	LinkRT *LinkRichTextService
	// LinkADF is the service for managing ADF issue links.
	LinkADF *LinkADFService
	// Metadata is the service for managing issue metadata.
	Metadata *MetadataService
	// Priority is the service for managing issue priorities.
	Priority *PriorityService
	// Resolution is the service for managing issue resolutions.
	Resolution *ResolutionService
	// SearchRT is the service for managing rich text issue searches.
	SearchRT *SearchRichTextService
	// SearchADF is the service for managing ADF issue searches.
	SearchADF *SearchADFService
	// Type is the service for managing issue types.
	Type *TypeService
	// Vote is the service for managing issue votes.
	Vote *VoteService
	// Watcher is the service for managing issue watchers.
	Watcher *WatcherService
	// WorklogAdf is the service for managing ADF worklogs.
	WorklogAdf *WorklogADFService
	// WorklogRichText is the service for managing rich text worklogs.
	WorklogRichText *WorklogRichTextService
	// Property is the service for managing issue properties.
	Property *IssuePropertyService
}

// NewIssueService creates new instances of IssueRichTextService and IssueADFService.
// It takes a service.Connector, a version string, and an optional IssueServices struct as input.
// Returns pointers to IssueRichTextService and IssueADFService, and an error if the version is not provided.
func NewIssueService(client service.Connector, version string, services *IssueServices) (*IssueRichTextService, *IssueADFService, error) {

	if version == "" {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	richTextService := &IssueRichTextService{
		internalClient: &internalRichTextServiceImpl{
			c:       client,
			version: version,
		},
	}

	adfService := &IssueADFService{
		internalClient: &internalIssueADFServiceImpl{
			c:       client,
			version: version,
		},
	}

	if services != nil {

		adfService.Attachment = services.Attachment
		adfService.Comment = services.CommentADF
		adfService.Field = services.Field
		adfService.Label = services.Label
		adfService.Link = services.LinkADF
		adfService.Metadata = services.Metadata
		adfService.Priority = services.Priority
		adfService.Resolution = services.Resolution
		adfService.Search = services.SearchADF
		adfService.Type = services.Type
		adfService.Vote = services.Vote
		adfService.Watcher = services.Watcher
		adfService.Worklog = services.WorklogAdf
		adfService.Property = services.Property

		richTextService.Comment = services.CommentRT
		richTextService.Attachment = services.Attachment
		richTextService.Field = services.Field
		richTextService.Label = services.Label
		richTextService.Link = services.LinkRT
		richTextService.Metadata = services.Metadata
		richTextService.Priority = services.Priority
		richTextService.Resolution = services.Resolution
		richTextService.Search = services.SearchRT
		richTextService.Type = services.Type
		richTextService.Vote = services.Vote
		richTextService.Watcher = services.Watcher
		richTextService.Worklog = services.WorklogRichText
		richTextService.Property = services.Property

	}

	return richTextService, adfService, nil
}

// -------------------------------------------
// These private functions are used on the Issue Services implementation, as that services is segmented in the ADF and Rich Text
// format, in order to avoid duplication, those function are injected on the ADF/Rich Text implementations.
// -------------------------------------------

func deleteIssue(ctx context.Context, client service.Connector, version, issueKeyOrID string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "deleteIssue", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_issue"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.Bool("jira.delete_subtasks", deleteSubTasks),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, err
	}

	params := url.Values{}
	params.Add("deleteSubtasks", fmt.Sprintf("%v", deleteSubTasks))

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?%v", version, issueKeyOrID, params.Encode())

	request, err := client.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := client.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func assignIssue(ctx context.Context, client service.Connector, version, issueKeyOrID, accountID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "assignIssue", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "assign_issue"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.account_id", accountID),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, err
	}

	if accountID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAccountID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("/rest/api/%v/issue/%v/assignee", version, issueKeyOrID)

	request, err := client.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"accountId": accountID})
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := client.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func sendNotification(ctx context.Context, client service.Connector, version, issueKeyOrID string, options *model.IssueNotifyOptionsScheme) (
	*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "sendNotification", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "send_issue_notification"),
		attribute.String("jira.issue.key", issueKeyOrID),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/notify", version, issueKeyOrID)

	request, err := client.NewRequest(ctx, http.MethodPost, endpoint, "", options)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := client.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func getTransitions(ctx context.Context, client service.Connector, version, issueKeyOrID string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "getTransitions", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_issue_transitions"),
		attribute.String("jira.issue.key", issueKeyOrID),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/transitions", version, issueKeyOrID)

	request, err := client.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	transitions := new(model.IssueTransitionsScheme)
	response, err := client.Call(request, transitions)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return transitions, response, nil
}
