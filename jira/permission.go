package jira

import (
	"context"
	"encoding/json"
	"net/http"
)

type PermissionService struct {
	client *Client
	Scheme *PermissionSchemeService
}

type GlobalPermissionsScheme struct {
	Permissions struct {
		ADDCOMMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ADD_COMMENTS"`
		ADMINISTER struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ADMINISTER"`
		ADMINISTERPROJECTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ADMINISTER_PROJECTS"`
		ASSIGNABLEUSER struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ASSIGNABLE_USER"`
		ASSIGNISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ASSIGN_ISSUES"`
		BROWSEPROJECTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"BROWSE_PROJECTS"`
		BULKCHANGE struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"BULK_CHANGE"`
		CLOSEISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CLOSE_ISSUES"`
		CREATEATTACHMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_ATTACHMENTS"`
		CREATEISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_ISSUES"`
		CREATEPROJECT struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_PROJECT"`
		CREATESHAREDOBJECTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_SHARED_OBJECTS"`
		DELETEALLATTACHMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_ALL_ATTACHMENTS"`
		DELETEALLCOMMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_ALL_COMMENTS"`
		DELETEALLWORKLOGS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_ALL_WORKLOGS"`
		DELETEISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_ISSUES"`
		DELETEOWNATTACHMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_OWN_ATTACHMENTS"`
		DELETEOWNCOMMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_OWN_COMMENTS"`
		DELETEOWNWORKLOGS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"DELETE_OWN_WORKLOGS"`
		EDITALLCOMMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"EDIT_ALL_COMMENTS"`
		EDITALLWORKLOGS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"EDIT_ALL_WORKLOGS"`
		EDITISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"EDIT_ISSUES"`
		EDITOWNCOMMENTS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"EDIT_OWN_COMMENTS"`
		EDITOWNWORKLOGS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"EDIT_OWN_WORKLOGS"`
		LINKISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"LINK_ISSUES"`
		MANAGEGROUPFILTERSUBSCRIPTIONS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"MANAGE_GROUP_FILTER_SUBSCRIPTIONS"`
		MANAGESPRINTSPERMISSION struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"MANAGE_SPRINTS_PERMISSION"`
		MANAGEWATCHERS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"MANAGE_WATCHERS"`
		MODIFYREPORTER struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"MODIFY_REPORTER"`
		MOVEISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"MOVE_ISSUES"`
		RESOLVEISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"RESOLVE_ISSUES"`
		SCHEDULEISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"SCHEDULE_ISSUES"`
		SERVICEDESKAGENT struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"SERVICEDESK_AGENT"`
		SETISSUESECURITY struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"SET_ISSUE_SECURITY"`
		SYSTEMADMIN struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"SYSTEM_ADMIN"`
		TRANSITIONISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"TRANSITION_ISSUES"`
		USERPICKER struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"USER_PICKER"`
		VIEWDEVTOOLS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"VIEW_DEV_TOOLS"`
		VIEWREADONLYWORKFLOW struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"VIEW_READONLY_WORKFLOW"`
		VIEWVOTERSANDWATCHERS struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"VIEW_VOTERS_AND_WATCHERS"`
		WORKONISSUES struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"WORK_ON_ISSUES"`
	} `json:"permissions"`
}

func (p *PermissionService) Gets(ctx context.Context) (result *GlobalPermissionsScheme, response *Response, err error) {

	var endpoint = "rest/api/3/permissions"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(GlobalPermissionsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
