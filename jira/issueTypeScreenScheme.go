package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/url"
	"strconv"
)

type IssueTypeScreenSchemeService struct{ client *Client }

type IssueTypeScreenSchemesScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"values"`
}

func (i *IssueTypeScreenSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (result *IssueTypeScreenSchemesScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var issueTypeSchemeIDs string
	for index, value := range ids {

		if index == 0 {
			issueTypeSchemeIDs = strconv.Itoa(value)
			continue
		}

		issueTypeSchemeIDs += "," + strconv.Itoa(value)
	}

	if len(issueTypeSchemeIDs) != 0 {
		params.Add("id", issueTypeSchemeIDs)
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeScreenSchemesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueTypeScreenSchemePayloadScheme struct {
	Name              string                                      `json:"name" validate:"required"`
	IssueTypeMappings []IssueTypeScreenSchemeMappingPayloadScheme `json:"issueTypeMappings" validate:"required"`
}

type IssueTypeScreenSchemeMappingPayloadScheme struct {
	IssueTypeID    string `json:"issueTypeId"`
	ScreenSchemeID string `json:"screenSchemeId"`
}

type issueTypeScreenScreenCreatedScheme struct {
	ID string `json:"id"`
}

// Creates an issue type screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-post
func (i *IssueTypeScreenSchemeService) Create(ctx context.Context, payload *IssueTypeScreenSchemePayloadScheme) (issueTypeScreenSchemeID int, response *Response, err error) {

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/issuetypescreenscheme"

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result := new(issueTypeScreenScreenCreatedScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	newIssueTypeScreenSchemeIDAsInt, err := strconv.Atoi(result.ID)
	if err != nil {
		return
	}

	issueTypeScreenSchemeID = newIssueTypeScreenSchemeIDAsInt
	return
}
