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

func Test_internalCustomContentServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		type_   string
		options *model.CustomContentOptionsScheme
		cursor  string
		limit   int
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
				ctx:   context.TODO(),
				type_: "custom_content_type_id",
				options: &model.CustomContentOptionsScheme{
					IDs:        []int{101, 102},
					SpaceIDs:   []int{10001, 10002},
					Sort:       "created-date",
					BodyFormat: "atlas_doc_format",
				},
				cursor: "cursor-id",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/custom-content?body-format=atlas_doc_format&cursor=cursor-id&id=101%2C102&limit=200&sort=created-date&space-id=101%2C102",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:   context.TODO(),
				type_: "custom_content_type_id",
				options: &model.CustomContentOptionsScheme{
					IDs:        []int{101, 102},
					SpaceIDs:   []int{10001, 10002},
					Sort:       "created-date",
					BodyFormat: "atlas_doc_format",
				},
				cursor: "cursor-id",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/custom-content?body-format=atlas_doc_format&cursor=cursor-id&id=101%2C102&limit=200&sort=created-date&space-id=101%2C102",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the type is not provided are correct",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoCustomContentTypeError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewCustomContentService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Gets(
				testCase.args.ctx,
				testCase.args.type_,
				testCase.args.options,
				testCase.args.cursor,
				testCase.args.limit,
			)

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

func Test_internalCustomContentServiceImpl_Create(t *testing.T) {

	payloadMocked := &model.CustomContentPayloadScheme{
		Type:            "",
		Status:          "",
		PageID:          "",
		CustomContentID: "",
		Title:           "",
		Body: &model.BodyTypeScheme{
			Representation: "",
			Value:          "",
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		payload *model.CustomContentPayloadScheme
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/api/v2/custom-content",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomContentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/api/v2/custom-content",
					"", payloadMocked).
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

			newService := NewCustomContentService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Create(
				testCase.args.ctx,
				testCase.args.payload,
			)

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

func Test_internalCustomContentServiceImpl_Update(t *testing.T) {

	payloadMocked := &model.CustomContentPayloadScheme{
		Type:            "",
		Status:          "",
		PageID:          "",
		CustomContentID: "",
		Title:           "",
		Body: &model.BodyTypeScheme{
			Representation: "",
			Value:          "",
		},
		Version: &model.CustomContentPayloadVersionScheme{
			Number:  1,
			Message: "message sample",
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx             context.Context
		customContentID int
		payload         *model.CustomContentPayloadScheme
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
				ctx:             context.TODO(),
				customContentID: 10001,
				payload:         payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/api/v2/custom-content/10001",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomContentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:             context.TODO(),
				customContentID: 10001,
				payload:         payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/api/v2/custom-content/10001",
					"", payloadMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the custom content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoCustomContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewCustomContentService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Update(
				testCase.args.ctx,
				testCase.args.customContentID,
				testCase.args.payload,
			)

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

func Test_internalCustomContentServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx             context.Context
		customContentID int
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
				ctx:             context.TODO(),
				customContentID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/api/v2/custom-content/10001",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:             context.TODO(),
				customContentID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/api/v2/custom-content/10001",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the custom content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoCustomContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewCustomContentService(testCase.fields.c)

			gotResponse, err := newService.Delete(
				testCase.args.ctx,
				testCase.args.customContentID)

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

func Test_internalCustomContentServiceImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx             context.Context
		customContentID int
		format          string
		versionID       int
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
				ctx:             context.TODO(),
				customContentID: 10001,
				format:          "anonymous_export_view",
				versionID:       2,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/custom-content/10001?body-format=anonymous_export_view&version=2",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomContentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:             context.TODO(),
				customContentID: 10001,
				format:          "anonymous_export_view",
				versionID:       2,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/custom-content/10001?body-format=anonymous_export_view&version=2",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the custom content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoCustomContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewCustomContentService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(
				testCase.args.ctx,
				testCase.args.customContentID,
				testCase.args.format,
				testCase.args.versionID,
			)

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
