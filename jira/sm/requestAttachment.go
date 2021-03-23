package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestAttachmentService struct{ client *Client }

func (r *RequestAttachmentService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *RequestAttachmentPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/attachment?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestAttachmentPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestAttachmentService) Create(ctx context.Context, issueKeyOrID string, temporaryAttachmentIDs []string, public bool) (result *RequestAttachmentCreationScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if len(temporaryAttachmentIDs) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid temporaryAttachmentIDs slice value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/attachment", issueKeyOrID)

	payload := struct {
		TemporaryAttachmentIds []string `json:"temporaryAttachmentIds"`
		Public                 bool     `json:"public"`
	}{
		TemporaryAttachmentIds: temporaryAttachmentIDs,
		Public:                 public,
	}

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestAttachmentCreationScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type RequestAttachmentPageScheme struct {
	Size       int                        `json:"size"`
	Start      int                        `json:"start"`
	Limit      int                        `json:"limit"`
	IsLastPage bool                       `json:"isLastPage"`
	Values     []*RequestAttachmentScheme `json:"values"`
	Expands    []string                   `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type RequestAttachmentScheme struct {
	Filename string `json:"filename"`
	Author   struct {
		AccountID    string `json:"accountId"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		TimeZone     string `json:"timeZone"`
		Links        struct {
		} `json:"_links"`
	} `json:"author"`
	Created struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int    `json:"epochMillis"`
	} `json:"created"`
	Size     int    `json:"size"`
	MimeType string `json:"mimeType"`
	Links    struct {
		Self      string `json:"self"`
		JiraRest  string `json:"jiraRest"`
		Content   string `json:"content"`
		Thumbnail string `json:"thumbnail"`
	} `json:"_links"`
}

type RequestAttachmentCreationScheme struct {
	Comment struct {
		Expands []string `json:"_expands"`
		ID      string   `json:"id"`
		Body    string   `json:"body"`
		Public  bool     `json:"public"`
		Author  struct {
			AccountID    string `json:"accountId"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			TimeZone     string `json:"timeZone"`
			Links        struct {
				JiraRest   string `json:"jiraRest"`
				AvatarUrls struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"author"`
		Created struct {
			Iso8601     string `json:"iso8601"`
			Jira        string `json:"jira"`
			Friendly    string `json:"friendly"`
			EpochMillis int64  `json:"epochMillis"`
		} `json:"created"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"comment"`
	Attachments struct {
		Expands    []interface{} `json:"_expands"`
		Size       int           `json:"size"`
		Start      int           `json:"start"`
		Limit      int           `json:"limit"`
		IsLastPage bool          `json:"isLastPage"`
		Links      struct {
			Base    string `json:"base"`
			Context string `json:"context"`
			Next    string `json:"next"`
			Prev    string `json:"prev"`
		} `json:"_links"`
		Values []struct {
			Filename string `json:"filename"`
			Author   struct {
				AccountID    string `json:"accountId"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				DisplayName  string `json:"displayName"`
				Active       bool   `json:"active"`
				TimeZone     string `json:"timeZone"`
				Links        struct {
					JiraRest   string `json:"jiraRest"`
					AvatarUrls struct {
						Four8X48  string `json:"48x48"`
						Two4X24   string `json:"24x24"`
						One6X16   string `json:"16x16"`
						Three2X32 string `json:"32x32"`
					} `json:"avatarUrls"`
					Self string `json:"self"`
				} `json:"_links"`
			} `json:"author"`
			Created struct {
				Iso8601     string `json:"iso8601"`
				Jira        string `json:"jira"`
				Friendly    string `json:"friendly"`
				EpochMillis int64  `json:"epochMillis"`
			} `json:"created"`
			Size     int    `json:"size"`
			MimeType string `json:"mimeType"`
			Links    struct {
				JiraRest  string `json:"jiraRest"`
				Content   string `json:"content"`
				Thumbnail string `json:"thumbnail"`
			} `json:"_links"`
		} `json:"values"`
	} `json:"attachments"`
}
