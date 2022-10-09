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

func Test_internalContentLabelImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                 context.Context
		contentID           string
		prefix              string
		startAt, maxResults int
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
				ctx:        context.TODO(),
				contentID:  "11727271",
				prefix:     "new-",
				startAt:    25,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/11727271/label?limit=50&prefix=new-&start=25",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentLabelPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:        context.TODO(),
				contentID:  "11727271",
				prefix:     "new-",
				startAt:    25,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/11727271/label?limit=50&prefix=new-&start=25",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewContentLabelService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.contentID, testCase.args.prefix,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalContentLabelImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                  context.Context
		contentID, labelName string
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
				ctx:       context.TODO(),
				contentID: "11727271",
				labelName: "test",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/11727271/label/test",
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
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				contentID: "11727271",
				labelName: "test",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/11727271/label/test",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},

		{
			name: "when the label name is not provided",
			args: args{
				ctx:       context.TODO(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentLabelError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewContentLabelService(testCase.fields.c)

			gotResponse, err := newService.Remove(testCase.args.ctx, testCase.args.contentID, testCase.args.labelName)

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

func Test_internalContentLabelImpl_Add(t *testing.T) {

	payloadMocked := []*model.ContentLabelPayloadScheme{
		{
			Prefix: "global",
			Name:   "label-02",
		},
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx             context.Context
		contentID       string
		payload         []*model.ContentLabelPayloadScheme
		want400Response bool
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
				contentID:       "11727271",
				payload:         payloadMocked,
				want400Response: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/11727271/label?use-400-error-response=true",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentLabelPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:             context.TODO(),
				contentID:       "11727271",
				payload:         payloadMocked,
				want400Response: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/11727271/label?use-400-error-response=true",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewContentLabelService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.contentID, testCase.args.payload,
				testCase.args.want400Response)

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
