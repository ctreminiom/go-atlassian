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

func Test_internalChildrenDescandantsImpl_Children(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		contentID     string
		expand        []string
		parentVersion int
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
				ctx:           context.Background(),
				contentID:     "100100101",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child?expand=attachment%2Ccomments&parentVersion=12",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentChildrenScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:           context.Background(),
				contentID:     "100100101",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child?expand=attachment%2Ccomments&parentVersion=12",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
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

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Children(testCase.args.ctx, testCase.args.contentID, testCase.args.expand,
				testCase.args.parentVersion)

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

func Test_internalChildrenDescandantsImpl_Move(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                        context.Context
		pageID, position, targetID string
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
			name: "append when the parameters are correct",
			args: args{
				ctx:      context.Background(),
				pageID:   "101010101",
				position: "append",
				targetID: "202020202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/101010101/move/append/202020202",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentMoveScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "position before when the parameters are correct",
			args: args{
				ctx:      context.Background(),
				pageID:   "101010101",
				position: "before",
				targetID: "202020202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/101010101/move/before/202020202",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentMoveScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "position after when the parameters are correct",
			args: args{
				ctx:      context.Background(),
				pageID:   "101010101",
				position: "after",
				targetID: "202020202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/101010101/move/after/202020202",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentMoveScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.Background(),
				pageID:   "100100101",
				position: "append",
				targetID: "200200202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100100101/move/append/200200202",
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
				ctx:      context.Background(),
				position: "append",
				targetID: "100100101",
			},
			wantErr: true,
			Err:     model.ErrNoPageID,
		},

		{
			name: "when the target id is not provided",
			args: args{
				ctx:      context.Background(),
				pageID:   "100100101",
				position: "append",
			},
			wantErr: true,
			Err:     model.ErrNoTargetID,
		},

		{
			name: "when the position is not provided",
			args: args{
				ctx:      context.Background(),
				pageID:   "100100101",
				targetID: "200200202",
			},
			wantErr: true,
			Err:     model.ErrNoPosition,
		},

		{
			name: "when the position is incorrect",
			args: args{
				ctx:      context.Background(),
				pageID:   "100100101",
				position: "gopher",
				targetID: "200200202",
			},
			wantErr: true,
			Err:     model.ErrInvalidPosition,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Move(testCase.args.ctx, testCase.args.pageID, testCase.args.position, testCase.args.targetID)

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

func Test_internalChildrenDescandantsImpl_ChildrenByType(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                    context.Context
		contentID, contentType string
		expand                 []string
		parentVersion          int
		startAt, maxResults    int
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
				ctx:           context.Background(),
				contentID:     "100100101",
				contentType:   "blogpost",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
				startAt:       50,
				maxResults:    25,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child/blogpost?expand=attachment%2Ccomments&limit=25&parentVersion=12&start=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:           context.Background(),
				contentID:     "100100101",
				contentType:   "blogpost",
				expand:        []string{"attachment", "comments"},
				parentVersion: 12,
				startAt:       50,
				maxResults:    25,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child/blogpost?expand=attachment%2Ccomments&limit=25&parentVersion=12&start=50",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},

		{
			name: "when the content type is not provided",
			args: args{
				ctx:       context.Background(),
				contentID: "11929292",
			},
			wantErr: true,
			Err:     model.ErrNoContentType,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.ChildrenByType(testCase.args.ctx, testCase.args.contentID,
				testCase.args.contentType, testCase.args.parentVersion, testCase.args.expand, testCase.args.startAt,
				testCase.args.maxResults)

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

func Test_internalChildrenDescandantsImpl_Descendants(t *testing.T) {

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
				contentID: "100100101",
				expand:    []string{"attachment", "comments"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant?expand=attachment%2Ccomments",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentChildrenScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				contentID: "100100101",
				expand:    []string{"attachment", "comments"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant?expand=attachment%2Ccomments",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
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

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Descendants(testCase.args.ctx, testCase.args.contentID, testCase.args.expand)

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

func Test_internalChildrenDescandantsImpl_DescendantsByType(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                           context.Context
		contentID, contentType, depth string
		expand                        []string
		startAt, maxResults           int
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
				ctx:         context.Background(),
				contentID:   "100100101",
				contentType: "blogpost",
				expand:      []string{"attachment", "comments"},
				startAt:     50,
				maxResults:  25,
				depth:       "root",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant/blogpost?depth=root&expand=attachment%2Ccomments&limit=25&start=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				contentID:   "100100101",
				contentType: "blogpost",
				expand:      []string{"attachment", "comments"},
				startAt:     50,
				maxResults:  25,
				depth:       "root",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/descendant/blogpost?depth=root&expand=attachment%2Ccomments&limit=25&start=50",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},

		{
			name: "when the content type is not provided",
			args: args{
				ctx:       context.Background(),
				contentID: "11929292",
			},
			wantErr: true,
			Err:     model.ErrNoContentType,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.DescendantsByType(testCase.args.ctx, testCase.args.contentID,
				testCase.args.contentType, testCase.args.depth, testCase.args.expand, testCase.args.startAt,
				testCase.args.maxResults)

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

func Test_internalChildrenDescandantsImpl_CopyHierarchy(t *testing.T) {

	payloadMocked := &model.CopyOptionsScheme{
		CopyAttachments:    true,
		CopyPermissions:    true,
		CopyProperties:     true,
		CopyLabels:         true,
		CopyCustomContents: true,
		DestinationPageID:  "223322",
		TitleOptions: &model.CopyTitleOptionScheme{
			Prefix:  "copy-",
			Replace: "test",
		},
		PageTitle: "Test Title",
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		contentID string
		options   *model.CopyOptionsScheme
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
				contentID: "100100101",
				options:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/pagehierarchy/copy",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.TaskScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				contentID: "100100101",
				options:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/pagehierarchy/copy",
					"", payloadMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
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

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.CopyHierarchy(testCase.args.ctx, testCase.args.contentID,
				testCase.args.options)

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

func Test_internalChildrenDescandantsImpl_CopyPage(t *testing.T) {

	payloadMocked := &model.CopyOptionsScheme{
		CopyAttachments:    true,
		CopyPermissions:    true,
		CopyProperties:     true,
		CopyLabels:         true,
		CopyCustomContents: true,
		DestinationPageID:  "223322",
		TitleOptions: &model.CopyTitleOptionScheme{
			Prefix:  "copy-",
			Replace: "test",
		},
		PageTitle: "Test Title",
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		contentID string
		expand    []string
		options   *model.CopyOptionsScheme
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
				contentID: "100100101",
				options:   payloadMocked,
				expand:    []string{"childTypes.all"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/copy?expand=childTypes.all",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				contentID: "100100101",
				options:   payloadMocked,
				expand:    []string{"childTypes.all"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/100100101/copy?expand=childTypes.all",
					"", payloadMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
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

			newService := NewChildrenDescandantsService(testCase.fields.c)

			gotResult, gotResponse, err := newService.CopyPage(testCase.args.ctx, testCase.args.contentID,
				testCase.args.expand, testCase.args.options)

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
