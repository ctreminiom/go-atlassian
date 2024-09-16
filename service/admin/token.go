package admin

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// UserTokenConnector represents the cloud admin user token endpoints.
// Use it to search, get, create, delete, and change tokens.
type UserTokenConnector interface {

	// Gets gets the API tokens owned by the specified user.
	//
	// GET /users/{accountID}/manage/api-tokens
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#get-api-tokens
	Gets(ctx context.Context, accountID string) ([]*model.UserTokensScheme, *model.ResponseScheme, error)

	// Delete deletes a specified API token by ID.
	//
	// DELETE /users/{accountID}/manage/api-tokens/{tokenID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#delete-api-token
	Delete(ctx context.Context, accountID, tokenID string) (*model.ResponseScheme, error)
}
