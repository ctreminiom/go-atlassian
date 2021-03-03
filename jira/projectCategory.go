package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectCategoryService struct{ client *Client }

type ProjectCategoryScheme struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Returns all project categories.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-all-project-categories
func (p *ProjectCategoryService) Gets(ctx context.Context) (result *[]ProjectCategoryScheme, response *Response, err error) {

	var endpoint = "rest/api/3/projectCategory"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectCategoryScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-project-category-by-id
func (p *ProjectCategoryService) Get(ctx context.Context, projectCategoryID int) (result *ProjectCategoryScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/projectCategory/%v", projectCategoryID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectCategoryScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#create-project-category
func (p *ProjectCategoryService) Create(ctx context.Context, name, description string) (result *ProjectCategoryScheme, response *Response, err error) {

	if len(name) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a project category name")
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = "rest/api/3/projectCategory"

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectCategoryScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#update-project-category
func (p *ProjectCategoryService) Update(ctx context.Context, projectCategoryID int, name, description string) (result *ProjectCategoryScheme, response *Response, err error) {

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/api/3/projectCategory/%v", projectCategoryID)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectCategoryScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#delete-project-category
func (p *ProjectCategoryService) Delete(ctx context.Context, projectCategoryID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/projectCategory/%v", projectCategoryID)

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}
