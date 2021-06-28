package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentCommentService struct {
	client *Client
}

// Gets returns the comments on a piece of content.
func (c *ContentCommentService) Gets(ctx context.Context, contentID string, expand, location []string,
	startAt, maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error){

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if len(location) != 0 {
		query.Add("location", strings.Join(location, ","))
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/child/comment?%v", contentID, query.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}