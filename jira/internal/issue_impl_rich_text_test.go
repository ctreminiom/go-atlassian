package internal

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalRichTextServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx            context.Context
		issueKeyOrID   string
		deleteSubTasks bool
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueKeyOrID:   "DUMMY-1",
				deleteSubTasks: true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/DUMMY-1?deleteSubtasks=true",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueKeyOrID:   "",
				deleteSubTasks: true,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueKeyOrID:   "DUMMY-1",
				deleteSubTasks: true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/DUMMY-1?deleteSubtasks=true",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := issueService.Delete(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.deleteSubTasks)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Assign(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrID, accountID string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				accountID:    "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"/rest/api/2/issue/DUMMY-1/assignee",
					"",
					map[string]interface{}{"accountId": "account-id-sample"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
				accountID:    "account-id-sample",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the account id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				accountID:    "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAccountID,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				accountID:    "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"/rest/api/2/issue/DUMMY-1/assignee",
					"",
					map[string]interface{}{"accountId": "account-id-sample"}).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := issueService.Assign(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Notify(t *testing.T) {

	optionsMocked := &model.IssueNotifyOptionsScheme{
		HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
		Subject:  "SUBJECT EMAIL EXAMPLE",
		To: &model.IssueNotifyToScheme{
			Reporter: true,
			Assignee: true,
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		options      *model.IssueNotifyOptionsScheme
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				options:      optionsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/notify",
					"",
					optionsMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				options:      optionsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/notify",
					"",
					optionsMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := issueService.Notify(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Transitions(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTransitionsScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := issueService.Transitions(testCase.args.ctx, testCase.args.issueKeyOrID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Create(t *testing.T) {

	payloadMocked := &model.IssueSchemeV2{
		Fields: &model.IssueFieldsSchemeV2{
			Summary:   "New summary test",
			Project:   &model.ProjectScheme{ID: "10000"},
			IssueType: &model.IssueTypeScheme{Name: "Story"},
		},
	}

	customFieldsMocked := &model.CustomFields{}

	// Add a new custom field
	err := customFieldsMocked.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldsMocked.Number("customfield_10042", 1000.2222)
	if err != nil {
		t.Fatal(err)
	}

	expectedPayloadWithCustomFields := map[string]interface{}{
		"fields": map[string]interface{}{
			"customfield_10042": 1000.2222,
			"customfield_10052": []map[string]interface{}{map[string]interface{}{"name": "jira-administrators"}, map[string]interface{}{"name": "jira-administrators-system"}},
			"issuetype":         map[string]interface{}{"name": "Story"},
			"project":           map[string]interface{}{"id": "10000"},
			"summary":           "New summary test"},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		payload      *model.IssueSchemeV2
		customFields *model.CustomFields
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				payload:      payloadMocked,
				customFields: customFieldsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue",
					"",
					expectedPayloadWithCustomFields).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the customfield are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				payload:      payloadMocked,
				customFields: nil,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				payload:      payloadMocked,
				customFields: customFieldsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue",
					"",
					expectedPayloadWithCustomFields).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := issueService.Create(testCase.args.ctx, testCase.args.payload, testCase.args.customFields)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Creates(t *testing.T) {

	customFieldsMocked := &model.CustomFields{}

	err := customFieldsMocked.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldsMocked.Number("customfield_10042", 1000.2222)
	if err != nil {
		t.Fatal(err)
	}

	payloadMocked := []*model.IssueBulkSchemeV2{
		{
			Payload: &model.IssueSchemeV2{
				Fields: &model.IssueFieldsSchemeV2{
					Summary:   "New summary test",
					Project:   &model.ProjectScheme{ID: "10000"},
					IssueType: &model.IssueTypeScheme{Name: "Story"},
				},
			},
			CustomFields: customFieldsMocked,
		},

		{
			Payload:      nil,
			CustomFields: nil,
		},

		{
			Payload: &model.IssueSchemeV2{
				Fields: &model.IssueFieldsSchemeV2{
					Summary:   "New summary test #2",
					Project:   &model.ProjectScheme{ID: "10000"},
					IssueType: &model.IssueTypeScheme{Name: "Story"},
				},
			},
			CustomFields: customFieldsMocked,
		},
	}

	expectedBulkWithCustomFieldsPayload := map[string]interface{}{

		"issueUpdates": []map[string]interface{}{map[string]interface{}{

			"fields": map[string]interface{}{
				"customfield_10042": 1000.2222,
				"customfield_10052": []map[string]interface{}{map[string]interface{}{"name": "jira-administrators"}, map[string]interface{}{"name": "jira-administrators-system"}},
				"issuetype":         map[string]interface{}{"name": "Story"},
				"project":           map[string]interface{}{"id": "10000"},
				"summary":           "New summary test"}}, map[string]interface{}{

			"fields": map[string]interface{}{
				"customfield_10042": 1000.2222,
				"customfield_10052": []map[string]interface{}{map[string]interface{}{"name": "jira-administrators"}, map[string]interface{}{"name": "jira-administrators-system"}},
				"issuetype":         map[string]interface{}{"name": "Story"},
				"project":           map[string]interface{}{"id": "10000"},
				"summary":           "New summary test #2"}}}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload []*model.IssueBulkSchemeV2
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/bulk",
					"",
					expectedBulkWithCustomFieldsPayload).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueBulkResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the payload is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				payload: nil,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoCreateIssues,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/bulk",
					"",
					expectedBulkWithCustomFieldsPayload).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := issueService.Creates(testCase.args.ctx, testCase.args.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Get(t *testing.T) {

	customFields := &model.CustomFields{}

	// Add a new custom field
	err := customFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFields.Number("customfield_10042", 1000.2222)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx            context.Context
		issueKeyOrID   string
		fields, expand []string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fields:       []string{"summary", "status"},
				expand:       []string{"operations", "changelogts"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1?expand=operations%2Cchangelogts&fields=summary%2Cstatus",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
				fields:       []string{"summary", "status"},
				expand:       []string{"operations", "changelogts"},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fields:       []string{"summary", "status"},
				expand:       []string{"operations", "changelogts"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1?expand=operations%2Cchangelogts&fields=summary%2Cstatus",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := issueService.Get(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.fields,
				testCase.args.expand)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Move(t *testing.T) {

	/*
		"customfield_10042": 1000.2222,
		"customfield_10052": [
			{
				"name": "jira-administrators"
			},
			{
				"name": "jira-administrators-system"
			}
		]
	*/
	customFieldsMocked := &model.CustomFields{}
	if err := customFieldsMocked.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"}); err != nil {
		t.Fatal(err)
	}
	if err := customFieldsMocked.Number("customfield_10042", 1000.2222); err != nil {
		t.Fatal(err)
	}

	/*
		 "update": {
			"labels": [
				{
					"remove": "triaged"
				}
			]
		}
	*/
	operationsMocked := &model.UpdateOperations{}
	if err := operationsMocked.AddArrayOperation("labels", map[string]string{"triaged": "remove"}); err != nil {
		t.Fatal(err)
	}

	expectedPayloadWithCustomFieldsAndOperations := map[string]interface{}{
		"fields": map[string]interface{}{
			"customfield_10042": 1000.2222,
			"customfield_10052": []map[string]interface{}{map[string]interface{}{
				"name": "jira-administrators"}, map[string]interface{}{
				"name": "jira-administrators-system"}},

			"issuetype":  map[string]interface{}{"name": "Story"},
			"project":    map[string]interface{}{"id": "10000"},
			"resolution": map[string]interface{}{"name": "Done"},
			"summary":    "New summary test"},

		"update": map[string]interface{}{
			"labels": []map[string]interface{}{map[string]interface{}{
				"remove": "triaged"}}},

		"transition": map[string]interface{}{"id": "10001"},
	}

	expectedPayloadWithCustomfields := map[string]interface{}{
		"fields": map[string]interface{}{
			"customfield_10042": 1000.2222,
			"customfield_10052": []map[string]interface{}{map[string]interface{}{
				"name": "jira-administrators"}, map[string]interface{}{
				"name": "jira-administrators-system"}},

			"issuetype": map[string]interface{}{"name": "Story"},
			"project":   map[string]interface{}{"id": "10000"},
			"summary":   "New summary test"},
		"transition": map[string]interface{}{"id": "10001"},
	}

	expectedPayloadWithOperations := map[string]interface{}{
		"fields": map[string]interface{}{
			"issuetype": map[string]interface{}{"name": "Story"},
			"project":   map[string]interface{}{"id": "10000"},
			"summary":   "New summary test"},

		"update": map[string]interface{}{
			"labels": []map[string]interface{}{map[string]interface{}{
				"remove": "triaged"}}},
		"transition": map[string]interface{}{"id": "10001"},
	}

	expectedPayloadWithNoOptions := map[string]interface{}{"transition": map[string]interface{}{"id": "10001"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                        context.Context
		issueKeyOrID, transitionID string
		options                    *model.IssueMoveOptionsV2
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "10001",
				options: &model.IssueMoveOptionsV2{
					Fields: &model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
							Resolution: &model.ResolutionScheme{
								Name: "Done",
							},
						},
					},
					CustomFields: customFieldsMocked,
					Operations:   operationsMocked,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					expectedPayloadWithCustomFieldsAndOperations).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the options are provided and the fields are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "10001",
				options: &model.IssueMoveOptionsV2{
					CustomFields: customFieldsMocked,
					Operations:   operationsMocked,
				},
			},
			wantErr: true,
			Err:     model.ErrNoIssueScheme,
		},

		{
			name:   "when the operations are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "10001",
				options: &model.IssueMoveOptionsV2{
					Fields: &model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFieldsMocked,
					Operations:   nil,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					expectedPayloadWithCustomfields).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the custom fields are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "10001",
				options: &model.IssueMoveOptionsV2{
					Fields: &model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: nil,
					Operations:   operationsMocked,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					expectedPayloadWithOperations).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the the issue comment options are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "10001",
				options:      nil,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					expectedPayloadWithNoOptions).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue key is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
				transitionID: "10001",
				options: &model.IssueMoveOptionsV2{
					Fields: &model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFieldsMocked,
					Operations:   operationsMocked,
				},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the transition id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "",
				options: &model.IssueMoveOptionsV2{
					Fields: &model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFieldsMocked,
					Operations:   operationsMocked,
				},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoTransitionID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				transitionID: "10001",
				options: &model.IssueMoveOptionsV2{
					Fields: &model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFieldsMocked,
					Operations:   operationsMocked,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/transitions",
					"",
					mock.Anything).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := issueService.Move(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.transitionID,
				testCase.args.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}

func Test_internalRichTextServiceImpl_Update(t *testing.T) {

	customFieldsMocked := &model.CustomFields{}

	err := customFieldsMocked.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldsMocked.Number("customfield_10042", 1000.2222)
	if err != nil {
		t.Fatal(err)
	}

	operations := &model.UpdateOperations{}
	err = operations.AddArrayOperation("labels", map[string]string{
		"triaged": "remove",
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedPayloadWithCustomFieldsAndOperations := map[string]interface{}{
		"fields": map[string]interface{}{
			"customfield_10042": 1000.2222,
			"customfield_10052": []map[string]interface{}{map[string]interface{}{
				"name": "jira-administrators"}, map[string]interface{}{
				"name": "jira-administrators-system"}}, "summary": "New summary test"},
		"update": map[string]interface{}{
			"labels": []map[string]interface{}{map[string]interface{}{
				"remove": "triaged"}}}}

	expectedPayloadWithCustomfields := map[string]interface{}{
		"fields": map[string]interface{}{
			"customfield_10042": 1000.2222,
			"customfield_10052": []map[string]interface{}{map[string]interface{}{
				"name": "jira-administrators"}, map[string]interface{}{
				"name": "jira-administrators-system"}},
			"summary": "New summary test"}}

	expectedPayloadWithOperations := map[string]interface{}{
		"fields": map[string]interface{}{
			"summary": "New summary test"},
		"update": map[string]interface{}{
			"labels": []map[string]interface{}{map[string]interface{}{"remove": "triaged"}}}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		notify       bool
		payload      *model.IssueSchemeV2
		customFields *model.CustomFields
		operations   *model.UpdateOperations
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				notify:       true,
				payload: &model.IssueSchemeV2{
					Fields: &model.IssueFieldsSchemeV2{
						Summary: "New summary test",
					},
				},
				customFields: customFieldsMocked,
				operations:   operations,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-1?notifyUsers=true",
					"",
					expectedPayloadWithCustomFieldsAndOperations).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
				notify:       true,
				payload: &model.IssueSchemeV2{
					Fields: &model.IssueFieldsSchemeV2{
						Summary: "New summary test",
					},
				},
				customFields: customFieldsMocked,
				operations:   operations,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				notify:       true,
				payload: &model.IssueSchemeV2{
					Fields: &model.IssueFieldsSchemeV2{
						Summary: "New summary test",
					},
				},
				customFields: customFieldsMocked,
				operations:   operations,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-1?notifyUsers=true",
					"",
					expectedPayloadWithCustomFieldsAndOperations).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name:   "when the operations are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				notify:       true,
				payload: &model.IssueSchemeV2{
					Fields: &model.IssueFieldsSchemeV2{
						Summary: "New summary test",
					},
				},
				customFields: customFieldsMocked,
				operations:   nil,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-1?notifyUsers=true",
					"",
					expectedPayloadWithCustomfields).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the custom fields are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				notify:       true,
				payload: &model.IssueSchemeV2{
					Fields: &model.IssueFieldsSchemeV2{
						Summary: "New summary test",
					},
				},
				customFields: nil,
				operations:   operations,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-1?notifyUsers=true",
					"",
					expectedPayloadWithOperations).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the operations are customfields are not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				notify:       true,
				payload: &model.IssueSchemeV2{
					Fields: &model.IssueFieldsSchemeV2{
						Summary: "New summary test",
					},
				},
				customFields: nil,
				operations:   nil,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-1?notifyUsers=true",
					"",
					&model.IssueSchemeV2{
						Fields: &model.IssueFieldsSchemeV2{
							Summary: "New summary test",
						},
					}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			issueService, _, err := NewIssueService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := issueService.Update(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.notify,
				testCase.args.payload, testCase.args.customFields, testCase.args.operations)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}
