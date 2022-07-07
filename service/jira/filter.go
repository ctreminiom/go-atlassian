package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// Filter is an interface that defines the methods available from Jira Filter API.
type Filter interface {

	// Create creates a filter. The filter is shared according to the default share scope.
	// The filter is not selected as a favorite.
	// POST /rest/api/{2-3}/filter
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#create-filter
	Create(ctx context.Context, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error)

	// Favorite returns the visible favorite filters of the user.
	// GET /rest/api/{2-3}/filter/favourite
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-favorites
	Favorite(ctx context.Context) ([]*model.FilterScheme, *model.ResponseScheme, error)

	// My returns the filters owned by the user. If includeFavourites is true,
	// the user's visible favorite filters are also returned.
	// GET /rest/api/{2-3}/filter/my
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-my-filters
	My(ctx context.Context, favorites bool, expand []string) ([]*model.FilterScheme, *model.ResponseScheme, error)

	// Search returns a paginated list of filters
	// GET /rest/api/{2-3}/filter/search
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#search-filters
	Search(ctx context.Context, options *model.FilterSearchOptionScheme, startAt, maxResults int) (*model.FilterSearchPageScheme,
		*model.ResponseScheme, error)

	// Get returns a filter.
	// GET /rest/api/{2-3}/filter/{id}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-filter
	Get(ctx context.Context, filterId int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error)

	// Update updates a filter. Use this operation to update a filter's name, description, JQL, or sharing.
	// PUT /rest/api/{2-3}/filter/{id}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#update-filter
	Update(ctx context.Context, filterId int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error)

	// Delete a filter.
	// DELETE /rest/api/{2-3}/filter/{id}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#delete-filter
	Delete(ctx context.Context, filterId int) (*model.ResponseScheme, error)

	// Change changes the owner of the filter.
	// PUT /rest/api/{2-3}/filter/{id}/owner
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#change-filter-owner
	Change(ctx context.Context, filterId int, accountId string) (*model.ResponseScheme, error)
}
