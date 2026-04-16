package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFolderScheme_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected FolderScheme
	}{
		{
			name: "full response from API",
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
			expected: FolderScheme{
				ID:         "123456",
				Type:       "folder",
				Status:     "current",
				Title:      "My Folder",
				ParentID:   "789",
				ParentType: "page",
				Position:   5,
				AuthorID:   "author-1",
				OwnerID:    "owner-1",
				CreatedAt:  "2024-09-23T20:17:35.607Z",
				SpaceID:    "space-42",
				Version: &FolderVersionScheme{
					CreatedAt: "2024-09-23T20:17:35.607Z",
					Message:   "initial version",
					Number:    1,
					MinorEdit: false,
					AuthorID:  "author-1",
				},
				Links: &FolderLinksScheme{
					Self: "https://example.atlassian.net/wiki/api/v2/folders/123456",
				},
			},
		},
		{
			name: "minimal response",
			json: `{
				"id": "123456",
				"title": "Minimal Folder",
				"createdAt": "2025-01-15T10:30:00.000Z"
			}`,
			expected: FolderScheme{
				ID:        "123456",
				Title:     "Minimal Folder",
				CreatedAt: "2025-01-15T10:30:00.000Z",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FolderScheme
			err := json.Unmarshal([]byte(tt.json), &got)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
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
				"createdAt": "2024-07-20T14:30:00.000Z"
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
	assert.Equal(t, "2024-06-15T08:00:00.000Z", got.Results[0].CreatedAt)
	assert.Equal(t, 1, got.Results[0].Version.Number)
	assert.Equal(t, "2024-06-15T08:00:00.000Z", got.Results[0].Version.CreatedAt)
	assert.Equal(t, "2", got.Results[1].ID)
	assert.Equal(t, "/wiki/api/v2/folders?cursor=abc123", got.Links.Next)
}
