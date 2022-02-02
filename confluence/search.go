package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SearchService struct{ client *Client }

// Content searches for content using the Confluence Query Language (CQL)
// Docs: https://docs.go-atlassian.io/confluence-cloud/search#search-content
func (s *SearchService) Content(ctx context.Context, cql string, options *models.SearchContentOptions) (result *models.SearchPageScheme, response *ResponseScheme, err error) {

	if cql == "" {
		return nil, nil, models.ErrNoCQLError
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
	}

	endpoint := fmt.Sprintf("wiki/rest/api/search?%v", query.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Users searches for users using user-specific queries from the Confluence Query Language (CQL).
// Docs: Searches for users using user-specific queries from the Confluence Query Language (CQL).
func (s *SearchService) Users(ctx context.Context, cql string, start, limit int, expand []string) (result *models.SearchPageScheme, response *ResponseScheme, err error) {

	if cql == "" {
		return nil, nil, models.ErrNoCQLError
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("start", strconv.Itoa(start))
	query.Add("cql", cql)

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/search/user?%v", query.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
