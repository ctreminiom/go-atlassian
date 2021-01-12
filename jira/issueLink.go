package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type IssueLinkService struct {
	client *Client
	Type   *IssueLinkTypeService
}

type issueLinkPayloadScheme struct {
	OutwardIssue struct {
		Key string `json:"key"`
	} `json:"outwardIssue"`
	InwardIssue struct {
		Key string `json:"key"`
	} `json:"inwardIssue"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

// Creates a link between two issues. Use this operation to indicate a relationship between two issues
// and optionally add a comment to the from (outward) issue.
// To use this resource the site must have Issue Linking enabled.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-post
func (i *IssueLinkService) Create(ctx context.Context, linkType, inwardIssue, outwardIssue string) (response *Response, err error) {

	payload := issueLinkPayloadScheme{
		OutwardIssue: struct {
			Key string `json:"key"`
		}{Key: outwardIssue},
		InwardIssue: struct {
			Key string `json:"key"`
		}{Key: inwardIssue},
		Type: struct {
			Name string `json:"name"`
		}{Name: linkType},
	}

	var endpoint = "rest/api/3/issueLink"
	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

type IssueLinkScheme struct {
	ID   string `json:"id"`
	Type struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Inward  string `json:"inward"`
		Outward string `json:"outward"`
		Self    string `json:"self"`
	} `json:"type"`
	InwardIssue struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Self   string `json:"self"`
		Fields struct {
			Status struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
				} `json:"statusCategory"`
			} `json:"status"`
			Priority struct {
				Self        string `json:"self"`
				StatusColor string `json:"statusColor"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				ID          string `json:"id"`
			} `json:"priority"`
			Issuetype struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
				AvatarID    int    `json:"avatarId"`
				EntityID    string `json:"entityId"`
				Scope       struct {
					Type    string `json:"type"`
					Project struct {
						ID   string `json:"id"`
						Key  string `json:"key"`
						Name string `json:"name"`
					} `json:"project"`
				} `json:"scope"`
			} `json:"issuetype"`
		} `json:"fields"`
	} `json:"inwardIssue"`
	OutwardIssue struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Self   string `json:"self"`
		Fields struct {
			Status struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Priority struct {
				Self        string `json:"self"`
				StatusColor string `json:"statusColor"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				ID          string `json:"id"`
			} `json:"priority"`
			Issuetype struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
				AvatarID    int    `json:"avatarId"`
			} `json:"issuetype"`
		} `json:"fields"`
	} `json:"outwardIssue"`
}

// Returns an issue link.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-get
func (i *IssueLinkService) Get(ctx context.Context, linkID string) (result *IssueLinkScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issueLink/%v", linkID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueLinkScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes an issue link.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-delete
func (i *IssueLinkService) Delete(ctx context.Context, linkID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issueLink/%v", linkID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}
