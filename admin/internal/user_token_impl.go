package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/admin"
	"net/http"
)

func NewUserTokenService(client service.Connector) *UserTokenService {
	return &UserTokenService{internalClient: &internalUserTokenImpl{c: client}}
}

type UserTokenService struct {
	internalClient admin.UserTokenConnector
}

// Gets gets the API tokens owned by the specified user.
//
// GET /users/{account_id}/manage/api-tokens
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#get-api-tokens
func (u *UserTokenService) Gets(ctx context.Context, accountID string) ([]*model.UserTokensScheme, *model.ResponseScheme, error) {
	return u.internalClient.Gets(ctx, accountID)
}

// Delete deletes a specified API token by ID.
//
// DELETE /users/{account_id}/manage/api-tokens/{tokenId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#delete-api-token
func (u *UserTokenService) Delete(ctx context.Context, accountID, tokenID string) (*model.ResponseScheme, error) {
	return u.internalClient.Delete(ctx, accountID, tokenID)
}

type internalUserTokenImpl struct {
	c service.Connector
}

func (i *internalUserTokenImpl) Gets(ctx context.Context, accountID string) ([]*model.UserTokensScheme, *model.ResponseScheme, error) {

	if accountID == "" {
		return nil, nil, model.ErrNoAdminAccountIDError
	}

	endpoint := fmt.Sprintf("users/%v/manage/api-tokens", accountID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var tokens []*model.UserTokensScheme
	response, err := i.c.Call(request, &tokens)
	if err != nil {
		return nil, response, err
	}

	return tokens, response, nil
}

func (i *internalUserTokenImpl) Delete(ctx context.Context, accountID, tokenID string) (*model.ResponseScheme, error) {

	if accountID == "" {
		return nil, model.ErrNoAdminAccountIDError
	}

	if tokenID == "" {
		return nil, model.ErrNoAdminUserTokenError
	}

	endpoint := fmt.Sprintf("users/%v/manage/api-tokens/%v", accountID, tokenID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
