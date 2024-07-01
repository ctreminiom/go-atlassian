// Package models provides the data structures used in the admin package.
package models

import "time"

// OrganizationPolicyPageScheme represents a page of organization policies.
type OrganizationPolicyPageScheme struct {
	Data  []*OrganizationPolicyData `json:"data,omitempty"`  // The organization policies on this page.
	Links *LinkPageModelScheme      `json:"links,omitempty"` // Links to other pages.
	Meta  struct {
		Next     string `json:"next,omitempty"`      // The next page.
		PageSize int    `json:"page_size,omitempty"` // The page size.
	} `json:"meta,omitempty"` // Metadata about the page.
}

// OrganizationPolicyScheme represents an organization policy.
type OrganizationPolicyScheme struct {
	Data OrganizationPolicyData `json:"data,omitempty"` // The organization policy data.
}

// OrganizationPolicyResource represents a resource in an organization policy.
type OrganizationPolicyResource struct {
	ID                string `json:"id,omitempty"`                // The ID of the resource.
	ApplicationStatus string `json:"applicationStatus,omitempty"` // The application status of the resource.
}

// OrganizationPolicyAttributes represents the attributes of an organization policy.
type OrganizationPolicyAttributes struct {
	Type      string                        `json:"type,omitempty"`      // The type of the policy.
	Name      string                        `json:"name,omitempty"`      // The name of the policy.
	Status    string                        `json:"status,omitempty"`    // The status of the policy.
	Resources []*OrganizationPolicyResource `json:"resources,omitempty"` // The resources of the policy.
	CreatedAt time.Time                     `json:"createdAt,omitempty"` // The creation time of the policy.
	UpdatedAt time.Time                     `json:"updatedAt,omitempty"` // The update time of the policy.
}

// OrganizationPolicyData represents the data of an organization policy.
type OrganizationPolicyData struct {
	ID         string                        `json:"id,omitempty"`         // The ID of the policy.
	Type       string                        `json:"type,omitempty"`       // The type of the policy.
	Attributes *OrganizationPolicyAttributes `json:"attributes,omitempty"` // The attributes of the policy.
}
