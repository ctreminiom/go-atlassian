package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserTokenService struct{ client *Client }

// Gets the API tokens owned by the specified user, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. accountID = The user account to manage (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-accountid-manage-api-tokens-get
// Library Example: N/A
func (u *UserTokenService) Gets(ctx context.Context, accountID string) (result *UserTokensScheme, response *Response, err error) {

	if len(accountID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid accountID value")
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/api-tokens", accountID)

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new(UserTokensScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type UserTokensScheme []struct {
	ID         string    `json:"id"`
	Label      string    `json:"label"`
	CreatedAt  time.Time `json:"createdAt"`
	LastAccess time.Time `json:"lastAccess"`
}

// Deletes a specified API token by ID, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. accountID = The user account to manage (REQUIRED)
// 3. tokenID = The ID of the API token
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-accountid-manage-api-tokens-tokenid-delete
func (u *UserTokenService) Delete(ctx context.Context, accountID, tokenID string) (response *Response, err error) {

	if len(accountID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid accountID value")
	}

	if len(tokenID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid tokenID value")
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/api-tokens/%v", accountID, tokenID)

	request, err := u.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	return
}
