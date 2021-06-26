package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentAttachmentService struct {
	client *Client
}

type GetContentAttachmentsOptionsScheme struct {
	Expand []string
	FileName string
	MediaType string
}

// Gets returns the attachments for a piece of content.
// By default, the following objects are expanded: metadata.
func (c *ContentAttachmentService) Gets(ctx context.Context, contentID string, startAt, maxResults int, options *GetContentAttachmentsOptionsScheme) (result *ContentPageScheme, response *ResponseScheme, err error)  {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			query.Add("expand", strings.Join(options.Expand, ","))
		}

		if options.FileName != "" {
			query.Add("filename", options.FileName)
		}

		if options.MediaType != "" {
			query.Add("mediaType", options.MediaType)
		}

	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/child/attachment?%v", contentID, query.Encode())

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
