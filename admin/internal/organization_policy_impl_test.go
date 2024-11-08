package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalOrganizationPolicyImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                context.Context
		organizationID, policyType, cursor string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyType:     "policy-type-sample",
				cursor:         "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-id-sample/policies?cursor=cursor-sample-uuid&type=policy-type-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPolicyPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyType:     "policy-type-sample",
				cursor:         "cursor-sample-uuid",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-id-sample/policies?cursor=cursor-sample-uuid&type=policy-type-sample",
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

			newOrganizationPolicyService := NewOrganizationPolicyService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationPolicyService.Gets(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.policyType, testCase.args.cursor)

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

func Test_internalOrganizationPolicyImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                      context.Context
		organizationID, policyID string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyID:       "policy-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-id-sample/policies/policy-id-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPolicyScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminPolicy,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyID:       "policy-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-id-sample/policies/policy-id-sample",
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

			newOrganizationPolicyService := NewOrganizationPolicyService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationPolicyService.Get(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.policyID)

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

func Test_internalOrganizationPolicyImpl_Create(t *testing.T) {

	payloadMocked := &model.OrganizationPolicyData{
		Type: "policy",
		Attributes: &model.OrganizationPolicyAttributes{
			Type:   "data-residency", //ip-allowlist
			Name:   "SCIMUserNameScheme of this Policy",
			Status: "enabled", //disabled
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID string
		payload        *model.OrganizationPolicyData
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"admin/v1/orgs/organization-id-sample/policies",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPolicyScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"admin/v1/orgs/organization-id-sample/policies",
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

			newOrganizationPolicyService := NewOrganizationPolicyService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationPolicyService.Create(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.payload)

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

func Test_internalOrganizationPolicyImpl_Update(t *testing.T) {

	payloadMocked := &model.OrganizationPolicyData{
		Type: "policy",
		Attributes: &model.OrganizationPolicyAttributes{
			Type:   "data-residency", //ip-allowlist
			Name:   "SCIMUserNameScheme of this Policy",
			Status: "enabled", //disabled
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                      context.Context
		organizationID, policyID string
		payload                  *model.OrganizationPolicyData
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyID:       "policy-id-sample",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"admin/v1/orgs/organization-id-sample/policies/policy-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPolicyScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the policy id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminPolicy,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyID:       "policy-id-sample",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"admin/v1/orgs/organization-id-sample/policies/policy-id-sample",
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

			newOrganizationPolicyService := NewOrganizationPolicyService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationPolicyService.Update(testCase.args.ctx, testCase.args.organizationID, testCase.args.policyID,
				testCase.args.payload)

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

func Test_internalOrganizationPolicyImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                      context.Context
		organizationID, policyID string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyID:       "policy-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"admin/v1/orgs/organization-id-sample/policies/policy-id-sample",
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
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the policy id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminPolicy,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				policyID:       "policy-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"admin/v1/orgs/organization-id-sample/policies/policy-id-sample",
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

			newOrganizationPolicyService := NewOrganizationPolicyService(testCase.fields.c)

			gotResponse, err := newOrganizationPolicyService.Delete(testCase.args.ctx, testCase.args.organizationID, testCase.args.policyID)

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
