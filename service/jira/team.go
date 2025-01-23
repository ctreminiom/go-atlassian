package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// TeamConnector is an interface that defines the methods available from Roadmap Teams API.
type TeamConnector interface {

	// Gets gets the Teams information from the Jira Advanced Roadmaps application.
	//
	// Teams in Advanced Roadmaps are different from the teams found in the rest of Jira Software Cloud.
	//
	// In Advanced Roadmaps, they act as a label applied to issues that designates which team will eventually.
	//
	// pick up the work on your timeline. By adding the Team field to your Jira issues.
	//
	// you can save this value back to your Jira issues, which makes sprint planning easier.
	//
	// POST /rest/teams/1.0/teams/find
	//
	// https://docs.go-atlassian.io/jira-software-cloud/teams#get-teams
	Gets(ctx context.Context, maxResults int) (*models.JiraTeamPageScheme, *models.ResponseScheme, error)

	// Create creates a team on the Advanced Roadmaps
	//
	// POST /rest/teams/1.0/teams/create
	//
	// https://docs.go-atlassian.io/jira-software-cloud/teams#create-team
	Create(ctx context.Context, payload *models.JiraTeamCreatePayloadScheme) (*models.JiraTeamCreateResponseScheme, *models.ResponseScheme, error)
}
