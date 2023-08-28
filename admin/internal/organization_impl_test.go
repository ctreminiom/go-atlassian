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
	"time"
)

func Test_internalOrganizationImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx    context.Context
		cursor string
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
				ctx:    context.TODO(),
				cursor: "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs?cursor=cursor-sample-uuid",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AdminOrganizationPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:    context.TODO(),
				cursor: "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs?cursor=cursor-sample-uuid",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Gets(testCase.args.ctx, testCase.args.cursor)

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

func Test_internalOrganizationImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AdminOrganizationScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganizationError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Get(testCase.args.ctx, testCase.args.organizationID)

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

func Test_internalOrganizationImpl_Users(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                    context.Context
		organizationID, cursor string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				cursor:         "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/users?cursor=cursor-sample-uuid",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationUserPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganizationError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				cursor:         "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/users?cursor=cursor-sample-uuid",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Users(testCase.args.ctx, testCase.args.organizationID, testCase.args.cursor)

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

func Test_internalOrganizationImpl_Domains(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                    context.Context
		organizationID, cursor string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				cursor:         "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/domains?cursor=cursor-sample-uuid",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationDomainPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganizationError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				cursor:         "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/domains?cursor=cursor-sample-uuid",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Domains(testCase.args.ctx, testCase.args.organizationID, testCase.args.cursor)

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

func Test_internalOrganizationImpl_Domain(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                      context.Context
		organizationID, domainID string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				domainID:       "domain-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/domains/domain-sample-uuid",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationDomainScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganizationError,
		},

		{
			name: "when the domain id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDomainIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				domainID:       "domain-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/domains/domain-sample-uuid",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Domain(testCase.args.ctx, testCase.args.organizationID, testCase.args.domainID)

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

func Test_internalOrganizationImpl_Events(t *testing.T) {

	fromMocked, err := time.Parse(time.RFC3339Nano, "2020-05-12T11:45:26.371Z")
	if err != nil {
		t.Fatal(err)
	}

	toMocked, err := time.Parse(time.RFC3339Nano, "2020-11-12T11:45:26.371Z")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID string
		options        *model.OrganizationEventOptScheme
		cursor         string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				options: &model.OrganizationEventOptScheme{
					Q:      "qq",
					From:   fromMocked.Add(time.Duration(-24) * time.Hour),
					To:     toMocked.Add(time.Duration(-1) * time.Hour),
					Action: "user_added_to_group",
				},
				cursor: "cursor-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/events?action=user_added_to_group&cursor=cursor-id-sample&from=1589197526&q=qq&to=1605177926",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationEventPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganizationError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				options: &model.OrganizationEventOptScheme{
					Q:      "qq",
					From:   fromMocked.Add(time.Duration(-24) * time.Hour),
					To:     toMocked.Add(time.Duration(-1) * time.Hour),
					Action: "user_added_to_group",
				},
				cursor: "cursor-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/events?action=user_added_to_group&cursor=cursor-id-sample&from=1589197526&q=qq&to=1605177926",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Events(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.options, testCase.args.cursor)

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

func Test_internalOrganizationImpl_Event(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                     context.Context
		organizationID, eventID string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				eventID:        "event-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/events/event-sample-uuid",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationEventScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganizationError,
		},

		{
			name: "when the event id is not provided",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
			},
			wantErr: true,
			Err:     model.ErrNoEventIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
				eventID:        "event-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/events/event-sample-uuid",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Event(testCase.args.ctx, testCase.args.organizationID, testCase.args.eventID)

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

func Test_internalOrganizationImpl_Actions(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID string
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
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/event-actions",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationEventActionScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.TODO(),
				organizationID: "organization-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-sample-uuid/event-actions",
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

			newOrganizationService := NewOrganizationService(testCase.fields.c, nil, nil)

			gotResult, gotResponse, err := newOrganizationService.Actions(testCase.args.ctx, testCase.args.organizationID)

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
