package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"path"
)

func NewIssueArchivalService(client service.Connector, version string) *IssueArchivalService {
	return &IssueArchivalService{
		internalClient: &internalIssueArchivalImpl{c: client, version: version},
	}
}

type IssueArchivalService struct {
	internalClient jira.ArchiveService
}

func (i *IssueArchivalService) Preserve(ctx context.Context, issueIdsOrKeys []string) (*model.IssueArchivalSyncResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Preserve(ctx, issueIdsOrKeys)
}

func (i *IssueArchivalService) PreserveByJQL(ctx context.Context, jql string) (string, *model.ResponseScheme, error) {
	return i.internalClient.PreserveByJQL(ctx, jql)
}

func (i *IssueArchivalService) Restore(ctx context.Context, issueIdsOrKeys []string) (*model.IssueArchivalSyncResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Restore(ctx, issueIdsOrKeys)
}

func (i *IssueArchivalService) Export(ctx context.Context, payload *model.IssueArchivalExportPayloadScheme) (string, *model.ResponseScheme, error) {
	return i.internalClient.Export(ctx, payload)
}

type internalIssueArchivalImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueArchivalImpl) Preserve(ctx context.Context, issueIdsOrKeys []string) (result *model.IssueArchivalSyncResponseScheme, response *model.ResponseScheme, err error) {

	if len(issueIdsOrKeys) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	payload := make(map[string]interface{})
	payload["issueIdsOrKeys"] = issueIdsOrKeys

	endpoint := fmt.Sprintf("rest/api/%s/issue/archive", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	report := new(model.IssueArchivalSyncResponseScheme)
	response, err = i.c.Call(request, report)
	if err != nil {
		return nil, response, err
	}

	return report, response, nil
}

func (i *internalIssueArchivalImpl) PreserveByJQL(ctx context.Context, jql string) (taskID string, response *model.ResponseScheme, err error) {

	if jql == "" {
		return "", nil, model.ErrNoJQL
	}

	payload := make(map[string]interface{})
	payload["jql"] = jql

	endpoint := fmt.Sprintf("rest/api/%s/issue/archive", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return "", nil, err
	}

	response, err = i.c.Call(request, nil)
	if err != nil {
		return "", response, err
	}

	return path.Base(response.Bytes.String()), response, nil
}

func (i *internalIssueArchivalImpl) Restore(ctx context.Context, issueIdsOrKeys []string) (result *model.IssueArchivalSyncResponseScheme, response *model.ResponseScheme, err error) {

	if len(issueIdsOrKeys) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	payload := make(map[string]interface{})
	payload["issueIdsOrKeys"] = issueIdsOrKeys

	endpoint := fmt.Sprintf("rest/api/%s/issue/unarchive", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	report := new(model.IssueArchivalSyncResponseScheme)
	response, err = i.c.Call(request, report)
	if err != nil {
		return nil, response, err
	}

	return report, response, nil
}

func (i *internalIssueArchivalImpl) Export(ctx context.Context, payload *model.IssueArchivalExportPayloadScheme) (taskID string, response *model.ResponseScheme, err error) {

	endpoint := fmt.Sprintf("rest/api/%s/issues/archive/export", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return "", nil, err
	}

	response, err = i.c.Call(request, nil)
	if err != nil {
		return "", response, err
	}

	return "", response, nil
}
