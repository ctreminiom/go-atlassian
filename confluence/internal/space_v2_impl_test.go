package internal

import (
	"context"
	"encoding/json"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_internalSpaceV2Impl_Bulk(t *testing.T) {

	optionsMocked := &model.GetSpacesOptionSchemeV2{
		IDs:               []string{"10001", "10002"},
		Keys:              []string{"DUMMY"},
		Type:              "global",
		Status:            "current",
		Labels:            []string{"test-label"},
		Sort:              "-name",
		DescriptionFormat: "view",
		SerializeIDs:      true,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		options *model.GetSpacesOptionSchemeV2
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
				ctx:     context.Background(),
				options: optionsMocked,
				cursor:  "cursor_sample_uuid",
				limit:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces?cursor=cursor_sample_uuid&description-format=view&ids=10001%2C10002&keys=DUMMY&labels=test-label&limit=50&serialize-ids-as-strings=true&sort=-name&status=current&type=global",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceChunkV2Scheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.Background(),
				options: optionsMocked,
				cursor:  "cursor_sample_uuid",
				limit:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces?cursor=cursor_sample_uuid&description-format=view&ids=10001%2C10002&keys=DUMMY&labels=test-label&limit=50&serialize-ids-as-strings=true&sort=-name&status=current&type=global",
					"", nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceV2Service(testCase.fields.c)

			gotResult, gotResponse, err := newService.Bulk(testCase.args.ctx, testCase.args.options, testCase.args.cursor,
				testCase.args.limit)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalSpaceV2Impl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx               context.Context
		spaceID           int
		descriptionFormat string
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
				ctx:               context.Background(),
				spaceID:           10001,
				descriptionFormat: "view",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/10001?description-format=view",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the space id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:               context.Background(),
				spaceID:           10001,
				descriptionFormat: "view",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/10001?description-format=view",
					"", nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceV2Service(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.spaceID, testCase.args.descriptionFormat)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalSpaceV2Impl_Permissions(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		spaceID int
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
				ctx:     context.Background(),
				spaceID: 10001,
				cursor:  "cursor_sample_uuid",
				limit:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/10001/permissions?cursor=cursor_sample_uuid&limit=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpacePermissionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.Background(),
				spaceID: 10001,
				cursor:  "cursor_sample_uuid",
				limit:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/10001/permissions?cursor=cursor_sample_uuid&limit=50",
					"", nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceV2Service(testCase.fields.c)

			gotResult, gotResponse, err := newService.Permissions(testCase.args.ctx, testCase.args.spaceID, testCase.args.cursor,
				testCase.args.limit)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}
