package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFolderScheme_UnmarshalJSON(t *testing.T) {

	ts := time.Date(2024, 9, 23, 20, 17, 35, 607000000, time.UTC)

	tests := []struct {
		name string
		json string
		check func(t *testing.T, got FolderScheme)
	}{
		{
			name: "full response with string timestamps",
			json: `{
				"id": "123456",
				"type": "folder",
				"status": "current",
				"title": "My Folder",
				"parentId": "789",
				"parentType": "page",
				"position": 5,
				"authorId": "author-1",
				"ownerId": "owner-1",
				"createdAt": "2024-09-23T20:17:35.607Z",
				"spaceId": "space-42",
				"version": {
					"createdAt": "2024-09-23T20:17:35.607Z",
					"message": "initial version",
					"number": 1,
					"minorEdit": false,
					"authorId": "author-1"
				},
				"_links": {
					"self": "https://example.atlassian.net/wiki/api/v2/folders/123456"
				}
			}`,
			check: func(t *testing.T, got FolderScheme) {
				assert.Equal(t, "123456", got.ID)
				assert.Equal(t, "folder", got.Type)
				assert.Equal(t, "current", got.Status)
				assert.Equal(t, "My Folder", got.Title)
				assert.Equal(t, "789", got.ParentID)
				assert.Equal(t, "page", got.ParentType)
				assert.Equal(t, 5, got.Position)
				assert.Equal(t, "author-1", got.AuthorID)
				assert.Equal(t, "owner-1", got.OwnerID)
				assert.Equal(t, "space-42", got.SpaceID)
				assert.True(t, ts.Equal(time.Time(*got.CreatedAt)))

				require.NotNil(t, got.Version)
				assert.True(t, ts.Equal(time.Time(*got.Version.CreatedAt)))
				assert.Equal(t, "initial version", got.Version.Message)
				assert.Equal(t, 1, got.Version.Number)
				assert.False(t, got.Version.MinorEdit)
				assert.Equal(t, "author-1", got.Version.AuthorID)

				require.NotNil(t, got.Links)
				assert.Equal(t, "https://example.atlassian.net/wiki/api/v2/folders/123456", got.Links.Self)
			},
		},
		{
			name: "numeric epoch millisecond timestamps",
			json: `{
				"id": "123456",
				"title": "Epoch Folder",
				"createdAt": 1727122655607,
				"version": {
					"number": 1,
					"createdAt": 1727122655607
				}
			}`,
			check: func(t *testing.T, got FolderScheme) {
				assert.Equal(t, "123456", got.ID)
				assert.Equal(t, "Epoch Folder", got.Title)
				assert.True(t, ts.Equal(time.Time(*got.CreatedAt)))
				assert.True(t, ts.Equal(time.Time(*got.Version.CreatedAt)))
			},
		},
		{
			name: "minimal response",
			json: `{
				"id": "123456",
				"title": "Minimal Folder"
			}`,
			check: func(t *testing.T, got FolderScheme) {
				assert.Equal(t, "123456", got.ID)
				assert.Equal(t, "Minimal Folder", got.Title)
				assert.Nil(t, got.CreatedAt)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FolderScheme
			err := json.Unmarshal([]byte(tt.json), &got)
			require.NoError(t, err)
			tt.check(t, got)
		})
	}
}

func TestFolderChunkScheme_UnmarshalJSON(t *testing.T) {
	payload := `{
		"results": [
			{
				"id": "1",
				"title": "Folder A",
				"createdAt": "2024-06-15T08:00:00.000Z",
				"version": {
					"number": 1,
					"createdAt": "2024-06-15T08:00:00.000Z"
				}
			},
			{
				"id": "2",
				"title": "Folder B",
				"createdAt": 1721486400000
			}
		],
		"_links": {
			"next": "/wiki/api/v2/folders?cursor=abc123"
		}
	}`

	var got FolderChunkScheme
	err := json.Unmarshal([]byte(payload), &got)
	require.NoError(t, err)

	assert.Len(t, got.Results, 2)
	assert.Equal(t, "1", got.Results[0].ID)
	assert.Equal(t, "Folder A", got.Results[0].Title)

	expectedTime := time.Date(2024, 6, 15, 8, 0, 0, 0, time.UTC)
	assert.True(t, expectedTime.Equal(time.Time(*got.Results[0].CreatedAt)))
	assert.True(t, expectedTime.Equal(time.Time(*got.Results[0].Version.CreatedAt)))

	assert.Equal(t, "2", got.Results[1].ID)
	assert.Equal(t, "/wiki/api/v2/folders?cursor=abc123", got.Links.Next)
}
