package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type IssueSecurityService struct {
	client *Client
	Scheme *IssueSecuritySchemeService
}

type IssueSecurityLevelScheme struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Returns details of an issue security level.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-security-level/#api-rest-api-3-securitylevel-id-get
func (i *IssueSecurityService) Level(ctx context.Context, issueSecurityLevelID string) (result *IssueSecurityLevelScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/securitylevel/%v", issueSecurityLevelID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSecurityLevelScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueSecurityMembers struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID                   int `json:"id"`
		IssueSecurityLevelID int `json:"issueSecurityLevelId"`
		Holder               struct {
			Type string `json:"type"`
			User struct {
				Self         string `json:"self"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"user"`
			Expand string `json:"expand"`
		} `json:"holder"`
	} `json:"values"`
}

func (i *IssueSecurityService) Members(ctx context.Context, issueSecuritySchemeID int, issueSecurityLevelIDs []int, expands []string, startAt, maxResults int) (result *IssueSecurityMembers, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var levelIDs string
	for index, value := range issueSecurityLevelIDs {

		if index == 0 {
			levelIDs = strconv.Itoa(value)
			continue
		}

		levelIDs += "," + strconv.Itoa(value)
	}

	if len(levelIDs) != 0 {
		params.Add("expand", levelIDs)
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuesecurityschemes/%v/members?%v", issueSecuritySchemeID, params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSecurityMembers)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
