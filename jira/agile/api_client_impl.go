package agile

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ctreminiom/go-atlassian/v2/jira/agile/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

func New(httpClient common.HTTPClient, site string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		return nil, model.ErrNoSite
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	u, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HTTP: httpClient,
		Site: u,
	}

	client.Board = internal.NewBoardService(client, "1.0")
	client.Epic = internal.NewEpicService(client, "1.0")
	client.Sprint = internal.NewSprintService(client, "1.0")
	client.Backlog = internal.NewBoardBacklogService(client, "1.0")
	client.Auth = internal.NewAuthenticationService(client)

	return client, nil
}

type Client struct {
	HTTP    common.HTTPClient
	Site    *url.URL
	Auth    common.Authentication
	Board   *internal.BoardService
	Backlog *internal.BoardBacklogService
	Epic    *internal.EpicService
	Sprint  *internal.SprintService
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.Site.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	if body != nil && contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if body != nil && contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	if c.Auth.HasBasicAuth() {
		req.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	if c.Auth.GetBearerToken() != "" && !c.Auth.HasBasicAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	return req, nil
}

func (c *Client) Call(request *http.Request, structure interface{}) (*model.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response, structure)
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*model.ResponseScheme, error) {

	defer response.Body.Close()

	res := &model.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	responseAsBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	res.Bytes.Write(responseAsBytes)

	wasSuccess := response.StatusCode >= 200 && response.StatusCode < 300

	if !wasSuccess {

		switch response.StatusCode {

		case http.StatusNotFound:
			return res, model.ErrNotFound

		case http.StatusUnauthorized:
			return res, model.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, model.ErrInternal

		case http.StatusBadRequest:
			return res, model.ErrBadRequest

		default:
			return res, model.ErrInvalidStatusCode
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
