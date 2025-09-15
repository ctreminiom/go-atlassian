package models

import (
	"errors"
)

var (
	// ErrNoSCIMComplexValue indicates that a required SCIM complex value was not provided
	ErrNoSCIMComplexValue = errors.New("no scim complex value set")

	// ErrNoSCIMValue indicates that a required SCIM value was not provided
	ErrNoSCIMValue = errors.New("no scim value set")

	// ErrNoSCIMPath indicates that a required SCIM path was not provided
	ErrNoSCIMPath = errors.New("no scim path set")

	// ErrNoSCIMOperation indicates that a required SCIM operation value was not provided
	ErrNoSCIMOperation = errors.New("no scim operation value set")

	// ErrNoAdminOrganization indicates that a required organization ID was not provided
	ErrNoAdminOrganization = errors.New("no organization id set")

	// ErrNoAdminDomainID indicates that a required domain ID was not provided
	ErrNoAdminDomainID = errors.New("no domain id set")

	// ErrNoEventID indicates that a required event ID was not provided
	ErrNoEventID = errors.New("no event id set")

	// ErrNoAdminPolicy indicates that a required organization policy ID was not provided
	ErrNoAdminPolicy = errors.New("no organization policy id set")

	// ErrNoAdminDirectoryID indicates that a required directory ID was not provided
	ErrNoAdminDirectoryID = errors.New("no directory id set")

	// ErrNoAdminGroupID indicates that a required group ID was not provided
	ErrNoAdminGroupID = errors.New("no group id set")

	// ErrNoAdminGroupName indicates that a required group name was not provided
	ErrNoAdminGroupName = errors.New("no group name set")

	// ErrNoAdminUserID indicates that a required user ID was not provided
	ErrNoAdminUserID = errors.New("no user id set")

	// ErrNoAdminAccountID indicates that a required account ID was not provided
	ErrNoAdminAccountID = errors.New("no account id set")

	// ErrNoAdminUserToken indicates that a required user token ID was not provided
	ErrNoAdminUserToken = errors.New("no user token id set")

	// ErrNoBoardID indicates that a required board ID was not provided
	ErrNoBoardID = errors.New("no board id set")

	// ErrNoFilterID indicates that a required filter ID was not provided
	ErrNoFilterID = errors.New("no filter id set")

	// ErrNoEpicID indicates that a required epic ID was not provided
	ErrNoEpicID = errors.New("no epic id set")

	// ErrNoSprintID indicates that a required sprint ID was not provided
	ErrNoSprintID = errors.New("no sprint id set")

	// ErrNoWorkspaceID indicates that a required workspace ID was not provided
	ErrNoWorkspaceID = errors.New("no workspace id set")

	// ErrNoAqlQuery indicates that a required AQL query was not provided
	ErrNoAqlQuery = errors.New("no aql query id set")

	// ErrNoIconID indicates that a required icon ID was not provided
	ErrNoIconID = errors.New("no icon id set")

	// ErrNoObjectID indicates that a required object ID was not provided
	ErrNoObjectID = errors.New("no object id set")

	// ErrNoObjectSchemaID indicates that a required object schema ID was not provided
	ErrNoObjectSchemaID = errors.New("no object schema id set")

	// ErrNoObjectTypeID indicates that a required object type ID was not provided
	ErrNoObjectTypeID = errors.New("no object type id set")

	// ErrNoObjectTypeAttributeID indicates that a required object type attribute ID was not provided
	ErrNoObjectTypeAttributeID = errors.New("no object type attribute id set")

	// ErrNoTaskID indicates that a required task ID was not provided
	ErrNoTaskID = errors.New("no task id set")

	// ErrNoWorkspace indicates that a required workspace was not provided
	ErrNoWorkspace = errors.New("no workspace set")

	// ErrNoMemberID indicates that a required member ID was not provided
	ErrNoMemberID = errors.New("no member id set")

	// ErrNoWebhookID indicates that a required webhook ID was not provided
	ErrNoWebhookID = errors.New("no webhook id set")

	// ErrNoRepository indicates that a required repository was not provided
	ErrNoRepository = errors.New("no repository set")

	// ErrNoVersionProvided indicates that a required module version was not provided
	ErrNoVersionProvided = errors.New("no module version set")

	// ErrInvalidStatusCode indicates that the HTTP response status code was invalid
	ErrInvalidStatusCode = errors.New("invalid http response status, please refer the response.body for more details")

	// ErrNotFound indicates that the requested Atlassian resource was not found
	ErrNotFound = errors.New("no atlassian resource found")

	// ErrUnauthorized indicates insufficient permissions for the requested operation
	ErrUnauthorized = errors.New("atlassian insufficient permissions")

	// ErrInternal indicates an internal Atlassian error occurred
	ErrInternal = errors.New("atlassian internal error")

	// ErrBadRequest indicates that the request payload was invalid
	ErrBadRequest = errors.New("atlassian invalid payload")

	// ErrNoSite indicates that no Atlassian site URL was provided
	ErrNoSite = errors.New("no atlassian site set")

	// ErrNoContentAttachmentID indicates that a required attachment ID was not provided
	ErrNoContentAttachmentID = errors.New("no attachment id set")

	// ErrNoContentAttachmentName indicates that a required attachment filename was not provided
	ErrNoContentAttachmentName = errors.New("no attachment filename set")

	// ErrNoContentReader indicates that a required content reader was not provided
	ErrNoContentReader = errors.New("no reader set")

	// ErrNoContentID indicates that a required content ID was not provided
	ErrNoContentID = errors.New("no content id set")

	// ErrNoCustomContentType indicates that a required custom content type was not provided
	ErrNoCustomContentType = errors.New("no custom content type set")

	// ErrNoCustomContentID indicates that a required custom content ID was not provided
	ErrNoCustomContentID = errors.New("no custom content id set")

	// ErrNoPageID indicates that a required page ID was not provided
	ErrNoPageID = errors.New("no page id set")

	// ErrNoSpaceID indicates that a required space ID was not provided
	ErrNoSpaceID = errors.New("no space id set")

	// ErrNoTargetID indicates that a required target ID was not provided
	ErrNoTargetID = errors.New("no target id set")

	// ErrNoPosition indicates that a required position value was not provided
	ErrNoPosition = errors.New("no position set")

	// ErrInvalidPosition indicates that the provided position value is invalid
	ErrInvalidPosition = errors.New("invalid position: (before, after, append)")

	// ValidPositions defines the valid position values for content positioning
	ValidPositions = map[string]bool{"before": true, "after": true, "append": true}

	// ErrNoLabelID indicates that a required label ID was not provided
	ErrNoLabelID = errors.New("no label id set")

	// ErrNoCQL indicates that a required CQL query was not provided
	ErrNoCQL = errors.New("no CQL query set")

	// ErrNoContentType indicates that a required content type was not provided
	ErrNoContentType = errors.New("no content type set")

	// ErrNoEntityID indicates that a required entity ID was not provided
	ErrNoEntityID = errors.New("no entity id set")

	// ValidEntityValues defines the valid entity types
	ValidEntityValues = []string{"blogposts", "custom-content", "labels", "pages"}

	// ErrNoEntityValue indicates that no valid entity value was provided
	ErrNoEntityValue = errors.New("no valid entity id set")

	// ErrNoContentLabel indicates that a required content label was not provided
	ErrNoContentLabel = errors.New("no content label set")

	// ErrNoContentProperty indicates that a required content property was not provided
	ErrNoContentProperty = errors.New("no content property set")

	// ErrNoSpaceName indicates that a required space name was not provided
	ErrNoSpaceName = errors.New("no space name set")

	// ErrNoSpaceKey indicates that a required space key was not provided
	ErrNoSpaceKey = errors.New("no space key set")

	// ErrNoContentRestrictionKey indicates that a required content restriction operation key was not provided
	ErrNoContentRestrictionKey = errors.New("no content restriction operation key set")

	// ErrNoConfluenceGroup indicates that neither group ID nor name was provided
	ErrNoConfluenceGroup = errors.New("no group id or name set")

	// ErrNoLabelName indicates that a required label name was not provided
	ErrNoLabelName = errors.New("no label name set")

	// Custom-field errors

	// ErrNoFloatType indicates that a required float type was not provided
	ErrNoFloatType = errors.New("no float type set")

	// ErrNoSprintType indicates that a required sprint type was not found
	ErrNoSprintType = errors.New("no sprint type found")

	// ErrNoMultiVersionType indicates that a required multi-version type was not found
	ErrNoMultiVersionType = errors.New("no multiversion type found")

	// ErrNoFieldInformation indicates that a valid fields object was not provided
	ErrNoFieldInformation = errors.New("please provide a buffer with a valid fields object")

	// ErrNoMultiSelectType indicates that a required multi-select type was not found
	ErrNoMultiSelectType = errors.New("no multiselect type found")

	// ErrNoAssetType indicates that a required asset type was not found
	ErrNoAssetType = errors.New("no asset type found")

	// ErrNoURLType indicates that a required URL type was not provided
	ErrNoURLType = errors.New("no url type set")

	// ErrNoTextType indicates that a required text type was not provided
	ErrNoTextType = errors.New("no text type set")

	// ErrNoDatePickerType indicates that a required date picker type was not provided
	ErrNoDatePickerType = errors.New("no datepicker type set")

	// ErrNoDateTimeType indicates that a required datetime type was not provided
	ErrNoDateTimeType = errors.New("no datetime type set")

	// ErrNoSelectType indicates that a required select type was not provided
	ErrNoSelectType = errors.New("no select type set")

	// ErrNoButtonType indicates that a required button type was not provided
	ErrNoButtonType = errors.New("no button type set")

	// ErrNoUserType indicates that a required user type was not provided
	ErrNoUserType = errors.New("no user type set")

	// ErrNoLabelsType indicates that a required labels type was not provided
	ErrNoLabelsType = errors.New("no labels type set")

	// ErrNoMultiUserType indicates that a required multi-user type was not provided
	ErrNoMultiUserType = errors.New("no multi-user type set")

	// ErrNoCheckBoxType indicates that a required checkbox type was not provided
	ErrNoCheckBoxType = errors.New("no check-box type set")

	// ErrNoCascadingParent indicates that a required cascading parent value was not provided
	ErrNoCascadingParent = errors.New("no cascading parent value set")

	// ErrNoCascadingChild indicates that a required cascading child value was not provided
	ErrNoCascadingChild = errors.New("no cascading child value set")

	// ErrNoValueType indicates that a required value was not provided
	ErrNoValueType = errors.New("no value set")

	// ErrNoRequestType indicates that a required request type value was not provided
	ErrNoRequestType = errors.New("no request type value set")

	// ErrNoTempoAccountType indicates that a required tempo account value was not provided
	ErrNoTempoAccountType = errors.New("no tempo account value set")

	// ErrNoNotificationSchemeID indicates that a required notification scheme ID was not provided
	ErrNoNotificationSchemeID = errors.New("no notification scheme id set")

	// ErrNoNotificationID indicates that a required notification ID was not provided
	ErrNoNotificationID = errors.New("no notification id set")

	// ErrNoApplicationRole indicates that a required application role key was not provided
	ErrNoApplicationRole = errors.New("no application role key set")

	// ErrNoDashboardID indicates that a required dashboard ID was not provided
	ErrNoDashboardID = errors.New("no dashboard id set")

	// ErrNoGroupName indicates that a required group name was not provided
	ErrNoGroupName = errors.New("no group name set")

	// ErrNoGroupsName indicates that required group names were not provided
	ErrNoGroupsName = errors.New("no groups names set")

	// ErrNoIssueKeyOrID indicates that neither issue key nor ID was provided
	ErrNoIssueKeyOrID = errors.New("no issue key/id set")

	// ErrNoRemoteLinkID indicates that a required remote link ID was not provided
	ErrNoRemoteLinkID = errors.New("no remote link id set")

	// ErrNoRemoteLinkGlobalID indicates that a required global remote link ID was not provided
	ErrNoRemoteLinkGlobalID = errors.New("no global remote link id set")

	// ErrNoTransitionID indicates that a required transition ID was not provided
	ErrNoTransitionID = errors.New("no transition id set")

	// ErrNoAttachmentID indicates that a required attachment ID was not provided
	ErrNoAttachmentID = errors.New("no attachment id set")

	// ErrNoAttachmentName indicates that a required attachment filename was not provided
	ErrNoAttachmentName = errors.New("no attachment filename set")

	// ErrNoReader indicates that a required reader was not provided
	ErrNoReader = errors.New("no reader set")

	// ErrNoCommentID indicates that a required comment ID was not provided
	ErrNoCommentID = errors.New("no comment id set")

	// ErrNoProjectID indicates that a required project ID was not provided
	ErrNoProjectID = errors.New("no project id set")

	// ErrNoProjectIDOrKey indicates that neither project ID nor key was provided
	ErrNoProjectIDOrKey = errors.New("no project id or key set")

	// ErrNoProjectRoleID indicates that a required project role ID was not provided
	ErrNoProjectRoleID = errors.New("no project role id set")

	// ErrNoProjectCategoryID indicates that a required project category ID was not provided
	ErrNoProjectCategoryID = errors.New("no project category id set")

	// ErrNoPropertyKey indicates that a required property key was not provided
	ErrNoPropertyKey = errors.New("no property key set")

	// ErrNoProjectFeatureKey indicates that a required project feature key was not provided
	ErrNoProjectFeatureKey = errors.New("no project feature key set")

	// ErrNoProjectFeatureState indicates that a required project state key was not provided
	ErrNoProjectFeatureState = errors.New("no project state key set")

	// ErrNoFieldID indicates that a required field ID was not provided
	ErrNoFieldID = errors.New("no field id set")

	// ErrInvalidCustomFieldUpdate represents an error indicating the custom field update payload contains an invalid type attribute.
	ErrInvalidCustomFieldUpdate = errors.New("invalid custom field update payload, type is not a valid attribute for update")

	// ErrNoEditOperator indicates that a required update operation was not provided
	ErrNoEditOperator = errors.New("no update operation set")

	// ErrNoEditValue indicates that a required update operation value was not provided
	ErrNoEditValue = errors.New("no update operation value set")

	// ErrNoCustomFieldID indicates that a required custom field ID was not provided
	ErrNoCustomFieldID = errors.New("no custom-field id set")

	// ErrNoWorkflowStatuses indicates that required workflow statuses were not provided
	ErrNoWorkflowStatuses = errors.New("no workflow statuses set")

	// ErrNoWorkflowScope indicates that a required workflow scope was not provided
	ErrNoWorkflowScope = errors.New("no workflow scope set")

	// ErrNoWorkflowStatusNameOrID indicates that neither workflow status name nor ID was provided
	ErrNoWorkflowStatusNameOrID = errors.New("no workflow status name or id set")

	// ErrNoFieldContextID indicates that a required field context ID was not provided
	ErrNoFieldContextID = errors.New("no field context id set")

	// ErrNoIssueTypes indicates that required issue type IDs were not provided
	ErrNoIssueTypes = errors.New("no issue types id's set")

	// ErrNoProjects indicates that required projects were not provided
	ErrNoProjects = errors.New("no projects set")

	// ErrNoContextOptionID indicates that a required field context option ID was not provided
	ErrNoContextOptionID = errors.New("no field context option id set")

	// ErrNoTypeID indicates that a required link ID was not provided
	ErrNoTypeID = errors.New("no link id set")

	// ErrNoLinkTypeID indicates that a required link type ID was not provided
	ErrNoLinkTypeID = errors.New("no link type id set")

	// ErrNoPriorityID indicates that a required priority ID was not provided
	ErrNoPriorityID = errors.New("no priority id set")

	// ErrNoResolutionID indicates that a required resolution ID was not provided
	ErrNoResolutionID = errors.New("no resolution id set")

	// ErrNoJQL indicates that a required JQL query was not provided
	ErrNoJQL = errors.New("no sql set")

	// ErrNoIssueTypeID indicates that a required issue type ID was not provided
	ErrNoIssueTypeID = errors.New("no issue type id set")

	// ErrNoIssueTypeScreenSchemeID indicates that a required issue type screen scheme ID was not provided
	ErrNoIssueTypeScreenSchemeID = errors.New("no issue type screen scheme id set")

	// ErrNoScreenSchemeID indicates that a required screen scheme ID was not provided
	ErrNoScreenSchemeID = errors.New("no screen scheme id set")

	// ErrNoAccountID indicates that a required account ID was not provided
	ErrNoAccountID = errors.New("no account id set")

	// ErrNoWorklogID indicates that a required worklog ID was not provided
	ErrNoWorklogID = errors.New("no worklog id set")

	// ErrNpWorklogs indicates that required worklog IDs were not provided
	ErrNpWorklogs = errors.New("no worklog's id set")

	// ErrNoPermissionSchemeID indicates that a required permission scheme ID was not provided
	ErrNoPermissionSchemeID = errors.New("no permission scheme id set")

	// ErrNoPermissionGrantID indicates that a required permission grant ID was not provided
	ErrNoPermissionGrantID = errors.New("no permission grant id set")

	// ErrNoPermissionKeys indicates that required permission keys were not provided
	ErrNoPermissionKeys = errors.New("no permission keys set")

	// ErrNoComponentID indicates that a required component ID was not provided
	ErrNoComponentID = errors.New("no component id set")

	// ErrProjectTypeKey indicates that a required project type key was not provided
	ErrProjectTypeKey = errors.New("no project type key set")

	// ErrNoProjectName indicates that a required project name was not provided
	ErrNoProjectName = errors.New("no project name set")

	// ErrNoVersionID indicates that a required version ID was not provided
	ErrNoVersionID = errors.New("no version id set")

	// ErrNoScreenName indicates that a required screen name was not provided
	ErrNoScreenName = errors.New("no screen name set")

	// ErrNoScreenTabName indicates that a required screen tab name was not provided
	ErrNoScreenTabName = errors.New("no screen tab name set")

	// ErrNoAccountSlice indicates that required account IDs were not provided
	ErrNoAccountSlice = errors.New("no account id's set")

	// ErrNoProjectKeySlice indicates that required project keys were not provided
	ErrNoProjectKeySlice = errors.New("no project key's set")

	// ErrNoProjectIDs indicates that required project IDs were not provided
	ErrNoProjectIDs = errors.New("no project id's set")

	// ErrNoWorkflowID indicates that a required workflow ID was not provided
	ErrNoWorkflowID = errors.New("no workflow id set")

	// ErrNoWorkflowSchemeID indicates that a required workflow scheme ID was not provided
	ErrNoWorkflowSchemeID = errors.New("no workflow scheme id set")

	// ErrNoScreenID indicates that a required screen ID was not provided
	ErrNoScreenID = errors.New("no screen id set")

	// ErrNoScreenTabID indicates that a required screen tab ID was not provided
	ErrNoScreenTabID = errors.New("no screen tab id set")

	// ErrNoFieldConfigurationName indicates that a required field configuration name was not provided
	ErrNoFieldConfigurationName = errors.New("no field configuration name set")

	// ErrNoFieldConfigurationID indicates that a required field configuration ID was not provided
	ErrNoFieldConfigurationID = errors.New("no field configuration id set")

	// ErrNoFieldConfigurationSchemeName indicates that a required field configuration scheme name was not provided
	ErrNoFieldConfigurationSchemeName = errors.New("no field configuration scheme name set")

	// ErrNoFieldConfigurationSchemeID indicates that a required field configuration scheme ID was not provided
	ErrNoFieldConfigurationSchemeID = errors.New("no field configuration scheme id set")

	// ErrNoQuery indicates that a required query was not provided
	ErrNoQuery = errors.New("no query set")

	// ErrNoIssueTypeSchemeID indicates that a required issue type scheme ID was not provided
	ErrNoIssueTypeSchemeID = errors.New("no issue type scheme id set")

	// ErrNoApprovalID indicates that a required approval ID was not provided
	ErrNoApprovalID = errors.New("no approval id set")

	// ErrNoKeyError indicates that a required key was not provided
	ErrNoKeyError = errors.New("no key set")

	// ErrNoCreateIssues indicates that required issues payload was not provided
	ErrNoCreateIssues = errors.New("no issues payload set")

	// ErrNoIssueScheme indicates that a required issue instance was not provided
	ErrNoIssueScheme = errors.New("no issue instance set")

	// ErrNoMapValues indicates that required map values were not provided
	ErrNoMapValues = errors.New("no map values set")

	// ErrNoIssuesSlice indicates that required issues object was not provided
	ErrNoIssuesSlice = errors.New("no issues object set")

	// ErrNoKBQuery indicates that a required knowledge base query was not provided
	ErrNoKBQuery = errors.New("no knowledge base query set")

	// ErrNoOrganizationName indicates that a required organization name was not provided
	ErrNoOrganizationName = errors.New("no organization name set")

	// ErrNoOrganizationID indicates that a required organization ID was not provided
	ErrNoOrganizationID = errors.New("no organization id set")

	// ErrNoServiceDeskID indicates that a required service desk ID was not provided
	ErrNoServiceDeskID = errors.New("no service desk id set")

	// ErrNoQueueID indicates that a required service desk queue ID was not provided
	ErrNoQueueID = errors.New("no service desk queue id set")

	// ErrNoRequestTypeID indicates that a required request type ID was not provided
	ErrNoRequestTypeID = errors.New("no request type id set")

	// ErrNoFileName indicates that a required file name was not provided
	ErrNoFileName = errors.New("no file name set")

	// ErrNoFileReader indicates that a required io.Reader was not provided
	ErrNoFileReader = errors.New("no io.Reader set")

	// ErrNoSLAMetricID indicates that a required SLA metric ID was not provided
	ErrNoSLAMetricID = errors.New("no sla metric id set")

	// ErrNoComponents indicates that required components were not provided
	ErrNoComponents = errors.New("no components set")

	// ErrNCoComponent indicates that a required component was not provided
	ErrNCoComponent = errors.New("no component set")

	// ErrNoCommentBody indicates that a required comment body was not provided
	ErrNoCommentBody = errors.New("no comment body set")

	// ErrCreateHttpReq represents an error indicating the failure to create an HTTP request. Used in unit tests
	ErrCreateHttpReq = errors.New("error, unable to create the http request")

	// ErrReqFailed represents an error indicating that a request has failed.
	ErrReqFailed = errors.New("error, request failed")

	// ErrNoExecHttpCall represents an error indicating the inability to execute the HTTP call.
	ErrNoExecHttpCall = errors.New("error, unable to execute the http call")

	// ErrNoHttpResponse indicates that no HTTP response was found in the client request
	ErrNoHttpResponse = errors.New("client: no http response found")

	// ErrNoAtlConnect indicates a failure to establish a connection with the Atlassian instance.
	ErrNoAtlConnect = errors.New("error, unable to connect with the Atlassian instance")

	// ErrApiExec represents an error indicating a failure to execute an API call.
	ErrApiExec = errors.New("error, unable to execute API call")

	// ErrHttpTransition indicates a failure or inability to execute the HTTP transition operation.
	ErrHttpTransition = errors.New("error, unable to execute the http transition")

	// ErrNoIssueTypeReorderAttr signifies that neither position nor after attribute is set for issue type scheme reorder.
	ErrNoIssueTypeReorderAttr = errors.New("no position or after attribute set for issue type scheme reorder. one must be set")

	// ErrInvalidIssueTypeSchemePosition indicates an invalid position value for the issue type scheme, which must be "First" or "Last".
	ErrInvalidIssueTypeSchemePosition = errors.New("invalid issue type scheme position. must be one of the following values: First, Last")

	// ErrInvalidIssueTypeSchemeAfter represents an error indicating an invalid 'after' attribute in the issue type scheme configuration.
	ErrInvalidIssueTypeSchemeAfter = errors.New("issue type scheme invalid 'after' attr, issue type id found in 'issueTypeIds'")

	// ErrNoFolderID indicates that a required folder ID was not provided
	ErrNoFolderID = errors.New("confluence: no folder id set")

	// ErrNoWhiteboardID indicates that a required whiteboard ID was not provided
	ErrNoWhiteboardID = errors.New("confluence: no whiteboard id set")

	// ErrNoDatabaseID indicates that a required database ID was not provided
	ErrNoDatabaseID = errors.New("confluence: no database id set")

	// ErrNoEmbedID indicates that a required embed ID was not provided
	ErrNoEmbedID = errors.New("confluence: no embed id set")
)
