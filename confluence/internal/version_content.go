package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewVersionService creates a new instance of VersionService.
// It takes a service.Connector as input and returns a pointer to VersionService.
func NewVersionService(client service.Connector) *VersionService {
	return &VersionService{
		internalClient: &internalVersionImpl{c: client},
	}
}

// VersionService provides methods to interact with version operations in Confluence.
type VersionService struct {
	// internalClient is the connector interface for version operations.
	internalClient confluence.VersionConnector
}

// Gets returns the versions for a piece of content in descending order.
//
// GET /wiki/rest/api/content/{id}/version
//
// https://docs.go-atlassian.io/confluence-cloud/content/versions#get-content-versions
func (v *VersionService) Gets(ctx context.Context, contentID string, expand []string, start, limit int) (*model.ContentVersionPageScheme, *model.ResponseScheme, error) {
	return v.internalClient.Gets(ctx, contentID, expand, start, limit)
}

// Get returns a version for a piece of content.
//
// GET /wiki/rest/api/content/{id}/version/{versionNumber}
//
// https://docs.go-atlassian.io/confluence-cloud/content/versions#get-content-version
func (v *VersionService) Get(ctx context.Context, contentID string, versionNumber int, expand []string) (*model.ContentVersionScheme, *model.ResponseScheme, error) {
	return v.internalClient.Get(ctx, contentID, versionNumber, expand)
}

// Restore restores a historical version to be the latest version.
//
// That is, a new version is created with the content of the historical version.
//
// https://docs.go-atlassian.io/confluence-cloud/content/versions#restore-content-version
func (v *VersionService) Restore(ctx context.Context, contentID string, payload *model.ContentRestorePayloadScheme, expand []string) (*model.ContentVersionScheme, *model.ResponseScheme, error) {
	return v.internalClient.Restore(ctx, contentID, payload, expand)
}

// Delete deletes a historical version.
//
// # This does not delete the changes made to the content in that version, rather the changes for the deleted version
//
// are rolled up into the next version. Note, you cannot delete the current version.
//
// DELETE /wiki/rest/api/content/{id}/version/{versionNumber}
//
// https://docs.go-atlassian.io/confluence-cloud/content/versions#delete-content-version
func (v *VersionService) Delete(ctx context.Context, contentID string, versionNumber int) (*model.ResponseScheme, error) {
	return v.internalClient.Delete(ctx, contentID, versionNumber)
}

type internalVersionImpl struct {
	c service.Connector
}

func (i *internalVersionImpl) Gets(ctx context.Context, contentID string, expand []string, start, limit int) (*model.ContentVersionPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/version?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentVersionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalVersionImpl) Get(ctx context.Context, contentID string, versionNumber int, expand []string) (*model.ContentVersionScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/version/%v", contentID, versionNumber))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	version := new(model.ContentVersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	return version, response, nil
}

func (i *internalVersionImpl) Restore(ctx context.Context, contentID string, payload *model.ContentRestorePayloadScheme, expand []string) (*model.ContentVersionScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/version", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	version := new(model.ContentVersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	return version, response, nil
}

func (i *internalVersionImpl) Delete(ctx context.Context, contentID string, versionNumber int) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/version/%v", contentID, versionNumber)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
