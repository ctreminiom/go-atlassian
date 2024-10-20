package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type SearchConnector interface {

	// Content searches for content using the Confluence Query Language (CQL)
	//
	// GET /wiki/rest/api/search
	//
	// https://docs.go-atlassian.io/confluence-cloud/search#search-content
	Content(ctx context.Context, cql string, options *model.SearchContentOptions) (*model.SearchPageScheme, *model.ResponseScheme, error)

	// Users searches for users using user-specific queries from the Confluence Query Language (CQL).
	//
	// Note that some user fields may be set to null depending on the user's privacy settings.
	//
	// These are: email, profilePicture, displayName, and timeZone.
	//
	// GET /wiki/rest/api/search/user
	//
	// https://docs.go-atlassian.io/confluence-cloud/search#search-users
	Users(ctx context.Context, cql string, start, limit int, expand []string) (*model.SearchPageScheme, *model.ResponseScheme, error)
}
