package internal

import (
	"context"
	"errors"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"github.com/imdario/mergo"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type IssueRichTextService struct {
	internalClient jira.IssueRichTextConnector
	Attachment     jira.Attachment
	Comment        jira.RichTextComment
	Field          jira.Field
}

func (i IssueRichTextService) Delete(ctx context.Context, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, issueKeyOrId, deleteSubTasks)
}

func (i IssueRichTextService) Assign(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {
	return i.internalClient.Assign(ctx, issueKeyOrId, accountId)
}

func (i IssueRichTextService) Notify(ctx context.Context, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Notify(ctx, issueKeyOrId, options)
}

func (i IssueRichTextService) Transitions(ctx context.Context, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	return i.internalClient.Transitions(ctx, issueKeyOrId)
}

func (i IssueRichTextService) Create(ctx context.Context, payload *model.IssueSchemeV2, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, payload, customFields)
}

func (i IssueRichTextService) Creates(ctx context.Context, payload []*model.IssueBulkSchemeV2) (*model.IssueBulkResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Creates(ctx, payload)
}

func (i IssueRichTextService) Get(ctx context.Context, issueKeyOrId string, fields, expand []string) (*model.IssueSchemeV2, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, issueKeyOrId, fields, expand)
}

func (i IssueRichTextService) Update(ctx context.Context, issueKeyOrId string, notify bool, payload *model.IssueSchemeV2, customFields *model.CustomFields, operations *model.UpdateOperations) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, issueKeyOrId, notify, payload, customFields, operations)
}

func (i IssueRichTextService) Move(ctx context.Context, issueKeyOrId, transitionId string, options *model.IssueMoveOptionsV2) (*model.ResponseScheme, error) {
	return i.internalClient.Move(ctx, issueKeyOrId, transitionId, options)
}

type internalRichTextServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalRichTextServiceImpl) Delete(ctx context.Context, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	return deleteIssue(ctx, i.c, i.version, issueKeyOrId, deleteSubTasks)
}

func (i *internalRichTextServiceImpl) Assign(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {
	return assignIssue(ctx, i.c, i.version, issueKeyOrId, accountId)
}

func (i *internalRichTextServiceImpl) Notify(ctx context.Context, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error) {
	return sendNotification(ctx, i.c, i.version, issueKeyOrId, options)
}

func (i *internalRichTextServiceImpl) Transitions(ctx context.Context, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	return getTransitions(ctx, i.c, i.version, issueKeyOrId)
}

func (i *internalRichTextServiceImpl) Create(ctx context.Context, payload *model.IssueSchemeV2, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error) {

	var reader io.Reader
	var err error

	if customFields != nil {

		payloadUpdated, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, nil, err
		}

		reader, err = i.c.TransformStructToReader(payloadUpdated)
		if err != nil {
			return nil, nil, err
		}

	} else {

		reader, err = i.c.TransformStructToReader(payload)
		if err != nil {
			return nil, nil, err
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	issue := new(model.IssueResponseScheme)
	response, err := i.c.Call(request, issue)
	if err != nil {
		return nil, response, err
	}

	return issue, response, nil
}

func (i *internalRichTextServiceImpl) Creates(ctx context.Context, payload []*model.IssueBulkSchemeV2) (*model.IssueBulkResponseScheme, *model.ResponseScheme, error) {

	if len(payload) == 0 {
		return nil, nil, errors.New("error, please provide a valid []*IssueBulkScheme slice of pointers")
		// TODO: The errors when the bulk creates does not contains values needs to be parsed and moved to the model package
	}

	var issuePayloads []map[string]interface{}
	for _, newIssue := range payload {

		if newIssue.Payload == nil {
			continue
		}

		issuePayload, err := newIssue.Payload.MergeCustomFields(newIssue.CustomFields)
		if err != nil {
			return nil, nil, err
		}

		issuePayloads = append(issuePayloads, issuePayload)
	}

	var bulkPayload = map[string]interface{}{}
	bulkPayload["issueUpdates"] = issuePayloads

	reader, err := i.c.TransformStructToReader(&bulkPayload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/bulk", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueBulkResponseScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

func (i *internalRichTextServiceImpl) Get(ctx context.Context, issueKeyOrId string, fields, expand []string) (*model.IssueSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if len(fields) != 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v", i.version, issueKeyOrId))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	issue := new(model.IssueSchemeV2)
	response, err := i.c.Call(request, issue)
	if err != nil {
		return nil, response, err
	}

	return issue, response, nil
}

func (i *internalRichTextServiceImpl) Update(ctx context.Context, issueKeyOrId string, notify bool, payload *model.IssueSchemeV2, customFields *model.CustomFields, operations *model.UpdateOperations) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("notifyUsers", fmt.Sprintf("%v", notify))
	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?%v", i.version, issueKeyOrId, params.Encode())

	var reader io.Reader
	var err error

	// Executed when customfields and operations are not provided
	if customFields == nil && operations == nil {

		reader, err = i.c.TransformStructToReader(payload)
		if err != nil {
			return nil, err
		}
	}

	// Executed when customfields and operation are provided
	if customFields != nil && operations != nil {

		payloadUpdated, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		payloadWithOperations, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		if err := mergo.Map(&payloadUpdated, &payloadWithOperations, mergo.WithOverride); err != nil {
			return nil, err
		}

		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	// Executed when customfields are provided but not the operations
	if customFields != nil && operations == nil {

		payloadUpdated, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	// Executed when operations are provided but not the customfields
	if customFields == nil && operations != nil {

		payloadUpdated, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRichTextServiceImpl) Move(ctx context.Context, issueKeyOrId, transitionId string, options *model.IssueMoveOptionsV2) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if transitionId == "" {
		return nil, model.ErrNoTransitionIDError
	}

	payloadUpdated := make(map[string]interface{})
	payloadUpdated["transition"] = map[string]interface{}{"id": transitionId}

	var reader io.Reader
	var err error

	if options != nil && options.Fields != nil {

		// Executed when customfields and operation are provided
		if options.CustomFields != nil && options.Operations != nil {

			payloadWithCustomFields, err := options.Fields.MergeCustomFields(options.CustomFields)
			if err != nil {
				return nil, err
			}

			payloadWithOperations, err := options.Fields.MergeOperations(options.Operations)
			if err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithCustomFields, &payloadWithOperations, mergo.WithOverride); err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithCustomFields, &payloadUpdated, mergo.WithOverride); err != nil {
				return nil, err
			}

			reader, err = i.c.TransformStructToReader(&payloadWithCustomFields)
			if err != nil {
				return nil, err
			}
		}

		// Executed when customfields are provided but not the operations
		if options.CustomFields != nil && options.Operations == nil {

			payloadWithCustomFields, err := options.Fields.MergeCustomFields(options.CustomFields)
			if err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithCustomFields, &payloadUpdated, mergo.WithOverride); err != nil {
				return nil, err
			}

			reader, err = i.c.TransformStructToReader(&payloadWithCustomFields)
			if err != nil {
				return nil, err
			}
		}

		// Executed when operations are provided but not the customfields
		if options.CustomFields == nil && options.Operations != nil {

			payloadWithOperations, err := options.Fields.MergeOperations(options.Operations)
			if err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithOperations, &payloadUpdated, mergo.WithOverride); err != nil {
				return nil, err
			}

			reader, err = i.c.TransformStructToReader(&payloadWithOperations)
			if err != nil {
				return nil, err
			}
		}
	} else {
		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/transitions", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
