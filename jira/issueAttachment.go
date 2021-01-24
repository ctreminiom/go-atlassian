package jira

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type AttachmentService struct{ client *Client }

type AttachmentSettingScheme struct {
	Enabled     bool `json:"enabled"`
	UploadLimit int  `json:"uploadLimit"`
}

// Returns the attachment settings, that is, whether attachments are enabled and the maximum attachment size allowed.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-meta-get
func (a *AttachmentService) Settings(ctx context.Context) (result *AttachmentSettingScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	var endpoint = "rest/api/3/attachment/meta"
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(AttachmentSettingScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type AttachmentMetadataScheme struct {
	ID       int    `json:"id"`
	Self     string `json:"self"`
	Filename string `json:"filename"`
	Author   struct {
		Self       string `json:"self"`
		Key        string `json:"key"`
		AccountID  string `json:"accountId"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"author"`
	Created   string `json:"created"`
	Size      int    `json:"size"`
	MimeType  string `json:"mimeType"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
}

// Returns the metadata for an attachment. Note that the attachment itself is not returned.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-id-get
func (a *AttachmentService) Metadata(ctx context.Context, attachmentID string) (result *AttachmentMetadataScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	var endpoint = fmt.Sprintf("rest/api/3/attachment/%v", attachmentID)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(AttachmentMetadataScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes an attachment from an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-id-delete
func (a *AttachmentService) Delete(ctx context.Context, attachmentID string) (response *Response, err error) {

	if ctx == nil {
		return nil, errors.New("the context param is nil, please provide a valid one")
	}

	var endpoint = fmt.Sprintf("rest/api/3/attachment/%v", attachmentID)
	request, err := a.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	return
}

type AttachmentHumanMetadataScheme struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Entries []struct {
		Path      string `json:"path"`
		Index     int    `json:"index"`
		Size      string `json:"size"`
		MediaType string `json:"mediaType"`
		Label     string `json:"label"`
	} `json:"entries"`
	TotalEntryCount int    `json:"totalEntryCount"`
	MediaType       string `json:"mediaType"`
}

// Returns the metadata for the contents of an attachment, if it is an archive, and metadata for the attachment itself.
// For example, if the attachment is a ZIP archive, then information about the files in the archive is returned and metadata for the ZIP archive.
// Currently, only the ZIP archive format is supported.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-attachment-id-expand-human-get
func (a *AttachmentService) Human(ctx context.Context, attachmentID string) (result *AttachmentHumanMetadataScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	var endpoint = fmt.Sprintf("rest/api/3/attachment/%v/expand/human", attachmentID)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(AttachmentHumanMetadataScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type AttachmentScheme struct {
	Self     string `json:"self"`
	ID       string `json:"id,omitempty"`
	Filename string `json:"filename"`
	Author   struct {
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
	} `json:"author"`
	Created   string `json:"created"`
	Size      int    `json:"size"`
	MimeType  string `json:"mimeType"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

// Adds one or more attachments to an issue. Attachments are posted as multipart/form-data (RFC 1867).
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-attachments/#api-rest-api-3-issue-issueidorkey-attachments-post
func (a *AttachmentService) Add(issueKeyOrID string, path string) (result *[]AttachmentScheme, response *Response, err error) {

	if !filepath.IsAbs(path) {
		return nil, nil, fmt.Errorf("the path provided is not an absolute path, please provide a valid one")
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	filePart, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return
	}

	_, err = io.Copy(filePart, file)
	if err != nil {
		return
	}

	if err = writer.Close(); err != nil {
		return
	}

	var endpoint = fmt.Sprintf("%v/rest/api/3/issue/%v/attachments", a.client.Site.String(), issueKeyOrID)
	request, err := http.NewRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}
	request.SetBasicAuth(a.client.Auth.mail, a.client.Auth.token)

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Accept", "application/json")
	request.Header.Set("X-Atlassian-Token", "no-check")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new([]AttachmentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
