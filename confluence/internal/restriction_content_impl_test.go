package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalRestrictionImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx        context.Context
		contentID  string
		expand     []string
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				contentID:  "100001",
				expand:     []string{"restrictions.user"},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user&limit=50&start=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentRestrictionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:        context.Background(),
				contentID:  "100001",
				expand:     []string{"restrictions.user"},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user&limit=50&start=50",
					"", nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.contentID, testCase.args.expand,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalRestrictionImpl_Add(t *testing.T) {

	payloadMocked := &model.ContentRestrictionUpdatePayloadScheme{}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		contentID string
		payload   *model.ContentRestrictionUpdatePayloadScheme
		expand    []string
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
				ctx:       context.Background(),
				contentID: "100001",
				payload:   payloadMocked,
				expand:    []string{"restrictions.user"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentRestrictionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				contentID: "100001",
				payload:   payloadMocked,
				expand:    []string{"restrictions.user"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user",
					"", payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.contentID, testCase.args.payload,
				testCase.args.expand)

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

func Test_internalRestrictionImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		contentID string
		expand    []string
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
				ctx:       context.Background(),
				contentID: "100001",
				expand:    []string{"restrictions.user"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentRestrictionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				contentID: "100001",
				expand:    []string{"restrictions.user"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user",
					"", nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.contentID, testCase.args.expand)

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

func Test_internalRestrictionImpl_Update(t *testing.T) {

	payloadMocked := &model.ContentRestrictionUpdatePayloadScheme{}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		contentID string
		payload   *model.ContentRestrictionUpdatePayloadScheme
		expand    []string
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
				ctx:       context.Background(),
				contentID: "100001",
				payload:   payloadMocked,
				expand:    []string{"restrictions.user"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentRestrictionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				contentID: "100001",
				payload:   payloadMocked,
				expand:    []string{"restrictions.user"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100001/restriction?expand=restrictions.user",
					"", payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.contentID, testCase.args.payload,
				testCase.args.expand)

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
