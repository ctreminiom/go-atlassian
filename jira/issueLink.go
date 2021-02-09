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
func (i *IssueLinkService) Create(ctx context.Context, linkType, inWardIssue, outWardIssue string) (response *Response, err error) {

	payload := issueLinkPayloadScheme{
		OutwardIssue: struct {
			Key string `json:"key"`
		}{Key: inWardIssue},
		InwardIssue: struct {
			Key string `json:"key"`
		}{Key: outWardIssue},
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
	ID   string `json:"id,omitempty"`
	Type struct {
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Inward  string `json:"inward,omitempty"`
		Outward string `json:"outward,omitempty"`
		Self    string `json:"self,omitempty"`
	} `json:"type,omitempty"`
	InwardIssue struct {
		ID     string `json:"id,omitempty"`
		Key    string `json:"key,omitempty"`
		Self   string `json:"self,omitempty"`
		Fields struct {
			Status struct {
				Self           string `json:"self,omitempty"`
				Description    string `json:"description,omitempty"`
				IconURL        string `json:"iconUrl,omitempty"`
				Name           string `json:"name,omitempty"`
				ID             string `json:"id,omitempty"`
				StatusCategory struct {
					Self      string `json:"self,omitempty"`
					ID        int    `json:"id,omitempty"`
					Key       string `json:"key,omitempty"`
					ColorName string `json:"colorName,omitempty"`
				} `json:"statusCategory,omitempty"`
			} `json:"status,omitempty"`
			Priority  PriorityScheme `json:"priority,omitempty"`
			Issuetype struct {
				Self        string `json:"self,omitempty"`
				ID          string `json:"id,omitempty"`
				Description string `json:"description,omitempty"`
				IconURL     string `json:"iconUrl,omitempty"`
				Name        string `json:"name,omitempty"`
				Subtask     bool   `json:"subtask,omitempty"`
				AvatarID    int    `json:"avatarId,omitempty"`
				EntityID    string `json:"entityId,omitempty"`
				Scope       struct {
					Type    string `json:"type,omitempty"`
					Project struct {
						ID   string `json:"id,omitempty"`
						Key  string `json:"key,omitempty"`
						Name string `json:"name,omitempty"`
					} `json:"project,omitempty"`
				} `json:"scope,omitempty"`
			} `json:"issuetype,omitempty"`
		} `json:"fields,omitempty"`
	} `json:"inwardIssue,omitempty"`
	OutwardIssue struct {
		ID     string `json:"id,omitempty"`
		Key    string `json:"key,omitempty"`
		Self   string `json:"self,omitempty"`
		Fields struct {
			Status struct {
				Self           string `json:"self,omitempty"`
				Description    string `json:"description,omitempty"`
				IconURL        string `json:"iconUrl,omitempty"`
				Name           string `json:"name,omitempty"`
				ID             string `json:"id,omitempty"`
				StatusCategory struct {
					Self      string `json:"self,omitempty"`
					ID        int    `json:"id,omitempty"`
					Key       string `json:"key,omitempty"`
					ColorName string `json:"colorName,omitempty"`
					Name      string `json:"name,omitempty"`
				} `json:"statusCategory,omitempty"`
			} `json:"status,omitempty"`
			Priority  PriorityScheme `json:"priority,omitempty"`
			Issuetype struct {
				Self        string `json:"self,omitempty"`
				ID          string `json:"id,omitempty"`
				Description string `json:"description,omitempty"`
				IconURL     string `json:"iconUrl,omitempty"`
				Name        string `json:"name,omitempty"`
				Subtask     bool   `json:"subtask,omitempty"`
				AvatarID    int    `json:"avatarId,omitempty"`
			} `json:"issuetype,omitempty"`
		} `json:"fields,omitempty"`
	} `json:"outwardIssue,omitempty"`
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

type IssueLinksScheme struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		IssueLinks []IssueLinkScheme `json:"issuelinks"`
	} `json:"fields"`
}

// Get the issue links ID's associated with a Jira Issue
func (i *IssueLinkService) Gets(ctx context.Context, issueKeyOrID string) (result *IssueLinksScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v?fields=issuelinks", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueLinksScheme)
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
