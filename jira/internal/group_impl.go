package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
)

func NewGroupService(client service.Client, version string) (*GroupService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &GroupService{
		internalClient: &internalGroupServiceImpl{c: client, version: version},
	}, nil
}

type GroupService struct {
	internalClient jira.GroupConnector
}

func (g *GroupService) Delete(ctx context.Context, groupName string) (*model.ResponseScheme, error) {
	return g.internalClient.Delete(ctx, groupName)
}

func (g *GroupService) Bulk(ctx context.Context, options *model.GroupBulkOptionsScheme, startAt, maxResults int) (*model.BulkGroupScheme, *model.ResponseScheme, error) {
	return g.internalClient.Bulk(ctx, options, startAt, maxResults)
}

func (g *GroupService) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (*model.GroupMemberPageScheme, *model.ResponseScheme, error) {
	return g.internalClient.Members(ctx, groupName, inactive, startAt, maxResults)
}

func (g *GroupService) Add(ctx context.Context, groupName, accountId string) (*model.GroupScheme, *model.ResponseScheme, error) {
	return g.internalClient.Add(ctx, groupName, accountId)
}

func (g *GroupService) Remove(ctx context.Context, groupName, accountId string) (*model.ResponseScheme, error) {
	return g.internalClient.Remove(ctx, groupName, accountId)
}

func (g *GroupService) Create(ctx context.Context, groupName string) (*model.GroupScheme, *model.ResponseScheme, error) {
	return g.internalClient.Create(ctx, groupName)
}

type internalGroupServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalGroupServiceImpl) Create(ctx context.Context, groupName string) (*model.GroupScheme, *model.ResponseScheme, error) {

	if groupName == "" {
		return nil, nil, model.ErrNoGroupNameError
	}

	payload := struct {
		Name string `json:"name"`
	}{
		Name: groupName,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/group", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.GroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalGroupServiceImpl) Delete(ctx context.Context, groupName string) (*model.ResponseScheme, error) {

	if groupName == "" {
		return nil, model.ErrNoGroupNameError
	}

	params := url.Values{}
	params.Add("groupname", groupName)

	endpoint := fmt.Sprintf("rest/api/%v/group?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i internalGroupServiceImpl) Bulk(ctx context.Context, options *model.GroupBulkOptionsScheme, startAt, maxResults int) (*model.BulkGroupScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, id := range options.GroupIDs {
			params.Add("groupId", id)
		}

		for _, name := range options.GroupNames {
			params.Add("groupName", name)
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/group/bulk?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BulkGroupScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i internalGroupServiceImpl) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (*model.GroupMemberPageScheme, *model.ResponseScheme, error) {

	if groupName == "" {
		return nil, nil, model.ErrNoGroupNameError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("groupname", groupName)
	params.Add("includeInactiveUsers", fmt.Sprintf("%v", inactive))

	endpoint := fmt.Sprintf("rest/api/%v/group/member?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.GroupMemberPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i internalGroupServiceImpl) Add(ctx context.Context, groupName, accountId string) (*model.GroupScheme, *model.ResponseScheme, error) {

	if groupName == "" {
		return nil, nil, model.ErrNoGroupNameError
	}

	if accountId == "" {
		return nil, nil, model.ErrNoAccountIDError
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{
		AccountID: accountId,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("groupname", groupName)
	endpoint := fmt.Sprintf("rest/api/%v/group/user?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.GroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i internalGroupServiceImpl) Remove(ctx context.Context, groupName, accountId string) (*model.ResponseScheme, error) {

	if groupName == "" {
		return nil, model.ErrNoGroupNameError
	}

	if accountId == "" {
		return nil, model.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("groupname", groupName)
	params.Add("accountId", accountId)
	endpoint := fmt.Sprintf("rest/api/%v/group/user?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
