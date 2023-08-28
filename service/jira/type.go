package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type TypeConnector interface {

	// Gets returns all issue types.
	//
	// GET /rest/api/{2-3}/issuetype
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-all-issue-types-for-user
	Gets(ctx context.Context) ([]*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Create creates an issue type and adds it to the default issue type scheme.
	//
	// POST /rest/api/{2-3}/issuetype
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#create-issue-type
	Create(ctx context.Context, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Get returns an issue type.
	//
	// GET /rest/api/{2-3}/issuetype/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-issue-type
	Get(ctx context.Context, issueTypeId string) (*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Update updates the issue type.
	//
	// PUT /rest/api/{2-3}/issuetype/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#update-issue-type
	Update(ctx context.Context, issueTypeId string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Delete deletes the issue type.
	//
	// If the issue type is in use, all uses are updated with the alternative issue type (alternativeIssueTypeId).
	// A list of alternative issue types are obtained from the Get alternative issue types resource.
	//
	// DELETE /rest/api/{2-3}/issuetype/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#delete-issue-type
	Delete(ctx context.Context, issueTypeId string) (*model.ResponseScheme, error)

	// Alternatives returns a list of issue types that can be used to replace the issue type.
	//
	// The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.
	//
	// GET /rest/api/{2-3}/issuetype/{id}/alternatives
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-alternative-issue-types
	Alternatives(ctx context.Context, issueTypeId string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error)
}

type TypeSchemeConnector interface {

	// Gets returns a paginated list of issue type schemes.
	//
	// GET /rest/api/{2-3}/issuetypescheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-all-issue-type-schemes
	Gets(ctx context.Context, issueTypeSchemeIds []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error)

	// Create creates an issue type scheme.
	//
	// POST /rest/api/{2-3}/issuetypescheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#create-issue-type-scheme
	Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error)

	// Items returns a paginated list of issue type scheme items.
	//
	// GET /rest/api/{2-3}/issuetypescheme/mapping
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-scheme-items
	Items(ctx context.Context, issueTypeSchemeIds []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error)

	// Projects returns a paginated list of issue type schemes and, for each issue type scheme, a list of the projects that use it.
	//
	// GET /rest/api/{2-3}/issuetypescheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-schemes-for-projects
	Projects(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error)

	// Assign assigns an issue type scheme to a project.
	//
	// PUT /rest/api/{2-3}/issuetypescheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#assign-issue-type-scheme-to-project
	Assign(ctx context.Context, issueTypeSchemeId, projectId string) (*model.ResponseScheme, error)

	// Update updates an issue type scheme.
	//
	// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#update-issue-type-scheme
	Update(ctx context.Context, issueTypeSchemeId int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error)

	// Delete deletes an issue type scheme.
	//
	// 1.Only issue type schemes used in classic projects can be deleted.
	//
	// 2.Any projects assigned to the scheme are reassigned to the default issue type scheme.
	//
	// DELETE /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#delete-issue-type-scheme
	Delete(ctx context.Context, issueTypeSchemeId int) (*model.ResponseScheme, error)

	// Append adds issue types to an issue type scheme.
	//
	// 1.The added issue types are appended to the issue types list.
	//
	// 2.If any of the issue types exist in the issue type scheme, the operation fails and no issue types are added.
	//
	// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeId}/issuetype
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#add-issue-types-to-issue-type-scheme
	Append(ctx context.Context, issueTypeSchemeId int, issueTypeIds []int) (*model.ResponseScheme, error)

	// Remove removes an issue type from an issue type scheme, this operation cannot remove:
	//
	// 1.any issue type used by issues.
	//
	// 2.any issue types from the default issue type scheme.
	//
	// 3.the last standard issue type from an issue type scheme.
	//
	// DELETE /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeId}/issuetype/{issueTypeId}
	//
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#remove-issue-type-from-issue-type-scheme
	Remove(ctx context.Context, issueTypeSchemeId, issueTypeId int) (*model.ResponseScheme, error)
}

type TypeScreenSchemeConnector interface {

	// Gets returns a paginated list of issue type screen schemes.
	//
	// Only issue type screen schemes used in classic projects are returned.
	//
	// GET /rest/api/{2-3}/issuetypescreenscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-schemes
	Gets(ctx context.Context, options *model.ScreenSchemeParamsScheme, startAt, maxResults int) (*model.IssueTypeScreenSchemePageScheme, *model.ResponseScheme, error)

	// Create creates an issue type screen scheme.
	//
	// POST /rest/api/{2-3}/issuetypescreenscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#create-issue-type-screen-scheme
	Create(ctx context.Context, payload *model.IssueTypeScreenSchemePayloadScheme) (*model.IssueTypeScreenScreenCreatedScheme, *model.ResponseScheme, error)

	// Assign assigns an issue type screen scheme to a project.
	//
	// Issue type screen schemes can only be assigned to classic projects.
	//
	// PUT /rest/api/{2-3}/issuetypescreenscheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#assign-issue-type-screen-scheme-to-project
	Assign(ctx context.Context, issueTypeScreenSchemeId, projectId string) (*model.ResponseScheme, error)

	// Projects returns a paginated list of issue type screen schemes and,
	// for each issue type screen scheme, a list of the projects that use it.
	//
	// GET /rest/api/{2-3}/issuetypescreenscheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#assign-issue-type-screen-scheme-to-project
	Projects(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.IssueTypeProjectScreenSchemePageScheme, *model.ResponseScheme, error)

	// Mapping returns a paginated list of issue type screen scheme items.
	//
	// Only issue type screen schemes used in classic projects are returned.
	//
	// GET /rest/api/{2-3}/issuetypescreenscheme/mapping
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-scheme-items
	Mapping(ctx context.Context, issueTypeScreenSchemeIds []int, startAt, maxResults int) (*model.IssueTypeScreenSchemeMappingScheme, *model.ResponseScheme, error)

	// Update updates an issue type screen scheme.
	//
	// PUT /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#update-issue-type-screen-scheme
	Update(ctx context.Context, issueTypeScreenSchemeId, name, description string) (*model.ResponseScheme, error)

	// Delete deletes an issue type screen scheme.
	//
	// DELETE /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#delete-issue-type-screen-scheme
	Delete(ctx context.Context, issueTypeScreenSchemeId string) (*model.ResponseScheme, error)

	// Append appends issue type to screen scheme mappings to an issue type screen scheme.
	//
	// PUT /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#append-mappings-to-issue-type-screen-scheme
	Append(ctx context.Context, issueTypeScreenSchemeId string, payload *model.IssueTypeScreenSchemePayloadScheme) (*model.ResponseScheme, error)

	// UpdateDefault updates the default screen scheme of an issue type screen scheme.
	// The default screen scheme is used for all unmapped issue types.
	//
	// PUT /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/default
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#update-issue-type-screen-scheme-default-screen-scheme
	UpdateDefault(ctx context.Context, issueTypeScreenSchemeId, screenSchemeId string) (*model.ResponseScheme, error)

	// Remove removes issue type to screen scheme mappings from an issue type screen scheme.
	//
	// POST /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/remove
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#remove-mappings-from-issue-type-screen-scheme
	Remove(ctx context.Context, issueTypeScreenSchemeId string, issueTypeIds []string) (*model.ResponseScheme, error)

	// SchemesByProject returns a paginated list of projects associated with an issue type screen scheme.
	//
	// GET /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-scheme-projects
	SchemesByProject(ctx context.Context, issueTypeScreenSchemeId, startAt, maxResults int) (*model.IssueTypeScreenSchemeByProjectPageScheme, *model.ResponseScheme, error)
}
