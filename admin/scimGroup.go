package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SCIMGroupService struct{ client *Client }

func (g *SCIMGroupService) Gets(ctx context.Context, directoryID, filter string, startAt, maxResults int) (result *ScimGroupPageScheme, response *Response, err error) {

	if directoryID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	params := url.Values{}
	params.Add("startIndex", strconv.Itoa(startAt))
	params.Add("count", strconv.Itoa(maxResults))

	if filter != "" {
		params.Add("filter", filter)
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups?%v", directoryID, params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScimGroupPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (g *SCIMGroupService) Get(ctx context.Context, directoryID, groupID string) (result *ScimGroupScheme, response *Response, err error) {

	if directoryID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if groupID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid groupID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScimGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (g *SCIMGroupService) Update(ctx context.Context, directoryID, groupID string, newGroupName string) (result *ScimGroupScheme, response *Response, err error) {

	if directoryID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if groupID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid groupID value")
	}

	if newGroupName == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid newGroupName value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	payload := struct {
		DisplayName string `json:"displayName"`
	}{
		DisplayName: newGroupName,
	}

	request, err := g.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScimGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (g *SCIMGroupService) Delete(ctx context.Context, directoryID, groupID string) (response *Response, err error) {

	if directoryID == "" {
		return nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if groupID == "" {
		return nil, fmt.Errorf("error!, please provide a valid groupID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := g.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	return
}

func (g *SCIMGroupService) Create(ctx context.Context, directoryID, groupName string) (result *ScimGroupScheme, response *Response, err error) {

	if directoryID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if groupName == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid groupName value")
	}

	payload := struct {
		DisplayName string `json:"displayName"`
	}{
		DisplayName: groupName,
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups", directoryID)

	request, err := g.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScimGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (g *SCIMGroupService) Path(ctx context.Context, directoryID, groupID string, payload *SCIMGroupPathScheme) (result *ScimGroupScheme, response *Response, err error) {

	if directoryID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if groupID == "" {
		return nil, nil, fmt.Errorf("error!, please provide a valid groupID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("erro!, please provide a SCIMGroupPathScheme pointer")
	}

	if len(payload.Operations) == 0 {
		return nil, nil, fmt.Errorf("erro!, the SCIMGroupPathScheme value must contains operations")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := g.client.newRequest(ctx, http.MethodPatch, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScimGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SCIMGroupPathScheme struct {
	Schemas    []string                    `json:"schemas,omitempty"`
	Operations []*SCIMGroupOperationScheme `json:"Operations,omitempty"`
}

type SCIMGroupOperationScheme struct {
	Op    string                           `json:"op,omitempty"`
	Path  string                           `json:"path,omitempty"`
	Value []*SCIMGroupOperationValueScheme `json:"value,omitempty"`
}

type SCIMGroupOperationValueScheme struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
}

type ScimGroupPageScheme struct {
	Schemas      []string           `json:"schemas,omitempty"`
	TotalResults int                `json:"totalResults,omitempty"`
	StartIndex   int                `json:"startIndex,omitempty"`
	ItemsPerPage int                `json:"itemsPerPage,omitempty"`
	Resources    []*ScimGroupScheme `json:"Resources,omitempty"`
}

type ScimGroupScheme struct {
	Schemas     []string                 `json:"schemas,omitempty"`
	ID          string                   `json:"id,omitempty"`
	ExternalID  string                   `json:"externalId,omitempty"`
	DisplayName string                   `json:"displayName,omitempty"`
	Members     []*ScimGroupMemberScheme `json:"members,omitempty"`
	Meta        *ScimMetadata            `json:"meta,omitempty"`
}

type ScimGroupMemberScheme struct {
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type ScimMetadata struct {
	ResourceType string `json:"resourceType,omitempty"`
	Location     string `json:"location,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	Created      string `json:"created,omitempty"`
}
