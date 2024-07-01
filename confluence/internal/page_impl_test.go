package internal

import (
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
		c service.Connector
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
				ctx:     context.Background(),
				pageID:  200001,
				format:  "atlas_doc_format",
				draft:   true,
				version: 2,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/200001?body-format=atlas_doc_format&get-draft=true&version=2",
					"", nil).
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
				ctx:     context.Background(),
				pageID:  200001,
				format:  "atlas_doc_format",
				draft:   true,
				version: 2,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/200001?body-format=atlas_doc_format&get-draft=true&version=2",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the page id is not provided",
			args: args{
				ctx: context.Background(),
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

func Test_internalPageImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		options *model.PageOptionsScheme
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
				ctx: context.Background(),
				options: &model.PageOptionsScheme{
					PageIDs:    []int{112, 1223},
					SpaceIDs:   []int{3040, 3040},
					Sort:       "-created-date",
					Status:     []string{"current", "trashed"},
					Title:      "Title sample!!",
					BodyFormat: "atlas_doc_format",
				},
				cursor: "cursor-sample",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages?body-format=atlas_doc_format&cursor=cursor-sample&id=112%2C1223&limit=200&sort=-created-date&space-id=3040%2C3040&status=current%2Ctrashed&title=Title+sample%21%21",
					"", nil).
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
				ctx:    context.Background(),
				cursor: "cursor-sample",
				options: &model.PageOptionsScheme{
					PageIDs:    []int{112, 1223},
					SpaceIDs:   []int{3040, 3040},
					Sort:       "-created-date",
					Status:     []string{"current", "trashed"},
					Title:      "Title sample!!",
					BodyFormat: "atlas_doc_format",
				},
				limit: 200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages?body-format=atlas_doc_format&cursor=cursor-sample&id=112%2C1223&limit=200&sort=-created-date&space-id=3040%2C3040&status=current%2Ctrashed&title=Title+sample%21%21",
					"", nil).
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

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.options, testCase.args.cursor, testCase.args.limit)

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
		c service.Connector
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
				ctx:    context.Background(),
				cursor: "cursor-sample",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages?cursor=cursor-sample&limit=200",
					"", nil).
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
				ctx:    context.Background(),
				cursor: "cursor-sample",
				limit:  200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages?cursor=cursor-sample&limit=200",
					"", nil).
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
		c service.Connector
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
				ctx:     context.Background(),
				labelID: 20001,
				sort:    "test-label",
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/labels/20001/pages?cursor=cursor-sample&limit=200&sort=test-label",
					"", nil).
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoLabelIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.Background(),
				labelID: 20001,
				sort:    "test-label",
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/labels/20001/pages?cursor=cursor-sample&limit=200&sort=test-label",
					"", nil).
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
				spaceID: 20001,
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/20001/pages?cursor=cursor-sample&limit=200",
					"", nil).
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.Background(),
				spaceID: 20001,
				cursor:  "cursor-sample",
				limit:   200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/20001/pages?cursor=cursor-sample&limit=200",
					"", nil).
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
		c service.Connector
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
				ctx:      context.Background(),
				parentID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/20001/children?cursor=cursor-sample&limit=200",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ChildPageChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the parent id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoPageIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.Background(),
				parentID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/20001/children?cursor=cursor-sample&limit=200",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the call fails",
			args: args{
				ctx:      context.Background(),
				parentID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/pages/20001/children?cursor=cursor-sample&limit=200",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ChildPageChunkScheme{}).
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
		c service.Connector
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
				ctx:    context.Background(),
				pageID: 200001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/api/v2/pages/200001",
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
				ctx:    context.Background(),
				pageID: 200001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/api/v2/pages/200001",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the page id is not provided",
			args: args{
				ctx: context.Background(),
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
		SpaceID: "203718658",
		Status:  "current",
		Title:   "Page create title test",
		Body: &model.PageBodyRepresentationScheme{
			Representation: "atlas_doc_format",
			Value:          string(mockedBodyValue),
		},
	}

	type fields struct {
		c service.Connector
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
				ctx:     context.Background(),
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/api/v2/pages",
					"", mockedPayload).
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
				ctx:     context.Background(),
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/api/v2/pages",
					"", mockedPayload).
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
		c service.Connector
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
				ctx:     context.Background(),
				pageID:  215646235,
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/api/v2/pages/215646235",
					"", mockedPayload).
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
				ctx:     context.Background(),
				pageID:  215646235,
				payload: mockedPayload,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/api/v2/pages/215646235",
					"", mockedPayload).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the page id is not provided",
			args: args{
				ctx: context.Background(),
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
