package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type UserTokenService struct{ client *Client }

// Gets the API tokens owned by the specified user, this func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-accountid-manage-api-tokens-get
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#get-api-tokens
func (u *UserTokenService) Gets(ctx context.Context, accountID string) (result *model.UserTokensScheme,
	response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/api-tokens", accountID)

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

// Delete deletes a specified API token by ID, this func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-accountid-manage-api-tokens-tokenid-delete
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#delete-api-token
func (u *UserTokenService) Delete(ctx context.Context, accountID, tokenID string) (response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, model.ErrNoAdminAccountIDError
	}

	if len(tokenID) == 0 {
		return nil, model.ErrNoAdminUserTokenError
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/api-tokens/%v", accountID, tokenID)

	request, err := u.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
