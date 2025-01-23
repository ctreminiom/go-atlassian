package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
)

func Test_internalServiceRequestImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		options      *model.ServiceRequestOptionScheme
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
				ctx: context.Background(),
				options: &model.ServiceRequestOptionScheme{
					ApprovalStatus:    "MY_PENDING_APPROVAL",
					RequestStatus:     "OPEN_REQUESTS",
					SearchTerm:        "IT Help",
					OrganizationID:    39933,
					ServiceDeskID:     10002,
					RequestTypeID:     1002,
					Expand:            []string{"serviceDesk", "action"},
					RequestOwnerships: []string{"ORGANIZATION"},
				},
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request?approvalStatus=MY_PENDING_APPROVAL&expand=serviceDesk%2Caction&limit=50&organizationId=39933&requestOwnership=ORGANIZATION&requestStatus=OPEN_REQUESTS&requestTypeId=1002&searchTerm=IT+Help&serviceDeskId=10002&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx: context.Background(),
				options: &model.ServiceRequestOptionScheme{
					ApprovalStatus:    "MY_PENDING_APPROVAL",
					RequestStatus:     "OPEN_REQUESTS",
					SearchTerm:        "IT Help",
					OrganizationID:    39933,
					ServiceDeskID:     10002,
					RequestTypeID:     1002,
					Expand:            []string{"serviceDesk", "action"},
					RequestOwnerships: []string{"ORGANIZATION"},
				},
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request?approvalStatus=MY_PENDING_APPROVAL&expand=serviceDesk%2Caction&limit=50&organizationId=39933&requestOwnership=ORGANIZATION&requestStatus=OPEN_REQUESTS&requestTypeId=1002&searchTerm=IT+Help&serviceDeskId=10002&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx: context.Background(),
				options: &model.ServiceRequestOptionScheme{
					ApprovalStatus:    "MY_PENDING_APPROVAL",
					RequestStatus:     "OPEN_REQUESTS",
					SearchTerm:        "IT Help",
					OrganizationID:    39933,
					ServiceDeskID:     10002,
					RequestTypeID:     1002,
					Expand:            []string{"serviceDesk", "action"},
					RequestOwnerships: []string{"ORGANIZATION"},
				},
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request?approvalStatus=MY_PENDING_APPROVAL&expand=serviceDesk%2Caction&limit=50&organizationId=39933&requestOwnership=ORGANIZATION&requestStatus=OPEN_REQUESTS&requestTypeId=1002&searchTerm=IT+Help&serviceDeskId=10002&start=100",
					"",
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

			smService, err := NewRequestService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Gets(testCase.args.ctx, testCase.args.options, testCase.args.start,
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

func Test_internalServiceRequestImpl_Transitions(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1/transition?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestTransitionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1/transition?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestTransitionPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1/transition?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewRequestService(testCase.fields.c, "latest", &ServiceRequestSubServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Transitions(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.start,
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

func Test_internalServiceRequestImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		expand       []string
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				expand:       []string{"serviceDesk"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1?expand=serviceDesk",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				expand:       []string{"serviceDesk"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1?expand=serviceDesk",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				expand:       []string{"serviceDesk"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1?expand=serviceDesk",
					"",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewRequestService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Get(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.expand)

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

func Test_internalServiceRequestImpl_Subscribe(t *testing.T) {

	type fields struct {
		c service.Connector
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
			name: "when the parameters are correct",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/servicedeskapi/request/DESK-1/notification",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/servicedeskapi/request/DESK-1/notification",
					"",
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/servicedeskapi/request/DESK-1/notification",
					"",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewRequestService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResponse, err := smService.Subscribe(testCase.args.ctx, testCase.args.issueKeyOrID)

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

func Test_internalServiceRequestImpl_Unsubscribe(t *testing.T) {

	type fields struct {
		c service.Connector
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
			name: "when the parameters are correct",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/request/DESK-1/notification",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/request/DESK-1/notification",
					"",
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/request/DESK-1/notification",
					"",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewRequestService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResponse, err := smService.Unsubscribe(testCase.args.ctx, testCase.args.issueKeyOrID)

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

func Test_internalServiceRequestImpl_Transition(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                 context.Context
		issueKeyOrID, transitionID, comment string
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				transitionID: "299991",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/transition",
					"",
					map[string]interface{}{"id": "299991"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the comment is provided",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				transitionID: "299991",
				comment:      "Hello there!",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/transition",
					"",
					map[string]interface{}{"additionalComment": map[string]interface{}{"body": "Hello there!"}, "id": "299991"}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				transitionID: "299991",
				comment:      "Hello there!",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/transition",
					"",
					map[string]interface{}{"additionalComment": map[string]interface{}{"body": "Hello there!"}, "id": "299991"}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				transitionID: "299991",
				comment:      "Hello there!",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/transition",
					"",
					map[string]interface{}{"additionalComment": map[string]interface{}{"body": "Hello there!"}, "id": "299991"}).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name: "when the transition id is not provided",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			wantErr: true,
			Err:     model.ErrNoTransitionID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewRequestService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResponse, err := smService.Transition(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.transitionID,
				testCase.args.comment)

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

func Test_internalServiceRequestImpl_Create(t *testing.T) {

	payloadMocked := &model.CreateCustomerRequestPayloadScheme{
		RequestParticipants: []string{"uuid-sample-1", "uuid-sample-2"},
		ServiceDeskID:       "29990",
		RequestTypeID:       "28881",
	}

	// Append the request form fields
	if err := payloadMocked.AddCustomField("summary", "Request JSD help via REST"); err != nil {
		t.Fatal(err)
	}

	if err := payloadMocked.AddCustomField("description", "Request JSD help via REST"); err != nil {
		t.Fatal(err)
	}

	if err := payloadMocked.RadioButtonOrSelectCustomField("priority", "Major"); err != nil {
		t.Fatal(err)
	}

	if err := payloadMocked.UsersCustomField("customfield_320239", []string{"account-id-sample", "account-id-sample"}); err != nil {
		t.Fatal(err)
	}

	if err := payloadMocked.AddCustomField("labels", []string{"label-00", "label-01"}); err != nil {
		t.Fatal(err)
	}

	if err := payloadMocked.AddCustomField("labels", []string{"label-00", "label-01"}); err != nil {
		t.Fatal(err)
	}

	if err := payloadMocked.CascadingCustomField("customfield_10002", "America", "Costa Rica"); err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		payload *model.CreateCustomerRequestPayloadScheme
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
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerRequestScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request",
					"",
					payloadMocked).
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

			smService, err := NewRequestService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Create(testCase.args.ctx, testCase.args.payload)

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
