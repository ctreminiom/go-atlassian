package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func Test_internalPageImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		pageID  int
		format  string
		draft   bool
		version int
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
				pageID:  200001,
				format:  "atlas_doc_format",
				draft:   true,
				version: 2,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/200001?body-format=atlas_doc_format&get-draft=true&version=2",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				pageID:  200001,
				format:  "atlas_doc_format",
				draft:   true,
				version: 2,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/200001?body-format=atlas_doc_format&get-draft=true&version=2",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the page id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoPageIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.pageID, testCase.args.format,
				testCase.args.draft, testCase.args.version)

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

func Test_internalPageImpl_Bulk(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx    context.Context
		cursor string
		limit  int
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
				cursor: "cursor-sample",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages?cursor=cursor-sample&limit=200",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:    context.TODO(),
				cursor: "cursor-sample",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages?cursor=cursor-sample&limit=200",
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

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Bulk(testCase.args.ctx, testCase.args.cursor, testCase.args.limit)

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

func Test_internalPageImpl_GetsByLabel(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		labelID int
		sort    string
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
				ctx:     context.TODO(),
				labelID: 20001,
				sort:    "test-label",
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/labels/20001/pages?cursor=cursor-sample&limit=200&sort=test-label",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the label id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoLabelIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				labelID: 20001,
				sort:    "test-label",
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/labels/20001/pages?cursor=cursor-sample&limit=200&sort=test-label",
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

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.GetsByLabel(testCase.args.ctx, testCase.args.labelID,
				testCase.args.sort, testCase.args.cursor, testCase.args.limit)

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

func Test_internalPageImpl_GetsBySpace(t *testing.T) {

	type fields struct {
		c service.Client
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
				ctx:     context.TODO(),
				spaceID: 20001,
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/20001/pages?cursor=cursor-sample&limit=200",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the space id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				spaceID: 20001,
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/20001/pages?cursor=cursor-sample&limit=200",
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

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.GetsBySpace(testCase.args.ctx, testCase.args.spaceID,
				testCase.args.cursor, testCase.args.limit)

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

func Test_internalPageImpl_GetsByParent(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		parentID int
		cursor   string
		limit    int
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
				ctx:      context.TODO(),
				parentID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/20001/children?cursor=cursor-sample&limit=200",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the parent id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoPageIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.TODO(),
				parentID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/20001/children?cursor=cursor-sample&limit=200",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the call fails",
			args: args{
				ctx:      context.TODO(),
				parentID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/20001/children?cursor=cursor-sample&limit=200",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageChunkScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, request failed"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, request failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.GetsByParent(testCase.args.ctx, testCase.args.parentID,
				testCase.args.cursor, testCase.args.limit)

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

func Test_internalPageImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx    context.Context
		pageID int
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
				pageID: 200001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/api/v2/pages/200001",
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
				ctx:    context.TODO(),
				pageID: 200001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/api/v2/pages/200001",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the page id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoPageIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewPageService(testCase.fields.c)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.pageID)

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

func Test_internalPageImpl_Create(t *testing.T) {

	//Create the ADF body
	mockedBody := model.CommentNodeScheme{}
	mockedBody.Version = 1
	mockedBody.Type = "doc"

	mockedBody.AppendNode(&model.CommentNodeScheme{
		Type: "paragraph",
		Content: []*model.CommentNodeScheme{
			{
				Type: "status",
				Attrs: map[string]interface{}{
					"color":   "neutral",
					"style":   "bold",
					"text":    "Status neutral status",
					"localId": "3378f5db-c151-4ef2-aec4-56c18a1fbd96",
				},
			},
		},
	})

	mockedBodyValue, err := json.Marshal(&mockedBody)
	if err != nil {
		log.Fatal(err)
	}

	mockedPayload := &model.PageCreatePayloadScheme{
		SpaceID: 203718658,
		Status:  "current",
		Title:   "Page create title test",
		Body: &model.PageBodyRepresentationScheme{
			Representation: "atlas_doc_format",
			Value:          string(mockedBodyValue),
		},
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		payload *model.PageCreatePayloadScheme
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
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					mockedPayload).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/api/v2/pages",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					mockedPayload).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/api/v2/pages",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the payload is not provided",
			args: args{
				ctx: context.TODO(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.PageCreatePayloadScheme)(nil)).
					Return(nil, model.ErrNilPayloadError)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrNilPayloadError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_internalPageImpl_Update(t *testing.T) {

	//Create the ADF body
	mockedBody := model.CommentNodeScheme{}
	mockedBody.Version = 1
	mockedBody.Type = "doc"

	mockedBody.AppendNode(&model.CommentNodeScheme{
		Type: "paragraph",
		Content: []*model.CommentNodeScheme{
			{
				Type: "status",
				Attrs: map[string]interface{}{
					"color":   "neutral",
					"style":   "bold",
					"text":    "Status neutral status",
					"localId": "3378f5db-c151-4ef2-aec4-56c18a1fbd96",
				},
			},
		},
	})

	mockedBodyValue, err := json.Marshal(&mockedBody)
	if err != nil {
		log.Fatal(err)
	}

	mockedPayload := &model.PageUpdatePayloadScheme{
		ID:      215646235,
		SpaceID: 203718658,
		Status:  "current",
		Title:   "Page create title test",
		Body: &model.PageBodyRepresentationScheme{
			Representation: "atlas_doc_format",
			Value:          string(mockedBodyValue),
		},
		Version: &model.PageUpdatePayloadVersionScheme{
			Number:  4,
			Message: "Version #4",
		},
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		pageID  int
		payload *model.PageUpdatePayloadScheme
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
				pageID:  215646235,
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					mockedPayload).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/api/v2/pages/215646235",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.PageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				pageID:  215646235,
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					mockedPayload).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/api/v2/pages/215646235",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the page id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoPageIDError,
		},

		{
			name: "when the payload is not provided",
			args: args{
				ctx:    context.TODO(),
				pageID: 215646235,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.PageUpdatePayloadScheme)(nil)).
					Return(nil, model.ErrNilPayloadError)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrNilPayloadError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewPageService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.pageID, testCase.args.payload)

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
