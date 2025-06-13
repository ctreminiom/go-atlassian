package internal

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
)

// NewTeamService creates a new instance of TeamService.
func NewTeamService(client service.Connector) *TeamService {

	return &TeamService{
		internalClient: &internalTeamServiceImpl{c: client},
	}
}

// TeamService provides methods to manage team information in Jira Advanced Roadmaps.
type TeamService struct {
	// internalClient is the connector interface for team operations.
	internalClient jira.TeamConnector
}

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
func (t *TeamService) Gets(ctx context.Context, maxResults int) (*model.JiraTeamPageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx, maxResults)
}

// Create creates a team on the Advanced Roadmaps
//
// POST /rest/teams/1.0/teams/create
//
// https://docs.go-atlassian.io/jira-software-cloud/teams#create-team
func (t *TeamService) Create(ctx context.Context, payload *model.JiraTeamCreatePayloadScheme) (*model.JiraTeamCreateResponseScheme, *model.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

type internalTeamServiceImpl struct {
	c service.Connector
}

func (i *internalTeamServiceImpl) Gets(ctx context.Context, maxResults int) (*model.JiraTeamPageScheme, *model.ResponseScheme, error) {

	endpoint := "rest/teams/1.0/teams/find"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"maxResults": maxResults})
	if err != nil {
		return nil, nil, err
	}

	teams := new(model.JiraTeamPageScheme)
	response, err := i.c.Call(request, teams)
	if err != nil {
		return nil, response, err
	}

	return teams, response, nil
}

func (i *internalTeamServiceImpl) Create(ctx context.Context, payload *model.JiraTeamCreatePayloadScheme) (*model.JiraTeamCreateResponseScheme, *model.ResponseScheme, error) {

	endpoint := "rest/teams/1.0/teams/create"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	team := new(model.JiraTeamCreateResponseScheme)
	response, err := i.c.Call(request, team)
	if err != nil {
		return nil, response, err
	}

	return team, response, nil
}
