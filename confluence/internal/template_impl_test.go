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

func Test_internalTemplateImpl_Create(t *testing.T) {
	type fields struct {
		c service.Connector
	}
	type args struct {
		ctx     context.Context
		payload *model.CreateTemplateScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "CreateTemplateScheme.Create returned error",
			args: args{
				ctx: context.Background(),
				payload: &model.CreateTemplateScheme{
					Name:         "Test Template",
					TemplateType: "page",
					Body: &model.ContentTemplateBodyCreateScheme{
						View: &model.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &model.SpaceScheme{
						Key: "TEST",
					},
				},
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/wiki/rest/api/template",
					"",
					&model.CreateTemplateScheme{
						Name:         "Test Template",
						TemplateType: "page",
						Body: &model.ContentTemplateBodyCreateScheme{
							View: &model.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &model.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
		{
			name: "CreateTemplateScheme.Create success",
			args: args{
				ctx: context.Background(),
				payload: &model.CreateTemplateScheme{
					Name:         "Test Template",
					TemplateType: "page",
					Body: &model.ContentTemplateBodyCreateScheme{
						View: &model.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &model.SpaceScheme{
						Key: "TEST",
					},
				},
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/wiki/rest/api/template",
					"",
					&model.CreateTemplateScheme{
						Name:         "Test Template",
						TemplateType: "page",
						Body: &model.ContentTemplateBodyCreateScheme{
							View: &model.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &model.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, nil)

				client.On("Call", &http.Request{}, &model.ContentTemplateScheme{}).
					Return(&model.ResponseScheme{Code: 200}, nil)

				fields.c = client
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewTemplateService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Create(testCase.args.ctx, testCase.args.payload)

			if testCase.wantErr {
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

func Test_internalTemplateImpl_Update(t *testing.T) {
	type fields struct {
		c service.Connector
	}
	type args struct {
		ctx     context.Context
		payload *model.UpdateTemplateScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "CreateTemplateScheme.Create returned error",
			args: args{
				ctx: context.Background(),
				payload: &model.UpdateTemplateScheme{
					TemplateID:   "1234567",
					Name:         "Test Template",
					TemplateType: "page",
					Body: &model.ContentTemplateBodyCreateScheme{
						View: &model.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &model.SpaceScheme{
						Key: "TEST",
					},
				},
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"/wiki/rest/api/template",
					"",
					&model.UpdateTemplateScheme{
						TemplateID:   "1234567",
						Name:         "Test Template",
						TemplateType: "page",
						Body: &model.ContentTemplateBodyCreateScheme{
							View: &model.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &model.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
		{
			name: "CreateTemplateScheme.Create success",
			args: args{
				ctx: context.Background(),
				payload: &model.UpdateTemplateScheme{
					TemplateID:   "123456789",
					Name:         "Test Template",
					TemplateType: "page",
					Body: &model.ContentTemplateBodyCreateScheme{
						View: &model.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &model.SpaceScheme{
						Key: "TEST",
					},
				},
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"/wiki/rest/api/template",
					"",
					&model.UpdateTemplateScheme{
						TemplateID:   "123456789",
						Name:         "Test Template",
						TemplateType: "page",
						Body: &model.ContentTemplateBodyCreateScheme{
							View: &model.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &model.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, nil)

				client.On("Call", &http.Request{}, &model.ContentTemplateScheme{}).
					Return(&model.ResponseScheme{Code: 200}, nil)

				fields.c = client
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewTemplateService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.payload)

			if testCase.wantErr {
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

func Test_internalTemplateImpl_Get(t *testing.T) {
	type fields struct {
		c service.Connector
	}
	type args struct {
		ctx        context.Context
		templateID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "CreateTemplateScheme.Create returned error",
			args: args{
				ctx:        context.Background(),
				templateID: "123456",
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/wiki/rest/api/template/123456",
					"",
					nil,
				).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
		{
			name: "CreateTemplateScheme.Create success",
			args: args{
				ctx:        context.Background(),
				templateID: "123456789",
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/wiki/rest/api/template/123456789",
					"",
					nil,
				).
					Return(&http.Request{}, nil)

				client.On("Call", &http.Request{}, &model.ContentTemplateScheme{}).
					Return(&model.ResponseScheme{Code: 200}, nil)

				fields.c = client
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewTemplateService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.templateID)

			if testCase.wantErr {
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
