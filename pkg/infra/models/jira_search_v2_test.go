package models

import (
	"encoding/json"
	"testing"
)

func TestIssueSearchJQLSchemeV2_SchemaUnmarshal(t *testing.T) {
	// Test data that simulates the API response with expanded schema
	testJSON := `{
		"startAt": 0,
		"maxResults": 10,
		"total": 1,
		"issues": [],
		"schema": {
			"summary": {
				"type": "string",
				"system": "summary"
			},
			"assignee": {
				"type": "user",
				"system": "assignee"
			},
			"status": {
				"type": "status",
				"system": "status"
			}
		}
	}`

	// Try to unmarshal into the V2 struct
	var result IssueSearchJQLSchemeV2
	err := json.Unmarshal([]byte(testJSON), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the schema was parsed correctly
	if result.Schema == nil {
		t.Fatal("Schema field is nil after unmarshaling")
	}

	// Check specific fields
	tests := []struct {
		field      string
		wantType   string
		wantSystem string
	}{
		{"summary", "string", "summary"},
		{"assignee", "user", "assignee"},
		{"status", "status", "status"},
	}

	for _, tc := range tests {
		t.Run(tc.field, func(t *testing.T) {
			schema, ok := result.Schema[tc.field]
			if !ok {
				t.Errorf("Schema field %s not found", tc.field)
				return
			}
			if schema.Type != tc.wantType {
				t.Errorf("Schema field %s: got type %s, want %s", tc.field, schema.Type, tc.wantType)
			}
			if schema.System != tc.wantSystem {
				t.Errorf("Schema field %s: got system %s, want %s", tc.field, schema.System, tc.wantSystem)
			}
		})
	}
}
