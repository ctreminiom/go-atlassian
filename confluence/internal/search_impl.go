package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewSearchService creates a new instance of SearchService.
// It takes a service.Connector as input and returns a pointer to SearchService.
func NewSearchService(client service.Connector) *SearchService {
	return &SearchService{
		internalClient: &internalSearchImpl{c: client},
	}
}

// SearchService provides methods to interact with search operations in Confluence.
type SearchService struct {
	// internalClient is the connector interface for search operations.
	internalClient confluence.SearchConnector
}

// Content searches for content using the Confluence Query Language (CQL)
//
// GET /wiki/rest/api/search
//
// https://docs.go-atlassian.io/confluence-cloud/search#search-content
func (s *SearchService) Content(ctx context.Context, cql string, options *model.SearchContentOptions) (*model.SearchPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Content(ctx, cql, options)
}

// Users searches for users using user-specific queries from the Confluence Query Language (CQL).
//
// Note that some user fields may be set to null depending on the user's privacy settings.
//
// These are: email, profilePicture, displayName, and timeZone.
//
// GET /wiki/rest/api/search/user
//
// https://docs.go-atlassian.io/confluence-cloud/search#search-users
func (s *SearchService) Users(ctx context.Context, cql string, start, limit int, expand []string) (*model.SearchPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Users(ctx, cql, start, limit, expand)
}

type internalSearchImpl struct {
	c service.Connector
}

func (i *internalSearchImpl) Content(ctx context.Context, cql string, options *model.SearchContentOptions) (*model.SearchPageScheme, *model.ResponseScheme, error) {

	if cql == "" {
		return nil, nil, model.ErrNoCQL
	}

	query := url.Values{}
	query.Add("cql", cql)

	if options != nil {

		if options.Context != "" {
			query.Add("cqlcontext", options.Context)
		}

		if options.Cursor != "" {
			query.Add("cursor", options.Cursor)
		}

		if options.Next {
			query.Add("next", "true")
		}

		if options.Prev {
			query.Add("prev", "true")
		}

		if options.Limit != 0 {
			query.Add("limit", strconv.Itoa(options.Limit))
		}

		if options.Start != 0 {
			query.Add("start", strconv.Itoa(options.Start))
		}

		if options.IncludeArchivedSpaces {
			query.Add("includeArchivedSpaces", "true")
		}

		if options.ExcludeCurrentSpaces {
			query.Add("excludeCurrentSpaces", "true")
		}

		if options.Excerpt != "" {
			query.Add("excerpt", options.Excerpt)
		}

		if options.SitePermissionTypeFilter != "" {
			query.Add("sitePermissionTypeFilter", options.SitePermissionTypeFilter)
		}

		if len(options.Expand) > 0 {
			for _, value := range options.Expand {
				query.Add("expand", value)
			}
		}
	}

	endpoint := fmt.Sprintf("wiki/rest/api/search?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.SearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalSearchImpl) Users(ctx context.Context, cql string, start, limit int, expand []string) (*model.SearchPageScheme, *model.ResponseScheme, error) {

	if cql == "" {
		return nil, nil, model.ErrNoCQL
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("start", strconv.Itoa(start))
	query.Add("cql", cql)

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/search/user?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.SearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
