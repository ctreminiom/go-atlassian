package internal

import (
	"bytes"
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalIssueADFServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		issueKeyOrId   string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				issueKeyOrId:   "DUMMY-1",
				deleteSubTasks: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-1?deleteSubtasks=true",
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
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				issueKeyOrId:   "",
				deleteSubTasks: true,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				issueKeyOrId:   "DUMMY-1",
				deleteSubTasks: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-1?deleteSubtasks=true",
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := issueService.Delete(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.deleteSubTasks)

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

func Test_internalIssueADFServiceImpl_Assign(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrId, accountId string
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				accountId:    "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						AccountID string "json:\"accountId\""
					}{AccountID: "account-id-sample"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"/rest/api/3/issue/DUMMY-1/assignee",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
				accountId:    "account-id-sample",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the account id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				accountId:    "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoAccountIDError,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				accountId:    "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						AccountID string "json:\"accountId\""
					}{AccountID: "account-id-sample"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"/rest/api/3/issue/DUMMY-1/assignee",
					bytes.NewReader([]byte{})).
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := issueService.Assign(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.accountId)

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

func Test_internalIssueADFServiceImpl_Notify(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrId string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				options: &model.IssueNotifyOptionsScheme{
					HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
					Subject:  "SUBJECT EMAIL EXAMPLE",
					To: &model.IssueNotifyToScheme{
						Reporter: true,
						Assignee: true,
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.IssueNotifyOptionsScheme{
						HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
						Subject:  "SUBJECT EMAIL EXAMPLE",
						To: &model.IssueNotifyToScheme{
							Reporter: true,
							Assignee: true,
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/notify",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				options: &model.IssueNotifyOptionsScheme{
					HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
					Subject:  "SUBJECT EMAIL EXAMPLE",
					To: &model.IssueNotifyToScheme{
						Reporter: true,
						Assignee: true,
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.IssueNotifyOptionsScheme{
						HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
						Subject:  "SUBJECT EMAIL EXAMPLE",
						To: &model.IssueNotifyToScheme{
							Reporter: true,
							Assignee: true,
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/notify",
					bytes.NewReader([]byte{})).
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := issueService.Notify(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.options)

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

func Test_internalIssueADFServiceImpl_Transitions(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrId string
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1/transitions",
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the request method cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1/transitions",
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := issueService.Transitions(testCase.args.ctx, testCase.args.issueKeyOrId)

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

func Test_internalIssueADFServiceImpl_Create(t *testing.T) {

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
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		payload      *model.IssueScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary:   "New summary test",
						Project:   &model.ProjectScheme{ID: "10000"},
						IssueType: &model.IssueTypeScheme{Name: "Story"},
					},
				},
				customFields: customFields,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					map[string]interface{}{
						"fields": map[string]interface{}{"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary:   "New summary test",
						Project:   &model.ProjectScheme{ID: "10000"},
						IssueType: &model.IssueTypeScheme{Name: "Story"},
					},
				},
				customFields: nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.IssueScheme{Fields: &model.IssueFieldsScheme{
						Summary:   "New summary test",
						Project:   &model.ProjectScheme{ID: "10000"},
						IssueType: &model.IssueTypeScheme{Name: "Story"},
					}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary:   "New summary test",
						Project:   &model.ProjectScheme{ID: "10000"},
						IssueType: &model.IssueTypeScheme{Name: "Story"},
					},
				},
				customFields: customFields,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					map[string]interface{}{
						"fields": map[string]interface{}{"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue",
					bytes.NewReader([]byte{})).
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
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

func Test_internalIssueADFServiceImpl_Creates(t *testing.T) {

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
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		payload []*model.IssueBulkSchemeV3
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
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: []*model.IssueBulkSchemeV3{
					{
						Payload: &model.IssueScheme{
							Fields: &model.IssueFieldsScheme{
								Summary:   "New summary test",
								Project:   &model.ProjectScheme{ID: "10000"},
								IssueType: &model.IssueTypeScheme{Name: "Story"},
							},
						},
						CustomFields: customFields,
					},

					{
						Payload:      nil,
						CustomFields: nil,
					},

					{
						Payload: &model.IssueScheme{
							Fields: &model.IssueFieldsScheme{
								Summary:   "New summary test #2",
								Project:   &model.ProjectScheme{ID: "10000"},
								IssueType: &model.IssueTypeScheme{Name: "Story"},
							},
						},
						CustomFields: customFields,
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{"issueUpdates": []map[string]interface{}{{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"}},

						{"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test #2"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/bulk",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				payload: nil,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     errors.New("error, please provide a valid []*IssueBulkScheme slice of pointers"),
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: []*model.IssueBulkSchemeV3{
					{
						Payload: &model.IssueScheme{
							Fields: &model.IssueFieldsScheme{
								Summary:   "New summary test",
								Project:   &model.ProjectScheme{ID: "10000"},
								IssueType: &model.IssueTypeScheme{Name: "Story"},
							},
						},
						CustomFields: customFields,
					},

					{
						Payload:      nil,
						CustomFields: nil,
					},

					{
						Payload: &model.IssueScheme{
							Fields: &model.IssueFieldsScheme{
								Summary:   "New summary test #2",
								Project:   &model.ProjectScheme{ID: "10000"},
								IssueType: &model.IssueTypeScheme{Name: "Story"},
							},
						},
						CustomFields: customFields,
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{"issueUpdates": []map[string]interface{}{{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"}},

						{"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test #2"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/bulk",
					bytes.NewReader([]byte{})).
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
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

func Test_internalIssueADFServiceImpl_Get(t *testing.T) {

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
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		issueKeyOrId   string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				fields:       []string{"summary", "status"},
				expand:       []string{"operations", "changelogts"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1?expand=operations%2Cchangelogts&fields=summary%2Cstatus",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
				fields:       []string{"summary", "status"},
				expand:       []string{"operations", "changelogts"},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				fields:       []string{"summary", "status"},
				expand:       []string{"operations", "changelogts"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1?expand=operations%2Cchangelogts&fields=summary%2Cstatus",
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := issueService.Get(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.fields,
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

func Test_internalIssueADFServiceImpl_Move(t *testing.T) {

	customFields := &model.CustomFields{}

	err := customFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFields.Number("customfield_10042", 1000.2222)
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

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                        context.Context
		issueKeyOrId, transitionId string
		options                    *model.IssueMoveOptionsV3
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				transitionId: "10001",
				options: &model.IssueMoveOptionsV3{
					Fields: &model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFields,
					Operations:   operations,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"},
						"transition": map[string]interface{}{"id": "10001"},
						"update": map[string]interface{}{
							"labels": []map[string]interface{}{{"remove": "triaged"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/transitions",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the operations are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				transitionId: "10001",
				options: &model.IssueMoveOptionsV3{
					Fields: &model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFields,
					Operations:   nil,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"},
						"transition": map[string]interface{}{"id": "10001"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/transitions",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				transitionId: "10001",
				options: &model.IssueMoveOptionsV3{
					Fields: &model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: nil,
					Operations:   operations,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"issuetype": map[string]interface{}{"name": "Story"},
							"project":   map[string]interface{}{"id": "10000"},
							"summary":   "New summary test"},
						"transition": map[string]interface{}{"id": "10001"},
						"update": map[string]interface{}{
							"labels": []map[string]interface{}{{"remove": "triaged"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/transitions",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				transitionId: "10001",
				options:      nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"transition": map[string]interface{}{"id": "10001"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/transitions",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
				transitionId: "10001",
				options: &model.IssueMoveOptionsV3{
					Fields: &model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFields,
					Operations:   operations,
				},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the transition id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				transitionId: "",
				options: &model.IssueMoveOptionsV3{
					Fields: &model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFields,
					Operations:   operations,
				},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoTransitionIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				transitionId: "10001",
				options: &model.IssueMoveOptionsV3{
					Fields: &model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary:   "New summary test",
							Project:   &model.ProjectScheme{ID: "10000"},
							IssueType: &model.IssueTypeScheme{Name: "Story"},
						},
					},
					CustomFields: customFields,
					Operations:   operations,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"issuetype":         map[string]interface{}{"name": "Story"},
							"project":           map[string]interface{}{"id": "10000"},
							"summary":           "New summary test"},
						"transition": map[string]interface{}{"id": "10001"},
						"update": map[string]interface{}{
							"labels": []map[string]interface{}{{"remove": "triaged"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/transitions",
					bytes.NewReader([]byte{})).
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := issueService.Move(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.transitionId,
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

func Test_internalIssueADFServiceImpl_Update(t *testing.T) {

	customFields := &model.CustomFields{}

	err := customFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFields.Number("customfield_10042", 1000.2222)
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

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrId string
		notify       bool
		payload      *model.IssueScheme
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				notify:       true,
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary: "New summary test",
					},
				},
				customFields: customFields,
				operations:   operations,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"summary":           "New summary test"},
						"update": map[string]interface{}{
							"labels": []map[string]interface{}{{"remove": "triaged"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-1?notifyUsers=true",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
				notify:       true,
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary: "New summary test",
					},
				},
				customFields: customFields,
				operations:   operations,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				notify:       true,
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary: "New summary test",
					},
				},
				customFields: customFields,
				operations:   operations,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"summary":           "New summary test"},
						"update": map[string]interface{}{
							"labels": []map[string]interface{}{{"remove": "triaged"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-1?notifyUsers=true",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name:   "when the operations are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				notify:       true,
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary: "New summary test",
					},
				},
				customFields: customFields,
				operations:   nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"customfield_10042": 1000.2222,
							"customfield_10052": []map[string]interface{}{{"name": "jira-administrators"}, {"name": "jira-administrators-system"}},
							"summary":           "New summary test"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-1?notifyUsers=true",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				notify:       true,
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary: "New summary test",
					},
				},
				customFields: nil,
				operations:   operations,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&map[string]interface{}{
						"fields": map[string]interface{}{
							"summary": "New summary test"},
						"update": map[string]interface{}{
							"labels": []map[string]interface{}{{"remove": "triaged"}}}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-1?notifyUsers=true",
					bytes.NewReader([]byte{})).
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
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				notify:       true,
				payload: &model.IssueScheme{
					Fields: &model.IssueFieldsScheme{
						Summary: "New summary test",
					},
				},
				customFields: nil,
				operations:   nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.IssueScheme{
						Fields: &model.IssueFieldsScheme{
							Summary: "New summary test",
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-1?notifyUsers=true",
					bytes.NewReader([]byte{})).
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

			_, issueService, err := NewIssueService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := issueService.Update(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.notify,
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
