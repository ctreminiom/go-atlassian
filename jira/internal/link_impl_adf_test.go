package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalLinkADFServiceImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx    context.Context
		linkId string
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
				ctx:    context.TODO(),
				linkId: "10002",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issueLink/10002",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueLinkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:    context.TODO(),
				linkId: "10002",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issueLink/10002",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueLinkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:    context.TODO(),
				linkId: "10002",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issueLink/10002",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name:   "when the issue link is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:    context.TODO(),
				linkId: "",
			},
			wantErr: true,
			Err:     model.ErrNoTypeIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			linkService, _, err := NewLinkService(testCase.fields.c, testCase.fields.version, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := linkService.Get(testCase.args.ctx, testCase.args.linkId)

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

func Test_internalLinkADFServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
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
				issueKeyOrId: "DUMMY-4",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-4?fields=issuelinks",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueLinkPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-4",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-4?fields=issuelinks",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueLinkPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-4",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-4?fields=issuelinks",
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

			linkService, _, err := NewLinkService(testCase.fields.c, testCase.fields.version, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := linkService.Gets(testCase.args.ctx, testCase.args.issueKeyOrId)

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

func Test_internalLinkADFServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx    context.Context
		linkId string
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
				ctx:    context.TODO(),
				linkId: "10002",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issueLink/10002",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:    context.TODO(),
				linkId: "10002",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issueLink/10002",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:    context.TODO(),
				linkId: "10002",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issueLink/10002",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name:   "when the issue link is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:    context.TODO(),
				linkId: "",
			},
			wantErr: true,
			Err:     model.ErrNoTypeIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			linkService, _, err := NewLinkService(testCase.fields.c, testCase.fields.version, nil, nil)
			assert.NoError(t, err)

			gotResponse, err := linkService.Delete(testCase.args.ctx, testCase.args.linkId)

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

func Test_internalLinkADFServiceImpl_Create(t *testing.T) {

	payloadMocked := &model.LinkPayloadSchemeV3{
		InwardIssue: &model.LinkedIssueScheme{
			Key: "KP-1",
		},
		OutwardIssue: &model.LinkedIssueScheme{
			Key: "KP-2",
		},
		Type: &model.LinkTypeScheme{
			Name: "Duplicate",
		},
		Comment: &model.CommentPayloadScheme{
			Body: &model.CommentNodeScheme{
				Version: 1,
				Type:    "doc",
				Content: []*model.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*model.CommentNodeScheme{
							{
								Type: "text",
								Text: "Carlos Test",
							},
							{
								Type: "emoji",
								Attrs: map[string]interface{}{
									"shortName": ":grin",
									"id":        "1f601",
									"text":      "😁",
								},
							},
						},
					},
				},
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.LinkPayloadSchemeV3
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issueLink",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issueLink",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issueLink",
					"",
					payloadMocked).
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

			linkService, _, err := NewLinkService(testCase.fields.c, testCase.fields.version, nil, nil)
			assert.NoError(t, err)

			gotResponse, err := linkService.Create(testCase.args.ctx, testCase.args.payload)

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
