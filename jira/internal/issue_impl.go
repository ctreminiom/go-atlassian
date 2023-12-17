package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"net/http"
	"net/url"
)

type IssueServices struct {
	Attachment      *IssueAttachmentService
	CommentRT       *CommentRichTextService
	CommentADF      *CommentADFService
	Field           *IssueFieldService
	Label           *LabelService
	LinkRT          *LinkRichTextService
	LinkADF         *LinkADFService
	Metadata        *MetadataService
	Priority        *PriorityService
	Resolution      *ResolutionService
	SearchRT        *SearchRichTextService
	SearchADF       *SearchADFService
	Type            *TypeService
	Vote            *VoteService
	Watcher         *WatcherService
	WorklogAdf      *WorklogADFService
	WorklogRichText *WorklogRichTextService
	Property        *IssuePropertyService
}

func NewIssueService(client service.Connector, version string, services *IssueServices) (*IssueRichTextService, *IssueADFService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
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

func deleteIssue(ctx context.Context, client service.Connector, version, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("deleteSubtasks", fmt.Sprintf("%v", deleteSubTasks))

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?%v", version, issueKeyOrId, params.Encode())

	request, err := client.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return client.Call(request, nil)
}

func assignIssue(ctx context.Context, client service.Connector, version, issueKeyOrID, accountID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if accountID == "" {
		return nil, model.ErrNoAccountIDError
	}

	endpoint := fmt.Sprintf("/rest/api/%v/issue/%v/assignee", version, issueKeyOrID)

	request, err := client.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"accountId": accountID})
	if err != nil {
		return nil, err
	}

	return client.Call(request, nil)
}

func sendNotification(ctx context.Context, client service.Connector, version, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (
	*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/notify", version, issueKeyOrId)

	request, err := client.NewRequest(ctx, http.MethodPost, endpoint, "", options)
	if err != nil {
		return nil, err
	}

	return client.Call(request, nil)
}

func getTransitions(ctx context.Context, client service.Connector, version, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/transitions", version, issueKeyOrId)

	request, err := client.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	transitions := new(model.IssueTransitionsScheme)
	response, err := client.Call(request, transitions)
	if err != nil {
		return nil, response, err
	}

	return transitions, response, nil
}
