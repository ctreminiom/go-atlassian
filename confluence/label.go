package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type LabelService struct{ client *Client }

// Get returns label information and a list of contents associated with the label.
func (l *LabelService) Get(ctx context.Context, labelName, labelType string, start, limit int) (result *models.LabelDetailsScheme,
	response *ResponseScheme, err error) {

	if labelName == "" {
		return nil, nil, models.ErrNoLabelNameError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))
	query.Add("name", labelName)
	query.Add("type", labelType)

	endpoint := fmt.Sprintf("wiki/rest/api/label?%v", query.Encode())

	request, err := l.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = l.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
