package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type LabelConnector interface {

	// Get returns label information and a list of contents associated with the label.
	//
	// GET /wiki/rest/api/label
	//
	// https://docs.go-atlassian.io/confluence-cloud/label#get-label-information
	Get(ctx context.Context, labelName, labelType string, start, limit int) (*model.LabelDetailsScheme, *model.ResponseScheme, error)
}

type LabelsConnector interface {

	// Gets returns the labels on a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/label
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/labels#get-labels-for-content
	Gets(ctx context.Context, contentID, prefix string, startAt, maxResults int) (*model.ContentLabelPageScheme, *model.ResponseScheme, error)

	// Add adds labels to a piece of content. Does not modify the existing labels.
	//
	// POST /wiki/rest/api/content/{id}/label
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/labels#add-labels-to-content
	Add(ctx context.Context, contentID string, payload []*model.ContentLabelPayloadScheme, want400Response bool) (*model.ContentLabelPageScheme, *model.ResponseScheme, error)

	// Remove removes a label from a piece of content
	//
	// DELETE /wiki/rest/api/content/{id}/label/{label}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/labels#remove-label-from-content
	Remove(ctx context.Context, contentID, labelName string) (*model.ResponseScheme, error)
}
