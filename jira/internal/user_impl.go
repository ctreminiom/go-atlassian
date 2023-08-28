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
	"strings"
)

func NewUserService(client service.Connector, version string, connector *UserSearchService) (*UserService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &UserService{
		internalClient: &internalUserImpl{c: client, version: version},
		Search:         connector,
	}, nil
}

type UserService struct {
	internalClient jira.UserConnector
	Search         *UserSearchService
}

// Get returns a user
//
// GET /rest/api/{2-3}/user
//
// https://docs.go-atlassian.io/jira-software-cloud/users#get-user
func (u *UserService) Get(ctx context.Context, accountId string, expand []string) (*model.UserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Get(ctx, accountId, expand)
}

// Create creates a user. This resource is retained for legacy compatibility.
//
// As soon as a more suitable alternative is available this resource will be deprecated.
//
// The option is provided to set or generate a password for the user.
//
// When using the option to generate a password, by omitting password from the request, include "notification": "true" to ensure the user is
//
// sent an email advising them that their account is created.
//
// This email includes a link for the user to set their password. If the notification isn't sent for a generated password,
//
// the user will need to be sent a reset password request from Jira.
//
// POST /rest/api/{2-3}user
//
// https://docs.go-atlassian.io/jira-software-cloud/users#create-user
func (u *UserService) Create(ctx context.Context, payload *model.UserPayloadScheme) (*model.UserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Create(ctx, payload)
}

// Delete deletes a user.
//
// DELETE /rest/api/{2-3}/user
//
// https://docs.go-atlassian.io/jira-software-cloud/users#delete-user
func (u *UserService) Delete(ctx context.Context, accountId string) (*model.ResponseScheme, error) {
	return u.internalClient.Delete(ctx, accountId)
}

// Find returns a paginated list of the users specified by one or more account IDs.
//
// GET /rest/api/{2-3}/user/bulk
//
// https://docs.go-atlassian.io/jira-software-cloud/users#bulk-get-users
func (u *UserService) Find(ctx context.Context, accountIds []string, startAt, maxResults int) (*model.UserSearchPageScheme, *model.ResponseScheme, error) {
	return u.internalClient.Find(ctx, accountIds, startAt, maxResults)
}

// Groups returns the groups to which a user belongs.
//
// GET /rest/api/{2-3}/user/groups
//
// https://docs.go-atlassian.io/jira-software-cloud/users#get-user-groups
func (u *UserService) Groups(ctx context.Context, accountIds string) ([]*model.UserGroupScheme, *model.ResponseScheme, error) {
	return u.internalClient.Groups(ctx, accountIds)
}

// Gets returns a list of all (active and inactive) users.
//
// GET /rest/api/{2-3}/users/search
//
// https://docs.go-atlassian.io/jira-software-cloud/users#get-all-users
func (u *UserService) Gets(ctx context.Context, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Gets(ctx, startAt, maxResults)
}

type internalUserImpl struct {
	c       service.Connector
	version string
}

func (i *internalUserImpl) Get(ctx context.Context, accountId string, expand []string) (*model.UserScheme, *model.ResponseScheme, error) {

	if accountId == "" {
		return nil, nil, model.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("accountId", accountId)

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("rest/api/%v/user?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.UserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalUserImpl) Create(ctx context.Context, payload *model.UserPayloadScheme) (*model.UserScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/user", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.UserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalUserImpl) Delete(ctx context.Context, accountId string) (*model.ResponseScheme, error) {

	if accountId == "" {
		return nil, model.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("accountId", accountId)
	endpoint := fmt.Sprintf("rest/api/%v/user?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalUserImpl) Find(ctx context.Context, accountIds []string, startAt, maxResults int) (*model.UserSearchPageScheme, *model.ResponseScheme, error) {

	if len(accountIds) == 0 {
		return nil, nil, model.ErrNoAccountSliceError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, accountID := range accountIds {
		params.Add("accountId", accountID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/user/bulk?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.UserSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalUserImpl) Groups(ctx context.Context, accountId string) ([]*model.UserGroupScheme, *model.ResponseScheme, error) {

	if accountId == "" {
		return nil, nil, model.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("accountId", accountId)
	endpoint := fmt.Sprintf("rest/api/%v/user/groups?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var groups []*model.UserGroupScheme
	response, err := i.c.Call(request, &groups)
	if err != nil {
		return nil, response, err
	}

	return groups, response, nil
}

func (i *internalUserImpl) Gets(ctx context.Context, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/users/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*model.UserScheme
	response, err := i.c.Call(request, &users)
	if err != nil {
		return nil, response, err
	}

	return users, response, nil
}
