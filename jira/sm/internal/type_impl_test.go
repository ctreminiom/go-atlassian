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

func Test_internalTypeImpl_Search(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx          context.Context
		query        string
		start, limit int
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
			name: "when the parameters are correct",
			args: args{
				ctx:   context.Background(),
				query: "Request Testing Environment",
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/requesttype?limit=50&searchQuery=Request+Testing+Environment&start=100",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:   context.Background(),
				query: "Request Testing Environment",
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/requesttype?limit=50&searchQuery=Request+Testing+Environment&start=100",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:   context.Background(),
				query: "Request Testing Environment",
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/requesttype?limit=50&searchQuery=Request+Testing+Environment&start=100",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewTypeService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Search(testCase.args.ctx, testCase.args.query, testCase.args.start,
				testCase.args.limit)

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

func Test_internalTypeImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                    context.Context
		serviceDeskID, groupID int
		start, limit           int
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				groupID:       38383,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype?groupId=38383&limit=50&start=100",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectRequestTypePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				groupID:       38383,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype?groupId=38383&limit=50&start=100",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectRequestTypePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				groupID:       38383,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype?groupId=38383&limit=50&start=100",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewTypeService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Gets(testCase.args.ctx, testCase.args.serviceDeskID,
				testCase.args.groupID, testCase.args.start, testCase.args.limit)

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

func Test_internalTypeImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                          context.Context
		serviceDeskID, requestTypeID int
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypeScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskIDError,
			wantErr: true,
		},

		{
			name: "when the request type id is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			Err:     model.ErrNoRequestTypeIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewTypeService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Get(testCase.args.ctx, testCase.args.serviceDeskID,
				testCase.args.requestTypeID)

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

func Test_internalTypeImpl_Fields(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                          context.Context
		serviceDeskID, requestTypeID int
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383/field",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypeFieldsScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383/field",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypeFieldsScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383/field",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskIDError,
			wantErr: true,
		},

		{
			name: "when the request type id is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			Err:     model.ErrNoRequestTypeIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewTypeService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Fields(testCase.args.ctx, testCase.args.serviceDeskID,
				testCase.args.requestTypeID)

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

func Test_internalTypeImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                          context.Context
		serviceDeskID, requestTypeID int
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				requestTypeID: 38383,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/requesttype/38383",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskIDError,
			wantErr: true,
		},

		{
			name: "when the request type id is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			Err:     model.ErrNoRequestTypeIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewTypeService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResponse, err := smService.Delete(testCase.args.ctx, testCase.args.serviceDeskID,
				testCase.args.requestTypeID)

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

func Test_internalTypeImpl_Create(t *testing.T) {

	payloadMocked := &struct {
		IssueTypeID string "json:\"issueTypeId,omitempty\""
		HelpText    string "json:\"helpText,omitempty\""
		Name        string "json:\"name,omitempty\""
		Description string "json:\"description,omitempty\""
	}{IssueTypeID: "12345", HelpText: "Please tell us clearly the problem you have within 100 words.", Name: "Get IT Help", Description: "Get IT Help"}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                                      context.Context
		serviceDeskID                            int
		issueTypeID, name, description, helpText string
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				issueTypeID:   "12345",
				name:          "Get IT Help",
				description:   "Get IT Help",
				helpText:      "Please tell us clearly the problem you have within 100 words.",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/requesttype",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				issueTypeID:   "12345",
				name:          "Get IT Help",
				description:   "Get IT Help",
				helpText:      "Please tell us clearly the problem you have within 100 words.",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/requesttype",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestTypeScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				issueTypeID:   "12345",
				name:          "Get IT Help",
				description:   "Get IT Help",
				helpText:      "Please tell us clearly the problem you have within 100 words.",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/requesttype",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewTypeService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Create(testCase.args.ctx, testCase.args.serviceDeskID,
				testCase.args.issueTypeID, testCase.args.name, testCase.args.description, testCase.args.helpText)

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
