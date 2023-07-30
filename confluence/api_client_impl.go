package confluence

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/confluence/internal"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func New(httpClient common.HttpClient, site string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		return nil, models.ErrNoSiteError
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

	contentSubServices := &internal.ContentSubServices{
		Attachment:         internal.NewContentAttachmentService(client),
		ChildrenDescendant: internal.NewChildrenDescandantsService(client),
		Comment:            internal.NewCommentService(client),
		Permission:         internal.NewPermissionService(client),
		Label:              internal.NewContentLabelService(client),
		Property:           internal.NewPropertyService(client),
		Version:            internal.NewVersionService(client),
		Restriction: internal.NewRestrictionService(client,
			internal.NewRestrictionOperationService(client,
				internal.NewRestrictionOperationGroupService(client),
				internal.NewRestrictionOperationUserService(client))),
	}

	client.Auth = internal.NewAuthenticationService(client)
	client.Content = internal.NewContentService(client, contentSubServices)
	client.Space = internal.NewSpaceService(client, internal.NewSpacePermissionService(client))
	client.Label = internal.NewLabelService(client)
	client.Search = internal.NewSearchService(client)
	client.LongTask = internal.NewTaskService(client)
	client.Analytics = internal.NewAnalyticsService(client)

	return client, nil
}

type Client struct {
	HTTP      common.HttpClient
	Site      *url.URL
	Auth      common.Authentication
	Content   *internal.ContentService
	Space     *internal.SpaceService
	Label     *internal.LabelService
	Search    *internal.SearchService
	LongTask  *internal.TaskService
	Analytics *internal.AnalyticsService
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr, type_ string, body interface{}) (*http.Request, error) {

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

	// If the body interface is a *bytes.Buffer type
	// it means the NewRequest() requires to handle the RFC 1867 ISO
	if attachBuffer, ok := body.(*bytes.Buffer); ok {
		buf = attachBuffer
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if type_ != "" {
		// When the type_ is provided, it means the request needs to be created to handle files
		req.Header.Set("Content-Type", type_)
		req.Header.Set("X-Atlassian-Token", "no-check")
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

func (c *Client) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response, structure)
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	defer response.Body.Close()

	res := &models.ResponseScheme{
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
			return res, models.ErrNotFound

		case http.StatusUnauthorized:
			return res, models.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, models.ErrInternalError

		case http.StatusBadRequest:
			return res, models.ErrBadRequestError

		default:
			return res, models.ErrInvalidStatusCodeError
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
