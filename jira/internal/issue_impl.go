package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
)

type IssueServices struct {
	Attachment      jira.Attachment
	CommentRichText jira.RichTextComment
	CommentADF      jira.AdfComment
}

func NewIssueService(client service.Client, version string, services *IssueServices) (*IssueRichTextService, *IssueADFService, error) {

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

		if services.Attachment != nil {
			adfService.Attachment = services.Attachment
			richTextService.Attachment = services.Attachment
		}

		if services.CommentADF != nil {
			adfService.Comment = services.CommentADF
		}

		if services.CommentRichText != nil {
			richTextService.Comment = services.CommentRichText
		}
	}

	return richTextService, adfService, nil
}

func deleteIssue(ctx context.Context, client service.Client, version, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("deleteSubtasks", fmt.Sprintf("%v", deleteSubTasks))

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?%v", version, issueKeyOrId, params.Encode())

	request, err := client.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return client.Call(request, nil)
}

func assignIssue(ctx context.Context, client service.Client, version, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if accountId == "" {
		return nil, model.ErrNoAccountIDError
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{
		AccountID: accountId,
	}

	reader, err := client.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/rest/api/%v/issue/%v/assignee", version, issueKeyOrId)

	request, err := client.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return client.Call(request, nil)
}

func sendNotification(ctx context.Context, client service.Client, version, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (
	*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	reader, err := client.TransformStructToReader(options)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/notify", version, issueKeyOrId)

	request, err := client.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return client.Call(request, nil)
}

func getTransitions(ctx context.Context, client service.Client, version, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/transitions", version, issueKeyOrId)

	request, err := client.NewRequest(ctx, http.MethodGet, endpoint, nil)
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
