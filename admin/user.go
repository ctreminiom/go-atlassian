package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strings"
)

type UserService struct {
	client *Client
	Token  *UserTokenService
}

// Permissions returns the set of permissions you have for managing the specified Atlassian account, this func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-user-management-permissions
func (u *UserService) Permissions(ctx context.Context, accountID string, privileges []string) (result *model.AdminUserPermissionScheme,
	response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	params := url.Values{}
	if len(privileges) != 0 {
		params.Add("privileges", strings.Join(privileges, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/users/%v/manage", accountID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns information about a single Atlassian account by ID, this func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-profile-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-profile
func (u *UserService) Get(ctx context.Context, accountID string) (result *model.AdminUserScheme, response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/profile", accountID)

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates fields in a user account. The profile.write privilege details which fields you can change
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-profile-patch
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#update-profile
func (u *UserService) Update(ctx context.Context, accountID string, payload map[string]interface{}) (
	result *model.AdminUserScheme, response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid payload map with keys")
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/profile", accountID)

	request, err := u.client.newRequest(ctx, http.MethodPatch, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Disable disables the specified user account.
// The permission to make use of this resource is exposed by the lifecycle.enablement privilege
// You can optionally set a message associated with the block that will be shown to the user on attempted authentication.
// If none is supplied, a default message will be used.
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-lifecycle-disable-post
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#disable-a-user
func (u *UserService) Disable(ctx context.Context, accountID, message string) (response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, model.ErrNoAdminAccountIDError
	}

	var (
		endpoint = fmt.Sprintf("/users/%v/manage/lifecycle/disable", accountID)
		request  *http.Request
	)

	if len(message) != 0 {

		payload := struct {
			Message string `json:"message"`
		}{
			Message: message,
		}

		payloadAsReader, _ := transformStructToReader(&payload)
		request, err = u.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
		if err != nil {
			return
		}

		request.Header.Set("Accept", "application/json")
		request.Header.Set("Content-Type", "application/json")

	} else {
		request, err = u.client.newRequest(ctx, http.MethodPost, endpoint, nil)
		if err != nil {
			return
		}

		request.Header.Set("Accept", "application/json")
	}

	response, err = u.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Enable enables the specified user account.
// The permission to make use of this resource is exposed by the lifecycle.enablement privilege.
// This func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-lifecycle-enable-post
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#enable-a-user
func (u *UserService) Enable(ctx context.Context, accountID string) (response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, model.ErrNoAdminAccountIDError
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/lifecycle/enable", accountID)

	request, err := u.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = u.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
