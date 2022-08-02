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

func Test_internalIssueFieldContextOptionServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldId             string
		contextId           int
		options             *model.FieldOptionContextParams
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/10001/option?maxResults=50&onlyOptions=false&optionId=3022&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextOptionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/10001/option?maxResults=50&onlyOptions=false&optionId=3022&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextOptionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "",
				contextId: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/10001/option?maxResults=50&onlyOptions=false&optionId=3022&startAt=50",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Gets(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.options, testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalIssueFieldContextOptionServiceImpl_Create(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx       context.Context
		fieldId   string
		contextId int
		payload   *model.FieldContextOptionListScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.FieldContextOptionListScheme{
					Options: []*model.CustomFieldContextOptionScheme{

						// Single/Multiple Choice example
						{
							Value:    "Option 2",
							Disabled: false,
						},
						{
							Value:    "Option 4",
							Disabled: false,
						},
					}},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextOptionListScheme{
						Options: []*model.CustomFieldContextOptionScheme{

							// Single/Multiple Choice example
							{
								Value:    "Option 2",
								Disabled: false,
							},
							{
								Value:    "Option 4",
								Disabled: false,
							},
						}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.FieldContextOptionListScheme{
					Options: []*model.CustomFieldContextOptionScheme{

						// Single/Multiple Choice example
						{
							Value:    "Option 2",
							Disabled: false,
						},
						{
							Value:    "Option 4",
							Disabled: false,
						},
					}},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextOptionListScheme{
						Options: []*model.CustomFieldContextOptionScheme{

							// Single/Multiple Choice example
							{
								Value:    "Option 2",
								Disabled: false,
							},
							{
								Value:    "Option 4",
								Disabled: false,
							},
						}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/option",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "",
				contextId: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.FieldContextOptionListScheme{
					Options: []*model.CustomFieldContextOptionScheme{

						// Single/Multiple Choice example
						{
							Value:    "Option 2",
							Disabled: false,
						},
						{
							Value:    "Option 4",
							Disabled: false,
						},
					}},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextOptionListScheme{
						Options: []*model.CustomFieldContextOptionScheme{

							// Single/Multiple Choice example
							{
								Value:    "Option 2",
								Disabled: false,
							},
							{
								Value:    "Option 4",
								Disabled: false,
							},
						}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Create(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
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

func Test_internalIssueFieldContextOptionServiceImpl_Update(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx       context.Context
		fieldId   string
		contextId int
		payload   *model.FieldContextOptionListScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.FieldContextOptionListScheme{
					Options: []*model.CustomFieldContextOptionScheme{

						// Single/Multiple Choice example
						{
							Value:    "Option 2",
							Disabled: false,
						},
						{
							Value:    "Option 4",
							Disabled: false,
						},
					}},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextOptionListScheme{
						Options: []*model.CustomFieldContextOptionScheme{

							// Single/Multiple Choice example
							{
								Value:    "Option 2",
								Disabled: false,
							},
							{
								Value:    "Option 4",
								Disabled: false,
							},
						}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.FieldContextOptionListScheme{
					Options: []*model.CustomFieldContextOptionScheme{

						// Single/Multiple Choice example
						{
							Value:    "Option 2",
							Disabled: false,
						},
						{
							Value:    "Option 4",
							Disabled: false,
						},
					}},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextOptionListScheme{
						Options: []*model.CustomFieldContextOptionScheme{

							// Single/Multiple Choice example
							{
								Value:    "Option 2",
								Disabled: false,
							},
							{
								Value:    "Option 4",
								Disabled: false,
							},
						}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/option",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "",
				contextId: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.FieldContextOptionListScheme{
					Options: []*model.CustomFieldContextOptionScheme{

						// Single/Multiple Choice example
						{
							Value:    "Option 2",
							Disabled: false,
						},
						{
							Value:    "Option 4",
							Disabled: false,
						},
					}},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextOptionListScheme{
						Options: []*model.CustomFieldContextOptionScheme{

							// Single/Multiple Choice example
							{
								Value:    "Option 2",
								Disabled: false,
							},
							{
								Value:    "Option 4",
								Disabled: false,
							},
						}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Update(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
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

func Test_internalIssueFieldContextOptionServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx       context.Context
		fieldId   string
		contextId int
		optionId  int
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				optionId:  1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/custom_field_10002/context/10001/option/1001",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				optionId:  1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/custom_field_10002/context/10001/option/1001",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "",
				contextId: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				optionId:  1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/custom_field_10002/context/10001/option/1001",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Delete(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.optionId)

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

func Test_internalIssueFieldContextOptionServiceImpl_Order(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx       context.Context
		fieldId   string
		contextId int
		payload   *model.OrderFieldOptionPayloadScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.OrderFieldOptionPayloadScheme{
					Position:             "Last",
					CustomFieldOptionIds: []string{"111"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.OrderFieldOptionPayloadScheme{
						Position:             "Last",
						CustomFieldOptionIds: []string{"111"},
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option/move",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.OrderFieldOptionPayloadScheme{
					Position:             "Last",
					CustomFieldOptionIds: []string{"111"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.OrderFieldOptionPayloadScheme{
						Position:             "Last",
						CustomFieldOptionIds: []string{"111"},
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/option/move",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "",
				contextId: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
				payload: &model.OrderFieldOptionPayloadScheme{
					Position:             "Last",
					CustomFieldOptionIds: []string{"111"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.OrderFieldOptionPayloadScheme{
						Position:             "Last",
						CustomFieldOptionIds: []string{"111"},
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option/move",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Order(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.payload)

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
