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

func Test_internalSpaceImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                 context.Context
		options             *model.GetSpacesOptionScheme
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
				ctx: context.TODO(),
				options: &model.GetSpacesOptionScheme{
					SpaceKeys:       []string{"DUMMY", "TEST"},
					SpaceIDs:        []int{1111, 2222, 3333},
					SpaceType:       "global",
					Status:          "archived",
					Labels:          []string{"label-09", "label-02"},
					Favorite:        true,
					FavoriteUserKey: "DUMMY",
					Expand:          []string{"operations"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY&spaceKey=TEST&start=0&status=archived&type=global",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpacePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx: context.TODO(),
				options: &model.GetSpacesOptionScheme{
					SpaceKeys:       []string{"DUMMY", "TEST"},
					SpaceIDs:        []int{1111, 2222, 3333},
					SpaceType:       "global",
					Status:          "archived",
					Labels:          []string{"label-09", "label-02"},
					Favorite:        true,
					FavoriteUserKey: "DUMMY",
					Expand:          []string{"operations"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY&spaceKey=TEST&start=0&status=archived&type=global",
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

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.options, testCase.args.startAt,
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

func Test_internalSpaceImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		spaceKey string
		expand   []string
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
				spaceKey: "DUMMY",
				expand:   []string{"childtypes.all", "operations"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.TODO(),
				spaceKey: "DUMMY",
				expand:   []string{"childtypes.all", "operations"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the space key is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceKeyError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.spaceKey, testCase.args.expand)

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

func Test_internalSpaceImpl_Content(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                 context.Context
		spaceKey, depth     string
		expand              []string
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
				spaceKey:   "DUMMY",
				depth:      "all",
				expand:     []string{"childtypes.all", "operations"},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space/DUMMY/content?depth=all&expand=childtypes.all%2Coperations&limit=50&start=0",
					nil).
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
				ctx:        context.TODO(),
				spaceKey:   "DUMMY",
				depth:      "all",
				expand:     []string{"childtypes.all", "operations"},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space/DUMMY/content?depth=all&expand=childtypes.all%2Coperations&limit=50&start=0",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the space key is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceKeyError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Content(testCase.args.ctx, testCase.args.spaceKey, testCase.args.depth,
				testCase.args.expand, testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalSpaceImpl_ContentByType(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                          context.Context
		spaceKey, depth, contentType string
		expand                       []string
		startAt, maxResults          int
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
				ctx:         context.TODO(),
				spaceKey:    "DUMMY",
				depth:       "all",
				contentType: "page",
				expand:      []string{"childtypes.all", "operations"},
				startAt:     0,
				maxResults:  50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space/DUMMY/content/page?depth=all&expand=childtypes.all%2Coperations&limit=50&start=0",
					nil).
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
				ctx:         context.TODO(),
				spaceKey:    "DUMMY",
				depth:       "all",
				contentType: "page",
				expand:      []string{"childtypes.all", "operations"},
				startAt:     0,
				maxResults:  50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/space/DUMMY/content/page?depth=all&expand=childtypes.all%2Coperations&limit=50&start=0",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the space key is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceKeyError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.ContentByType(testCase.args.ctx, testCase.args.spaceKey,
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

func Test_internalSpaceImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		spaceKey string
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
				spaceKey: "DUMMY",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/space/DUMMY",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentTaskScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.TODO(),
				spaceKey: "DUMMY",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/space/DUMMY",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the space key is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceKeyError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.spaceKey)

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

func Test_internalSpaceImpl_Create(t *testing.T) {

	payloadMocked := &model.CreateSpaceScheme{
		Key:              "DUMMY",
		Name:             "DUMMY Space",
		Description:      nil,
		AnonymousAccess:  false,
		UnlicensedAccess: true,
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		payload *model.CreateSpaceScheme
		private bool
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
				private: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/space/_private",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
				private: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/space/_private",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the space name is not provided",
			args: args{
				ctx:     context.TODO(),
				payload: &model.CreateSpaceScheme{},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.CreateSpaceScheme{}).
					Return(bytes.NewReader([]byte{}), nil)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrNoSpaceNameError,
		},

		{
			name: "when the space key is not provided",
			args: args{
				ctx: context.TODO(),
				payload: &model.CreateSpaceScheme{
					Name: "DUMMY Space",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.CreateSpaceScheme{Name: "DUMMY Space"}).
					Return(bytes.NewReader([]byte{}), nil)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrNoSpaceKeyError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Create(testCase.args.ctx, testCase.args.payload, testCase.args.private)

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

func Test_internalSpaceImpl_Update(t *testing.T) {

	payloadMocked := &model.UpdateSpaceScheme{
		Name:        "DUMMY Space",
		Description: nil,
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		spaceKey string
		payload  *model.UpdateSpaceScheme
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
				spaceKey: "DUMMY",
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/space/DUMMY",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.TODO(),
				spaceKey: "DUMMY",
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/space/DUMMY",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the space key is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceKeyError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceService(testCase.fields.c, nil)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.spaceKey, testCase.args.payload)

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
