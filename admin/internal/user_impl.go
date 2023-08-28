package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/admin"
	"net/http"
	"net/url"
	"strings"
)

func NewUserService(client service.Connector, token *UserTokenService) *UserService {
	return &UserService{
		internalClient: &internalUserImpl{c: client},
		Token:          token,
	}
}

type UserService struct {
	internalClient admin.UserConnector
	Token          *UserTokenService
}

// Permissions returns the set of permissions you have for managing the specified Atlassian account
//
// GET /users/{account_id}/manage
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-user-management-permissions
func (u *UserService) Permissions(ctx context.Context, accountID string, privileges []string) (*model.AdminUserPermissionScheme, *model.ResponseScheme, error) {
	return u.internalClient.Permissions(ctx, accountID, privileges)
}

// Get returns information about a single Atlassian account by ID
//
// GET /users/{account_id}/manage/profile
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-profile
func (u *UserService) Get(ctx context.Context, accountID string) (*model.AdminUserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Get(ctx, accountID)
}

// Update updates fields in a user account. The profile.write privilege details which fields you can change.
//
// PATCH /users/{account_id}/manage/profile
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user#update-profile
func (u *UserService) Update(ctx context.Context, accountID string, payload map[string]interface{}) (*model.AdminUserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Update(ctx, accountID, payload)
}

// Disable deactivate the specified user account.
//
// The permission to make use of this resource is exposed by the lifecycle.enablement privilege.
//
// You can optionally set a message associated with the block.
//
// If none is supplied, a default message will be used.
//
// POST /users/{account_id}/manage/lifecycle/disable
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user#disable-a-user
func (u *UserService) Disable(ctx context.Context, accountID, message string) (*model.ResponseScheme, error) {
	return u.internalClient.Disable(ctx, accountID, message)
}

// Enable activates the specified user account.
//
// The permission to make use of this resource is exposed by the lifecycle.enablement privilege.
//
// POST /users/{account_id}/manage/lifecycle/enable
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user#enable-a-user
func (u *UserService) Enable(ctx context.Context, accountID string) (*model.ResponseScheme, error) {
	return u.internalClient.Enable(ctx, accountID)
}

type internalUserImpl struct {
	c service.Connector
}

func (i *internalUserImpl) Permissions(ctx context.Context, accountID string, privileges []string) (*model.AdminUserPermissionScheme, *model.ResponseScheme, error) {

	if accountID == "" {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("users/%v/manage", accountID))

	if len(privileges) != 0 {

		params := url.Values{}
		params.Add("privileges", strings.Join(privileges, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(model.AdminUserPermissionScheme)
	response, err := i.c.Call(request, permissions)
	if err != nil {
		return nil, response, err
	}

	return permissions, response, nil
}

func (i *internalUserImpl) Get(ctx context.Context, accountID string) (*model.AdminUserScheme, *model.ResponseScheme, error) {

	if accountID == "" {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	endpoint := fmt.Sprintf("users/%v/manage/profile", accountID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.AdminUserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalUserImpl) Update(ctx context.Context, accountID string, payload map[string]interface{}) (*model.AdminUserScheme, *model.ResponseScheme, error) {

	if accountID == "" {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	endpoint := fmt.Sprintf("users/%v/manage/profile", accountID)

	request, err := i.c.NewRequest(ctx, http.MethodPatch, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.AdminUserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalUserImpl) Disable(ctx context.Context, accountID, message string) (*model.ResponseScheme, error) {

	if accountID == "" {
		return nil, model.ErrNoAdminAccountIDError
	}

	endpoint := fmt.Sprintf("users/%v/manage/lifecycle/disable", accountID)

	payload := map[string]interface{}{}
	if message != "" {
		payload["message"] = message
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalUserImpl) Enable(ctx context.Context, accountID string) (*model.ResponseScheme, error) {

	if accountID == "" {
		return nil, model.ErrNoAdminAccountIDError
	}

	endpoint := fmt.Sprintf("users/%v/manage/lifecycle/enable", accountID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
