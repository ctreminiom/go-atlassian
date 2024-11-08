package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// AnnouncementBannerConnector resource represents the Jira announcement banner.
// Use it to retrieve and update banner configuration.
type AnnouncementBannerConnector interface {

	// Get returns the current announcement banner configuration.
	//
	// GET /rest/api/{2-3}/announcementBanner
	//
	// https://docs.go-atlassian.io/jira-software-cloud/announcement-banner#get-announcement-banner-configuration
	Get(ctx context.Context) (*models.AnnouncementBannerScheme, *models.ResponseScheme, error)

	// Update updates the announcement banner configuration.
	//
	// PUT /rest/api/{2-3}/announcementBanner
	//
	// https://docs.go-atlassian.io/jira-software-cloud/announcement-banner#get-announcement-banner-configuration
	Update(ctx context.Context, payload *models.AnnouncementBannerPayloadScheme) (*models.ResponseScheme, error)
}
