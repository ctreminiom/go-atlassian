package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalTemplateImpl_Create(t *testing.T) {
	type fields struct {
		c service.Connector
	}
	type args struct {
		ctx     context.Context
		payload *models.CreateTemplateScheme
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
				payload: &models.CreateTemplateScheme{
					Name:         "Test Template",
					TemplateType: "page",
					Body: &models.ContentTemplateBodyCreateScheme{
						View: &models.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &models.SpaceScheme{
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
					&models.CreateTemplateScheme{
						Name:         "Test Template",
						TemplateType: "page",
						Body: &models.ContentTemplateBodyCreateScheme{
							View: &models.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &models.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
		{
			name: "CreateTemplateScheme.Create success",
			args: args{
				ctx: context.Background(),
				payload: &models.CreateTemplateScheme{
					Name:         "Test Template",
					TemplateType: "page",
					Body: &models.ContentTemplateBodyCreateScheme{
						View: &models.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &models.SpaceScheme{
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
					&models.CreateTemplateScheme{
						Name:         "Test Template",
						TemplateType: "page",
						Body: &models.ContentTemplateBodyCreateScheme{
							View: &models.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &models.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, nil)

				client.On("Call", &http.Request{}, &models.ContentTemplateScheme{}).
					Return(&models.ResponseScheme{Code: 200}, nil)

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
				assert.EqualError(t, err, testCase.Err.Error())
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
		payload *models.UpdateTemplateScheme
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
				payload: &models.UpdateTemplateScheme{
					TemplateID:   "1234567",
					Name:         "Test Template",
					TemplateType: "page",
					Body: &models.ContentTemplateBodyCreateScheme{
						View: &models.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &models.SpaceScheme{
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
					&models.UpdateTemplateScheme{
						TemplateID:   "1234567",
						Name:         "Test Template",
						TemplateType: "page",
						Body: &models.ContentTemplateBodyCreateScheme{
							View: &models.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &models.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
		{
			name: "CreateTemplateScheme.Create success",
			args: args{
				ctx: context.Background(),
				payload: &models.UpdateTemplateScheme{
					TemplateID:   "123456789",
					Name:         "Test Template",
					TemplateType: "page",
					Body: &models.ContentTemplateBodyCreateScheme{
						View: &models.ContentBodyCreateScheme{
							Value:          "<h1>Test Template</h1>",
							Representation: "storage",
						},
					},
					Description: "This is a test template",
					Space: &models.SpaceScheme{
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
					&models.UpdateTemplateScheme{
						TemplateID:   "123456789",
						Name:         "Test Template",
						TemplateType: "page",
						Body: &models.ContentTemplateBodyCreateScheme{
							View: &models.ContentBodyCreateScheme{
								Value:          "<h1>Test Template</h1>",
								Representation: "storage",
							},
						},
						Description: "This is a test template",
						Space: &models.SpaceScheme{
							Key: "TEST",
						},
					},
				).
					Return(&http.Request{}, nil)

				client.On("Call", &http.Request{}, &models.ContentTemplateScheme{}).
					Return(&models.ResponseScheme{Code: 200}, nil)

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
				assert.EqualError(t, err, testCase.Err.Error())
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
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
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

				client.On("Call", &http.Request{}, &models.ContentTemplateScheme{}).
					Return(&models.ResponseScheme{Code: 200}, nil)

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
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}
